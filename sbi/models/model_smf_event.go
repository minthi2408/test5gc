/*
Nsmf_EventExposure

Session Management Event Exposure Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type SmfEvent string

// List of SmfEvent
const (
	SMFEVENT_AC_TY_CH    SmfEvent = "AC_TY_CH"
	SMFEVENT_UP_PATH_CH  SmfEvent = "UP_PATH_CH"
	SMFEVENT_PDU_SES_REL SmfEvent = "PDU_SES_REL"
	SMFEVENT_PLMN_CH     SmfEvent = "PLMN_CH"
	SMFEVENT_UE_IP_CH    SmfEvent = "UE_IP_CH"
	SMFEVENT_DDDS        SmfEvent = "DDDS"
	SMFEVENT_COMM_FAIL   SmfEvent = "COMM_FAIL"
	SMFEVENT_PDU_SES_EST SmfEvent = "PDU_SES_EST"
	SMFEVENT_QFI_ALLOC   SmfEvent = "QFI_ALLOC"
	SMFEVENT_QOS_MON     SmfEvent = "QOS_MON"
)
