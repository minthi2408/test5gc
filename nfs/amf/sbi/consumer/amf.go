package consumer

import (
	"fmt"
	"bytes"
	org_context	"context"
	"etri5gc/nfs/amf/context"
	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/Namf_Communication"
)

type AmfConsumer interface {
	CreateUEContextRequest(*context.AmfUe, models.UeContextCreateData) (*models.UeContextCreatedData, *models.ProblemDetails, error)
	ReleaseUEContextRequest(*context.AmfUe, models.NgApCause) (*models.ProblemDetails, error)
	UEContextTransferRequest(*context.AmfUe, models.AccessType, models.TransferReason) (*models.UeContextTransferRspData, *models.ProblemDetails, error)
	RegistrationStatusUpdate(*context.AmfUe, models.UeRegStatusUpdateReqData) (bool, *models.ProblemDetails, error)
}


type amfConsumer struct {
	context		*context.AMFContext
}

func newAmfConsumer(c *context.AMFContext) AmfConsumer {
	return &amfConsumer {
		context: c,
	}
}

func amf_comm_client(ue *context.AmfUe) (*Namf_Communication.APIClient) {
	conf := Namf_Communication.NewConfiguration()
	conf.SetBasePath(ue.TargetAmfUri)
	return Namf_Communication.NewAPIClient(conf)
}

func (c *amfConsumer) CreateUEContextRequest(ue *context.AmfUe, ueContextCreateData models.UeContextCreateData) (
	ueContextCreatedData *models.UeContextCreatedData, problemDetails *models.ProblemDetails, err error) {
	client := amf_comm_client(ue)

	req := models.CreateUeContextRequest{
		JsonData: &ueContextCreateData,
	}
	res, httpResp, localErr := client.IndividualUeContextDocumentApi.CreateUEContext(org_context.TODO(), ue.Guti, req)
	if localErr == nil {
		ueContextCreatedData = res.JsonData
		//logger.ConsumerLog.Debugf("UeContextCreatedData: %+v", *ueContextCreatedData)
	} else if httpResp != nil {
		if httpResp.Status != localErr.Error() {
			err = localErr
			return
		}
		problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
		problemDetails = &problem
	} else {
		err = openapi.ReportError("%s: server no response", ue.TargetAmfUri)
	}
	return
}

func (c *amfConsumer) ReleaseUEContextRequest(ue *context.AmfUe, ngapCause models.NgApCause) (
	problemDetails *models.ProblemDetails, err error) {
	
	client := amf_comm_client(ue)

	var ueContextId string
	if ue.Supi != "" {
		ueContextId = ue.Supi
	} else {
		ueContextId = ue.Pei
	}

	ueContextRelease := models.UeContextRelease{
		NgapCause: &ngapCause,
	}
	if ue.RegistrationType5GS == nasMessage.RegistrationType5GSEmergencyRegistration && ue.UnauthenticatedSupi {
		ueContextRelease.Supi = ue.Supi
		ueContextRelease.UnauthenticatedSupi = true
	}

	httpResp, localErr := client.IndividualUeContextDocumentApi.ReleaseUEContext(
		org_context.TODO(), ueContextId, ueContextRelease)
	if localErr == nil {
		return
	} else if httpResp != nil {
		if httpResp.Status != localErr.Error() {
			err = localErr
			return
		}
		problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
		problemDetails = &problem
	} else {
		err = openapi.ReportError("%s: server no response", ue.TargetAmfUri)
	}
	return problemDetails, err
}

func (c *amfConsumer) UEContextTransferRequest(
	ue *context.AmfUe, accessType models.AccessType, transferReason models.TransferReason) (
	ueContextTransferRspData *models.UeContextTransferRspData, problemDetails *models.ProblemDetails, err error) {

	client := amf_comm_client(ue)

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
		problemDetails = &problem
	} else {
		err = openapi.ReportError("%s: server no response", ue.TargetAmfUri)
	}
	return ueContextTransferRspData, problemDetails, err
}

// This operation is called "RegistrationCompleteNotify" at TS 23.502
func (c *amfConsumer) RegistrationStatusUpdate(ue *context.AmfUe, request models.UeRegStatusUpdateReqData) (
	regStatusTransferComplete bool, problemDetails *models.ProblemDetails, err error) {

	client := amf_comm_client(ue)

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
		problemDetails = &problem
	} else {
		err = openapi.ReportError("%s: server no response", ue.TargetAmfUri)
	}
	return
}

/*
func BuildUeContextCreateData(ue *context.AmfUe, targetRanId models.NgRanTargetId,
	sourceToTargetData models.N2InfoContent, pduSessionList []models.N2SmInformation,
	n2NotifyUri string, ngapCause *models.NgApCause) models.UeContextCreateData {
	var ueContextCreateData models.UeContextCreateData

	ueContext := BuildUeContextModel(ue)
	ueContextCreateData.UeContext = &ueContext
	ueContextCreateData.TargetId = &targetRanId
	ueContextCreateData.SourceToTargetData = &sourceToTargetData
	ueContextCreateData.PduSessionList = pduSessionList
	ueContextCreateData.N2NotifyUri = n2NotifyUri

	if ue.UeRadioCapability != "" {
		ueContextCreateData.UeRadioCapability = &models.N2InfoContent{
			NgapData: &models.RefToBinaryData{
				ContentId: ue.UeRadioCapability,
			},
		}
	}
	ueContextCreateData.NgapCause = ngapCause
	return ueContextCreateData
}

func BuildUeContextModel(ue *context.AmfUe) (ueContext models.UeContext) {
	ueContext.Supi = ue.Supi
	ueContext.SupiUnauthInd = ue.UnauthenticatedSupi

	if ue.Gpsi != "" {
		ueContext.GpsiList = append(ueContext.GpsiList, ue.Gpsi)
	}

	if ue.Pei != "" {
		ueContext.Pei = ue.Pei
	}

	if ue.UdmGroupId != "" {
		ueContext.UdmGroupId = ue.UdmGroupId
	}

	if ue.AusfGroupId != "" {
		ueContext.AusfGroupId = ue.AusfGroupId
	}

	if ue.RoutingIndicator != "" {
		ueContext.RoutingIndicator = ue.RoutingIndicator
	}

	if ue.AccessAndMobilitySubscriptionData != nil {
		if ue.AccessAndMobilitySubscriptionData.SubscribedUeAmbr != nil {
			ueContext.SubUeAmbr = &models.Ambr{
				Uplink:   ue.AccessAndMobilitySubscriptionData.SubscribedUeAmbr.Uplink,
				Downlink: ue.AccessAndMobilitySubscriptionData.SubscribedUeAmbr.Downlink,
			}
		}
		if ue.AccessAndMobilitySubscriptionData.RfspIndex != 0 {
			ueContext.SubRfsp = ue.AccessAndMobilitySubscriptionData.RfspIndex
		}
	}

	if ue.PcfId != "" {
		ueContext.PcfId = ue.PcfId
	}

	if ue.AmPolicyUri != "" {
		ueContext.PcfAmPolicyUri = ue.AmPolicyUri
	}

	if ue.AmPolicyAssociation != nil {
		if len(ue.AmPolicyAssociation.Triggers) > 0 {
			ueContext.AmPolicyReqTriggerList = buildAmPolicyReqTriggers(ue.AmPolicyAssociation.Triggers)
		}
	}

	for _, eventSub := range ue.EventSubscriptionsInfo {
		if eventSub.EventSubscription != nil {
			ueContext.EventSubscriptionList = append(ueContext.EventSubscriptionList, *eventSub.EventSubscription)
		}
	}

	if ue.TraceData != nil {
		ueContext.TraceData = ue.TraceData
	}
	return ueContext
}

func buildAmPolicyReqTriggers(triggers []models.RequestTrigger) (amPolicyReqTriggers []models.AmPolicyReqTrigger) {
	for _, trigger := range triggers {
		switch trigger {
		case models.RequestTrigger_LOC_CH:
			amPolicyReqTriggers = append(amPolicyReqTriggers, models.AmPolicyReqTrigger_LOCATION_CHANGE)
		case models.RequestTrigger_PRA_CH:
			amPolicyReqTriggers = append(amPolicyReqTriggers, models.AmPolicyReqTrigger_PRA_CHANGE)
		case models.RequestTrigger_SERV_AREA_CH:
			amPolicyReqTriggers = append(amPolicyReqTriggers, models.AmPolicyReqTrigger_SARI_CHANGE)
		case models.RequestTrigger_RFSP_CH:
			amPolicyReqTriggers = append(amPolicyReqTriggers, models.AmPolicyReqTrigger_RFSP_INDEX_CHANGE)
		}
	}
	return
}

*/