/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type FlowDirection string

// List of FlowDirectionAnyOf
const (
	FLOWDIRECTIONANYOF_DOWNLINK      FlowDirection = "DOWNLINK"
	FLOWDIRECTIONANYOF_UPLINK        FlowDirection = "UPLINK"
	FLOWDIRECTIONANYOF_BIDIRECTIONAL FlowDirection = "BIDIRECTIONAL"
	FLOWDIRECTIONANYOF_UNSPECIFIED   FlowDirection = "UNSPECIFIED"
)
