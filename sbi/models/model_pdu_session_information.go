/*
Npcf_EventExposure

PCF Event Exposure Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type PduSessionInformation struct {

	Snssai Snssai `json:"snssai"`

	Dnn string `json:"dnn"`

	UeIpv4 string `json:"ueIpv4,omitempty"`

	UeIpv6 Ipv6Prefix `json:"ueIpv6,omitempty"`

	IpDomain string `json:"ipDomain,omitempty"`

	UeMac string `json:"ueMac,omitempty"`
}
