package producer

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
)

func (p *Producer) NIDD_HandleDeliver(pduSessionRef string, body models.DeliverRequest) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("NIDD_HandleDeliver has not been implemented")
	return
}
