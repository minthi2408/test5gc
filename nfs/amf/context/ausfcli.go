package context

import (
	"encoding/binary"
	"encoding/hex"
	"regexp"

	"github.com/free5gc/nas/nasType"
	"github.com/free5gc/nas/security"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/ueauth"
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

type ausfClient struct {
	ue   *AmfUe
	info AusfInfo
}

func (c *ausfClient) Info() *AusfInfo {
	return &c.info
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
	P1 := []byte{c.ue.seccon.CipheringAlg}
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
	copy(c.ue.seccon.KnasEnc[:], kenc[16:32])

	// Integrity Key
	P0 = []byte{security.NNASIntAlg}
	L0 = ueauth.KDFLen(P0)
	P1 = []byte{c.ue.seccon.IntegrityAlg}
	L1 = ueauth.KDFLen(P1)

	kint, err := ueauth.GetKDFValue(KamfBytes, ueauth.FC_FOR_ALGORITHM_KEY_DERIVATION, P0, L0, P1, L1)
	if err != nil {
		//	logger.ContextLog.Error(err)
		return
	}
	copy(c.ue.seccon.KnasInt[:], kint[16:32])
}

// Access Network key Derivation function defined in TS 33.501 Annex A.9
func (c *ausfClient) DerivateAnKey(anType models.AccessType) {
	accessType := security.AccessType3GPP // Defalut 3gpp
	P0 := make([]byte, 4)
	binary.BigEndian.PutUint32(P0, c.ue.seccon.ULCount.Get())
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
		c.ue.seccon.Kgnb = key
	case security.AccessTypeNon3GPP:
		c.ue.seccon.Kn3iwf = key
	}
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
	c.ue.seccon.NH, err = ueauth.GetKDFValue(KamfBytes, ueauth.FC_FOR_NH_DERIVATION, P0, L0)
	if err != nil {
		//logger.ContextLog.Error(err)
		return
	}
}

func (c *ausfClient) clear() {
	c.info.AuthFailureCauseSynchFailureTimes = 0
	c.info.IdentityRequestSendTimes = 0
}
func (c *ausfClient) Select() {
}

func (c *ausfClient) SendUEAuthenticationAuthenticateRequest(resynchronizationInfo *models.ResynchronizationInfo) (*models.UeAuthenticationCtx, *models.ProblemDetails, error) {
	/*
		client := ausf_auth_client(ue.GetAusfInfo().AusfUri)
		servedGuami := c.amf.ServedGuami()

		var authInfo models.AuthenticationInfo
		authInfo.SupiOrSuci = ue.Suci
		if mnc, err := strconv.Atoi(servedGuami.PlmnId.Mnc); err != nil {
			return nil, nil, err
		} else {
			authInfo.ServingNetworkName = fmt.Sprintf("5G:mnc%03d.mcc%s.3gppnetwork.org", mnc, servedGuami.PlmnId.Mcc)
		}
		if resynchronizationInfo != nil {
			authInfo.ResynchronizationInfo = resynchronizationInfo
		}

		ueAuthenticationCtx, httpResponse, err := client.DefaultApi.UeAuthenticationsPost(org_context.Background(), authInfo)
		if err == nil {
			return &ueAuthenticationCtx, nil, nil
		} else if httpResponse != nil {
			if httpResponse.Status != err.Error() {
				return nil, nil, err
			}
			problem := err.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
			return nil, &problem, nil
		} else {
			return nil, nil, openapi.ReportError("server no response")
		}
	*/
	return nil, nil, nil
}

func (c *ausfClient) SendAuth5gAkaConfirmRequest(resStar string) (
	*models.ConfirmationDataResponse, *models.ProblemDetails, error) {
	/*
		var ausfUri string
		if confirmUri, err := url.Parse(ue.GetAusfInfo().AuthenticationCtx.Links["5g-aka"].Href); err != nil {
			return nil, nil, err
		} else {
			ausfUri = fmt.Sprintf("%s://%s", confirmUri.Scheme, confirmUri.Host)
		}

		client := ausf_auth_client(ausfUri)

		confirmData := &Nausf_UEAuthentication.UeAuthenticationsAuthCtxId5gAkaConfirmationPutParamOpts{
			ConfirmationData: optional.NewInterface(models.ConfirmationData{
				ResStar: resStar,
			}),
		}

		confirmResult, httpResponse, err := client.DefaultApi.UeAuthenticationsAuthCtxId5gAkaConfirmationPut(
			org_context.Background(), ue.Suci, confirmData)
		if err == nil {
			return &confirmResult, nil, nil
		} else if httpResponse != nil {
			if httpResponse.Status != err.Error() {
				return nil, nil, err
			}
			switch httpResponse.StatusCode {
			case 400, 500:
				problem := err.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
				return nil, &problem, nil
			}
			return nil, nil, nil
		} else {
			return nil, nil, openapi.ReportError("server no response")
		}
	*/
	return nil, nil, nil
}

func (c *ausfClient) SendEapAuthConfirmRequest(eapMsg nasType.EAPMessage) (
	response *models.EapSession, problemDetails *models.ProblemDetails, err1 error) {
	/*
		confirmUri, err := url.Parse(ue.GetAusfInfo().AuthenticationCtx.Links["eap-session"].Href)
		if err != nil {
			//logger.ConsumerLog.Errorf("url Parse failed: %+v", err)
		}
		ausfUri := fmt.Sprintf("%s://%s", confirmUri.Scheme, confirmUri.Host)

		client := ausf_auth_client(ausfUri)

		eapSessionReq := &Nausf_UEAuthentication.EapAuthMethodParamOpts{
			EapSession: optional.NewInterface(models.EapSession{
				EapPayload: base64.StdEncoding.EncodeToString(eapMsg.GetEAPMessage()),
			}),
		}

		eapSession, httpResponse, err := client.DefaultApi.EapAuthMethod(org_context.Background(), ue.Suci, eapSessionReq)
		if err == nil {
			response = &eapSession
		} else if httpResponse != nil {
			if httpResponse.Status != err.Error() {
				err1 = err
				return
			}
			switch httpResponse.StatusCode {
			case 400, 500:
				problem := err.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
				problemDetails = &problem
			}
		} else {
			err1 = openapi.ReportError("server no response")
		}

		return response, problemDetails, err1
	*/
	return
}
