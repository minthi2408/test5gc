/*
Namf_Communication

AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type PolicyReqTrigger string

// List of PolicyReqTrigger
const (
	POLICYREQTRIGGERANYOF_LOCATION_CHANGE      PolicyReqTrigger = "LOCATION_CHANGE"
	POLICYREQTRIGGERANYOF_PRA_CHANGE           PolicyReqTrigger = "PRA_CHANGE"
	POLICYREQTRIGGERANYOF_ALLOWED_NSSAI_CHANGE PolicyReqTrigger = "ALLOWED_NSSAI_CHANGE"
	POLICYREQTRIGGERANYOF_PLMN_CHANGE          PolicyReqTrigger = "PLMN_CHANGE"
	POLICYREQTRIGGERANYOF_CON_STATE_CHANGE     PolicyReqTrigger = "CON_STATE_CHANGE"
	POLICYREQTRIGGERANYOF_SMF_SELECT_CHANGE    PolicyReqTrigger = "SMF_SELECT_CHANGE"
	POLICYREQTRIGGERANYOF_ACCESS_TYPE_CHANGE   PolicyReqTrigger = "ACCESS_TYPE_CHANGE"
)
