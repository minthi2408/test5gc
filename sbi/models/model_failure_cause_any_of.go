/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type FailureCauseAnyOf string

// List of FailureCauseAnyOf
const (
	FAILURECAUSEANYOF_PCC_RULE_EVENT FailureCauseAnyOf = "PCC_RULE_EVENT"
	FAILURECAUSEANYOF_PCC_QOS_FLOW_EVENT FailureCauseAnyOf = "PCC_QOS_FLOW_EVENT"
	FAILURECAUSEANYOF_RULE_PERMANENT_ERROR FailureCauseAnyOf = "RULE_PERMANENT_ERROR"
	FAILURECAUSEANYOF_RULE_TEMPORARY_ERROR FailureCauseAnyOf = "RULE_TEMPORARY_ERROR"
	FAILURECAUSEANYOF_POL_DEC_ERROR FailureCauseAnyOf = "POL_DEC_ERROR"
)
