/*
Nsmf_EventExposure

Session Management Event Exposure Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type EventSubscription struct {
	Event SmfEvent `json:"event"`

	DnaiChgType DnaiChangeType `json:"dnaiChgType,omitempty"`

	DddTraDescriptors []DddTrafficDescriptor `json:"dddTraDescriptors,omitempty"`

	DddStati []DlDataDeliveryStatus `json:"dddStati,omitempty"`

	AppIds []string `json:"appIds,omitempty"`
}
