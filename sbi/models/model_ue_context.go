/*
Namf_Communication

AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

import (
	"time"
)

type UeContext struct {

	Supi string `json:"supi,omitempty"`

	SupiUnauthInd bool `json:"supiUnauthInd,omitempty"`

	GpsiList []string `json:"gpsiList,omitempty"`

	Pei string `json:"pei,omitempty"`

	UdmGroupId string `json:"udmGroupId,omitempty"`

	AusfGroupId string `json:"ausfGroupId,omitempty"`

	PcfGroupId string `json:"pcfGroupId,omitempty"`

	RoutingIndicator string `json:"routingIndicator,omitempty"`

	GroupList []string `json:"groupList,omitempty"`

	DrxParameter string `json:"drxParameter,omitempty"`

	SubRfsp int32 `json:"subRfsp,omitempty"`

	UsedRfsp int32 `json:"usedRfsp,omitempty"`

	SubUeAmbr Ambr `json:"subUeAmbr,omitempty"`

	SmsfId string `json:"smsfId,omitempty"`

	SeafData SeafData `json:"seafData,omitempty"`

	Var5gMmCapability string `json:"5gMmCapability,omitempty"`

	PcfId string `json:"pcfId,omitempty"`

	PcfSetId string `json:"pcfSetId,omitempty"`

	PcfAmpServiceSetId string `json:"pcfAmpServiceSetId,omitempty"`

	PcfUepServiceSetId string `json:"pcfUepServiceSetId,omitempty"`

	PcfBinding SbiBindingLevel `json:"pcfBinding,omitempty"`

	PcfAmPolicyUri string `json:"pcfAmPolicyUri,omitempty"`

	AmPolicyReqTriggerList []PolicyReqTrigger `json:"amPolicyReqTriggerList,omitempty"`

	PcfUePolicyUri string `json:"pcfUePolicyUri,omitempty"`

	UePolicyReqTriggerList []PolicyReqTrigger `json:"uePolicyReqTriggerList,omitempty"`

	HpcfId string `json:"hpcfId,omitempty"`

	HpcfSetId string `json:"hpcfSetId,omitempty"`

	RestrictedRatList []RatType `json:"restrictedRatList,omitempty"`

	ForbiddenAreaList []Area `json:"forbiddenAreaList,omitempty"`

	ServiceAreaRestriction ServiceAreaRestriction `json:"serviceAreaRestriction,omitempty"`

	RestrictedCoreNwTypeList []CoreNetworkType `json:"restrictedCoreNwTypeList,omitempty"`

	EventSubscriptionList []ExtAmfEventSubscription `json:"eventSubscriptionList,omitempty"`

	MmContextList []MmContext `json:"mmContextList,omitempty"`

	SessionContextList []PduSessionContext `json:"sessionContextList,omitempty"`

	TraceData *TraceData `json:"traceData,omitempty"`

	ServiceGapExpiryTime time.Time `json:"serviceGapExpiryTime,omitempty"`

	StnSr string `json:"stnSr,omitempty"`

	CMsisdn string `json:"cMsisdn,omitempty"`

	MsClassmark2 string `json:"msClassmark2,omitempty"`

	SupportedCodecList []string `json:"supportedCodecList,omitempty"`

	SmallDataRateStatusInfos []SmallDataRateStatusInfo `json:"smallDataRateStatusInfos,omitempty"`

	RestrictedPrimaryRatList []RatType `json:"restrictedPrimaryRatList,omitempty"`

	RestrictedSecondaryRatList []RatType `json:"restrictedSecondaryRatList,omitempty"`

	V2xContext V2xContext `json:"v2xContext,omitempty"`

	LteCatMInd bool `json:"lteCatMInd,omitempty"`

	MoExpDataCounter MoExpDataCounter `json:"moExpDataCounter,omitempty"`

	CagData CagData `json:"cagData,omitempty"`

	ManagementMdtInd bool `json:"managementMdtInd,omitempty"`

	ImmediateMdtConf ImmediateMdtConf `json:"immediateMdtConf,omitempty"`

	EcRestrictionDataWb EcRestrictionDataWb `json:"ecRestrictionDataWb,omitempty"`

	EcRestrictionDataNb bool `json:"ecRestrictionDataNb,omitempty"`

	IabOperationAllowed bool `json:"iabOperationAllowed,omitempty"`

	UsedServiceAreaRestriction ServiceAreaRestriction `json:"usedServiceAreaRestriction,omitempty"`

	// A map(list of key-value pairs) where praId serves as key.
	PraInAmPolicy map[string]PresenceInfo `json:"praInAmPolicy,omitempty"`

	// A map(list of key-value pairs) where praId serves as key.
	PraInUePolicy map[string]PresenceInfo `json:"praInUePolicy,omitempty"`

	UpdpSubscriptionData UpdpSubscriptionData `json:"updpSubscriptionData,omitempty"`

	SmfSelInfo *SmfSelectionData `json:"smfSelInfo,omitempty"`

	PcfAmpBindingInfo string `json:"pcfAmpBindingInfo,omitempty"`

	PcfUepBindingInfo string `json:"pcfUepBindingInfo,omitempty"`
}