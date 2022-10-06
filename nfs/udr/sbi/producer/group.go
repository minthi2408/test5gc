package producer

import (
	"etri5gc/sbi"
	"etri5gc/sbi/models"
)

func (p *Producer) GROUP_HandleGetNfGroupIDs(nfType []models.NFType, subscriberId string) (successCode int32, result map[string]string, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("GROUP_HandleGetNfGroupIDs has not been implemented")
	return
}
