package producer

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
)

func (p *Producer) UEAU_HandleConfirmAuth(supi string, body models.AuthEvent) (successCode int32, result models.AuthEvent, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UEAU_HandleConfirmAuth has not been implemented")
	return
}
func (p *Producer) UEAU_HandleDeleteAuth(supi string, authEventId string, body models.AuthEvent) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UEAU_HandleDeleteAuth has not been implemented")
	return
}
func (p *Producer) UEAU_HandleGenerateAuthData(supiOrSuci string, body models.AuthenticationInfoRequest) (successCode int32, result models.AuthenticationInfoResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UEAU_HandleGenerateAuthData has not been implemented")
	return
}
func (p *Producer) UEAU_HandleGenerateAv(supi string, hssAuthType models.HssAuthTypeInUri, body models.HssAuthenticationInfoRequest) (successCode int32, result models.HssAuthenticationInfoResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UEAU_HandleGenerateAv has not been implemented")
	return
}
func (p *Producer) UEAU_HandleGetRgAuthData(supiOrSuci string, authenticatedInd bool, supportedFeatures string, plmnId *models.PlmnId, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.RgAuthCtx, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UEAU_HandleGetRgAuthData has not been implemented")
	return
}
