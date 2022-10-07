/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type ReportingTrigger string

// List of ReportingTriggerAnyOf
const (
	REPORTINGTRIGGERANYOF_PERIODICAL             ReportingTrigger = "PERIODICAL"
	REPORTINGTRIGGERANYOF_EVENT_A2               ReportingTrigger = "EVENT_A2"
	REPORTINGTRIGGERANYOF_EVENT_A2_PERIODIC      ReportingTrigger = "EVENT_A2_PERIODIC"
	REPORTINGTRIGGERANYOF_ALL_RRM_EVENT_TRIGGERS ReportingTrigger = "ALL_RRM_EVENT_TRIGGERS"
)
