/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type AccumulatedUsage struct {

	// Unsigned integer identifying a period of time in units of seconds.
	Duration int32 `json:"duration,omitempty"`

	// Unsigned integer identifying a volume in units of bytes.
	TotalVolume int64 `json:"totalVolume,omitempty"`

	// Unsigned integer identifying a volume in units of bytes.
	DownlinkVolume int64 `json:"downlinkVolume,omitempty"`

	// Unsigned integer identifying a volume in units of bytes.
	UplinkVolume int64 `json:"uplinkVolume,omitempty"`
}
