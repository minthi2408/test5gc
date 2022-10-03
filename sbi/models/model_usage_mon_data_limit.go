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

// UsageMonDataLimit - Contains usage monitoring control data for a subscriber.
type UsageMonDataLimit struct {

	LimitId string `json:"limitId"`

	Scopes map[string]UsageMonDataScope `json:"scopes,omitempty"`

	UmLevel UsageMonLevel `json:"umLevel,omitempty"`

	StartDate time.Time `json:"startDate,omitempty"`

	EndDate time.Time `json:"endDate,omitempty"`

	UsageLimit UsageThreshold `json:"usageLimit,omitempty"`

	ResetPeriod TimePeriod `json:"resetPeriod,omitempty"`
}
