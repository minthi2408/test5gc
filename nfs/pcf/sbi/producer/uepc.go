package producer

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
)

func (p *Producer) UEPC_HandleDeleteIndividualUEPolicyAssociation(polAssoId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UEPC_HandleDeleteIndividualUEPolicyAssociation has not been implemented")
	return
}
func (p *Producer) UEPC_HandleReadIndividualUEPolicyAssociation(polAssoId string) (successCode int32, result models.PolicyAssociation, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UEPC_HandleReadIndividualUEPolicyAssociation has not been implemented")
	return
}
func (p *Producer) UEPC_HandleReportObservedEventTriggersForIndividualUEPolicyAssociation(polAssoId string, body models.PolicyAssociationUpdateRequest) (successCode int32, result models.PolicyUpdate, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UEPC_HandleReportObservedEventTriggersForIndividualUEPolicyAssociation has not been implemented")
	return
}
func (p *Producer) UEPC_HandleCreateIndividualUEPolicyAssociation(body models.PolicyAssociationRequest) (successCode int32, result models.PolicyAssociation, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UEPC_HandleCreateIndividualUEPolicyAssociation has not been implemented")
	return
}
