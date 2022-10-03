/*
Npcf_BDTPolicyControl Service API

PCF BDT Policy Control Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

// Notification - Describes a BDT notification.
type Notification struct {

	// string identifying a BDT Reference ID as defined in subclause 5.3.3 of 3GPP TS 29.154.
	BdtRefId string `json:"bdtRefId"`

	// Contains a list of the candidate transfer policies from which the AF may select a new transfer policy due to a network performance is below the criteria set by the operator.
	CandPolicies []TransferPolicy `json:"candPolicies,omitempty"`

	NwAreaInfo NetworkAreaInfo `json:"nwAreaInfo,omitempty"`

	TimeWindow TimeWindow `json:"timeWindow,omitempty"`
}
