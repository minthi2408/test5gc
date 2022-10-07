/*
Namf_Location

AMF Location Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type LdrType string

// List of LdrType
const (
	LDRTYPEANYOF_UE_AVAILABLE       LdrType = "UE_AVAILABLE"
	LDRTYPEANYOF_PERIODIC           LdrType = "PERIODIC"
	LDRTYPEANYOF_ENTERING_INTO_AREA LdrType = "ENTERING_INTO_AREA"
	LDRTYPEANYOF_LEAVING_FROM_AREA  LdrType = "LEAVING_FROM_AREA"
	LDRTYPEANYOF_BEING_INSIDE_AREA  LdrType = "BEING_INSIDE_AREA"
	LDRTYPEANYOF_MOTION             LdrType = "MOTION"
)
