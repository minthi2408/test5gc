/*
Nsmf_PDUSession

SMF PDU Session Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type PsaInformation struct {
	PsaInd PsaIndication `json:"psaInd,omitempty"`

	DnaiList []string `json:"dnaiList,omitempty"`

	UeIpv6Prefix Ipv6Prefix `json:"ueIpv6Prefix,omitempty"`

	PsaUpfId string `json:"psaUpfId,omitempty"`
}
