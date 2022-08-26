package context

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"etri5gc/openapi"
	"fmt"
	"reflect"
	"regexp"
	"strconv"

	"etri5gc/openapi/consumers/ausf"
	"etri5gc/openapi/models"
	"etri5gc/openapi/utils/nasConvert"

	"github.com/free5gc/nas"
	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/nas/nasType"
	"github.com/free5gc/nas/security"
	"github.com/free5gc/ngap/ngapType"
	"github.com/free5gc/util/ueauth"
	"github.com/mitchellh/mapstructure"
)

type AusfInfo struct {
	AusfGroupId                       string
	AusfId                            string
	AusfUri                           string
	RoutingIndicator                  string
	AuthenticationCtx                 *models.UeAuthenticationCtx
	AuthFailureCauseSynchFailureTimes int
	IdentityRequestSendTimes          int
	ABBA                              []uint8
	Kseaf                             string
	Kamf                              string
}

type SecContext struct {
	SecurityContextAvailable bool
	UESecurityCapability     nasType.UESecurityCapability // for security command
	NgKsi                    models.NgKsi
	MacFailed                bool      // set to true if the integrity check of current NAS message is failed
	KnasInt                  [16]uint8 // 16 byte
	KnasEnc                  [16]uint8 // 16 byte
	Kgnb                     []uint8   // 32 byte
	Kn3iwf                   []uint8   // 32 byte
	NH                       []uint8   // 32 byte
	NCC                      uint8     // 0..7
	ULCount                  security.Count
	DLCount                  security.Count
	CipheringAlg             uint8
	IntegrityAlg             uint8
}

func (s *SecContext) isValid() bool {
	return s.SecurityContextAvailable && s.NgKsi.Ksi != nasMessage.NasKeySetIdentifierNoKeyIsAvailable && !s.MacFailed
}

type ausfClient struct {
	ue       *AmfUe
	info     AusfInfo
	seccon   SecContext
	consumer ausf.AusfConsumer
	//sender requestSender
}

func (c *ausfClient) Info() *AusfInfo {
	return &c.info
}
func (c *ausfClient) SecInfo() *SecContext {
	return &c.seccon
}

func (c *ausfClient) IsMacFailed() bool {
	return c.seccon.MacFailed
}
func (c *ausfClient) SecurityContextIsValid() bool {
	return c.seccon.isValid()
}

func (c *ausfClient) SecurityContextAvailable() bool {
	return c.seccon.SecurityContextAvailable
}

func (c *ausfClient) SetSecurityContextAvailable(f bool) {
	c.seccon.SecurityContextAvailable = f
}

// Kamf Derivation function defined in TS 33.501 Annex A.7
func (c *ausfClient) DerivateKamf() {
	supiRegexp, err := regexp.Compile("(?:imsi|supi)-([0-9]{5,15})")
	if err != nil {
		//	logger.ContextLog.Error(err)
		return
	}
	groups := supiRegexp.FindStringSubmatch(c.ue.Supi)
	if groups == nil {
		//	logger.NasLog.Errorln("supi is not correct")
		return
	}

	P0 := []byte(groups[1])
	L0 := ueauth.KDFLen(P0)
	P1 := c.info.ABBA
	L1 := ueauth.KDFLen(P1)

	KseafDecode, err := hex.DecodeString(c.info.Kseaf)
	if err != nil {
		//	logger.ContextLog.Error(err)
		return
	}
	KamfBytes, err := ueauth.GetKDFValue(KseafDecode, ueauth.FC_FOR_KAMF_DERIVATION, P0, L0, P1, L1)
	if err != nil {
		//	logger.ContextLog.Error(err)
		return
	}
	c.info.Kamf = hex.EncodeToString(KamfBytes)
}

// Algorithm key Derivation function defined in TS 33.501 Annex A.9
func (c *ausfClient) DerivateAlgKey() {
	// Security Key
	P0 := []byte{security.NNASEncAlg}
	L0 := ueauth.KDFLen(P0)
	P1 := []byte{c.seccon.CipheringAlg}
	L1 := ueauth.KDFLen(P1)

	KamfBytes, err := hex.DecodeString(c.info.Kamf)
	if err != nil {
		//logger.ContextLog.Error(err)
		return
	}
	kenc, err := ueauth.GetKDFValue(KamfBytes, ueauth.FC_FOR_ALGORITHM_KEY_DERIVATION, P0, L0, P1, L1)
	if err != nil {
		//logger.ContextLog.Error(err)
		return
	}
	copy(c.seccon.KnasEnc[:], kenc[16:32])

	// Integrity Key
	P0 = []byte{security.NNASIntAlg}
	L0 = ueauth.KDFLen(P0)
	P1 = []byte{c.seccon.IntegrityAlg}
	L1 = ueauth.KDFLen(P1)

	kint, err := ueauth.GetKDFValue(KamfBytes, ueauth.FC_FOR_ALGORITHM_KEY_DERIVATION, P0, L0, P1, L1)
	if err != nil {
		//	logger.ContextLog.Error(err)
		return
	}
	copy(c.seccon.KnasInt[:], kint[16:32])
}

// Access Network key Derivation function defined in TS 33.501 Annex A.9
func (c *ausfClient) DerivateAnKey(anType models.AccessType) {
	accessType := security.AccessType3GPP // Defalut 3gpp
	P0 := make([]byte, 4)
	binary.BigEndian.PutUint32(P0, c.seccon.ULCount.Get())
	L0 := ueauth.KDFLen(P0)
	if anType == models.AccessType_NON_3_GPP_ACCESS {
		accessType = security.AccessTypeNon3GPP
	}
	P1 := []byte{accessType}
	L1 := ueauth.KDFLen(P1)

	KamfBytes, err := hex.DecodeString(c.info.Kamf)
	if err != nil {
		//logger.ContextLog.Error(err)
		return
	}
	key, err := ueauth.GetKDFValue(KamfBytes, ueauth.FC_FOR_KGNB_KN3IWF_DERIVATION, P0, L0, P1, L1)
	if err != nil {
		//	logger.ContextLog.Error(err)
		return
	}
	switch accessType {
	case security.AccessType3GPP:
		c.seccon.Kgnb = key
	case security.AccessTypeNon3GPP:
		c.seccon.Kn3iwf = key
	}
}

func (c *ausfClient) UpdateNH() {
	c.seccon.NCC++
	c.DerivateNH(c.seccon.NH)
}

// NH Derivation function defined in TS 33.501 Annex A.10
func (c *ausfClient) DerivateNH(syncInput []byte) {
	P0 := syncInput
	L0 := ueauth.KDFLen(P0)

	KamfBytes, err := hex.DecodeString(c.info.Kamf)
	if err != nil {
		//logger.ContextLog.Error(err)
		return
	}
	c.seccon.NH, err = ueauth.GetKDFValue(KamfBytes, ueauth.FC_FOR_NH_DERIVATION, P0, L0)
	if err != nil {
		//logger.ContextLog.Error(err)
		return
	}
}

func (c *ausfClient) UpdateSecurityContext(anType models.AccessType) {
	c.DerivateAnKey(anType)
	switch anType {
	case models.AccessType__3_GPP_ACCESS:
		c.DerivateNH(c.seccon.Kgnb)
	case models.AccessType_NON_3_GPP_ACCESS:
		c.DerivateNH(c.seccon.Kn3iwf)
	}
	c.seccon.NCC = 1
}

func (c ausfClient) SetUeSecCap(seccap *ngapType.UESecurityCapabilities) {
	if seccap != nil {
		c.seccon.UESecurityCapability.SetEA1_128_5G(seccap.NRencryptionAlgorithms.Value.Bytes[0] & 0x80)
		c.seccon.UESecurityCapability.SetEA2_128_5G(seccap.NRencryptionAlgorithms.Value.Bytes[0] & 0x40)
		c.seccon.UESecurityCapability.SetEA3_128_5G(seccap.NRencryptionAlgorithms.Value.Bytes[0] & 0x20)
		c.seccon.UESecurityCapability.SetIA1_128_5G(seccap.NRintegrityProtectionAlgorithms.Value.Bytes[0] & 0x80)
		c.seccon.UESecurityCapability.SetIA2_128_5G(seccap.NRintegrityProtectionAlgorithms.Value.Bytes[0] & 0x40)
		c.seccon.UESecurityCapability.SetIA3_128_5G(seccap.NRintegrityProtectionAlgorithms.Value.Bytes[0] & 0x20)
		// not support any E-UTRA algorithms
	}

}
func (c *ausfClient) SelectSecurityAlg(intOrder, encOrder []uint8) {
	c.seccon.CipheringAlg = security.AlgCiphering128NEA0
	c.seccon.IntegrityAlg = security.AlgIntegrity128NIA0

	ueSupported := uint8(0)
	for _, intAlg := range intOrder {
		switch intAlg {
		case security.AlgIntegrity128NIA0:
			ueSupported = c.seccon.UESecurityCapability.GetIA0_5G()
		case security.AlgIntegrity128NIA1:
			ueSupported = c.seccon.UESecurityCapability.GetIA1_128_5G()
		case security.AlgIntegrity128NIA2:
			ueSupported = c.seccon.UESecurityCapability.GetIA2_128_5G()
		case security.AlgIntegrity128NIA3:
			ueSupported = c.seccon.UESecurityCapability.GetIA3_128_5G()
		}
		if ueSupported == 1 {
			c.seccon.IntegrityAlg = intAlg
			break
		}
	}

	ueSupported = uint8(0)
	for _, encAlg := range encOrder {
		switch encAlg {
		case security.AlgCiphering128NEA0:
			ueSupported = c.seccon.UESecurityCapability.GetEA0_5G()
		case security.AlgCiphering128NEA1:
			ueSupported = c.seccon.UESecurityCapability.GetEA1_128_5G()
		case security.AlgCiphering128NEA2:
			ueSupported = c.seccon.UESecurityCapability.GetEA2_128_5G()
		case security.AlgCiphering128NEA3:
			ueSupported = c.seccon.UESecurityCapability.GetEA3_128_5G()
		}
		if ueSupported == 1 {
			c.seccon.CipheringAlg = encAlg
			break
		}
	}
}

func (c *ausfClient) clear() {
	c.info.AuthFailureCauseSynchFailureTimes = 0
	c.info.IdentityRequestSendTimes = 0
}

///////////////////////////////////// consumer /////////////////////
//build a query to select an AUSF producer
func (c *ausfClient) Select() {
}

func (c *ausfClient) SendUEAuthRequest(resynchronizationInfo *models.ResynchronizationInfo) (*models.UeAuthenticationCtx, *models.ProblemDetails, error) {

	//1. Prepare parameters
	servedGuami := c.ue.amf.ServedGuami()

	var authInfo models.AuthenticationInfo
	authInfo.SupiOrSuci = c.ue.Suci
	if mnc, err := strconv.Atoi(servedGuami.PlmnId.Mnc); err != nil {
		return nil, nil, err
	} else {
		authInfo.ServingNetworkName = fmt.Sprintf("5G:mnc%03d.mcc%s.3gppnetwork.org", mnc, servedGuami.PlmnId.Mcc)
	}
	if resynchronizationInfo != nil {
		authInfo.ResynchronizationInfo = resynchronizationInfo
	}
	if authCtx, _, err := c.consumer.UeAuthPost(authInfo); err == nil {
		return &authCtx, nil, nil
	} else if errEx, ok := err.(*openapi.Error); ok {
		return nil, errEx.Problem(), err
	} else {
		return nil, nil, err
	}
}

func (c *ausfClient) SendAuth5gAkaConfirmRequest(resStar string) (
	*models.ConfirmationDataResponse, *models.ProblemDetails, error) {
	//NOTE: the ausf producer should return a relative path instead of a
	//absolute path
	//it seems it is not neccessary to extract the path for confirmation
	//path := c.info.AuthenticationCtx.Links["5g-aka"].Href)

	if result, _, err := c.consumer.UeAuthAuthCtxId5gAkaConfirmationPut(c.ue.Suci, resStar); err == nil {
		return &result, nil, nil
	} else if errEx, ok := err.(*openapi.Error); ok {
		return nil, errEx.Problem(), err
	} else {
		return nil, nil, err
	}
}

func (c *ausfClient) SendEapAuthConfirmRequest(eapMsg nasType.EAPMessage) (*models.EapSession, *models.ProblemDetails, error) {
	//confirmUri, err := url.Parse(ue.GetAusfInfo().AuthenticationCtx.Links["eap-session"].Href)

	eapin := &models.EapSession{
		EapPayload: base64.StdEncoding.EncodeToString(eapMsg.GetEAPMessage()),
	}

	if eapout, _, err := c.consumer.EapAuthMethod(c.ue.Suci, eapin); err == nil {
		return &eapout, nil, nil
	} else if errEx, ok := err.(*openapi.Error); ok {
		return nil, errEx.Problem(), err
	} else {
		return nil, nil, err
	}
}

func (c *ausfClient) NasEncode(msg *nas.Message, accessType models.AccessType) ([]byte, error) {
	if msg == nil {
		return nil, fmt.Errorf("Nas Message is empty")
	}

	// Plain NAS message
	if !c.seccon.SecurityContextAvailable {
		return msg.PlainNasEncode()
	} else {
		// Security protected NAS Message
		// a security protected NAS message must be integrity protected, and ciphering is optional
		needCiphering := false
		switch msg.SecurityHeader.SecurityHeaderType {
		case nas.SecurityHeaderTypeIntegrityProtected:
			//secinfo.NASLog.Debugln("Security header type: Integrity Protected")
		case nas.SecurityHeaderTypeIntegrityProtectedAndCiphered:
			//secinfo.NASLog.Debugln("Security header type: Integrity Protected And Ciphered")
			needCiphering = true
		case nas.SecurityHeaderTypeIntegrityProtectedWithNew5gNasSecurityContext:
			//secinfo.NASLog.Debugln("Security header type: Integrity Protected With New 5G Security Context")
			c.seccon.ULCount.Set(0, 0)
			c.seccon.DLCount.Set(0, 0)
		default:
			return nil, fmt.Errorf("Wrong security header type: 0x%0x", msg.SecurityHeader.SecurityHeaderType)
		}

		// encode plain nas first
		payload, err := msg.PlainNasEncode()
		if err != nil {
			return nil, fmt.Errorf("Plain NAS encode error: %+v", err)
		}

		//secinfo.NASLog.Tracef("plain payload:\n%+v", hex.Dump(payload))
		if needCiphering {
			//secinfo.NASLog.Debugf("Encrypt NAS message (algorithm: %+v, DLCount: 0x%0x)", secinfo.CipheringAlg, secinfo.DLCount.Get())
			//secinfo.NASLog.Tracef("NAS ciphering key: %0x", secinfo.KnasEnc)
			if err = security.NASEncrypt(c.seccon.CipheringAlg, c.seccon.KnasEnc, c.seccon.DLCount.Get(),
				GetBearerType(accessType), security.DirectionDownlink, payload); err != nil {
				return nil, fmt.Errorf("Encrypt error: %+v", err)
			}
		}

		// add sequece number
		payload = append([]byte{c.seccon.DLCount.SQN()}, payload[:]...)

		//secinfo.NASLog.Debugf("Calculate NAS MAC (algorithm: %+v, DLCount: 0x%0x)", secinfo.IntegrityAlg, secinfo.DLCount.Get())
		//secinfo.NASLog.Tracef("NAS integrity key: %0x", secinfo.KnasInt)
		mac32, err := security.NASMacCalculate(c.seccon.IntegrityAlg, c.seccon.KnasInt, c.seccon.DLCount.Get(),
			GetBearerType(accessType), security.DirectionDownlink, payload)
		if err != nil {
			return nil, fmt.Errorf("MAC calcuate error: %+v", err)
		}
		// Add mac value
		//secinfo.NASLog.Tracef("MAC: 0x%08x", mac32)
		payload = append(mac32, payload[:]...)

		// Add EPD and Security Type
		msgSecurityHeader := []byte{msg.SecurityHeader.ProtocolDiscriminator, msg.SecurityHeader.SecurityHeaderType}
		payload = append(msgSecurityHeader, payload[:]...)

		// Increase DL Count
		c.seccon.DLCount.AddOne()
		return payload, nil
	}
}

/*
payload either a security protected 5GS NAS message or a plain 5GS NAS message which
format is followed TS 24.501 9.1.1
*/
func (c *ausfClient) NasDecode(accessType models.AccessType, payload []byte) (*nas.Message, error) {
	if payload == nil {
		return nil, fmt.Errorf("Nas payload is empty")
	}

	msg := new(nas.Message)
	msg.SecurityHeaderType = nas.GetSecurityHeaderType(payload) & 0x0f
	if msg.SecurityHeaderType == nas.SecurityHeaderTypePlainNas {
		// RRCEstablishmentCause 0 is for emergency service
		if c.seccon.SecurityContextAvailable && c.ue.RanUe[accessType].RRCEstablishmentCause != "0" {
			c.seccon.MacFailed = false
			if err := msg.PlainNasDecode(&payload); err != nil {
				return nil, err
			}

			if msg.GmmMessage == nil {
				return nil, fmt.Errorf("Gmm Message is nil")
			}

			// TS 24.501 4.4.4.3: Except the messages listed below, no NAS signalling messages shall be processed
			// by the receiving 5GMM entity in the AMF or forwarded to the 5GSM entity, unless the secure exchange
			// of NAS messages has been established for the NAS signalling connection
			switch msg.GmmHeader.GetMessageType() {
			case nas.MsgTypeRegistrationRequest:
				return msg, nil
			case nas.MsgTypeIdentityResponse:
				return msg, nil
			case nas.MsgTypeAuthenticationResponse:
				return msg, nil
			case nas.MsgTypeAuthenticationFailure:
				return msg, nil
			case nas.MsgTypeSecurityModeReject:
				return msg, nil
			case nas.MsgTypeDeregistrationRequestUEOriginatingDeregistration:
				return msg, nil
			case nas.MsgTypeDeregistrationAcceptUETerminatedDeregistration:
				return msg, nil
			default:
				return nil, fmt.Errorf(
					"UE can not send plain nas for non-emergency service when there is a valid security context")
			}
		} else {
			c.seccon.MacFailed = false
			err := msg.PlainNasDecode(&payload)
			return msg, err
		}
	} else { // Security protected NAS message
		securityHeader := payload[0:6]
		//secinfo.NASLog.Traceln("securityHeader is ", securityHeader)
		sequenceNumber := payload[6]
		//secinfo.NASLog.Traceln("sequenceNumber", sequenceNumber)

		receivedMac32 := securityHeader[2:]
		// remove security Header except for sequece Number
		payload = payload[6:]

		// a security protected NAS message must be integrity protected, and ciphering is optional
		ciphered := false
		switch msg.SecurityHeaderType {
		case nas.SecurityHeaderTypeIntegrityProtected:
			//secinfo.NASLog.Debugln("Security header type: Integrity Protected")
		case nas.SecurityHeaderTypeIntegrityProtectedAndCiphered:
			//secinfo.NASLog.Debugln("Security header type: Integrity Protected And Ciphered")
			ciphered = true
		case nas.SecurityHeaderTypeIntegrityProtectedAndCipheredWithNew5gNasSecurityContext:
			//secinfo.NASLog.Debugln("Security header type: Integrity Protected And Ciphered With New 5G Security Context")
			ciphered = true
			c.seccon.ULCount.Set(0, 0)
		default:
			return nil, fmt.Errorf("Wrong security header type: 0x%0x", msg.SecurityHeader.SecurityHeaderType)
		}

		if c.seccon.ULCount.SQN() > sequenceNumber {
			//secinfo.NASLog.Debugf("set ULCount overflow")
			c.seccon.ULCount.SetOverflow(c.seccon.ULCount.Overflow() + 1)
		}
		c.seccon.ULCount.SetSQN(sequenceNumber)

		//secinfo.NASLog.Tracef("NAS integrity key0x: %0x", secinfo.KnasInt)
		mac32, err := security.NASMacCalculate(c.seccon.IntegrityAlg, c.seccon.KnasInt, c.seccon.ULCount.Get(),
			GetBearerType(accessType), security.DirectionUplink, payload)
		if err != nil {
			return nil, fmt.Errorf("MAC calcuate error: %+v", err)
		}

		if !reflect.DeepEqual(mac32, receivedMac32) {
			//secinfo.NASLog.Warnf("NAS MAC verification failed(received: 0x%08x, expected: 0x%08x)", receivedMac32, mac32)
			c.seccon.MacFailed = true
		} else {
			//secinfo.NASLog.Tracef("cmac value: 0x%08x", mac32)
			c.seccon.MacFailed = false
		}

		if ciphered {
			//secinfo.NASLog.Debugf("Decrypt NAS message (algorithm: %+v, ULCount: 0x%0x)", secinfo.CipheringAlg, secinfo.ULCount.Get())
			//secinfo.NASLog.Tracef("NAS ciphering key: %0x", secinfo.KnasEnc)
			// decrypt payload without sequence number (payload[1])
			if err = security.NASEncrypt(c.seccon.CipheringAlg, c.seccon.KnasEnc, c.seccon.ULCount.Get(), GetBearerType(accessType),
				security.DirectionUplink, payload[1:]); err != nil {
				return nil, fmt.Errorf("Encrypt error: %+v", err)
			}
		}

		// remove sequece Number
		payload = payload[1:]
		err = msg.PlainNasDecode(&payload)
		return msg, err
	}
}

func GetBearerType(accessType models.AccessType) uint8 {
	if accessType == models.AccessType__3_GPP_ACCESS {
		return security.Bearer3GPP
	} else if accessType == models.AccessType_NON_3_GPP_ACCESS {
		return security.BearerNon3GPP
	} else {
		return security.OnlyOneBearer
	}
}

func (c *ausfClient) BuildAuthResult(content *nasMessage.AuthenticationResult, eapSuccess bool, eapMsg string) error {
	content.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	content.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	content.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	content.AuthenticationResultMessageIdentity.SetMessageType(nas.MsgTypeAuthenticationResult)
	content.SpareHalfOctetAndNgksi = nasConvert.SpareHalfOctetAndNgksiToNas(c.seccon.NgKsi)
	rawEapMsg, err := base64.StdEncoding.DecodeString(eapMsg)
	if err != nil {
		return err
	}

	content.EAPMessage.SetLen(uint16(len(rawEapMsg)))
	content.EAPMessage.SetEAPMessage(rawEapMsg)
	if eapSuccess {
		content.ABBA = nasType.NewABBA(nasMessage.AuthenticationResultABBAType)
		content.ABBA.SetLen(uint8(len(c.info.ABBA)))
		content.ABBA.SetABBAContents(c.info.ABBA)
	}
	return nil
}

func (c *ausfClient) BuildAuthRequest(req *nasMessage.AuthenticationRequest) error {
	req.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	req.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	req.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	req.AuthenticationRequestMessageIdentity.SetMessageType(nas.MsgTypeAuthenticationRequest)
	req.SpareHalfOctetAndNgksi = nasConvert.SpareHalfOctetAndNgksiToNas(c.seccon.NgKsi)
	req.ABBA.SetLen(uint8(len(c.info.ABBA)))
	req.ABBA.SetABBAContents(c.info.ABBA)

	switch c.info.AuthenticationCtx.AuthType {
	case models.AuthType__5_G_AKA:
		var tmpArray [16]byte
		var av5gAka models.Av5gAka

		if err := mapstructure.Decode(c.info.AuthenticationCtx.Var5gAuthData, &av5gAka); err != nil {
			log.Error("Var5gAuthData Convert Type Error")
			return err
		}

		rand, err := hex.DecodeString(av5gAka.Rand)
		if err != nil {
			return err
		}
		req.AuthenticationParameterRAND =
			nasType.NewAuthenticationParameterRAND(nasMessage.AuthenticationRequestAuthenticationParameterRANDType)
		copy(tmpArray[:], rand[0:16])
		req.AuthenticationParameterRAND.SetRANDValue(tmpArray)

		autn, err := hex.DecodeString(av5gAka.Autn)
		if err != nil {
			return err
		}
		req.AuthenticationParameterAUTN =
			nasType.NewAuthenticationParameterAUTN(nasMessage.AuthenticationRequestAuthenticationParameterAUTNType)
		req.AuthenticationParameterAUTN.SetLen(uint8(len(autn)))
		copy(tmpArray[:], autn[0:16])
		req.AuthenticationParameterAUTN.SetAUTN(tmpArray)
	case models.AuthType_EAP_AKA_PRIME:
		eapMsg := c.info.AuthenticationCtx.Var5gAuthData.(string)
		rawEapMsg, err := base64.StdEncoding.DecodeString(eapMsg)
		if err != nil {
			return err
		}
		req.EAPMessage = nasType.NewEAPMessage(nasMessage.AuthenticationRequestEAPMessageType)
		req.EAPMessage.SetLen(uint16(len(rawEapMsg)))
		req.EAPMessage.SetEAPMessage(rawEapMsg)
	}
	return nil
}

func (c *ausfClient) BuildSecModeCmd(cmd *nasMessage.SecurityModeCommand, eapSuccess bool, eapMessage string) error {
	cmd.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	cmd.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nas.SecurityHeaderTypePlainNas)
	cmd.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	cmd.SecurityModeCommandMessageIdentity.SetMessageType(nas.MsgTypeSecurityModeCommand)

	cmd.SelectedNASSecurityAlgorithms.SetTypeOfCipheringAlgorithm(c.seccon.CipheringAlg)
	cmd.SelectedNASSecurityAlgorithms.SetTypeOfIntegrityProtectionAlgorithm(c.seccon.IntegrityAlg)

	cmd.SpareHalfOctetAndNgksi = nasConvert.SpareHalfOctetAndNgksiToNas(c.seccon.NgKsi)

	cmd.ReplayedUESecurityCapabilities.SetLen(c.seccon.UESecurityCapability.GetLen())
	cmd.ReplayedUESecurityCapabilities.Buffer = c.seccon.UESecurityCapability.Buffer

	if c.ue.Pei != "" {
		cmd.IMEISVRequest = nasType.NewIMEISVRequest(nasMessage.SecurityModeCommandIMEISVRequestType)
		cmd.IMEISVRequest.SetIMEISVRequestValue(nasMessage.IMEISVNotRequested)
	} else {
		cmd.IMEISVRequest = nasType.NewIMEISVRequest(nasMessage.SecurityModeCommandIMEISVRequestType)
		cmd.IMEISVRequest.SetIMEISVRequestValue(nasMessage.IMEISVRequested)
	}

	cmd.Additional5GSecurityInformation =
		nasType.NewAdditional5GSecurityInformation(nasMessage.SecurityModeCommandAdditional5GSecurityInformationType)
	cmd.Additional5GSecurityInformation.SetLen(1)
	if c.ue.RetransmissionOfInitialNASMsg {
		cmd.Additional5GSecurityInformation.SetRINMR(1)
	} else {
		cmd.Additional5GSecurityInformation.SetRINMR(0)
	}

	if c.ue.RegistrationType5GS == nasMessage.RegistrationType5GSPeriodicRegistrationUpdating ||
		c.ue.RegistrationType5GS == nasMessage.RegistrationType5GSMobilityRegistrationUpdating {
		cmd.Additional5GSecurityInformation.SetHDP(1)
	} else {
		cmd.Additional5GSecurityInformation.SetHDP(0)
	}

	if eapMessage != "" {
		cmd.EAPMessage = nasType.NewEAPMessage(nasMessage.SecurityModeCommandEAPMessageType)
		rawEapMsg, err := base64.StdEncoding.DecodeString(eapMessage)
		if err != nil {
			return err
		}
		cmd.EAPMessage.SetLen(uint16(len(rawEapMsg)))
		cmd.EAPMessage.SetEAPMessage(rawEapMsg)

		if eapSuccess {
			cmd.ABBA = nasType.NewABBA(nasMessage.SecurityModeCommandABBAType)
			cmd.ABBA.SetLen(uint8(len(c.info.ABBA)))
			cmd.ABBA.SetABBAContents(c.info.ABBA)
		}
	}
	return nil
}
