package producer

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
)

func (p *Producer) PDU_HandleReleasePduSession(pduSessionRef string, body *models.ReleaseData) (successCode int32, result models.ReleasedData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PDU_HandleReleasePduSession has not been implemented")
	return
}
func (p *Producer) PDU_HandleRetrievePduSession(pduSessionRef string, body models.RetrieveData) (successCode int32, result models.RetrievedData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PDU_HandleRetrievePduSession has not been implemented")
	return
}
func (p *Producer) PDU_HandleTransferMoData(pduSessionRef string, body models.TransferMoDataRequest) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PDU_HandleTransferMoData has not been implemented")
	return
}
func (p *Producer) PDU_HandleUpdatePduSession(pduSessionRef string, body models.HsmfUpdateData) (successCode int32, result models.HsmfUpdatedData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PDU_HandleUpdatePduSession has not been implemented")
	return
}
func (p *Producer) PDU_HandleReleaseSmContext(smContextRef string, body *models.SmContextReleaseData) (successCode int32, result models.SmContextReleasedData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PDU_HandleReleaseSmContext has not been implemented")
	return
}
func (p *Producer) PDU_HandleRetrieveSmContext(smContextRef string, body *models.SmContextRetrieveData) (successCode int32, result models.SmContextRetrievedData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PDU_HandleRetrieveSmContext has not been implemented")
	return
}
func (p *Producer) PDU_HandleSendMoData(smContextRef string, body models.SendMoDataRequest) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PDU_HandleSendMoData has not been implemented")
	return
}
func (p *Producer) PDU_HandleUpdateSmContext(smContextRef string, body models.SmContextUpdateData) (successCode int32, result models.SmContextUpdatedData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PDU_HandleUpdateSmContext has not been implemented")
	return
}
func (p *Producer) PDU_HandlePostPduSessions(body models.PduSessionCreateData) (successCode int32, result models.PduSessionCreatedData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PDU_HandlePostPduSessions has not been implemented")
	return
}
func (p *Producer) PDU_HandlePostSmContexts(body models.PostSmContextsRequest) (successCode int32, result models.SmContextCreatedData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PDU_HandlePostSmContexts has not been implemented")
	return
}
