/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type SubscribedDefaultQos struct {
	Var5qi int32 `json:"5qi"`

	Arp Arp `json:"arp"`

	PriorityLevel int32 `json:"priorityLevel,omitempty"`
}
