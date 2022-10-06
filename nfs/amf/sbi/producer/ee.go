package producer

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
)

func (p *Producer) EE_HandleDeleteSubscription(subscriptionId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("EE_HandleDeleteSubscription has not been implemented")
	return
}
func (p *Producer) EE_HandleModifySubscription(subscriptionId string, body models.ModifySubscriptionRequest) (successCode int32, result models.AmfUpdatedEventSubscription, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("EE_HandleModifySubscription has not been implemented")
	return
}
func (p *Producer) EE_HandleCreateSubscription(body models.AmfCreateEventSubscription) (successCode int32, result models.AmfCreatedEventSubscription, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("EE_HandleCreateSubscription has not been implemented")
	return
}
