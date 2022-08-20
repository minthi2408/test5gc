package communication

import (
	"strings"
	"etri5gc/fabric"
	"github.com/free5gc/util/httpwrapper"
)
const (
	SERVICE_NAME="/namf-comm/v1"
)

type Backend interface {
	HandleCreateUEContextRequest(*httpwrapper.Request) *httpwrapper.Response
	HandleReleaseUEContextRequest(*httpwrapper.Request) *httpwrapper.Response
	HandleUEContextTransferRequest(*httpwrapper.Request) *httpwrapper.Response 
	HandleAssignEbiDataRequest(*httpwrapper.Request) *httpwrapper.Response 
	HandleRegistrationStatusUpdateRequest(*httpwrapper.Request) *httpwrapper.Response 
	HandleN1N2MessageTransferRequest(*httpwrapper.Request) *httpwrapper.Response 
	HandleN1N2MessageTransferStatusRequest(*httpwrapper.Request) *httpwrapper.Response 
	HandleN1N2MessageSubscirbeRequest(*httpwrapper.Request) *httpwrapper.Response 
	HandleN1N2MessageUnSubscribeRequest(*httpwrapper.Request) *httpwrapper.Response 
	HandleAMFStatusChangeSubscribeRequest(*httpwrapper.Request) *httpwrapper.Response 
	HandleAMFStatusChangeUnSubscribeRequest(*httpwrapper.Request) *httpwrapper.Response 
	HandleAMFStatusChangeSubscribeModify(*httpwrapper.Request) *httpwrapper.Response 
}

type Handler struct {
	prod	Backend	
}

func MakeService(p Backend) (service fabric.HttpService) {
	h := &Handler{
		prod:	p,
	}
	service.Routes = fabric.HttpRoutes{
		{
			"Index",
			"GET",
			"/",
			fabric.HttpIndexHandler,
		},

		{
			"AMFStatusChangeSubscribeModify",
			strings.ToUpper("Put"),
			"/subscriptions/:subscriptionId",
			h.HTTPAMFStatusChangeSubscribeModify,
		},

		{
			"AMFStatusChangeUnSubscribe",
			strings.ToUpper("Delete"),
			"/subscriptions/:subscriptionId",
			h.HTTPAMFStatusChangeUnSubscribe,
		},

		{
			"CreateUEContext",
			strings.ToUpper("Put"),
			"/ue-contexts/:ueContextId",
			h.HTTPCreateUEContext,
		},

		{
			"EBIAssignment",
			strings.ToUpper("Post"),
			"/ue-contexts/:ueContextId/assign-ebi",
			h.HTTPEBIAssignment,
		},

		{
			"RegistrationStatusUpdate",
			strings.ToUpper("Post"),
			"/ue-contexts/:ueContextId/transfer-update",
			h.HTTPRegistrationStatusUpdate,
		},

		{
			"ReleaseUEContext",
			strings.ToUpper("Post"),
			"/ue-contexts/:ueContextId/release",
			h.HTTPReleaseUEContext,
		},

		{
			"UEContextTransfer",
			strings.ToUpper("Post"),
			"/ue-contexts/:ueContextId/transfer",
			h.HTTPUEContextTransfer,
		},

		{
			"N1N2MessageUnSubscribe",
			strings.ToUpper("Delete"),
			"/ue-contexts/:ueContextId/n1-n2-messages/subscriptions/:subscriptionId",
			h.HTTPN1N2MessageUnSubscribe,
		},

		{
			"N1N2MessageTransfer",
			strings.ToUpper("Post"),
			"/ue-contexts/:ueContextId/n1-n2-messages",
			h.HTTPN1N2MessageTransfer,
		},

		{
			"N1N2MessageTransferStatus",
			strings.ToUpper("Get"),
			"/ue-contexts/:ueContextId/n1-n2-messages/:n1N2MessageId",
			h.HTTPN1N2MessageTransferStatus,
		},

		{
			"N1N2MessageSubscribe",
			strings.ToUpper("Post"),
			"/ue-contexts/:ueContextId/n1-n2-messages/subscriptions",
			h.HTTPN1N2MessageSubscribe,
		},

		{
			"NonUeN2InfoUnSubscribe",
			strings.ToUpper("Delete"),
			"/non-ue-n2-messages/subscriptions/:n2NotifySubscriptionId",
			h.HTTPNonUeN2InfoUnSubscribe,
		},

		{
			"NonUeN2MessageTransfer",
			strings.ToUpper("Post"),
			"/non-ue-n2-messages/transfer",
			h.HTTPNonUeN2MessageTransfer,
		},

		{
			"NonUeN2InfoSubscribe",
			strings.ToUpper("Post"),
			"/non-ue-n2-messages/subscriptions",
			h.HTTPNonUeN2InfoSubscribe,
		},

		{
			"AMFStatusChangeSubscribe",
			strings.ToUpper("Post"),
			"/subscriptions",
			h.HTTPAMFStatusChangeSubscribe,
		},
	}
	service.Group = SERVICE_NAME
	return
}
