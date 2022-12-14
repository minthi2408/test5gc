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

type GeraLocation struct {
	LocationNumber string `json:"locationNumber,omitempty"`

	Cgi CellGlobalId `json:"cgi,omitempty"`

	Rai RoutingAreaId `json:"rai,omitempty"`

	Sai ServiceAreaId `json:"sai,omitempty"`

	Lai LocationAreaId `json:"lai,omitempty"`

	VlrNumber string `json:"vlrNumber,omitempty"`

	MscNumber string `json:"mscNumber,omitempty"`

	AgeOfLocationInformation int32 `json:"ageOfLocationInformation,omitempty"`

	UeLocationTimestamp time.Time `json:"ueLocationTimestamp,omitempty"`

	GeographicalInformation string `json:"geographicalInformation,omitempty"`

	GeodeticInformation string `json:"geodeticInformation,omitempty"`
}
