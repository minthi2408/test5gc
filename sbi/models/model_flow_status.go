/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type FlowStatus string

// List of FlowStatusAnyOf
const (
	FLOWSTATUSANYOF_ENABLED_UPLINK   FlowStatus = "ENABLED-UPLINK"
	FLOWSTATUSANYOF_ENABLED_DOWNLINK FlowStatus = "ENABLED-DOWNLINK"
	FLOWSTATUSANYOF_ENABLED          FlowStatus = "ENABLED"
	FLOWSTATUSANYOF_DISABLED         FlowStatus = "DISABLED"
	FLOWSTATUSANYOF_REMOVED          FlowStatus = "REMOVED"
)
