/*
Namf_Location

AMF Location Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type HorizontalWithVerticalVelocityAndUncertainty struct {

	HSpeed float32 `json:"hSpeed"`

	Bearing int32 `json:"bearing"`

	VSpeed float32 `json:"vSpeed"`

	VDirection VerticalDirection `json:"vDirection"`

	HUncertainty float32 `json:"hUncertainty"`

	VUncertainty float32 `json:"vUncertainty"`
}
