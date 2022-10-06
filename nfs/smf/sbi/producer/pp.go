package producer

import (
	"etri5gc/sbi"
	"etri5gc/sbi/models"
)

func (p *Producer) PP_HandleCreate5GVNGroup(extGroupId string, body models.Model5GVnGroupConfiguration) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PP_HandleCreate5GVNGroup has not been implemented")
	return
}
func (p *Producer) PP_HandleDelete5GVNGroup(extGroupId string, mtcProviderInfo string, afId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PP_HandleDelete5GVNGroup has not been implemented")
	return
}
func (p *Producer) PP_HandleGet5GVNGroup(extGroupId string) (successCode int32, result models.Model5GVnGroupConfiguration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PP_HandleGet5GVNGroup has not been implemented")
	return
}
func (p *Producer) PP_HandleModify5GVNGroup(extGroupId string, supportedFeatures string, body models.Model5GVnGroupConfiguration) (successCode int32, result models.PatchResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PP_HandleModify5GVNGroup has not been implemented")
	return
}
func (p *Producer) PP_HandleUpdate(ueId models.UpdateUeIdParameter, supportedFeatures string, body models.PpData) (successCode int32, result models.PatchResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("PP_HandleUpdate has not been implemented")
	return
}
