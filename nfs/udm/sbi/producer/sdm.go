package producer

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
)

func (p *Producer) SDM_HandleGetAmData(supi string, supportedFeatures string, plmnId *models.PlmnId, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.AccessAndMobilitySubscriptionData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleGetAmData has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetEcrData(supi string, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.EnhancedCoverageRestrictionData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleGetEcrData has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetSupiOrGpsi(ueId string, supportedFeatures string, appPortId *models.AppPortId, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.IdTranslationResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleGetSupiOrGpsi has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetGroupIdentifiers(extGroupId string, intGroupId string, ueIdInd *bool, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.GroupIdentifiers, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleGetGroupIdentifiers has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetLcsBcaData(supi string, supportedFeatures string, plmnId *models.PlmnId, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.LcsBroadcastAssistanceTypesData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleGetLcsBcaData has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetLcsMoData(supi string, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.LcsMoData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleGetLcsMoData has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetLcsPrivacyData(ueId string, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.LcsPrivacyData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleGetLcsPrivacyData has not been implemented")
	return
}
func (p *Producer) SDM_HandleCAGAck(supi string, body *models.AcknowledgeInfo) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleCAGAck has not been implemented")
	return
}
func (p *Producer) SDM_HandleSNSSAIsAck(supi string, body *models.AcknowledgeInfo) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleSNSSAIsAck has not been implemented")
	return
}
func (p *Producer) SDM_HandleSorAckInfo(supi string, body *models.AcknowledgeInfo) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleSorAckInfo has not been implemented")
	return
}
func (p *Producer) SDM_HandleUpuAck(supi string, body *models.AcknowledgeInfo) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleUpuAck has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetDataSets(supi string, datasetNames []models.DataSetName, plmnId *models.PlmnId, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.SubscriptionDataSets, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleGetDataSets has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetSharedData(sharedDataIds []string, supportedFeatures string, supportedFeatures2 string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result []models.SharedData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleGetSharedData has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetIndividualSharedData(sharedDataId []string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.SharedData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleGetIndividualSharedData has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetSmfSelData(supi string, supportedFeatures string, plmnId *models.PlmnId, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.SmfSelectionSubscriptionData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleGetSmfSelData has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetSmsMngtData(supi string, supportedFeatures string, plmnId *models.PlmnId, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.SmsManagementSubscriptionData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleGetSmsMngtData has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetSmsData(supi string, supportedFeatures string, plmnId *models.PlmnId, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.SmsSubscriptionData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleGetSmsData has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetSmData(supi string, supportedFeatures string, singleNssai *models.Snssai, dnn string, plmnId *models.PlmnId, ifNoneMatch string, ifModifiedSince string) (successCode int32, result []models.SessionManagementSubscriptionData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleGetSmData has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetNSSAI(supi string, supportedFeatures string, plmnId *models.PlmnId, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.Nssai, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleGetNSSAI has not been implemented")
	return
}
func (p *Producer) SDM_HandleSubscribe(ueId string, body models.SdmSubscription) (successCode int32, result models.SdmSubscription, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleSubscribe has not been implemented")
	return
}
func (p *Producer) SDM_HandleSubscribeToSharedData(body models.SdmSubscription) (successCode int32, result models.SdmSubscription, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleSubscribeToSharedData has not been implemented")
	return
}
func (p *Producer) SDM_HandleUnsubscribe(ueId string, subscriptionId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleUnsubscribe has not been implemented")
	return
}
func (p *Producer) SDM_HandleUnsubscribeForSharedData(subscriptionId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleUnsubscribeForSharedData has not been implemented")
	return
}
func (p *Producer) SDM_HandleModify(ueId string, subscriptionId string, supportedFeatures string, body models.SdmSubsModification) (successCode int32, result models.Modify200Response, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleModify has not been implemented")
	return
}
func (p *Producer) SDM_HandleModifySharedDataSubs(subscriptionId string, supportedFeatures string, body models.SdmSubsModification) (successCode int32, result models.Modify200Response, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleModifySharedDataSubs has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetTraceConfigData(supi string, supportedFeatures string, plmnId *models.PlmnId, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.TraceDataResponse, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleGetTraceConfigData has not been implemented")
	return
}
func (p *Producer) SDM_HandleUpdateSORInfo(supi string, body *models.SorUpdateInfo) (successCode int32, result models.SorInfo, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleUpdateSORInfo has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetUeCtxInAmfData(supi string, supportedFeatures string) (successCode int32, result models.UeContextInAmfData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleGetUeCtxInAmfData has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetUeCtxInSmfData(supi string, supportedFeatures string) (successCode int32, result models.UeContextInSmfData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleGetUeCtxInSmfData has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetUeCtxInSmsfData(supi string, supportedFeatures string) (successCode int32, result models.UeContextInSmsfData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleGetUeCtxInSmsfData has not been implemented")
	return
}
func (p *Producer) SDM_HandleGetV2xData(supi string, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.V2xSubscriptionData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("SDM_HandleGetV2xData has not been implemented")
	return
}
