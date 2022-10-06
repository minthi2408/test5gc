package producer

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
)

func (p *Producer) SOR_HandleSupiUeSorPost(supi string, body models.SorInfo) (successCode int32, result models.SorSecurityInfo, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SOR_HandleSupiUeSorPost has not been implemented")
	return
}
