/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type AfEvent string

// List of AfEventAnyOf
const (
	AFEVENTANYOF_ACCESS_TYPE_CHANGE              AfEvent = "ACCESS_TYPE_CHANGE"
	AFEVENTANYOF_ANI_REPORT                      AfEvent = "ANI_REPORT"
	AFEVENTANYOF_CHARGING_CORRELATION            AfEvent = "CHARGING_CORRELATION"
	AFEVENTANYOF_EPS_FALLBACK                    AfEvent = "EPS_FALLBACK"
	AFEVENTANYOF_FAILED_RESOURCES_ALLOCATION     AfEvent = "FAILED_RESOURCES_ALLOCATION"
	AFEVENTANYOF_OUT_OF_CREDIT                   AfEvent = "OUT_OF_CREDIT"
	AFEVENTANYOF_PLMN_CHG                        AfEvent = "PLMN_CHG"
	AFEVENTANYOF_QOS_MONITORING                  AfEvent = "QOS_MONITORING"
	AFEVENTANYOF_QOS_NOTIF                       AfEvent = "QOS_NOTIF"
	AFEVENTANYOF_RAN_NAS_CAUSE                   AfEvent = "RAN_NAS_CAUSE"
	AFEVENTANYOF_REALLOCATION_OF_CREDIT          AfEvent = "REALLOCATION_OF_CREDIT"
	AFEVENTANYOF_SUCCESSFUL_RESOURCES_ALLOCATION AfEvent = "SUCCESSFUL_RESOURCES_ALLOCATION"
	AFEVENTANYOF_TSN_BRIDGE_INFO                 AfEvent = "TSN_BRIDGE_INFO"
	AFEVENTANYOF_USAGE_REPORT                    AfEvent = "USAGE_REPORT"
)
