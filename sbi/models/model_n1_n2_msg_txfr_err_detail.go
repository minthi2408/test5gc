/*
Namf_Communication

AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type N1N2MsgTxfrErrDetail struct {
	RetryAfter int32 `json:"retryAfter,omitempty"`

	HighestPrioArp Arp `json:"highestPrioArp,omitempty"`

	MaxWaitingTime int32 `json:"maxWaitingTime,omitempty"`
}
