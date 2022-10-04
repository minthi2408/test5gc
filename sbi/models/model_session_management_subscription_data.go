/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type SessionManagementSubscriptionData struct {
	SingleNssai Snssai `json:"singleNssai"`

	// A map (list of key-value pairs where Dnn, or optionally the Wildcard DNN, serves as key) of DnnConfigurations
	DnnConfigurations map[string]DnnConfiguration `json:"dnnConfigurations,omitempty"`

	InternalGroupIds []string `json:"internalGroupIds,omitempty"`

	SharedVnGroupDataIds map[string]string `json:"sharedVnGroupDataIds,omitempty"`

	SharedDnnConfigurationsId string `json:"sharedDnnConfigurationsId,omitempty"`

	OdbPacketServices OdbPacketServices `json:"odbPacketServices,omitempty"`

	TraceData *TraceData1 `json:"traceData,omitempty"`

	SharedTraceDataId string `json:"sharedTraceDataId,omitempty"`

	ExpectedUeBehavioursList map[string]ExpectedUeBehaviourData `json:"expectedUeBehavioursList,omitempty"`

	SuggestedPacketNumDlList map[string]SuggestedPacketNumDl `json:"suggestedPacketNumDlList,omitempty"`

	Var3gppChargingCharacteristics string `json:"3gppChargingCharacteristics,omitempty"`
}
