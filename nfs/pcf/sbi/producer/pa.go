package producer

import (
	"etri5gc/sbi"
	"etri5gc/sbi/models"
)

func (p *Producer) PA_HandlePostAppSessions(body models.AppSessionContext) (successCode int32, result models.AppSessionContext, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PA_HandlePostAppSessions has not been implemented")
	return
}
func (p *Producer) PA_HandleDeleteEventsSubsc(appSessionId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PA_HandleDeleteEventsSubsc has not been implemented")
	return
}
func (p *Producer) PA_HandleUpdateEventsSubsc(appSessionId string, body models.EventsSubscReqData) (successCode int32, result models.EventsSubscPutData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PA_HandleUpdateEventsSubsc has not been implemented")
	return
}
func (p *Producer) PA_HandleDeleteAppSession(appSessionId string, body *models.EventsSubscReqData) (successCode int32, result models.AppSessionContext, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PA_HandleDeleteAppSession has not been implemented")
	return
}
func (p *Producer) PA_HandleGetAppSession(appSessionId string) (successCode int32, result models.AppSessionContext, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PA_HandleGetAppSession has not been implemented")
	return
}
func (p *Producer) PA_HandleModAppSession(appSessionId string, body models.AppSessionContextUpdateDataPatch) (successCode int32, result models.AppSessionContext, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PA_HandleModAppSession has not been implemented")
	return
}
func (p *Producer) PA_HandlePcscfRestoration(body models.PcscfRestorationRequestData) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PA_HandlePcscfRestoration has not been implemented")
	return
}
