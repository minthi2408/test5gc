package amf

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
)

type AmfProducer interface {
	//communication
	HandleStatusChangeSubscribe(*models.SubscriptionData) (models.SubscriptionData, string, *models.ProblemDetails)
	HandleStatusChangeSubscribeModify(string, *models.SubscriptionData) (models.SubscriptionData, *models.ProblemDetails)
	HandleStatusChangeUnSubscribe(string) *models.ProblemDetails
	HandleCreateUEContext(string, *models.CreateUeContextRequest) (models.CreateUeContextResponse, *openapi.ApiError, *models.ProblemDetails)
	HandleEBIAssignment(string, *models.AssignEbiData) (models.AssignedEbiData, *openapi.ApiError, *models.ProblemDetails)
	HandleRegistrationStatusUpdate(string, *models.UeRegStatusUpdateReqData) (models.UeRegStatusUpdateRspData, *models.ProblemDetails)

	HandleReleaseUEContext(string, *models.UeContextRelease) *models.ProblemDetails

	HandleUEContextTransfer(string, *models.UeContextTransferRequest) (models.UeContextTransferResponse, *models.ProblemDetails)
	HandleN1N2MessageUnSubscribe(string, string) *models.ProblemDetails
	HandleN1N2MessageTransfer(string, *models.N1N2MessageTransferRequest) (models.N1N2MessageTransferRspData, *openapi.ApiError, *models.ProblemDetails)
	HandleN1N2MessageTransferStatus(string, string) *models.ProblemDetails
	HandleN1N2MessageSubscribe(string, *models.UeN1N2InfoSubscriptionCreateData) (models.UeN1N2InfoSubscriptionCreatedData, *models.ProblemDetails)
	HandleN1N2TransferFailureNotification(*models.N1N2MsgTxfrFailureNotification) *models.ProblemDetails
	HandleNonUeN2InfoUnSubscribe(string) *models.ProblemDetails
	HandleNonUeN2MessageTransfer(*models.NonUeN2MessageTransferRequest) (models.N2InformationTransferRspData, *openapi.ApiError, *models.ProblemDetails)
	HandleNonUeN2InfoSubscribe(*models.NonUeN2InfoSubscriptionCreateData) (models.NonUeN2InfoSubscriptionCreatedData, *models.ProblemDetails)

	//event exposure
	HandleCreateSubscription(*models.AmfCreateEventSubscription) (models.AmfCreatedEventSubscription, *models.ProblemDetails)
	HandleDeleteSubscription(string) *models.ProblemDetails
	HandleModifySubscription(string, *models.ModifySubscriptionRequest) (models.AmfUpdatedEventSubscription, *models.ProblemDetails)

	//location
	HandleProvideLocationInfo(string, *models.RequestLocInfo) (models.ProvideLocInfo, *models.ProblemDetails)
	HandleProvidePositioningInfo(string, *models.RequestPosInfo) (models.ProvidePosInfo, *models.ProblemDetails)

	//mt
	HandleProvideDomainSelectionInfo(string, string, string) (models.UeContextInfo, *models.ProblemDetails)
	HandleEnableUeReachability(string, *models.EnableUeReachabilityReqData) (models.EnableUeReachabilityRspData, *models.ProblemDetails)

    //callback
	HandleN1MessageNotify(*models.N1MessageNotify) *models.ProblemDetails
	HandleAmPolicyControlUpdateNotifyTerminate(string, *models.TerminationNotification) *models.ProblemDetails
	HandleAmPolicyControlUpdateNotifyUpdate(string, *models.PolicyUpdate) *models.ProblemDetails
    HandleSmContextStatusNotify(string, string, *models.SmContextStatusNotification) *models.ProblemDetails
}
