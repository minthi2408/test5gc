/*
Namf_Communication

AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type UeContextRelocateData struct {
	UeContext UeContext `json:"ueContext"`

	TargetId NgRanTargetId `json:"targetId"`

	SourceToTargetData N2InfoContent `json:"sourceToTargetData"`

	ForwardRelocationRequest RefToBinaryData `json:"forwardRelocationRequest"`

	PduSessionList []N2SmInformation `json:"pduSessionList,omitempty"`

	UeRadioCapability N2InfoContent `json:"ueRadioCapability,omitempty"`

	NgapCause NgApCause `json:"ngapCause,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}
