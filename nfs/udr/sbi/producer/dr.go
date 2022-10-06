package producer

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
)

func (p *Producer) DR_HandleAmfContext3gpp(ueId string, supportedFeatures string, body []models.PatchItem) (successCode int32, result models.PatchResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleAmfContext3gpp has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateAmfContext3gpp(ueId string, body models.Amf3GppAccessRegistration) (successCode int32, result models.Amf3GppAccessRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateAmfContext3gpp has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryAmfContext3gpp(ueId string, fields []string, supportedFeatures string) (successCode int32, result models.Amf3GppAccessRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryAmfContext3gpp has not been implemented")
	return
}
func (p *Producer) DR_HandleAmfContextNon3gpp(ueId string, supportedFeatures string, body []models.PatchItem) (successCode int32, result models.PatchResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleAmfContextNon3gpp has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateAmfContextNon3gpp(ueId string, body models.AmfNon3GppAccessRegistration) (successCode int32, result models.Amf3GppAccessRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateAmfContextNon3gpp has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryAmfContextNon3gpp(ueId string, fields []string, supportedFeatures string) (successCode int32, result models.AmfNon3GppAccessRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryAmfContextNon3gpp has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateAMFSubscriptions(ueId string, subsId string, body []models.AmfSubscriptionInfo) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateAMFSubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateOrReplaceAccessAndMobilityData(ueId string, body models.AccessAndMobilityData) (successCode int32, result models.AccessAndMobilityData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateOrReplaceAccessAndMobilityData has not been implemented")
	return
}
func (p *Producer) DR_HandleDeleteAccessAndMobilityData(ueId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDeleteAccessAndMobilityData has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryAccessAndMobilityData(ueId string, suppFeat string) (successCode int32, result models.AccessAndMobilityData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryAccessAndMobilityData has not been implemented")
	return
}
func (p *Producer) DR_HandleUpdateAccessAndMobilityData(ueId string, body models.AccessAndMobilityData) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleUpdateAccessAndMobilityData has not been implemented")
	return
}
func (p *Producer) DR_HandleReadAccessAndMobilityPolicyData(ueId string) (successCode int32, result models.AmPolicyData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReadAccessAndMobilityPolicyData has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryAmData(ueId string, servingPlmnId string, fields []string, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.AccessAndMobilitySubscriptionData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryAmData has not been implemented")
	return
}
func (p *Producer) DR_HandleModifyAmfSubscriptionInfo(ueId string, subsId string, supportedFeatures string, body []models.PatchItem) (successCode int32, result models.PatchResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleModifyAmfSubscriptionInfo has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateIndividualApplicationDataSubscription(body models.ApplicationDataSubs) (successCode int32, result models.ApplicationDataSubs, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateIndividualApplicationDataSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleReadApplicationDataChangeSubscriptions(dataFilter *models.DataFilter) (successCode int32, result []models.ApplicationDataSubs, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReadApplicationDataChangeSubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleDeleteAuthenticationStatus(ueId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDeleteAuthenticationStatus has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryAuthenticationStatus(ueId string, fields []string, supportedFeatures string) (successCode int32, result models.AuthEvent, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryAuthenticationStatus has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryAuthSubsData(ueId string, supportedFeatures string) (successCode int32, result models.AuthenticationSubscription, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryAuthSubsData has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateAuthenticationSoR(ueId string, supportedFeatures string, body models.SorData) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateAuthenticationSoR has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryAuthSoR(ueId string, supportedFeatures string) (successCode int32, result models.SorData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryAuthSoR has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateAuthenticationStatus(ueId string, body models.AuthEvent) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateAuthenticationStatus has not been implemented")
	return
}
func (p *Producer) DR_HandleModifyAuthenticationSubscription(ueId string, supportedFeatures string, body []models.PatchItem) (successCode int32, result models.PatchResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleModifyAuthenticationSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateAuthenticationUPU(ueId string, supportedFeatures string, body *models.UpuData) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateAuthenticationUPU has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryAuthUPU(ueId string, supportedFeatures string) (successCode int32, result models.UpuData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryAuthUPU has not been implemented")
	return
}
func (p *Producer) DR_HandleReadBdtData(bdtRefIds []string, suppFeat string) (successCode int32, result []models.BdtData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReadBdtData has not been implemented")
	return
}
func (p *Producer) DR_HandleReadBdtPolicyData(bdtPolicyIds []string, internalGroupIds []string, supis []string) (successCode int32, result []models.BdtPolicyData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReadBdtPolicyData has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryCagAck(ueId string, supportedFeatures string) (successCode int32, result models.CagAckData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryCagAck has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateCagUpdateAck(ueId string, supportedFeatures string, body models.CagAckData) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateCagUpdateAck has not been implemented")
	return
}
func (p *Producer) DR_HandleQuery5GVnGroupInternal(internalGroupIds []string) (successCode int32, result map[string]models.Model5GVnGroupConfiguration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQuery5GVnGroupInternal has not been implemented")
	return
}
func (p *Producer) DR_HandleQuery5GVnGroup(gpsis []string) (successCode int32, result map[string]models.Model5GVnGroupConfiguration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQuery5GVnGroup has not been implemented")
	return
}
func (p *Producer) DR_HandleCreate5GVnGroup(externalGroupId string, body models.Model5GVnGroupConfiguration) (successCode int32, result models.Model5GVnGroupConfiguration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreate5GVnGroup has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryContextData(ueId string, contextDatasetNames []models.ContextDataSetName) (successCode int32, result models.ContextDataSets, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryContextData has not been implemented")
	return
}
func (p *Producer) DR_HandleDelete5GVnGroup(externalGroupId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDelete5GVnGroup has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryCoverageRestrictionData(ueId string, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.EnhancedCoverageRestrictionData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryCoverageRestrictionData has not been implemented")
	return
}
func (p *Producer) DR_HandleRemoveAmfSubscriptionsInfo(ueId string, subsId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleRemoveAmfSubscriptionsInfo has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryEEData(ueId string, fields []string, supportedFeatures string) (successCode int32, result models.EeProfileData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryEEData has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryGroupEEData(ueGroupId string, supportedFeatures string) (successCode int32, result models.EeGroupProfileData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryGroupEEData has not been implemented")
	return
}
func (p *Producer) DR_HandleModifyEeGroupSubscription(ueGroupId string, subsId string, supportedFeatures string, body []models.PatchItem) (successCode int32, result models.PatchResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleModifyEeGroupSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryEeGroupSubscription(ueGroupId string, subsId string) (successCode int32, result map[string]interface{}, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryEeGroupSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleRemoveEeGroupSubscriptions(ueGroupId string, subsId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleRemoveEeGroupSubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleUpdateEeGroupSubscriptions(ueGroupId string, subsId string, body models.EeSubscription) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleUpdateEeGroupSubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateEeGroupSubscriptions(ueGroupId string, body models.EeSubscription) (successCode int32, result models.EeSubscription, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateEeGroupSubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryEeGroupSubscriptions(ueGroupId string, supportedFeatures string) (successCode int32, result []models.EeSubscription, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryEeGroupSubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleModifyEesubscription(ueId string, subsId string, supportedFeatures string, body []models.PatchItem) (successCode int32, result models.PatchResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleModifyEesubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryeeSubscription(ueId string, subsId string) (successCode int32, result map[string]interface{}, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryeeSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleRemoveeeSubscriptions(ueId string, subsId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleRemoveeeSubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleUpdateEesubscriptions(ueId string, subsId string, body models.EeSubscription) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleUpdateEesubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateEeSubscriptions(ueId string, body models.EeSubscription) (successCode int32, result models.EeSubscription, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateEeSubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryeesubscriptions(ueId string, supportedFeatures string) (successCode int32, result []models.EeSubscription, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryeesubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateIndividualExposureDataSubscription(body models.ExposureDataSubscription) (successCode int32, result models.ExposureDataSubscription, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateIndividualExposureDataSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleGetGroupIdentifiers(extGroupId string, intGroupId string, ueIdInd *bool, supportedFeatures string) (successCode int32, result models.GroupIdentifiers, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleGetGroupIdentifiers has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateHSSSubscriptions(ueId string, subsId string, body models.HssSubscriptionInfo) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateHSSSubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleGetHssSubscriptionInfo(ueId string, subsId string) (successCode int32, result models.SmfSubscriptionInfo, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleGetHssSubscriptionInfo has not been implemented")
	return
}
func (p *Producer) DR_HandleModifyHssSubscriptionInfo(ueId string, subsId string, supportedFeatures string, body []models.PatchItem) (successCode int32, result models.PatchResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleModifyHssSubscriptionInfo has not been implemented")
	return
}
func (p *Producer) DR_HandleRemoveHssSubscriptionsInfo(ueId string, subsId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleRemoveHssSubscriptionsInfo has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateHSSSDMSubscriptions(ueId string, subsId string, body models.HssSubscriptionInfo) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateHSSSDMSubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleGetHssSDMSubscriptionInfo(ueId string, subsId string) (successCode int32, result models.SmfSubscriptionInfo, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleGetHssSDMSubscriptionInfo has not been implemented")
	return
}
func (p *Producer) DR_HandleModifyHssSDMSubscriptionInfo(ueId string, subsId string, supportedFeatures string, body []models.PatchItem) (successCode int32, result models.PatchResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleModifyHssSDMSubscriptionInfo has not been implemented")
	return
}
func (p *Producer) DR_HandleRemoveHssSDMSubscriptionsInfo(ueId string, subsId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleRemoveHssSDMSubscriptionsInfo has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateIpSmGwContext(ueId string, body models.IpSmGwRegistration) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateIpSmGwContext has not been implemented")
	return
}
func (p *Producer) DR_HandleDeleteIpSmGwContext(ueId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDeleteIpSmGwContext has not been implemented")
	return
}
func (p *Producer) DR_HandleModifyIpSmGwContext(ueId string, body []models.PatchItem) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleModifyIpSmGwContext has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryIpSmGwContext(ueId string, fields []string, supportedFeatures string) (successCode int32, result models.IpSmGwRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryIpSmGwContext has not been implemented")
	return
}
func (p *Producer) DR_HandleReadIPTVCongifurationData(configIds []string, dnns []string, snssais []models.Snssai, supis []string, interGroupIds []string) (successCode int32, result []models.IptvConfigData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReadIPTVCongifurationData has not been implemented")
	return
}
func (p *Producer) DR_HandleDeleteIndividualApplicationDataSubscription(subsId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDeleteIndividualApplicationDataSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleReadIndividualApplicationDataSubscription(subsId string) (successCode int32, result models.ApplicationDataSubs, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReadIndividualApplicationDataSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleReplaceIndividualApplicationDataSubscription(subsId string, body models.ApplicationDataSubs) (successCode int32, result models.ApplicationDataSubs, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReplaceIndividualApplicationDataSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateIndividualAppliedBdtPolicyData(bdtPolicyId string, body models.BdtPolicyData) (successCode int32, result models.BdtPolicyData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateIndividualAppliedBdtPolicyData has not been implemented")
	return
}
func (p *Producer) DR_HandleDeleteIndividualAppliedBdtPolicyData(bdtPolicyId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDeleteIndividualAppliedBdtPolicyData has not been implemented")
	return
}
func (p *Producer) DR_HandleUpdateIndividualAppliedBdtPolicyData(bdtPolicyId string, body models.BdtPolicyDataPatch) (successCode int32, result models.BdtPolicyData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleUpdateIndividualAppliedBdtPolicyData has not been implemented")
	return
}
func (p *Producer) DR_HandleDeleteIndividualAuthenticationStatus(ueId string, servingNetworkName string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDeleteIndividualAuthenticationStatus has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryIndividualAuthenticationStatus(ueId string, servingNetworkName string, fields []string, supportedFeatures string) (successCode int32, result models.AuthEvent, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryIndividualAuthenticationStatus has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateIndividualAuthenticationStatus(ueId string, servingNetworkName string, body models.AuthEvent) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateIndividualAuthenticationStatus has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateIndividualBdtData(bdtReferenceId string, body models.BdtData) (successCode int32, result models.BdtData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateIndividualBdtData has not been implemented")
	return
}
func (p *Producer) DR_HandleDeleteIndividualBdtData(bdtReferenceId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDeleteIndividualBdtData has not been implemented")
	return
}
func (p *Producer) DR_HandleReadIndividualBdtData(bdtReferenceId string, suppFeat string) (successCode int32, result models.BdtData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReadIndividualBdtData has not been implemented")
	return
}
func (p *Producer) DR_HandleUpdateIndividualBdtData(bdtReferenceId string, body models.BdtDataPatch) (successCode int32, result models.BdtData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleUpdateIndividualBdtData has not been implemented")
	return
}
func (p *Producer) DR_HandleDeleteIndividualExposureDataSubscription(subId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDeleteIndividualExposureDataSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleReplaceIndividualExposureDataSubscription(subId string, body models.ExposureDataSubscription) (successCode int32, result models.ExposureDataSubscription, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReplaceIndividualExposureDataSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandlePartialReplaceIndividualIPTVConfigurationData(configurationId string, body models.IptvConfigDataPatch) (successCode int32, result models.IptvConfigData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandlePartialReplaceIndividualIPTVConfigurationData has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateOrReplaceIndividualIPTVConfigurationData(configurationId string, body models.IptvConfigData) (successCode int32, result models.IptvConfigData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateOrReplaceIndividualIPTVConfigurationData has not been implemented")
	return
}
func (p *Producer) DR_HandleDeleteIndividualIPTVConfigurationData(configurationId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDeleteIndividualIPTVConfigurationData has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateOrReplaceIndividualInfluenceData(influenceId string, body models.TrafficInfluData) (successCode int32, result models.TrafficInfluData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateOrReplaceIndividualInfluenceData has not been implemented")
	return
}
func (p *Producer) DR_HandleDeleteIndividualInfluenceData(influenceId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDeleteIndividualInfluenceData has not been implemented")
	return
}
func (p *Producer) DR_HandleUpdateIndividualInfluenceData(influenceId string, body models.TrafficInfluDataPatch) (successCode int32, result models.TrafficInfluData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleUpdateIndividualInfluenceData has not been implemented")
	return
}
func (p *Producer) DR_HandleDeleteIndividualInfluenceDataSubscription(subscriptionId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDeleteIndividualInfluenceDataSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleReadIndividualInfluenceDataSubscription(subscriptionId string) (successCode int32, result models.TrafficInfluSub, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReadIndividualInfluenceDataSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleReplaceIndividualInfluenceDataSubscription(subscriptionId string, body models.TrafficInfluSub) (successCode int32, result models.TrafficInfluSub, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReplaceIndividualInfluenceDataSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateOrReplaceIndividualPFDData(appId string, body models.PfdDataForAppExt) (successCode int32, result models.PfdDataForAppExt, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateOrReplaceIndividualPFDData has not been implemented")
	return
}
func (p *Producer) DR_HandleDeleteIndividualPFDData(appId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDeleteIndividualPFDData has not been implemented")
	return
}
func (p *Producer) DR_HandleReadIndividualPFDData(appId string) (successCode int32, result models.PfdDataForAppExt, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReadIndividualPFDData has not been implemented")
	return
}
func (p *Producer) DR_HandleDeleteIndividualPolicyDataSubscription(subsId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDeleteIndividualPolicyDataSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleReplaceIndividualPolicyDataSubscription(subsId string, body models.PolicyDataSubscription) (successCode int32, result models.PolicyDataSubscription, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReplaceIndividualPolicyDataSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateOrReplaceServiceParameterData(serviceParamId string, body models.ServiceParameterData) (successCode int32, result models.ServiceParameterData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateOrReplaceServiceParameterData has not been implemented")
	return
}
func (p *Producer) DR_HandleDeleteIndividualServiceParameterData(serviceParamId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDeleteIndividualServiceParameterData has not been implemented")
	return
}
func (p *Producer) DR_HandleUpdateIndividualServiceParameterData(serviceParamId string, body models.ServiceParameterDataPatch) (successCode int32, result models.ServiceParameterData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleUpdateIndividualServiceParameterData has not been implemented")
	return
}
func (p *Producer) DR_HandleReadInfluenceData(influenceIds []string, dnns []string, snssais []models.Snssai, internalGroupIds []string, supis []string, suppFeat string) (successCode int32, result []models.TrafficInfluData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReadInfluenceData has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateIndividualInfluenceDataSubscription(body models.TrafficInfluSub) (successCode int32, result models.TrafficInfluSub, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateIndividualInfluenceDataSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleReadInfluenceDataSubscriptions(dnn string, snssai *models.Snssai, internalGroupId string, supi string) (successCode int32, result []models.TrafficInfluSub, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReadInfluenceDataSubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryLcsBcaData(ueId string, servingPlmnId string, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.LcsBroadcastAssistanceTypesData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryLcsBcaData has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryLcsMoData(ueId string, fields []string, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.LcsMoData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryLcsMoData has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryLcsPrivacyData(ueId string, fields []string, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.LcsPrivacyData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryLcsPrivacyData has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateMessageWaitingData(ueId string, body models.MessageWaitingData) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateMessageWaitingData has not been implemented")
	return
}
func (p *Producer) DR_HandleDeleteMessageWaitingData(ueId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDeleteMessageWaitingData has not been implemented")
	return
}
func (p *Producer) DR_HandleModifyMessageWaitingData(ueId string, body []models.PatchItem) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleModifyMessageWaitingData has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryMessageWaitingData(ueId string, fields []string, supportedFeatures string) (successCode int32, result models.MessageWaitingData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryMessageWaitingData has not been implemented")
	return
}
func (p *Producer) DR_HandleModify5GVnGroup(externalGroupId string, supportedFeatures string, body []models.PatchItem) (successCode int32, result models.PatchResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleModify5GVnGroup has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryNssaiAck(ueId string, supportedFeatures string) (successCode int32, result models.NssaiAckData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryNssaiAck has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateOrUpdateNssaiAck(ueId string, supportedFeatures string, body models.NssaiAckData) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateOrUpdateNssaiAck has not been implemented")
	return
}
func (p *Producer) DR_HandleModifyOperSpecData(ueId string, supportedFeatures string, body []models.PatchItem) (successCode int32, result models.PatchResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleModifyOperSpecData has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryOperSpecData(ueId string, fields []string, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result map[string]models.OperatorSpecificDataContainer, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryOperSpecData has not been implemented")
	return
}
func (p *Producer) DR_HandleReadOperatorSpecificData(ueId string, fields []string, suppFeat string) (successCode int32, result map[string]models.OperatorSpecificDataContainer, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReadOperatorSpecificData has not been implemented")
	return
}
func (p *Producer) DR_HandleReplaceOperatorSpecificData(ueId string, body map[string]models.OperatorSpecificDataContainer) (successCode int32, result map[string]models.OperatorSpecificDataContainer, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReplaceOperatorSpecificData has not been implemented")
	return
}
func (p *Producer) DR_HandleUpdateOperatorSpecificData(ueId string, body []models.PatchItem) (successCode int32, result models.PatchResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleUpdateOperatorSpecificData has not been implemented")
	return
}
func (p *Producer) DR_HandleReadPFDData(appId []string) (successCode int32, result []models.PfdDataForAppExt, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReadPFDData has not been implemented")
	return
}
func (p *Producer) DR_HandleGetppData(ueId string, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.PpData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleGetppData has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryPPData(ueId string, supportedFeatures string) (successCode int32, result models.PpProfileData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryPPData has not been implemented")
	return
}
func (p *Producer) DR_HandleQuery5GVNGroupPPData(extGroupIds []string, supportedFeatures string) (successCode int32, result models.Pp5gVnGroupProfileData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQuery5GVNGroupPPData has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateOrReplaceSessionManagementData(ueId string, pduSessionId int32, body models.PduSessionManagementData) (successCode int32, result models.AccessAndMobilityData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateOrReplaceSessionManagementData has not been implemented")
	return
}
func (p *Producer) DR_HandleDeleteSessionManagementData(ueId string, pduSessionId int32) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDeleteSessionManagementData has not been implemented")
	return
}
func (p *Producer) DR_HandleQuerySessionManagementData(ueId string, pduSessionId int32, ipv4Addr string, ipv6Prefix models.Ipv6Prefix, dnn string, fields []string, suppFeat string) (successCode int32, result models.PduSessionManagementData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQuerySessionManagementData has not been implemented")
	return
}
func (p *Producer) DR_HandleReadPlmnUePolicySet(plmnId string) (successCode int32, result models.UePolicySet, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReadPlmnUePolicySet has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateIndividualPolicyDataSubscription(body models.PolicyDataSubscription) (successCode int32, result models.PolicyDataSubscription, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateIndividualPolicyDataSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryProvisionedData(ueId string, servingPlmnId string, datasetNames []models.DataSetName) (successCode int32, result models.ProvisionedDataSets, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryProvisionedData has not been implemented")
	return
}
func (p *Producer) DR_HandleModifyPpData(ueId string, supportedFeatures string, body []models.PatchItem) (successCode int32, result models.PatchResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleModifyPpData has not been implemented")
	return
}
func (p *Producer) DR_HandleGet5GVnGroupConfiguration(externalGroupId string) (successCode int32, result models.Model5GVnGroupConfiguration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleGet5GVnGroupConfiguration has not been implemented")
	return
}
func (p *Producer) DR_HandleGetAmfSubscriptionInfo(ueId string, subsId string) (successCode int32, result []models.AmfSubscriptionInfo, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleGetAmfSubscriptionInfo has not been implemented")
	return
}
func (p *Producer) DR_HandleGetIdentityData(ueId string, appPortId *models.AppPortId, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.IdentityData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleGetIdentityData has not been implemented")
	return
}
func (p *Producer) DR_HandleGetNiddAuData(ueId string, singleNssai models.Snssai, dnn string, mtcProviderInformation string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.AuthorizationData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleGetNiddAuData has not been implemented")
	return
}
func (p *Producer) DR_HandleGetOdbData(ueId string) (successCode int32, result models.OdbData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleGetOdbData has not been implemented")
	return
}
func (p *Producer) DR_HandleGetIndividualSharedData(sharedDataId string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.SharedData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleGetIndividualSharedData has not been implemented")
	return
}
func (p *Producer) DR_HandleGetSharedData(sharedDataIds []string, supportedFeatures string) (successCode int32, result []models.SharedData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleGetSharedData has not been implemented")
	return
}
func (p *Producer) DR_HandleModifysdmSubscription(ueId string, subsId string, supportedFeatures string, body []models.PatchItem) (successCode int32, result models.PatchResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleModifysdmSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleQuerysdmSubscription(ueId string, subsId string) (successCode int32, result map[string]interface{}, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQuerysdmSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleRemovesdmSubscriptions(ueId string, subsId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleRemovesdmSubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleUpdatesdmsubscriptions(ueId string, subsId string, body models.SdmSubscription) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleUpdatesdmsubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateSdmSubscriptions(ueId string, body models.SdmSubscription) (successCode int32, result models.SdmSubscription, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateSdmSubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleQuerysdmsubscriptions(ueId string, supportedFeatures string) (successCode int32, result []models.SdmSubscription, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQuerysdmsubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateSMFSubscriptions(ueId string, subsId string, body models.SmfSubscriptionInfo) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateSMFSubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleGetSmfSubscriptionInfo(ueId string, subsId string) (successCode int32, result models.SmfSubscriptionInfo, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleGetSmfSubscriptionInfo has not been implemented")
	return
}
func (p *Producer) DR_HandleModifySmfSubscriptionInfo(ueId string, subsId string, supportedFeatures string, body []models.PatchItem) (successCode int32, result models.PatchResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleModifySmfSubscriptionInfo has not been implemented")
	return
}
func (p *Producer) DR_HandleRemoveSmfSubscriptionsInfo(ueId string, subsId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleRemoveSmfSubscriptionsInfo has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateOrUpdateSmfRegistration(ueId string, pduSessionId int32, body models.SmfRegistration) (successCode int32, result models.SmfRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateOrUpdateSmfRegistration has not been implemented")
	return
}
func (p *Producer) DR_HandleDeleteSmfRegistration(ueId string, pduSessionId int32) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDeleteSmfRegistration has not been implemented")
	return
}
func (p *Producer) DR_HandleQuerySmfRegistration(ueId string, pduSessionId int32, fields []string, supportedFeatures string) (successCode int32, result models.SmfRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQuerySmfRegistration has not been implemented")
	return
}
func (p *Producer) DR_HandleQuerySmfRegList(ueId string, supportedFeatures string) (successCode int32, result []models.SmfRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQuerySmfRegList has not been implemented")
	return
}
func (p *Producer) DR_HandleQuerySmfSelectData(ueId string, servingPlmnId string, fields []string, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.SmfSelectionSubscriptionData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQuerySmfSelectData has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateSmsfContext3gpp(ueId string, body models.SmsfRegistration) (successCode int32, result models.SmsfRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateSmsfContext3gpp has not been implemented")
	return
}
func (p *Producer) DR_HandleDeleteSmsfContext3gpp(ueId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDeleteSmsfContext3gpp has not been implemented")
	return
}
func (p *Producer) DR_HandleQuerySmsfContext3gpp(ueId string, fields []string, supportedFeatures string) (successCode int32, result models.SmsfRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQuerySmsfContext3gpp has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateSmsfContextNon3gpp(ueId string, body models.SmsfRegistration) (successCode int32, result models.SmsfRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateSmsfContextNon3gpp has not been implemented")
	return
}
func (p *Producer) DR_HandleDeleteSmsfContextNon3gpp(ueId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDeleteSmsfContextNon3gpp has not been implemented")
	return
}
func (p *Producer) DR_HandleQuerySmsfContextNon3gpp(ueId string, fields []string, supportedFeatures string) (successCode int32, result models.SmsfRegistration, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQuerySmsfContextNon3gpp has not been implemented")
	return
}
func (p *Producer) DR_HandleQuerySmsMngData(ueId string, servingPlmnId string, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.SmsManagementSubscriptionData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQuerySmsMngData has not been implemented")
	return
}
func (p *Producer) DR_HandleQuerySmsData(ueId string, servingPlmnId string, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.SmsSubscriptionData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQuerySmsData has not been implemented")
	return
}
func (p *Producer) DR_HandleReadServiceParameterData(serviceParamIds []string, dnns []string, snssais []models.Snssai, internalGroupIds []string, supis []string, ueIpv4s []string, ueIpv6s []models.Ipv6Addr, ueMacs []string, suppFeat string) (successCode int32, result []models.ServiceParameterData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReadServiceParameterData has not been implemented")
	return
}
func (p *Producer) DR_HandleReadSessionManagementPolicyData(ueId string, snssai *models.Snssai, dnn string, fields []string, suppFeat string) (successCode int32, result models.SmPolicyData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReadSessionManagementPolicyData has not been implemented")
	return
}
func (p *Producer) DR_HandleUpdateSessionManagementPolicyData(ueId string, body models.SmPolicyDataPatch) (successCode int32, result models.SmPolicyData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleUpdateSessionManagementPolicyData has not been implemented")
	return
}
func (p *Producer) DR_HandleQuerySmData(ueId string, servingPlmnId string, singleNssai *models.Snssai, dnn string, fields []string, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result []models.SessionManagementSubscriptionData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQuerySmData has not been implemented")
	return
}
func (p *Producer) DR_HandleReadSponsorConnectivityData(sponsorId string) (successCode int32, result models.SponsorConnectivityData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReadSponsorConnectivityData has not been implemented")
	return
}
func (p *Producer) DR_HandleQuerySubsToNotify(ueId string, supportedFeatures string) (successCode int32, result []models.SubscriptionDataSubscriptions, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQuerySubsToNotify has not been implemented")
	return
}
func (p *Producer) DR_HandleRemoveMultipleSubscriptionDataSubscriptions(ueId string, nfInstanceId string, deleteAllNfs *bool, implicitUnsubscribeIndication *bool) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleRemoveMultipleSubscriptionDataSubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleSubscriptionDataSubscriptions(body models.SubscriptionDataSubscriptions) (successCode int32, result models.SubscriptionDataSubscriptions, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleSubscriptionDataSubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleModifysubscriptionDataSubscription(subsId string, supportedFeatures string, body []models.PatchItem) (successCode int32, result models.PatchResult, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleModifysubscriptionDataSubscription has not been implemented")
	return
}
func (p *Producer) DR_HandleQuerySubscriptionDataSubscriptions(subsId string) (successCode int32, result map[string]interface{}, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQuerySubscriptionDataSubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleRemovesubscriptionDataSubscriptions(subsId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleRemovesubscriptionDataSubscriptions has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryTraceData(ueId string, servingPlmnId string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.TraceData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryTraceData has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateOrReplaceUEPolicySet(ueId string, body models.UePolicySet) (successCode int32, result models.UePolicySet, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateOrReplaceUEPolicySet has not been implemented")
	return
}
func (p *Producer) DR_HandleReadUEPolicySet(ueId string, suppFeat string) (successCode int32, result models.UePolicySet, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReadUEPolicySet has not been implemented")
	return
}
func (p *Producer) DR_HandleUpdateUEPolicySet(ueId string, body models.UePolicySetPatch) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleUpdateUEPolicySet has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryUeLocation(ueId string, supportedFeatures string) (successCode int32, result models.LocationInfo, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryUeLocation has not been implemented")
	return
}
func (p *Producer) DR_HandleCreateUsageMonitoringResource(ueId string, usageMonId string, body models.UsageMonData) (successCode int32, result models.UsageMonData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleCreateUsageMonitoringResource has not been implemented")
	return
}
func (p *Producer) DR_HandleDeleteUsageMonitoringInformation(ueId string, usageMonId string) (successCode int32, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleDeleteUsageMonitoringInformation has not been implemented")
	return
}
func (p *Producer) DR_HandleReadUsageMonitoringInformation(ueId string, usageMonId string, suppFeat string) (successCode int32, result models.UsageMonData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleReadUsageMonitoringInformation has not been implemented")
	return
}
func (p *Producer) DR_HandleQueryV2xData(ueId string, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.V2xSubscriptionData, err *sbi.ApiError) {
	//TODO: to be implemented
	panic("DR_HandleQueryV2xData has not been implemented")
	return
}
