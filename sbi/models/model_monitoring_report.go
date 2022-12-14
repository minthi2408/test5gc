/*
Nudm_EE

Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

import (
	"time"
)

type MonitoringReport struct {
	ReferenceId int32 `json:"referenceId"`

	EventType EventType `json:"eventType"`

	Report Report `json:"report,omitempty"`

	ReachabilityForSmsReport ReachabilityForSmsReport `json:"reachabilityForSmsReport,omitempty"`

	Gpsi string `json:"gpsi,omitempty"`

	TimeStamp time.Time `json:"timeStamp"`
}
