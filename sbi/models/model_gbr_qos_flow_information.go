/*
Nsmf_PDUSession

SMF PDU Session Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type GbrQosFlowInformation struct {
	MaxFbrDl string `json:"maxFbrDl"`

	MaxFbrUl string `json:"maxFbrUl"`

	GuaFbrDl string `json:"guaFbrDl"`

	GuaFbrUl string `json:"guaFbrUl"`

	NotifControl NotificationControl `json:"notifControl,omitempty"`

	MaxPacketLossRateDl int32 `json:"maxPacketLossRateDl,omitempty"`

	MaxPacketLossRateUl int32 `json:"maxPacketLossRateUl,omitempty"`

	AlternativeQosProfileList []AlternativeQosProfile `json:"alternativeQosProfileList,omitempty"`
}
