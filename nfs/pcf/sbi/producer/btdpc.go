package producer

import (
	"etri5gc/sbi"
	"etri5gc/sbi/models"
)

func (p *Producer) BTDPC_HandleCreateBDTPolicy(body models.BdtReqData) (successCode int32, result models.BdtPolicy, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("BTDPC_HandleCreateBDTPolicy has not been implemented")
	return
}
func (p *Producer) BTDPC_HandleGetBDTPolicy(bdtPolicyId string) (successCode int32, result models.BdtPolicy, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("BTDPC_HandleGetBDTPolicy has not been implemented")
	return
}
func (p *Producer) BTDPC_HandleUpdateBDTPolicy(bdtPolicyId string, body models.PatchBdtPolicy) (successCode int32, result models.BdtPolicy, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("BTDPC_HandleUpdateBDTPolicy has not been implemented")
	return
}
