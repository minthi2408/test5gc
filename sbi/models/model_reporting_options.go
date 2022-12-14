/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

import (
	"time"
)

type ReportingOptions struct {
	ReportMode EventReportMode `json:"reportMode,omitempty"`

	MaxNumOfReports int32 `json:"maxNumOfReports,omitempty"`

	Expiry time.Time `json:"expiry,omitempty"`

	SamplingRatio int32 `json:"samplingRatio,omitempty"`

	GuardTime int32 `json:"guardTime,omitempty"`

	ReportPeriod int32 `json:"reportPeriod,omitempty"`
}
