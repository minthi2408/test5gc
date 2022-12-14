/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type UnrelatedClass1 struct {
	DefaultUnrelatedClass DefaultUnrelatedClass1 `json:"defaultUnrelatedClass"`

	ExternalUnrelatedClass ExternalUnrelatedClass `json:"externalUnrelatedClass,omitempty"`

	ServiceTypeUnrelatedClasses []ServiceTypeUnrelatedClass1 `json:"serviceTypeUnrelatedClasses,omitempty"`
}
