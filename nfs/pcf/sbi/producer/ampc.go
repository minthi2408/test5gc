package producer

import (
	"etri5gc/sbi"
	"etri5gc/sbi/models"
)

func (p *Producer) AMPC_HandleCreateIndividualAMPolicyAssociation(body models.PolicyAssociationRequest) (successCode int32, result models.PolicyAssociation, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("AMPC_HandleCreateIndividualAMPolicyAssociation has not been implemented")
	return
}
func (p *Producer) AMPC_HandleDeleteIndividualAMPolicyAssociation(polAssoId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("AMPC_HandleDeleteIndividualAMPolicyAssociation has not been implemented")
	return
}
func (p *Producer) AMPC_HandleReadIndividualAMPolicyAssociation(polAssoId string) (successCode int32, result models.PolicyAssociation, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("AMPC_HandleReadIndividualAMPolicyAssociation has not been implemented")
	return
}
func (p *Producer) AMPC_HandleReportObservedEventTriggersForIndividualAMPolicyAssociation(polAssoId string, body models.PolicyAssociationUpdateRequest) (successCode int32, result models.PolicyUpdate, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("AMPC_HandleReportObservedEventTriggersForIndividualAMPolicyAssociation has not been implemented")
	return
}
