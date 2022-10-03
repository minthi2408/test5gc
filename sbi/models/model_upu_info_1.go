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

type UpuInfo1 struct {

	UpuDataList []UpuData2 `json:"upuDataList"`

	UpuRegInd bool `json:"upuRegInd"`

	UpuAckInd bool `json:"upuAckInd"`

	UpuMacIausf string `json:"upuMacIausf,omitempty"`

	CounterUpu string `json:"counterUpu,omitempty"`

	ProvisioningTime time.Time `json:"provisioningTime"`
}
