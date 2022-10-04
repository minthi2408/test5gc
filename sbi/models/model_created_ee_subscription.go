/*
Nudm_EE

Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type CreatedEeSubscription struct {
	EeSubscription EeSubscription `json:"eeSubscription"`

	NumberOfUes int32 `json:"numberOfUes,omitempty"`

	EventReports []MonitoringReport `json:"eventReports,omitempty"`

	EpcStatusInd bool `json:"epcStatusInd,omitempty"`
}
