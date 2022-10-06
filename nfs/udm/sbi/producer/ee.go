package producer

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
)

func (p *Producer) EE_HandleCreateEeSubscription(ueIdentity string, body models.EeSubscription) (successCode int32, result models.CreatedEeSubscription, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("EE_HandleCreateEeSubscription has not been implemented")
	return
}
func (p *Producer) EE_HandleDeleteEeSubscription(ueIdentity string, subscriptionId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("EE_HandleDeleteEeSubscription has not been implemented")
	return
}
func (p *Producer) EE_HandleUpdateEeSubscription(ueIdentity string, subscriptionId string, supportedFeatures string, body []models.PatchItem) (successCode int32, result models.PatchResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("EE_HandleUpdateEeSubscription has not been implemented")
	return
}
