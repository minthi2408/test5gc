/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

import (
	"time"
)

// TscaiInputContainer - Indicates TSC Traffic pattern.
type TscaiInputContainer struct {
	Periodicity int32 `json:"periodicity,omitempty"`

	BurstArrivalTime time.Time `json:"burstArrivalTime,omitempty"`
}
