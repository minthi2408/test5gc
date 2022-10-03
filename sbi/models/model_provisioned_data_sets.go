/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type ProvisionedDataSets struct {

	AmData AccessAndMobilitySubscriptionData `json:"amData,omitempty"`

	SmfSelData SmfSelectionSubscriptionData `json:"smfSelData,omitempty"`

	SmsSubsData SmsSubscriptionData `json:"smsSubsData,omitempty"`

	SmData []SessionManagementSubscriptionData `json:"smData,omitempty"`

	TraceData *TraceData `json:"traceData,omitempty"`

	SmsMngData SmsManagementSubscriptionData `json:"smsMngData,omitempty"`

	LcsPrivacyData LcsPrivacyData `json:"lcsPrivacyData,omitempty"`

	LcsMoData LcsMoData `json:"lcsMoData,omitempty"`

	LcsBcaData LcsBroadcastAssistanceTypesData `json:"lcsBcaData,omitempty"`

	V2xData V2xSubscriptionData `json:"v2xData,omitempty"`
}
