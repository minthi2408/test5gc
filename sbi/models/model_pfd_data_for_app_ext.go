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

type PfdDataForAppExt struct {

	ApplicationId string `json:"applicationId"`

	Pfds []PfdContent `json:"pfds"`

	CachingTime time.Time `json:"cachingTime,omitempty"`

	SuppFeat string `json:"suppFeat,omitempty"`
}
