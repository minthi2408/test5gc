package producer

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
)

func (p *Producer) COMM_HandleAMFStatusChangeSubscribeModfy(subscriptionId string, body models.SubscriptionData) (successCode int32, result models.SubscriptionData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("COMM_HandleAMFStatusChangeSubscribeModfy has not been implemented")
	return
}
func (p *Producer) COMM_HandleAMFStatusChangeUnSubscribe(subscriptionId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("COMM_HandleAMFStatusChangeUnSubscribe has not been implemented")
	return
}
func (p *Producer) COMM_HandleCancelRelocateUEContext(ueContextId string, body models.CancelRelocateUEContextRequest) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("COMM_HandleCancelRelocateUEContext has not been implemented")
	return
}
func (p *Producer) COMM_HandleCreateUEContext(ueContextId string, body models.CreateUEContextRequest) (successCode int32, result models.UeContextCreatedData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("COMM_HandleCreateUEContext has not been implemented")
	return
}
func (p *Producer) COMM_HandleEBIAssignment(ueContextId string, body models.AssignEbiData) (successCode int32, result models.AssignedEbiData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("COMM_HandleEBIAssignment has not been implemented")
	return
}
func (p *Producer) COMM_HandleRegistrationStatusUpdate(ueContextId string, body models.UeRegStatusUpdateReqData) (successCode int32, result models.UeRegStatusUpdateRspData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("COMM_HandleRegistrationStatusUpdate has not been implemented")
	return
}
func (p *Producer) COMM_HandleReleaseUEContext(ueContextId string, body models.UEContextRelease) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("COMM_HandleReleaseUEContext has not been implemented")
	return
}
func (p *Producer) COMM_HandleRelocateUEContext(ueContextId string, body models.RelocateUEContextRequest) (successCode int32, result models.UeContextRelocatedData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("COMM_HandleRelocateUEContext has not been implemented")
	return
}
func (p *Producer) COMM_HandleUEContextTransfer(ueContextId string, body models.UeContextTransferReqData) (successCode int32, result models.UeContextTransferRspData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("COMM_HandleUEContextTransfer has not been implemented")
	return
}
func (p *Producer) COMM_HandleN1N2MessageUnSubscribe(ueContextId string, subscriptionId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("COMM_HandleN1N2MessageUnSubscribe has not been implemented")
	return
}
func (p *Producer) COMM_HandleN1N2MessageTransfer(ueContextId string, body models.N1N2MessageTransferReqData) (successCode int32, result models.N1N2MessageTransferRspData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("COMM_HandleN1N2MessageTransfer has not been implemented")
	return
}
func (p *Producer) COMM_HandleN1N2MessageSubscribe(ueContextId string, body models.UeN1N2InfoSubscriptionCreateData) (successCode int32, result models.UeN1N2InfoSubscriptionCreatedData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("COMM_HandleN1N2MessageSubscribe has not been implemented")
	return
}
func (p *Producer) COMM_HandleNonUeN2InfoUnSubscribe(n2NotifySubscriptionId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("COMM_HandleNonUeN2InfoUnSubscribe has not been implemented")
	return
}
func (p *Producer) COMM_HandleNonUeN2MessageTransfer(body models.N2InformationTransferReqData) (successCode int32, result models.N2InformationTransferRspData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("COMM_HandleNonUeN2MessageTransfer has not been implemented")
	return
}
func (p *Producer) COMM_HandleNonUeN2InfoSubscribe(body models.NonUeN2InfoSubscriptionCreateData) (successCode int32, result models.NonUeN2InfoSubscriptionCreatedData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("COMM_HandleNonUeN2InfoSubscribe has not been implemented")
	return
}
func (p *Producer) COMM_HandleAMFStatusChangeSubscribe(body models.SubscriptionData) (successCode int32, result models.SubscriptionData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("COMM_HandleAMFStatusChangeSubscribe has not been implemented")
	return
}
