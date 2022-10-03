/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type FlowStatusAnyOf string

// List of FlowStatusAnyOf
const (
	FLOWSTATUSANYOF_ENABLED_UPLINK FlowStatusAnyOf = "ENABLED-UPLINK"
	FLOWSTATUSANYOF_ENABLED_DOWNLINK FlowStatusAnyOf = "ENABLED-DOWNLINK"
	FLOWSTATUSANYOF_ENABLED FlowStatusAnyOf = "ENABLED"
	FLOWSTATUSANYOF_DISABLED FlowStatusAnyOf = "DISABLED"
	FLOWSTATUSANYOF_REMOVED FlowStatusAnyOf = "REMOVED"
)
