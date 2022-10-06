package producer

import (
	"etri5gc/sbi"
	"etri5gc/sbi/models"
)

func (p *Producer) UECM_HandleGet3GppRegistration(ueId string, supportedFeatures string) (successCode int32, result models.Amf3GppAccessRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleGet3GppRegistration has not been implemented")
	return
}
func (p *Producer) UECM_HandleGetNon3GppRegistration(ueId string, supportedFeatures string) (successCode int32, result models.AmfNon3GppAccessRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleGetNon3GppRegistration has not been implemented")
	return
}
func (p *Producer) UECM_HandleCall3GppRegistration(ueId string, body models.Amf3GppAccessRegistration) (successCode int32, result models.Amf3GppAccessRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleCall3GppRegistration has not been implemented")
	return
}
func (p *Producer) UECM_HandleNon3GppRegistration(ueId string, body models.AmfNon3GppAccessRegistration) (successCode int32, result models.AmfNon3GppAccessRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleNon3GppRegistration has not been implemented")
	return
}
func (p *Producer) UECM_HandleIpSmGwDeregistration(ueId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleIpSmGwDeregistration has not been implemented")
	return
}
func (p *Producer) UECM_HandleIpSmGwRegistration(ueId string, body models.IpSmGwRegistration) (successCode int32, result models.IpSmGwRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleIpSmGwRegistration has not been implemented")
	return
}
func (p *Producer) UECM_HandleGetIpSmGwRegistration(ueId string) (successCode int32, result models.IpSmGwRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleGetIpSmGwRegistration has not been implemented")
	return
}
func (p *Producer) UECM_HandlePeiUpdate(ueId string, body models.PeiUpdateInfo) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandlePeiUpdate has not been implemented")
	return
}
func (p *Producer) UECM_HandleUpdate3GppRegistration(ueId string, supportedFeatures string, body models.Amf3GppAccessRegistrationModification) (successCode int32, result models.PatchResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleUpdate3GppRegistration has not been implemented")
	return
}
func (p *Producer) UECM_HandleUpdateNon3GppRegistration(ueId string, supportedFeatures string, body models.AmfNon3GppAccessRegistrationModification) (successCode int32, result models.PatchResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleUpdateNon3GppRegistration has not been implemented")
	return
}
func (p *Producer) UECM_HandleRetrieveSmfRegistration(ueId string, pduSessionId int32) (successCode int32, result models.SmfRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleRetrieveSmfRegistration has not been implemented")
	return
}
func (p *Producer) UECM_HandleSmfDeregistration(ueId string, pduSessionId int32, smfSetId string, smfInstanceId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleSmfDeregistration has not been implemented")
	return
}
func (p *Producer) UECM_HandleGetSmfRegistration(ueId string, singleNssai *models.Snssai, dnn string, supportedFeatures string) (successCode int32, result models.SmfRegistrationInfo, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleGetSmfRegistration has not been implemented")
	return
}
func (p *Producer) UECM_HandleRegistration(ueId string, pduSessionId int32, body models.SmfRegistration) (successCode int32, result models.SmfRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleRegistration has not been implemented")
	return
}
func (p *Producer) UECM_HandleGet3GppSmsfRegistration(ueId string, supportedFeatures string) (successCode int32, result models.SmsfRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleGet3GppSmsfRegistration has not been implemented")
	return
}
func (p *Producer) UECM_HandleCall3GppSmsfDeregistration(ueId string, smsfSetId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleCall3GppSmsfDeregistration has not been implemented")
	return
}
func (p *Producer) UECM_HandleNon3GppSmsfDeregistration(ueId string, smsfSetId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleNon3GppSmsfDeregistration has not been implemented")
	return
}
func (p *Producer) UECM_HandleGetNon3GppSmsfRegistration(ueId string, supportedFeatures string) (successCode int32, result models.SmsfRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleGetNon3GppSmsfRegistration has not been implemented")
	return
}
func (p *Producer) UECM_HandleCall3GppSmsfRegistration(ueId string, body models.SmsfRegistration) (successCode int32, result models.SmsfRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleCall3GppSmsfRegistration has not been implemented")
	return
}
func (p *Producer) UECM_HandleNon3GppSmsfRegistration(ueId string, body models.SmsfRegistration) (successCode int32, result models.SmsfRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleNon3GppSmsfRegistration has not been implemented")
	return
}
func (p *Producer) UECM_HandleDeregAMF(ueId string, body models.AmfDeregInfo) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleDeregAMF has not been implemented")
	return
}
func (p *Producer) UECM_HandleTriggerPCSCFRestoration(body models.TriggerRequest) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleTriggerPCSCFRestoration has not been implemented")
	return
}
func (p *Producer) UECM_HandleGetRegistrations(ueId string, registrationDatasetNames []models.RegistrationDataSetName, supportedFeatures string, singleNssai *models.Snssai, dnn string) (successCode int32, result models.RegistrationDataSets, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleGetRegistrations has not been implemented")
	return
}
func (p *Producer) UECM_HandleGetLocationInfo(ueId string, supportedFeatures string) (successCode int32, result models.LocationInfo, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("UECM_HandleGetLocationInfo has not been implemented")
	return
}
