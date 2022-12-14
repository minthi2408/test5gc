/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

// DnnRouteSelectionDescriptor - Contains the route selector parameters (PDU session types, SSC modes and ATSSS information) per DNN
type DnnRouteSelectionDescriptor struct {
	Dnn string `json:"dnn"`

	SscModes []SscMode `json:"sscModes,omitempty"`

	PduSessTypes []PduSessionType `json:"pduSessTypes,omitempty"`

	// Indicates whether MA PDU session establishment is allowed for this DNN. When set to value true MA PDU session establishment is allowed for this DNN.
	AtsssInfo bool `json:"atsssInfo,omitempty"`
}
