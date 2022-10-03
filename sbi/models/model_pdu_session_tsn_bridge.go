/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

// PduSessionTsnBridge - Contains the new 5GS Bridge information and may contain the DS-TT port and/or NW-TT port management information.
type PduSessionTsnBridge struct {

	TsnBridgeInfo TsnBridgeInfo `json:"tsnBridgeInfo"`

	TsnBridgeManCont BridgeManagementContainer `json:"tsnBridgeManCont,omitempty"`

	TsnPortManContDstt PortManagementContainer `json:"tsnPortManContDstt,omitempty"`

	TsnPortManContNwtts []PortManagementContainer `json:"tsnPortManContNwtts,omitempty"`
}
