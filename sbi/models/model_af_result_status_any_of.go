/*
Nsmf_EventExposure

Session Management Event Exposure Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type AfResultStatusAnyOf string

// List of AfResultStatusAnyOf
const (
	AFRESULTSTATUSANYOF_SUCCESS AfResultStatusAnyOf = "SUCCESS"
	AFRESULTSTATUSANYOF_TEMPORARY_CONGESTION AfResultStatusAnyOf = "TEMPORARY_CONGESTION"
	AFRESULTSTATUSANYOF_RELOC_NO_ALLOWED AfResultStatusAnyOf = "RELOC_NO_ALLOWED"
	AFRESULTSTATUSANYOF_OTHER AfResultStatusAnyOf = "OTHER"
)
