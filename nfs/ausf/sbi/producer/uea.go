package producer

import (
	"etri5gc/sbi"
	"etri5gc/sbi/models"
)

func (p *Producer) UEA_HandleDelete5gAkaAuthenticationResult(authCtxId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UEA_HandleDelete5gAkaAuthenticationResult has not been implemented")
	return
}
func (p *Producer) UEA_HandleDeleteEapAuthenticationResult(authCtxId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UEA_HandleDeleteEapAuthenticationResult has not been implemented")
	return
}
func (p *Producer) UEA_HandleEapAuthMethod(authCtxId string, body *models.EapSession) (successCode int32, result models.EapSession, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UEA_HandleEapAuthMethod has not been implemented")
	return
}
func (p *Producer) UEA_HandleRgAuthenticationsPost(body models.RgAuthenticationInfo) (successCode int32, result models.RgAuthCtx, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UEA_HandleRgAuthenticationsPost has not been implemented")
	return
}
func (p *Producer) UEA_HandleUeAuthenticationsAuthCtxId5gAkaConfirmationPut(authCtxId string, body *models.ConfirmationData) (successCode int32, result models.ConfirmationDataResponse, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UEA_HandleUeAuthenticationsAuthCtxId5gAkaConfirmationPut has not been implemented")
	return
}
func (p *Producer) UEA_HandleUeAuthenticationsDeregisterPost(body models.DeregistrationInfo) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UEA_HandleUeAuthenticationsDeregisterPost has not been implemented")
	return
}
func (p *Producer) UEA_HandleUeAuthenticationsPost(body models.AuthenticationInfo) (successCode int32, result models.UEAuthenticationCtx, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UEA_HandleUeAuthenticationsPost has not been implemented")
	return
}
