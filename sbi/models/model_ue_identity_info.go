/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

// UeIdentityInfo - Represents 5GS-Level UE identities.
type UeIdentityInfo struct {

	Gpsi string `json:"gpsi,omitempty"`

	Pei string `json:"pei,omitempty"`

	Supi string `json:"supi,omitempty"`
}
