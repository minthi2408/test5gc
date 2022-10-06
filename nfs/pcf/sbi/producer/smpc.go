package producer

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
)

func (p *Producer) SMPC_HandleDeleteSMPolicy(smPolicyId string, body models.SmPolicyDeleteData) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SMPC_HandleDeleteSMPolicy has not been implemented")
	return
}
func (p *Producer) SMPC_HandleGetSMPolicy(smPolicyId string) (successCode int32, result models.SmPolicyControl, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SMPC_HandleGetSMPolicy has not been implemented")
	return
}
func (p *Producer) SMPC_HandleUpdateSMPolicy(smPolicyId string, body models.SmPolicyUpdateContextData) (successCode int32, result models.SmPolicyDecision, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SMPC_HandleUpdateSMPolicy has not been implemented")
	return
}
func (p *Producer) SMPC_HandleCreateSMPolicy(body models.SmPolicyContextData) (successCode int32, result models.SmPolicyDecision, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SMPC_HandleCreateSMPolicy has not been implemented")
	return
}
