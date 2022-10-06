package producer

import (
	"etri5gc/sbi"
	"etri5gc/sbi/models"
)

func (p *Producer) MT_HandleProvideLocationInfo(supi string, body models.LocationInfoRequest) (successCode int32, result models.LocationInfoResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("MT_HandleProvideLocationInfo has not been implemented")
	return
}
func (p *Producer) MT_HandleQueryUeInfo(supi string, fields []string, supportedFeatures string) (successCode int32, result models.UeInfo, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("MT_HandleQueryUeInfo has not been implemented")
	return
}
