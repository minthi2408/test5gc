/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type JobTypeAnyOf string

// List of JobTypeAnyOf
const (
	JOBTYPEANYOF_IMMEDIATE_MDT_ONLY JobTypeAnyOf = "IMMEDIATE_MDT_ONLY"
	JOBTYPEANYOF_LOGGED_MDT_ONLY JobTypeAnyOf = "LOGGED_MDT_ONLY"
	JOBTYPEANYOF_TRACE_ONLY JobTypeAnyOf = "TRACE_ONLY"
	JOBTYPEANYOF_IMMEDIATE_MDT_AND_TRACE JobTypeAnyOf = "IMMEDIATE_MDT_AND_TRACE"
	JOBTYPEANYOF_RLF_REPORTS_ONLY JobTypeAnyOf = "RLF_REPORTS_ONLY"
	JOBTYPEANYOF_RCEF_REPORTS_ONLY JobTypeAnyOf = "RCEF_REPORTS_ONLY"
	JOBTYPEANYOF_LOGGED_MBSFN_MDT JobTypeAnyOf = "LOGGED_MBSFN_MDT"
)
