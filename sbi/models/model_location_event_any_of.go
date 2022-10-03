/*
Namf_Location

AMF Location Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type LocationEventAnyOf string

// List of LocationEventAnyOf
const (
	LOCATIONEVENTANYOF_EMERGENCY_CALL_ORIGINATION LocationEventAnyOf = "EMERGENCY_CALL_ORIGINATION"
	LOCATIONEVENTANYOF_EMERGENCY_CALL_RELEASE LocationEventAnyOf = "EMERGENCY_CALL_RELEASE"
	LOCATIONEVENTANYOF_EMERGENCY_CALL_HANDOVER LocationEventAnyOf = "EMERGENCY_CALL_HANDOVER"
	LOCATIONEVENTANYOF_ACTIVATION_OF_DEFERRED_LOCATION LocationEventAnyOf = "ACTIVATION_OF_DEFERRED_LOCATION"
	LOCATIONEVENTANYOF_UE_MOBILITY_FOR_DEFERRED_LOCATION LocationEventAnyOf = "UE_MOBILITY_FOR_DEFERRED_LOCATION"
	LOCATIONEVENTANYOF_CANCELLATION_OF_DEFERRED_LOCATION LocationEventAnyOf = "CANCELLATION_OF_DEFERRED_LOCATION"
)
