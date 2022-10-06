package producer

import (
	"etri5gc/sbi"
	"etri5gc/sbi/models"
)

func (p *Producer) NIDDAU_HandleAuthorizeNiddData(ueIdentity string, body models.AuthorizationInfo) (successCode int32, result models.AuthorizationData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("NIDDAU_HandleAuthorizeNiddData has not been implemented")
	return
}
