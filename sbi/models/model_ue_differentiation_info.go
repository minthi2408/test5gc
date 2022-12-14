/*
Namf_Communication

AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

import (
	"time"
)

type UeDifferentiationInfo struct {
	PeriodicComInd PeriodicCommunicationIndicator `json:"periodicComInd,omitempty"`

	PeriodicTime int32 `json:"periodicTime,omitempty"`

	ScheduledComTime ScheduledCommunicationTime `json:"scheduledComTime,omitempty"`

	StationaryInd StationaryIndication `json:"stationaryInd,omitempty"`

	TrafficProfile TrafficProfile `json:"trafficProfile,omitempty"`

	BatteryInd BatteryIndication `json:"batteryInd,omitempty"`

	ValidityTime time.Time `json:"validityTime,omitempty"`
}
