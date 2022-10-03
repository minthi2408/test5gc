/*
Nsmf_PDUSession

SMF PDU Session Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

import (
	"time"
)

type Trigger struct {

	TriggerType TriggerType `json:"triggerType"`

	TriggerCategory TriggerCategory `json:"triggerCategory"`

	TimeLimit int32 `json:"timeLimit,omitempty"`

	VolumeLimit int32 `json:"volumeLimit,omitempty"`

	VolumeLimit64 int32 `json:"volumeLimit64,omitempty"`

	EventLimit int32 `json:"eventLimit,omitempty"`

	MaxNumberOfccc int32 `json:"maxNumberOfccc,omitempty"`

	TariffTimeChange time.Time `json:"tariffTimeChange,omitempty"`
}
