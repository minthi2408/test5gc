package consumer

import (
	org_context "context"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"

	"github.com/free5gc/nas/nasType"
	"etri5gc/nfs/amf/context"
	"github.com/antihax/optional"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/Nausf_UEAuthentication"
)

type AusfConsumer interface {
}

type ausfConsumer struct {
	amf		*context.AMFContext
}


func newAusfConsumer(amf *context.AMFContext) AusfConsumer {
	return &ausfConsumer{amf: amf}
}

func ausf_auth_client(uri string) (*Nausf_UEAuthentication.APIClient) {
	conf := Nausf_UEAuthentication.NewConfiguration()
	conf.SetBasePath(uri)
	return Nausf_UEAuthentication.NewAPIClient(conf)
}

func (c *ausfConsumer) SendUEAuthenticationAuthenticateRequest(ue *context.AmfUe,
	resynchronizationInfo *models.ResynchronizationInfo) (*models.UeAuthenticationCtx, *models.ProblemDetails, error) {
	client := ausf_auth_client(ue.AusfUri)
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
}

func (c *ausfConsumer) SendAuth5gAkaConfirmRequest(ue *context.AmfUe, resStar string) (
	*models.ConfirmationDataResponse, *models.ProblemDetails, error) {
	var ausfUri string
	if confirmUri, err := url.Parse(ue.AuthenticationCtx.Links["5g-aka"].Href); err != nil {
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
}

func (c *ausfConsumer) SendEapAuthConfirmRequest(ue *context.AmfUe, eapMsg nasType.EAPMessage) (
	response *models.EapSession, problemDetails *models.ProblemDetails, err1 error) {
	confirmUri, err := url.Parse(ue.AuthenticationCtx.Links["eap-session"].Href)
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
}

