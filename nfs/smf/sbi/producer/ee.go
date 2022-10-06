package producer

import (
	"etri5gc/sbi"
	"etri5gc/sbi/models"
)

func (p *Producer) EE_HandleDeleteIndividualSubcription(subId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("EE_HandleDeleteIndividualSubcription has not been implemented")
	return
}
func (p *Producer) EE_HandleGetIndividualSubcription(subId string) (successCode int32, result models.NsmfEventExposure, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("EE_HandleGetIndividualSubcription has not been implemented")
	return
}
func (p *Producer) EE_HandleReplaceIndividualSubcription(subId string, body models.NsmfEventExposure) (successCode int32, result models.NsmfEventExposure, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("EE_HandleReplaceIndividualSubcription has not been implemented")
	return
}
func (p *Producer) EE_HandleCreateIndividualSubcription(body models.NsmfEventExposure) (successCode int32, result models.NsmfEventExposure, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("EE_HandleCreateIndividualSubcription has not been implemented")
	return
}
