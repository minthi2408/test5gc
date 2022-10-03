/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type SmsManagementSubscriptionData1 struct {

	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	MtSmsSubscribed bool `json:"mtSmsSubscribed,omitempty"`

	MtSmsBarringAll bool `json:"mtSmsBarringAll,omitempty"`

	MtSmsBarringRoaming bool `json:"mtSmsBarringRoaming,omitempty"`

	MoSmsSubscribed bool `json:"moSmsSubscribed,omitempty"`

	MoSmsBarringAll bool `json:"moSmsBarringAll,omitempty"`

	MoSmsBarringRoaming bool `json:"moSmsBarringRoaming,omitempty"`

	SharedSmsMngDataIds []string `json:"sharedSmsMngDataIds,omitempty"`

	TraceData *TraceData `json:"traceData,omitempty"`
}
