package producer

import (
	"etri5gc/sbi"
	"etri5gc/sbi/models"
)

func (p *Producer) LOC_HandleCancelLocation(ueContextId string, body models.CancelPosInfo) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("LOC_HandleCancelLocation has not been implemented")
	return
}
func (p *Producer) LOC_HandleProvideLocationInfo(ueContextId string, body models.RequestLocInfo) (successCode int32, result models.ProvideLocInfo, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("LOC_HandleProvideLocationInfo has not been implemented")
	return
}
func (p *Producer) LOC_HandleProvidePositioningInfo(ueContextId string, body models.RequestPosInfo) (successCode int32, result models.ProvidePosInfo, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("LOC_HandleProvidePositioningInfo has not been implemented")
	return
}
