package producer

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
)

func (p *Producer) MT_HandleProvideDomainSelectionInfo(ueContextId string, infoClass *models.UeContextInfoClass, supportedFeatures string, oldGuami *models.Guami) (successCode int32, result models.UeContextInfo, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("MT_HandleProvideDomainSelectionInfo has not been implemented")
	return
}
func (p *Producer) MT_HandleEnableUeReachability(ueContextId string, body models.EnableUeReachabilityReqData) (successCode int32, result models.EnableUeReachabilityRspData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("MT_HandleEnableUeReachability has not been implemented")
	return
}
