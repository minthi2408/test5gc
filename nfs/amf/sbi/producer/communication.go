package producer

import (
	"etri5gc/fabric/httpdp"
	"etri5gc/openapi"
	amfprod "etri5gc/openapi/producers/amf"
	"fmt"
	"strings"

	//openapi_http "etri5gc/openapi/httpdp"
	"etri5gc/openapi/models"
)


func communicationService(p *Producer) (service httpdp.HttpService) {
    fn := ginHandler
	subscriptions := openapi.AMF_COMMUNICATION_SUBSCRIPTIONS
	uecontexts := openapi.AMF_COMMUNICATION_UE_CONTEXTS
	nonuen2 := openapi.AMF_COMMUNICATION_NONUE_N2_MESSAGES

	service.Routes = httpdp.HttpRoutes{
		{
			"Index",
			"GET",
			"/",
			httpdp.HttpIndexHandler,
		},

		{
			"AMFStatusChangeSubscribeModify",
			strings.ToUpper("Put"),
			fmt.Sprintf("/%s/:subscriptionId", subscriptions),
			fn(amfprod.AMFStatusChangeSubscribeModify, p),
		},
		{
			"AMFStatusChangeUnSubscribe",
			strings.ToUpper("Delete"),
			fmt.Sprintf("/%s/:subscriptionId", subscriptions),
			fn(amfprod.AMFStatusChangeUnSubscribe, p),
		},

		{
			"CreateUEContext",
			strings.ToUpper("Put"),
			fmt.Sprintf("/%s/:ueContextId", uecontexts),
			fn(amfprod.CreateUEContext, p),
		},

		{
			"EBIAssignment",
			strings.ToUpper("Post"),
			fmt.Sprintf("/%s/:ueContextId/assign-ebi", uecontexts),
			fn(amfprod.EBIAssignment, p),
		},

		{
			"RegistrationStatusUpdate",
			strings.ToUpper("Post"),
			fmt.Sprintf("/%s/:ueContextId/transfer-update", uecontexts),
			fn(amfprod.RegistrationStatusUpdate, p),
		},

		{
			"ReleaseUEContext",
			strings.ToUpper("Post"),
			fmt.Sprintf("/%s/:ueContextId/release", uecontexts),
			fn(amfprod.ReleaseUEContext, p),
		},

		{
			"UEContextTransfer",
			strings.ToUpper("Post"),
			fmt.Sprintf("/%s/:ueContextId/transfer", uecontexts),
			fn(amfprod.UEContextTransfer, p),
		},
		{
			"N1N2MessageUnSubscribe",
			strings.ToUpper("Delete"),
			fmt.Sprintf("/%s/:ueContextId/n1-n2-messages/subscriptions/:subscriptionId", uecontexts),
			fn(amfprod.N1N2MessageUnSubscribe, p),
		},

		{
			"N1N2MessageTransfer",
			strings.ToUpper("Post"),
			fmt.Sprintf("/%s/:ueContextId/n1-n2-messages", uecontexts),
			fn(amfprod.N1N2MessageTransfer, p),
		},
		{
			"N1N2MessageTransferStatus",
			strings.ToUpper("Get"),
			fmt.Sprintf("/%s/:ueContextId/n1-n2-messages/:n1N2MessageId", uecontexts),
			fn(amfprod.N1N2MessageTransferStatus, p),
		},
		{
			"N1N2MessageSubscribe",
			strings.ToUpper("Post"),
			fmt.Sprintf("/%s/:ueContextId/n1-n2-messages/subscriptions", uecontexts),
			fn(amfprod.N1N2MessageSubscribe, p),
		},

		{
			"NonUeN2InfoUnSubscribe",
			strings.ToUpper("Delete"),
			fmt.Sprintf("/%s/subscriptions/:n2NotifySubscriptionId", nonuen2),
			fn(amfprod.NonUeN2InfoUnSubscribe, p),
		},

		{
			"NonUeN2MessageTransfer",
			strings.ToUpper("Post"),
			fmt.Sprintf("/%s/transfer", nonuen2),
			fn(amfprod.NonUeN2MessageTransfer, p),
		},

		{
			"NonUeN2InfoSubscribe",
			strings.ToUpper("Post"),
			fmt.Sprintf("/%s/subscriptions", nonuen2),
			fn(amfprod.NonUeN2InfoSubscribe, p),
		},
		{
			"AMFStatusChangeSubscribe",
			strings.ToUpper("Post"),
			fmt.Sprintf("/%s", subscriptions),
			fn(amfprod.AMFStatusChangeSubscribe, p),
		},
	}
	service.Group = openapi.AMF_COMMUNICATION
	return
}

func (p *Producer) HandleStatusChangeSubscribe(input *models.SubscriptionData) (result models.SubscriptionData, uri string, prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleStatusChangeSubscribeModify(subId string, input *models.SubscriptionData) (result models.SubscriptionData, prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleStatusChangeUnSubscribe(subId string) (prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleCreateUEContext(ueContextId string, input *models.CreateUeContextRequest) (result models.CreateUeContextResponse, err *openapi.ApiError, prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleEBIAssignment(ueContextId string, input *models.AssignEbiData) (result models.AssignedEbiData, err *openapi.ApiError, prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleRegistrationStatusUpdate(ueContextId string, input *models.UeRegStatusUpdateReqData) (result models.UeRegStatusUpdateRspData, prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleReleaseUEContext(ueContextId string, input *models.UeContextRelease) (prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleUEContextTransfer(ueContextId string, input *models.UeContextTransferRequest) (result models.UeContextTransferResponse, prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleN1N2MessageUnSubscribe(ueContextId string, subscriptionId string) (prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleN1N2MessageTransfer(ueContextId string, input *models.N1N2MessageTransferRequest) (result models.N1N2MessageTransferRspData, err *openapi.ApiError, prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleN1N2MessageSubscribe(ueContextId string, input *models.UeN1N2InfoSubscriptionCreateData) (result models.UeN1N2InfoSubscriptionCreatedData, prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleN1N2TransferFailureNotification(input *models.N1N2MsgTxfrFailureNotification) (prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleN1N2MessageTransferStatus(ueContextId string, reqUri string) (prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleNonUeN2MessageTransfer(input *models.NonUeN2MessageTransferRequest) (result models.N2InformationTransferRspData, err *openapi.ApiError, prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleNonUeN2InfoSubscribe(input *models.NonUeN2InfoSubscriptionCreateData) (result models.NonUeN2InfoSubscriptionCreatedData, prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleNonUeN2InfoUnSubscribe(subId string) (prob *models.ProblemDetails) {
	return
}
