/*
Npcf_UEPolicyControl

UE Policy Control Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type RequestTriggerAnyOf string

// List of RequestTriggerAnyOf
const (
	REQUESTTRIGGERANYOF_LOC_CH            RequestTriggerAnyOf = "LOC_CH"
	REQUESTTRIGGERANYOF_PRA_CH            RequestTriggerAnyOf = "PRA_CH"
	REQUESTTRIGGERANYOF_UE_POLICY         RequestTriggerAnyOf = "UE_POLICY"
	REQUESTTRIGGERANYOF_PLMN_CH           RequestTriggerAnyOf = "PLMN_CH"
	REQUESTTRIGGERANYOF_CON_STATE_CH      RequestTriggerAnyOf = "CON_STATE_CH"
	REQUESTTRIGGERANYOF_GROUP_ID_LIST_CHG RequestTriggerAnyOf = "GROUP_ID_LIST_CHG"
)
