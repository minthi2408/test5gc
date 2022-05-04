package nas

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/mitchellh/mapstructure"
	
	"etri5gc/nfs/amf/context"
	"etri5gc/nfs/amf/nas/nas_security"
	nasproto "github.com/free5gc/nas"
	"github.com/free5gc/nas/nasConvert"
	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/nas/nasType"
	"github.com/free5gc/openapi/models"
)

func (builder *Nas) BuildDLNASTransport(ue *context.AmfUe, accessType models.AccessType, payloadContainerType uint8, nasPdu []byte,
	pduSessionId uint8, cause *uint8, backoffTimerUint *uint8, backoffTimer uint8) ([]byte, error) {
	m := nasproto.NewMessage()
	m.GmmMessage = nasproto.NewGmmMessage()
	m.GmmHeader.SetMessageType(nasproto.MsgTypeDLNASTransport)

	m.SecurityHeader = nasproto.SecurityHeader{
		ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		SecurityHeaderType:    nasproto.SecurityHeaderTypeIntegrityProtectedAndCiphered,
	}

	dLNASTransport := nasMessage.NewDLNASTransport(0)
	dLNASTransport.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nasproto.SecurityHeaderTypePlainNas)
	dLNASTransport.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	dLNASTransport.SetMessageType(nasproto.MsgTypeDLNASTransport)
	dLNASTransport.SpareHalfOctetAndPayloadContainerType.SetPayloadContainerType(payloadContainerType)
	dLNASTransport.PayloadContainer.SetLen(uint16(len(nasPdu)))
	dLNASTransport.PayloadContainer.SetPayloadContainerContents(nasPdu)

	if pduSessionId != 0 {
		dLNASTransport.PduSessionID2Value = new(nasType.PduSessionID2Value)
		dLNASTransport.PduSessionID2Value.SetIei(nasMessage.DLNASTransportPduSessionID2ValueType)
		dLNASTransport.PduSessionID2Value.SetPduSessionID2Value(pduSessionId)
	}
	if cause != nil {
		dLNASTransport.Cause5GMM = new(nasType.Cause5GMM)
		dLNASTransport.Cause5GMM.SetIei(nasMessage.DLNASTransportCause5GMMType)
		dLNASTransport.Cause5GMM.SetCauseValue(*cause)
	}
	if backoffTimerUint != nil {
		dLNASTransport.BackoffTimerValue = new(nasType.BackoffTimerValue)
		dLNASTransport.BackoffTimerValue.SetIei(nasMessage.DLNASTransportBackoffTimerValueType)
		dLNASTransport.BackoffTimerValue.SetLen(1)
		dLNASTransport.BackoffTimerValue.SetUnitTimerValue(*backoffTimerUint)
		dLNASTransport.BackoffTimerValue.SetTimerValue(backoffTimer)
	}

	m.GmmMessage.DLNASTransport = dLNASTransport

	return nas_security.Encode(ue, m, accessType)
}

func (builder *Nas) BuildNotification(ue *context.AmfUe, accessType models.AccessType) ([]byte, error) {
	m := nasproto.NewMessage()
	m.GmmMessage = nasproto.NewGmmMessage()
	m.GmmHeader.SetMessageType(nasproto.MsgTypeNotification)

	m.SecurityHeader = nasproto.SecurityHeader{
		ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		SecurityHeaderType:    nasproto.SecurityHeaderTypeIntegrityProtectedAndCiphered,
	}

	notification := nasMessage.NewNotification(0)
	notification.SetSecurityHeaderType(nasproto.SecurityHeaderTypePlainNas)
	notification.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	notification.SetMessageType(nasproto.MsgTypeNotification)
	if accessType == models.AccessType__3_GPP_ACCESS {
		notification.SetAccessType(nasMessage.AccessType3GPP)
	} else {
		notification.SetAccessType(nasMessage.AccessTypeNon3GPP)
	}

	m.GmmMessage.Notification = notification

	return nas_security.Encode(ue, m, accessType)
}

func (builder *Nas) BuildIdentityRequest(ue *context.AmfUe, accessType models.AccessType, typeOfIdentity uint8) ([]byte, error) {
	m := nasproto.NewMessage()
	m.GmmMessage = nasproto.NewGmmMessage()
	m.GmmHeader.SetMessageType(nasproto.MsgTypeIdentityRequest)

	if ue.GetSecInfo().SecurityContextAvailable {
		m.SecurityHeader = nasproto.SecurityHeader{
			ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
			SecurityHeaderType:    nasproto.SecurityHeaderTypeIntegrityProtectedAndCiphered,
		}
	}

	identityRequest := nasMessage.NewIdentityRequest(0)
	identityRequest.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	identityRequest.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nasproto.SecurityHeaderTypePlainNas)
	identityRequest.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	identityRequest.IdentityRequestMessageIdentity.SetMessageType(nasproto.MsgTypeIdentityRequest)
	identityRequest.SpareHalfOctetAndIdentityType.SetTypeOfIdentity(typeOfIdentity)

	m.GmmMessage.IdentityRequest = identityRequest

	return nas_security.Encode(ue, m, accessType)
}

func (builder *Nas) BuildAuthenticationRequest(ue *context.AmfUe) ([]byte, error) {
	m := nasproto.NewMessage()
	m.GmmMessage = nasproto.NewGmmMessage()
	m.GmmHeader.SetMessageType(nasproto.MsgTypeAuthenticationRequest)

	ausfinfo := ue.GetAusfInfo()

	authenticationRequest := nasMessage.NewAuthenticationRequest(0)
	authenticationRequest.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	authenticationRequest.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nasproto.SecurityHeaderTypePlainNas)
	authenticationRequest.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	authenticationRequest.AuthenticationRequestMessageIdentity.SetMessageType(nasproto.MsgTypeAuthenticationRequest)
	authenticationRequest.SpareHalfOctetAndNgksi = nasConvert.SpareHalfOctetAndNgksiToNas(ue.GetSecInfo().NgKsi)
	authenticationRequest.ABBA.SetLen(uint8(len(ausfinfo.ABBA)))
	authenticationRequest.ABBA.SetABBAContents(ausfinfo.ABBA)

	switch ausfinfo.AuthenticationCtx.AuthType {
	case models.AuthType__5_G_AKA:
		var tmpArray [16]byte
		var av5gAka models.Av5gAka

		if err := mapstructure.Decode(ausfinfo.AuthenticationCtx.Var5gAuthData, &av5gAka); err != nil {
			log.Error("Var5gAuthData Convert Type Error")
			return nil, err
		}

		rand, err := hex.DecodeString(av5gAka.Rand)
		if err != nil {
			return nil, err
		}
		authenticationRequest.AuthenticationParameterRAND =
			nasType.NewAuthenticationParameterRAND(nasMessage.AuthenticationRequestAuthenticationParameterRANDType)
		copy(tmpArray[:], rand[0:16])
		authenticationRequest.AuthenticationParameterRAND.SetRANDValue(tmpArray)

		autn, err := hex.DecodeString(av5gAka.Autn)
		if err != nil {
			return nil, err
		}
		authenticationRequest.AuthenticationParameterAUTN =
			nasType.NewAuthenticationParameterAUTN(nasMessage.AuthenticationRequestAuthenticationParameterAUTNType)
		authenticationRequest.AuthenticationParameterAUTN.SetLen(uint8(len(autn)))
		copy(tmpArray[:], autn[0:16])
		authenticationRequest.AuthenticationParameterAUTN.SetAUTN(tmpArray)
	case models.AuthType_EAP_AKA_PRIME:
		eapMsg := ausfinfo.AuthenticationCtx.Var5gAuthData.(string)
		rawEapMsg, err := base64.StdEncoding.DecodeString(eapMsg)
		if err != nil {
			return nil, err
		}
		authenticationRequest.EAPMessage = nasType.NewEAPMessage(nasMessage.AuthenticationRequestEAPMessageType)
		authenticationRequest.EAPMessage.SetLen(uint16(len(rawEapMsg)))
		authenticationRequest.EAPMessage.SetEAPMessage(rawEapMsg)
	}

	m.GmmMessage.AuthenticationRequest = authenticationRequest

	return m.PlainNasEncode()
}

func (builder *Nas) BuildServiceAccept(ue *context.AmfUe, accessType models.AccessType, pDUSessionStatus *[16]bool,
	reactivationResult *[16]bool, errPduSessionId, errCause []uint8) ([]byte, error) {
	m := nasproto.NewMessage()
	m.GmmMessage = nasproto.NewGmmMessage()
	m.GmmHeader.SetMessageType(nasproto.MsgTypeServiceAccept)

	m.SecurityHeader = nasproto.SecurityHeader{
		ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		SecurityHeaderType:    nasproto.SecurityHeaderTypeIntegrityProtectedAndCiphered,
	}

	serviceAccept := nasMessage.NewServiceAccept(0)
	serviceAccept.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	serviceAccept.SetSecurityHeaderType(nasproto.SecurityHeaderTypePlainNas)
	serviceAccept.SetMessageType(nasproto.MsgTypeServiceAccept)
	if pDUSessionStatus != nil {
		serviceAccept.PDUSessionStatus = new(nasType.PDUSessionStatus)
		serviceAccept.PDUSessionStatus.SetIei(nasMessage.ServiceAcceptPDUSessionStatusType)
		serviceAccept.PDUSessionStatus.SetLen(2)
		serviceAccept.PDUSessionStatus.Buffer = nasConvert.PSIToBuf(*pDUSessionStatus)
	}
	if reactivationResult != nil {
		serviceAccept.PDUSessionReactivationResult = new(nasType.PDUSessionReactivationResult)
		serviceAccept.PDUSessionReactivationResult.SetIei(nasMessage.ServiceAcceptPDUSessionReactivationResultType)
		serviceAccept.PDUSessionReactivationResult.SetLen(2)
		serviceAccept.PDUSessionReactivationResult.Buffer = nasConvert.PSIToBuf(*reactivationResult)
	}
	if errPduSessionId != nil {
		serviceAccept.PDUSessionReactivationResultErrorCause = new(nasType.PDUSessionReactivationResultErrorCause)
		serviceAccept.PDUSessionReactivationResultErrorCause.SetIei(
			nasMessage.ServiceAcceptPDUSessionReactivationResultErrorCauseType)
		buf := nasConvert.PDUSessionReactivationResultErrorCauseToBuf(errPduSessionId, errCause)
		serviceAccept.PDUSessionReactivationResultErrorCause.SetLen(uint16(len(buf)))
		serviceAccept.PDUSessionReactivationResultErrorCause.Buffer = buf
	}
	m.GmmMessage.ServiceAccept = serviceAccept

	return nas_security.Encode(ue, m, accessType)
}

func (builder *Nas) BuildAuthenticationReject(ue *context.AmfUe, eapMsg string) ([]byte, error) {
	m := nasproto.NewMessage()
	m.GmmMessage = nasproto.NewGmmMessage()
	m.GmmHeader.SetMessageType(nasproto.MsgTypeAuthenticationReject)

	authenticationReject := nasMessage.NewAuthenticationReject(0)
	authenticationReject.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	authenticationReject.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nasproto.SecurityHeaderTypePlainNas)
	authenticationReject.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	authenticationReject.AuthenticationRejectMessageIdentity.SetMessageType(nasproto.MsgTypeAuthenticationReject)

	if eapMsg != "" {
		rawEapMsg, err := base64.StdEncoding.DecodeString(eapMsg)
		if err != nil {
			return nil, err
		}
		authenticationReject.EAPMessage = nasType.NewEAPMessage(nasMessage.AuthenticationRejectEAPMessageType)
		authenticationReject.EAPMessage.SetLen(uint16(len(rawEapMsg)))
		authenticationReject.EAPMessage.SetEAPMessage(rawEapMsg)
	}

	m.GmmMessage.AuthenticationReject = authenticationReject

	return m.PlainNasEncode()
}

func (builder *Nas) BuildAuthenticationResult(ue *context.AmfUe, eapSuccess bool, eapMsg string) ([]byte, error) {
	m := nasproto.NewMessage()
	m.GmmMessage = nasproto.NewGmmMessage()
	m.GmmHeader.SetMessageType(nasproto.MsgTypeAuthenticationResult)

	authenticationResult := nasMessage.NewAuthenticationResult(0)
	authenticationResult.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	authenticationResult.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nasproto.SecurityHeaderTypePlainNas)
	authenticationResult.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	authenticationResult.AuthenticationResultMessageIdentity.SetMessageType(nasproto.MsgTypeAuthenticationResult)
	authenticationResult.SpareHalfOctetAndNgksi = nasConvert.SpareHalfOctetAndNgksiToNas(ue.GetSecInfo().NgKsi)
	rawEapMsg, err := base64.StdEncoding.DecodeString(eapMsg)
	if err != nil {
		return nil, err
	}
	authenticationResult.EAPMessage.SetLen(uint16(len(rawEapMsg)))
	authenticationResult.EAPMessage.SetEAPMessage(rawEapMsg)
	ausfinfo := ue.GetAusfInfo()
	if eapSuccess {
		authenticationResult.ABBA = nasType.NewABBA(nasMessage.AuthenticationResultABBAType)
		authenticationResult.ABBA.SetLen(uint8(len(ausfinfo.ABBA)))
		authenticationResult.ABBA.SetABBAContents(ausfinfo.ABBA)
	}

	m.GmmMessage.AuthenticationResult = authenticationResult

	return m.PlainNasEncode()
}

// T3346 Timer and EAP are not Supported
func (builder *Nas) BuildServiceReject(pDUSessionStatus *[16]bool, cause uint8) ([]byte, error) {
	m := nasproto.NewMessage()
	m.GmmMessage = nasproto.NewGmmMessage()
	m.GmmHeader.SetMessageType(nasproto.MsgTypeServiceReject)

	serviceReject := nasMessage.NewServiceReject(0)
	serviceReject.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	serviceReject.SetSecurityHeaderType(nasproto.SecurityHeaderTypePlainNas)
	serviceReject.SetMessageType(nasproto.MsgTypeServiceReject)
	serviceReject.SetCauseValue(cause)
	if pDUSessionStatus != nil {
		serviceReject.PDUSessionStatus = new(nasType.PDUSessionStatus)
		serviceReject.PDUSessionStatus.SetIei(nasMessage.ServiceAcceptPDUSessionStatusType)
		serviceReject.PDUSessionStatus.SetLen(2)
		serviceReject.PDUSessionStatus.Buffer = nasConvert.PSIToBuf(*pDUSessionStatus)
	}

	m.GmmMessage.ServiceReject = serviceReject

	return m.PlainNasEncode()
}

// T3346 timer are not supported
func (builder *Nas) BuildRegistrationReject(ue *context.AmfUe, cause5GMM uint8, eapMessage string) ([]byte, error) {
	m := nasproto.NewMessage()
	m.GmmMessage = nasproto.NewGmmMessage()
	m.GmmHeader.SetMessageType(nasproto.MsgTypeRegistrationReject)

	registrationReject := nasMessage.NewRegistrationReject(0)
	registrationReject.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	registrationReject.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nasproto.SecurityHeaderTypePlainNas)
	registrationReject.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	registrationReject.RegistrationRejectMessageIdentity.SetMessageType(nasproto.MsgTypeRegistrationReject)
	registrationReject.Cause5GMM.SetCauseValue(cause5GMM)

	if ue.T3502Value != 0 {
		registrationReject.T3502Value = nasType.NewT3502Value(nasMessage.RegistrationRejectT3502ValueType)
		registrationReject.T3502Value.SetLen(1)
		t3502 := nasConvert.GPRSTimer2ToNas(ue.T3502Value)
		registrationReject.T3502Value.SetGPRSTimer2Value(t3502)
	}

	if eapMessage != "" {
		registrationReject.EAPMessage = nasType.NewEAPMessage(nasMessage.RegistrationRejectEAPMessageType)
		rawEapMsg, err := base64.StdEncoding.DecodeString(eapMessage)
		if err != nil {
			return nil, err
		}
		registrationReject.EAPMessage.SetLen(uint16(len(rawEapMsg)))
		registrationReject.EAPMessage.SetEAPMessage(rawEapMsg)
	}

	m.GmmMessage.RegistrationReject = registrationReject

	return m.PlainNasEncode()
}

// TS 24.501 8.2.25
func (builder *Nas) BuildSecurityModeCommand(ue *context.AmfUe, accessType models.AccessType, eapSuccess bool, eapMessage string) (
	[]byte, error) {
	m := nasproto.NewMessage()
	m.GmmMessage = nasproto.NewGmmMessage()
	m.GmmHeader.SetMessageType(nasproto.MsgTypeSecurityModeCommand)

	m.SecurityHeader = nasproto.SecurityHeader{
		ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		SecurityHeaderType:    nasproto.SecurityHeaderTypeIntegrityProtectedWithNew5gNasSecurityContext,
	}
	secinfo := ue.GetSecInfo()

	securityModeCommand := nasMessage.NewSecurityModeCommand(0)
	securityModeCommand.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	securityModeCommand.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nasproto.SecurityHeaderTypePlainNas)
	securityModeCommand.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	securityModeCommand.SecurityModeCommandMessageIdentity.SetMessageType(nasproto.MsgTypeSecurityModeCommand)

	securityModeCommand.SelectedNASSecurityAlgorithms.SetTypeOfCipheringAlgorithm(secinfo.CipheringAlg)
	securityModeCommand.SelectedNASSecurityAlgorithms.SetTypeOfIntegrityProtectionAlgorithm(secinfo.IntegrityAlg)

	securityModeCommand.SpareHalfOctetAndNgksi = nasConvert.SpareHalfOctetAndNgksiToNas(secinfo.NgKsi)

	securityModeCommand.ReplayedUESecurityCapabilities.SetLen(secinfo.UESecurityCapability.GetLen())
	securityModeCommand.ReplayedUESecurityCapabilities.Buffer = secinfo.UESecurityCapability.Buffer

	if ue.Pei != "" {
		securityModeCommand.IMEISVRequest = nasType.NewIMEISVRequest(nasMessage.SecurityModeCommandIMEISVRequestType)
		securityModeCommand.IMEISVRequest.SetIMEISVRequestValue(nasMessage.IMEISVNotRequested)
	} else {
		securityModeCommand.IMEISVRequest = nasType.NewIMEISVRequest(nasMessage.SecurityModeCommandIMEISVRequestType)
		securityModeCommand.IMEISVRequest.SetIMEISVRequestValue(nasMessage.IMEISVRequested)
	}

	securityModeCommand.Additional5GSecurityInformation =
		nasType.NewAdditional5GSecurityInformation(nasMessage.SecurityModeCommandAdditional5GSecurityInformationType)
	securityModeCommand.Additional5GSecurityInformation.SetLen(1)
	if ue.RetransmissionOfInitialNASMsg {
		securityModeCommand.Additional5GSecurityInformation.SetRINMR(1)
	} else {
		securityModeCommand.Additional5GSecurityInformation.SetRINMR(0)
	}

	if ue.RegistrationType5GS == nasMessage.RegistrationType5GSPeriodicRegistrationUpdating ||
		ue.RegistrationType5GS == nasMessage.RegistrationType5GSMobilityRegistrationUpdating {
		securityModeCommand.Additional5GSecurityInformation.SetHDP(1)
	} else {
		securityModeCommand.Additional5GSecurityInformation.SetHDP(0)
	}

	if eapMessage != "" {
		securityModeCommand.EAPMessage = nasType.NewEAPMessage(nasMessage.SecurityModeCommandEAPMessageType)
		rawEapMsg, err := base64.StdEncoding.DecodeString(eapMessage)
		if err != nil {
			return nil, err
		}
		securityModeCommand.EAPMessage.SetLen(uint16(len(rawEapMsg)))
		securityModeCommand.EAPMessage.SetEAPMessage(rawEapMsg)

		if eapSuccess {
			ausfinfo := ue.GetAusfInfo()
			securityModeCommand.ABBA = nasType.NewABBA(nasMessage.SecurityModeCommandABBAType)
			securityModeCommand.ABBA.SetLen(uint8(len(ausfinfo.ABBA)))
			securityModeCommand.ABBA.SetABBAContents(ausfinfo.ABBA)
		}
	}

	secinfo.SecurityContextAvailable = true
	m.GmmMessage.SecurityModeCommand = securityModeCommand
	payload, err := nas_security.Encode(ue, m, accessType)
	if err != nil {
		secinfo.SecurityContextAvailable = false
		return nil, err
	} else {
		return payload, nil
	}
}

// T3346 timer are not supported
func (builder *Nas) BuildDeregistrationRequest(ue *context.RanUe, accessType uint8, reRegistrationRequired bool,
	cause5GMM uint8) ([]byte, error) {
	m := nasproto.NewMessage()
	m.GmmMessage = nasproto.NewGmmMessage()
	m.GmmHeader.SetMessageType(nasproto.MsgTypeDeregistrationRequestUETerminatedDeregistration)

	deregistrationRequest := nasMessage.NewDeregistrationRequestUETerminatedDeregistration(0)
	deregistrationRequest.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	deregistrationRequest.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nasproto.SecurityHeaderTypePlainNas)
	deregistrationRequest.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	deregistrationRequest.SetMessageType(nasproto.MsgTypeDeregistrationRequestUETerminatedDeregistration)

	deregistrationRequest.SetAccessType(accessType)
	deregistrationRequest.SetSwitchOff(0)
	if reRegistrationRequired {
		deregistrationRequest.SetReRegistrationRequired(nasMessage.ReRegistrationRequired)
	} else {
		deregistrationRequest.SetReRegistrationRequired(nasMessage.ReRegistrationNotRequired)
	}

	if cause5GMM != 0 {
		deregistrationRequest.Cause5GMM = nasType.NewCause5GMM(
			nasMessage.DeregistrationRequestUETerminatedDeregistrationCause5GMMType)
		deregistrationRequest.Cause5GMM.SetCauseValue(cause5GMM)
	}
	m.GmmMessage.DeregistrationRequestUETerminatedDeregistration = deregistrationRequest

	if ue != nil && ue.AmfUe != nil {
		m.SecurityHeader = nasproto.SecurityHeader{
			ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
			SecurityHeaderType:    nasproto.SecurityHeaderTypeIntegrityProtectedAndCiphered,
		}
		var anType models.AccessType
		if accessType == 0x01 {
			anType = models.AccessType__3_GPP_ACCESS
		} else if accessType == 0x02 {
			anType = models.AccessType_NON_3_GPP_ACCESS
		}
		return nas_security.Encode(ue.AmfUe, m, anType)
	}
	return m.PlainNasEncode()
}

func (builder *Nas) BuildDeregistrationAccept() ([]byte, error) {
	m := nasproto.NewMessage()
	m.GmmMessage = nasproto.NewGmmMessage()
	m.GmmHeader.SetMessageType(nasproto.MsgTypeDeregistrationAcceptUEOriginatingDeregistration)

	deregistrationAccept := nasMessage.NewDeregistrationAcceptUEOriginatingDeregistration(0)
	deregistrationAccept.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	deregistrationAccept.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nasproto.SecurityHeaderTypePlainNas)
	deregistrationAccept.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	deregistrationAccept.SetMessageType(nasproto.MsgTypeDeregistrationAcceptUEOriginatingDeregistration)

	m.GmmMessage.DeregistrationAcceptUEOriginatingDeregistration = deregistrationAccept

	return m.PlainNasEncode()
}

func (builder *Nas) BuildRegistrationAccept(
	ue *context.AmfUe,
	anType models.AccessType,
	pDUSessionStatus *[16]bool,
	reactivationResult *[16]bool,
	errPduSessionId, errCause []uint8) ([]byte, error) {
	m := nasproto.NewMessage()
	m.GmmMessage = nasproto.NewGmmMessage()
	m.GmmHeader.SetMessageType(nasproto.MsgTypeRegistrationAccept)

	m.SecurityHeader = nasproto.SecurityHeader{
		ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		SecurityHeaderType:    nasproto.SecurityHeaderTypeIntegrityProtectedAndCiphered,
	}

	registrationAccept := nasMessage.NewRegistrationAccept(0)
	registrationAccept.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	registrationAccept.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nasproto.SecurityHeaderTypePlainNas)
	registrationAccept.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	registrationAccept.RegistrationAcceptMessageIdentity.SetMessageType(nasproto.MsgTypeRegistrationAccept)

	registrationAccept.RegistrationResult5GS.SetLen(1)
	registrationResult := uint8(0)
	if anType == models.AccessType__3_GPP_ACCESS {
		registrationResult |= nasMessage.AccessType3GPP
		if ue.State[models.AccessType_NON_3_GPP_ACCESS].Is(context.Registered) {
			registrationResult |= nasMessage.AccessTypeNon3GPP
		}
	} else {
		registrationResult |= nasMessage.AccessTypeNon3GPP
		if ue.State[models.AccessType__3_GPP_ACCESS].Is(context.Registered) {
			registrationResult |= nasMessage.AccessType3GPP
		}
	}
	registrationAccept.RegistrationResult5GS.SetRegistrationResultValue5GS(registrationResult)
	// TODO: set smsAllowed value of RegistrationResult5GS if need

	if ue.Guti != "" {
		gutiNas := nasConvert.GutiToNas(ue.Guti)
		registrationAccept.GUTI5G = &gutiNas
		registrationAccept.GUTI5G.SetIei(nasMessage.RegistrationAcceptGUTI5GType)
	}

	amf := builder.amf()
	plmnlist := amf.PlmnSupportList()
	if len(plmnlist) > 1 {
		registrationAccept.EquivalentPlmns = nasType.NewEquivalentPlmns(nasMessage.RegistrationAcceptEquivalentPlmnsType)
		var buf []uint8
		for _, plmnSupportItem := range plmnlist {
			buf = append(buf, nasConvert.PlmnIDToNas(*plmnSupportItem.PlmnId)...)
		}
		registrationAccept.EquivalentPlmns.SetLen(uint8(len(buf)))
		copy(registrationAccept.EquivalentPlmns.Octet[:], buf)
	}

	if len(ue.RegistrationArea[anType]) > 0 {
		registrationAccept.TAIList = nasType.NewTAIList(nasMessage.RegistrationAcceptTAIListType)
		taiListNas := nasConvert.TaiListToNas(ue.RegistrationArea[anType])
		registrationAccept.TAIList.SetLen(uint8(len(taiListNas)))
		registrationAccept.TAIList.SetPartialTrackingAreaIdentityList(taiListNas)
	}
	nssfinfo := ue.GetNssfInfo()
	if len(nssfinfo.AllowedNssai[anType]) > 0 {
		registrationAccept.AllowedNSSAI = nasType.NewAllowedNSSAI(nasMessage.RegistrationAcceptAllowedNSSAIType)
		var buf []uint8
		for _, allowedSnssai := range nssfinfo.AllowedNssai[anType] {
			buf = append(buf, nasConvert.SnssaiToNas(*allowedSnssai.AllowedSnssai)...)
		}
		registrationAccept.AllowedNSSAI.SetLen(uint8(len(buf)))
		registrationAccept.AllowedNSSAI.SetSNSSAIValue(buf)
	}

	if nssfinfo.NetworkSliceInfo != nil {
		if len(nssfinfo.NetworkSliceInfo.RejectedNssaiInPlmn) != 0 || len(nssfinfo.NetworkSliceInfo.RejectedNssaiInTa) != 0 {
			rejectedNssaiNas := nasConvert.RejectedNssaiToNas(
				nssfinfo.NetworkSliceInfo.RejectedNssaiInPlmn, nssfinfo.NetworkSliceInfo.RejectedNssaiInTa)
			registrationAccept.RejectedNSSAI = &rejectedNssaiNas
			registrationAccept.RejectedNSSAI.SetIei(nasMessage.RegistrationAcceptRejectedNSSAIType)
		}
	}

	if includeConfiguredNssaiCheck(ue) {
		registrationAccept.ConfiguredNSSAI = nasType.NewConfiguredNSSAI(nasMessage.RegistrationAcceptConfiguredNSSAIType)
		var buf []uint8
		for _, snssai := range nssfinfo.ConfiguredNssai {
			buf = append(buf, nasConvert.SnssaiToNas(*snssai.ConfiguredSnssai)...)
		}
		registrationAccept.ConfiguredNSSAI.SetLen(uint8(len(buf)))
		registrationAccept.ConfiguredNSSAI.SetSNSSAIValue(buf)
	}

	// 5gs network feature support
	cfg := builder.backend.Config().Configuration
	if cfg.Get5gsNwFeatSuppEnable() {
		registrationAccept.NetworkFeatureSupport5GS =
			nasType.NewNetworkFeatureSupport5GS(nasMessage.RegistrationAcceptNetworkFeatureSupport5GSType)
		registrationAccept.NetworkFeatureSupport5GS.SetLen(2)
		if anType == models.AccessType__3_GPP_ACCESS {
			registrationAccept.SetIMSVoPS3GPP(cfg.Get5gsNwFeatSuppImsVoPS())
		} else {
			registrationAccept.SetIMSVoPSN3GPP(cfg.Get5gsNwFeatSuppImsVoPS())
		}
		registrationAccept.SetEMC(cfg.Get5gsNwFeatSuppEmc())
		registrationAccept.SetEMF(cfg.Get5gsNwFeatSuppEmf())
		registrationAccept.SetIWKN26(cfg.Get5gsNwFeatSuppIwkN26())
		registrationAccept.SetMPSI(cfg.Get5gsNwFeatSuppMpsi())
		registrationAccept.SetEMCN(cfg.Get5gsNwFeatSuppEmcN3())
		registrationAccept.SetMCSI(cfg.Get5gsNwFeatSuppMcsi())
	}

	if pDUSessionStatus != nil {
		registrationAccept.PDUSessionStatus = nasType.NewPDUSessionStatus(nasMessage.RegistrationAcceptPDUSessionStatusType)
		registrationAccept.PDUSessionStatus.SetLen(2)
		registrationAccept.PDUSessionStatus.Buffer = nasConvert.PSIToBuf(*pDUSessionStatus)
	}

	if reactivationResult != nil {
		registrationAccept.PDUSessionReactivationResult =
			nasType.NewPDUSessionReactivationResult(nasMessage.RegistrationAcceptPDUSessionReactivationResultType)
		registrationAccept.PDUSessionReactivationResult.SetLen(2)
		registrationAccept.PDUSessionReactivationResult.Buffer = nasConvert.PSIToBuf(*reactivationResult)
	}

	if errPduSessionId != nil {
		registrationAccept.PDUSessionReactivationResultErrorCause = nasType.NewPDUSessionReactivationResultErrorCause(
			nasMessage.RegistrationAcceptPDUSessionReactivationResultErrorCauseType)
		buf := nasConvert.PDUSessionReactivationResultErrorCauseToBuf(errPduSessionId, errCause)
		registrationAccept.PDUSessionReactivationResultErrorCause.SetLen(uint16(len(buf)))
		registrationAccept.PDUSessionReactivationResultErrorCause.Buffer = buf
	}

	if ue.LadnInfo != nil {
		registrationAccept.LADNInformation = nasType.NewLADNInformation(nasMessage.RegistrationAcceptLADNInformationType)
		buf := make([]uint8, 0)
		for _, ladn := range ue.LadnInfo {
			ladnNas := nasConvert.LadnToNas(ladn.Dnn, ladn.TaiLists)
			buf = append(buf, ladnNas...)
		}
		registrationAccept.LADNInformation.SetLen(uint16(len(buf)))
		registrationAccept.LADNInformation.SetLADND(buf)
	}

	if nssfinfo.NetworkSlicingSubscriptionChanged {
		registrationAccept.NetworkSlicingIndication =
			nasType.NewNetworkSlicingIndication(nasMessage.RegistrationAcceptNetworkSlicingIndicationType)
		registrationAccept.NetworkSlicingIndication.SetNSSCI(1)
		registrationAccept.NetworkSlicingIndication.SetDCNI(0)
		nssfinfo.NetworkSlicingSubscriptionChanged = false // reset the value
	}
	pcfinfo := ue.GetPcfInfo()
	if anType == models.AccessType__3_GPP_ACCESS && pcfinfo.AmPolicyAssociation != nil &&
		pcfinfo.AmPolicyAssociation.ServAreaRes != nil {
		registrationAccept.ServiceAreaList = nasType.NewServiceAreaList(nasMessage.RegistrationAcceptServiceAreaListType)
		partialServiceAreaList := nasConvert.PartialServiceAreaListToNas(ue.PlmnId, *pcfinfo.AmPolicyAssociation.ServAreaRes)
		registrationAccept.ServiceAreaList.SetLen(uint8(len(partialServiceAreaList)))
		registrationAccept.ServiceAreaList.SetPartialServiceAreaList(partialServiceAreaList)
	}

	if anType == models.AccessType__3_GPP_ACCESS && ue.T3512Value != 0 {
		registrationAccept.T3512Value = nasType.NewT3512Value(nasMessage.RegistrationAcceptT3512ValueType)
		registrationAccept.T3512Value.SetLen(1)
		t3512 := nasConvert.GPRSTimer3ToNas(ue.T3512Value)
		registrationAccept.T3512Value.Octet = t3512
	}

	if anType == models.AccessType_NON_3_GPP_ACCESS {
		registrationAccept.Non3GppDeregistrationTimerValue =
			nasType.NewNon3GppDeregistrationTimerValue(nasMessage.RegistrationAcceptNon3GppDeregistrationTimerValueType)
		registrationAccept.Non3GppDeregistrationTimerValue.SetLen(1)
		timerValue := nasConvert.GPRSTimer2ToNas(ue.Non3gppDeregistrationTimerValue)
		registrationAccept.Non3GppDeregistrationTimerValue.SetGPRSTimer2Value(timerValue)
	}

	if ue.T3502Value != 0 {
		registrationAccept.T3502Value = nasType.NewT3502Value(nasMessage.RegistrationAcceptT3502ValueType)
		registrationAccept.T3502Value.SetLen(1)
		t3502 := nasConvert.GPRSTimer2ToNas(ue.T3502Value)
		registrationAccept.T3502Value.SetGPRSTimer2Value(t3502)
	}

	if ue.UESpecificDRX != nasMessage.DRXValueNotSpecified {
		registrationAccept.NegotiatedDRXParameters =
			nasType.NewNegotiatedDRXParameters(nasMessage.RegistrationAcceptNegotiatedDRXParametersType)
		registrationAccept.NegotiatedDRXParameters.SetLen(1)
		registrationAccept.NegotiatedDRXParameters.SetDRXValue(ue.UESpecificDRX)
	}

	m.GmmMessage.RegistrationAccept = registrationAccept

	return nas_security.Encode(ue, m, anType)
}

func includeConfiguredNssaiCheck(ue *context.AmfUe) bool {
	nssfinfo := ue.GetNssfInfo()
	if len(nssfinfo.ConfiguredNssai) == 0 {
		return false
	}

	registrationRequest := ue.RegistrationRequest
	if registrationRequest.RequestedNSSAI == nil {
		return true
	}
	if nssfinfo.NetworkSliceInfo != nil && len(nssfinfo.NetworkSliceInfo.RejectedNssaiInPlmn) != 0 {
		return true
	}
	if registrationRequest.NetworkSlicingIndication != nil && registrationRequest.NetworkSlicingIndication.GetDCNI() == 1 {
		return true
	}
	return false
}

func (builder *Nas) BuildStatus5GMM(cause uint8) ([]byte, error) {
	m := nasproto.NewMessage()
	m.GmmMessage = nasproto.NewGmmMessage()
	m.GmmHeader.SetMessageType(nasproto.MsgTypeStatus5GMM)

	status5GMM := nasMessage.NewStatus5GMM(0)
	status5GMM.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	status5GMM.SetSecurityHeaderType(nasproto.SecurityHeaderTypePlainNas)
	status5GMM.SetMessageType(nasproto.MsgTypeStatus5GMM)
	status5GMM.SetCauseValue(cause)

	m.GmmMessage.Status5GMM = status5GMM

	return m.PlainNasEncode()
}

func (builder *Nas) BuildConfigurationUpdateCommand(ue *context.AmfUe, anType models.AccessType,
	networkSlicingIndication *nasType.NetworkSlicingIndication) ([]byte, error) {
	m := nasproto.NewMessage()
	m.GmmMessage = nasproto.NewGmmMessage()
	m.GmmHeader.SetMessageType(nasproto.MsgTypeConfigurationUpdateCommand)

	configurationUpdateCommand := nasMessage.NewConfigurationUpdateCommand(0)
	configurationUpdateCommand.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	configurationUpdateCommand.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(nasproto.SecurityHeaderTypePlainNas)
	configurationUpdateCommand.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	configurationUpdateCommand.SetMessageType(nasproto.MsgTypeConfigurationUpdateCommand)

	if ue.ConfigurationUpdateIndication.Octet != 0 {
		configurationUpdateCommand.ConfigurationUpdateIndication =
			nasType.NewConfigurationUpdateIndication(nasMessage.ConfigurationUpdateCommandConfigurationUpdateIndicationType)
		configurationUpdateCommand.ConfigurationUpdateIndication = &ue.ConfigurationUpdateIndication
	}

	if networkSlicingIndication != nil {
		configurationUpdateCommand.NetworkSlicingIndication =
			nasType.NewNetworkSlicingIndication(nasMessage.ConfigurationUpdateCommandNetworkSlicingIndicationType)
		configurationUpdateCommand.NetworkSlicingIndication = networkSlicingIndication
	}

	if ue.Guti != "" {
		gutiNas := nasConvert.GutiToNas(ue.Guti)
		configurationUpdateCommand.GUTI5G = &gutiNas
		configurationUpdateCommand.GUTI5G.SetIei(nasMessage.ConfigurationUpdateCommandGUTI5GType)
	}

	if len(ue.RegistrationArea[anType]) > 0 {
		configurationUpdateCommand.TAIList = nasType.NewTAIList(nasMessage.ConfigurationUpdateCommandTAIListType)
		taiListNas := nasConvert.TaiListToNas(ue.RegistrationArea[anType])
		configurationUpdateCommand.TAIList.SetLen(uint8(len(taiListNas)))
		configurationUpdateCommand.TAIList.SetPartialTrackingAreaIdentityList(taiListNas)
	}
	nssfinfo := ue.GetNssfInfo()
	if len(nssfinfo.AllowedNssai[anType]) > 0 {
		configurationUpdateCommand.AllowedNSSAI =
			nasType.NewAllowedNSSAI(nasMessage.ConfigurationUpdateCommandAllowedNSSAIType)
		var buf []uint8
		for _, allowedSnssai := range nssfinfo.AllowedNssai[anType] {
			buf = append(buf, nasConvert.SnssaiToNas(*allowedSnssai.AllowedSnssai)...)
		}
		configurationUpdateCommand.AllowedNSSAI.SetLen(uint8(len(buf)))
		configurationUpdateCommand.AllowedNSSAI.SetSNSSAIValue(buf)
	}

	if len(nssfinfo.ConfiguredNssai) > 0 {
		configurationUpdateCommand.ConfiguredNSSAI =
			nasType.NewConfiguredNSSAI(nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType)
		var buf []uint8
		for _, snssai := range nssfinfo.ConfiguredNssai {
			buf = append(buf, nasConvert.SnssaiToNas(*snssai.ConfiguredSnssai)...)
		}
		configurationUpdateCommand.ConfiguredNSSAI.SetLen(uint8(len(buf)))
		configurationUpdateCommand.ConfiguredNSSAI.SetSNSSAIValue(buf)
	}

	if nssfinfo.NetworkSliceInfo != nil {
		if len(nssfinfo.NetworkSliceInfo.RejectedNssaiInPlmn) != 0 || len(nssfinfo.NetworkSliceInfo.RejectedNssaiInTa) != 0 {
			rejectedNssaiNas := nasConvert.RejectedNssaiToNas(
				nssfinfo.NetworkSliceInfo.RejectedNssaiInPlmn, nssfinfo.NetworkSliceInfo.RejectedNssaiInTa)
			configurationUpdateCommand.RejectedNSSAI = &rejectedNssaiNas
			configurationUpdateCommand.RejectedNSSAI.SetIei(nasMessage.ConfigurationUpdateCommandRejectedNSSAIType)
		}
	}

	// TODO: UniversalTimeAndLocalTimeZone
	pcfinfo := ue.GetPcfInfo()
	if anType == models.AccessType__3_GPP_ACCESS && pcfinfo.AmPolicyAssociation != nil &&
		pcfinfo.AmPolicyAssociation.ServAreaRes != nil {
		configurationUpdateCommand.ServiceAreaList =
			nasType.NewServiceAreaList(nasMessage.ConfigurationUpdateCommandServiceAreaListType)
		partialServiceAreaList := nasConvert.PartialServiceAreaListToNas(ue.PlmnId, *pcfinfo.AmPolicyAssociation.ServAreaRes)
		configurationUpdateCommand.ServiceAreaList.SetLen(uint8(len(partialServiceAreaList)))
		configurationUpdateCommand.ServiceAreaList.SetPartialServiceAreaList(partialServiceAreaList)
	}

	amf := builder.amf()
	if amf.NetworkName().Full != "" {
		fullNetworkName := nasConvert.FullNetworkNameToNas(amf.NetworkName().Full)
		configurationUpdateCommand.FullNameForNetwork = &fullNetworkName
		configurationUpdateCommand.FullNameForNetwork.SetIei(nasMessage.ConfigurationUpdateCommandFullNameForNetworkType)
	}

	if amf.NetworkName().Short != "" {
		shortNetworkName := nasConvert.ShortNetworkNameToNas(amf.NetworkName().Short)
		configurationUpdateCommand.ShortNameForNetwork = &shortNetworkName
		configurationUpdateCommand.ShortNameForNetwork.SetIei(nasMessage.ConfigurationUpdateCommandShortNameForNetworkType)
	}

	if ue.TimeZone != "" {
		localTimeZone := nasConvert.LocalTimeZoneToNas(ue.TimeZone)
		localTimeZone.SetIei(nasMessage.ConfigurationUpdateCommandLocalTimeZoneType)
		configurationUpdateCommand.LocalTimeZone =
			nasType.NewLocalTimeZone(nasMessage.ConfigurationUpdateCommandLocalTimeZoneType)
		configurationUpdateCommand.LocalTimeZone = &localTimeZone
	}

	if ue.TimeZone != "" {
		daylightSavingTime := nasConvert.DaylightSavingTimeToNas(ue.TimeZone)
		daylightSavingTime.SetIei(nasMessage.ConfigurationUpdateCommandNetworkDaylightSavingTimeType)
		configurationUpdateCommand.NetworkDaylightSavingTime =
			nasType.NewNetworkDaylightSavingTime(nasMessage.ConfigurationUpdateCommandNetworkDaylightSavingTimeType)
		configurationUpdateCommand.NetworkDaylightSavingTime = &daylightSavingTime
	}

	if len(ue.LadnInfo) > 0 {
		configurationUpdateCommand.LADNInformation =
			nasType.NewLADNInformation(nasMessage.ConfigurationUpdateCommandLADNInformationType)
		var buf []uint8
		for _, ladn := range ue.LadnInfo {
			ladnNas := nasConvert.LadnToNas(ladn.Dnn, ladn.TaiLists)
			buf = append(buf, ladnNas...)
		}
		configurationUpdateCommand.LADNInformation.SetLen(uint16(len(buf)))
		configurationUpdateCommand.LADNInformation.SetLADND(buf)
	}

	m.GmmMessage.ConfigurationUpdateCommand = configurationUpdateCommand

	return m.PlainNasEncode()
}
