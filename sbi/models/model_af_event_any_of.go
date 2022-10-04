/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type AfEventAnyOf string

// List of AfEventAnyOf
const (
	AFEVENTANYOF_ACCESS_TYPE_CHANGE AfEventAnyOf = "ACCESS_TYPE_CHANGE"
	AFEVENTANYOF_ANI_REPORT AfEventAnyOf = "ANI_REPORT"
	AFEVENTANYOF_CHARGING_CORRELATION AfEventAnyOf = "CHARGING_CORRELATION"
	AFEVENTANYOF_EPS_FALLBACK AfEventAnyOf = "EPS_FALLBACK"
	AFEVENTANYOF_FAILED_RESOURCES_ALLOCATION AfEventAnyOf = "FAILED_RESOURCES_ALLOCATION"
	AFEVENTANYOF_OUT_OF_CREDIT AfEventAnyOf = "OUT_OF_CREDIT"
	AFEVENTANYOF_PLMN_CHG AfEventAnyOf = "PLMN_CHG"
	AFEVENTANYOF_QOS_MONITORING AfEventAnyOf = "QOS_MONITORING"
	AFEVENTANYOF_QOS_NOTIF AfEventAnyOf = "QOS_NOTIF"
	AFEVENTANYOF_RAN_NAS_CAUSE AfEventAnyOf = "RAN_NAS_CAUSE"
	AFEVENTANYOF_REALLOCATION_OF_CREDIT AfEventAnyOf = "REALLOCATION_OF_CREDIT"
	AFEVENTANYOF_SUCCESSFUL_RESOURCES_ALLOCATION AfEventAnyOf = "SUCCESSFUL_RESOURCES_ALLOCATION"
	AFEVENTANYOF_TSN_BRIDGE_INFO AfEventAnyOf = "TSN_BRIDGE_INFO"
	AFEVENTANYOF_USAGE_REPORT AfEventAnyOf = "USAGE_REPORT"
)