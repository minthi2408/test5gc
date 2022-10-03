/*
Nsmf_PDUSession

SMF PDU Session Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type NonDynamic5Qi struct {

	PriorityLevel int32 `json:"priorityLevel,omitempty"`

	AverWindow int32 `json:"averWindow,omitempty"`

	MaxDataBurstVol int32 `json:"maxDataBurstVol,omitempty"`

	ExtMaxDataBurstVol int32 `json:"extMaxDataBurstVol,omitempty"`

	CnPacketDelayBudgetDl int32 `json:"cnPacketDelayBudgetDl,omitempty"`

	CnPacketDelayBudgetUl int32 `json:"cnPacketDelayBudgetUl,omitempty"`
}
