package consumer

import (
	org_context	"context"
	"etri5gc/nfs/amf/context"
)

type AmfConsumer interface {
	CreateUEContextRequest(*context.AmfUe, models.UeContextCreateData) (*models.UeContextCreatedData, *models.ProblemDetails, error)
	ReleaseUEContextRequest(*context.AmfUe, models.NgApCause) (*models.ProblemDetails, error)
	UEContextTransferRequest(*context.AmfUe, models.AccessType, models.TransferReason) (*models.UeContextTransferRspData, *models.ProblemDetails, error)
	RegistrationStatusUpdate(*context.AmfUe, models.UeRegStatusUpdateReqData) (bool, *models.ProblemDetails, error) {
}


type amfConsumer struct {
	context		*context.AmfContext
}

func newAmfConsumer(c *context.AmfContext) AmfConsumer {
	return &amfConsumer {
		context: c,
	}
}

func comm_client(ue *context.AmfUe) NamfCommunication.APIClient {
	conf := Namf_Communication.NewConfiguration()
	conf.SetBasePath(ue.TargetAmfUri)
	return Namf_Communication.NewAPIClient(conf)
}

func (c *amfConsumer) CreateUEContextRequest(ue *context.AmfUe, dat models.UeContextCreateData) (cdat *models.UeContextCreatedData, prob *models.ProblemDetails, err error) {
	client := comm_client(ue)

	req := models.CreateUeContextRequest{
		JsonData: &dat,
	}

	res, httpResp, localErr := client.IndividualUeContextDocumentApi.CreateUEContext(org_context.TODO(), ue.Guti, req)
	if localErr == nil {
		cdat = res.JsonData
		//logger.ConsumerLog.Debugf("UeContextCreatedData: %+v", *ueContextCreatedData)
	} else if httpResp != nil {
		if httpResp.Status != localErr.Error() {
			err = localErr
			return
		}
		problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
		prob = &problem
	} else {
		err = openapi.ReportError("%s: server no response", ue.TargetAmfUri)
	}
	return
}

func (c *amfConsumer) ReleaseUEContextRequest(ue *context.AmfUe, cause models.NgApCause) (
	prob *models.ProblemDetails, err error) {
	client := comm_client(ue) 

	var ueContextId string
	if ue.Supi != "" {
		ueContextId = ue.Supi
	} else {
		ueContextId = ue.Pei
	}

	ueContextRelease := models.UeContextRelease{
		NgapCause: &cause,
	}
	if ue.RegistrationType5GS == nasMessage.RegistrationType5GSEmergencyRegistration && ue.UnauthenticatedSupi {
		ueContextRelease.Supi = ue.Supi
		ueContextRelease.UnauthenticatedSupi = true
	}

	httpResp, localErr := client.IndividualUeContextDocumentApi.ReleaseUEContext(
		context.TODO(), ueContextId, ueContextRelease)
	if localErr == nil {
		return
	} else if httpResp != nil {
		if httpResp.Status != localErr.Error() {
			err = localErr
			return
		}
		problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
		prob = &problem
	} else {
		err = openapi.ReportError("%s: server no response", ue.TargetAmfUri)
	}
	return prob, err
}

func (c *amfConsumer) UEContextTransferRequest(
	ue *amf_context.AmfUe, accessType models.AccessType, transferReason models.TransferReason) (
	ueContextTransferRspData *models.UeContextTransferRspData, problemDetails *models.ProblemDetails, err error) {
	
	client := comm_client(ue) 

	ueContextTransferReqData := models.UeContextTransferReqData{
		Reason:     transferReason,
		AccessType: accessType,
	}

	req := models.UeContextTransferRequest{
		JsonData: &ueContextTransferReqData,
	}
	if transferReason == models.TransferReason_INIT_REG || transferReason == models.TransferReason_MOBI_REG {
		var buf bytes.Buffer
		ue.RegistrationRequest.EncodeRegistrationRequest(&buf)
		ueContextTransferReqData.RegRequest = &models.N1MessageContainer{
			N1MessageClass: models.N1MessageClass__5_GMM,
			N1MessageContent: &models.RefToBinaryData{
				ContentId: "n1Msg",
			},
		}
		req.BinaryDataN1Message = buf.Bytes()
	}

	// guti format is defined at TS 29.518 Table 6.1.3.2.2-1 5g-guti-[0-9]{5,6}[0-9a-fA-F]{14}
	ueContextId := fmt.Sprintf("5g-guti-%s", ue.Guti)

	res, httpResp, localErr := client.IndividualUeContextDocumentApi.UEContextTransfer(org_context.TODO(), ueContextId, req)
	if localErr == nil {
		ueContextTransferRspData = res.JsonData
		//logger.ConsumerLog.Debugf("UeContextTransferRspData: %+v", *ueContextTransferRspData)
	} else if httpResp != nil {
		if httpResp.Status != localErr.Error() {
			err = localErr
			return
		}
		problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
		prob = &problem
	} else {
		err = openapi.ReportError("%s: server no response", ue.TargetAmfUri)
	}
	return ueContextTransferRspData, prob, err
}

// This operation is called "RegistrationCompleteNotify" at TS 23.502
func (c *amfConsumer) RegistrationStatusUpdate(ue *amf_context.AmfUe, request models.UeRegStatusUpdateReqData) (
	regStatusTransferComplete bool, problemDetails *models.ProblemDetails, err error) {
	client := comm_client(ue) 

	ueContextId := fmt.Sprintf("5g-guti-%s", ue.Guti)
	res, httpResp, localErr :=
		client.IndividualUeContextDocumentApi.RegistrationStatusUpdate(org_context.TODO(), ueContextId, request)
	if localErr == nil {
		regStatusTransferComplete = res.RegStatusTransferComplete
	} else if httpResp != nil {
		if httpResp.Status != localErr.Error() {
			err = localErr
			return
		}
		problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
		prob = &problem
	} else {
		err = openapi.ReportError("%s: server no response", ue.TargetAmfUri)
	}
	return
}
