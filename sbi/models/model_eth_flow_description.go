/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

// EthFlowDescription - Identifies an Ethernet flow
type EthFlowDescription struct {

	DestMacAddr string `json:"destMacAddr,omitempty"`

	EthType string `json:"ethType"`

	// Defines a packet filter of an IP flow.
	FDesc string `json:"fDesc,omitempty"`

	FDir FlowDirection `json:"fDir,omitempty"`

	SourceMacAddr string `json:"sourceMacAddr,omitempty"`

	VlanTags []string `json:"vlanTags,omitempty"`

	SrcMacAddrEnd string `json:"srcMacAddrEnd,omitempty"`

	DestMacAddrEnd string `json:"destMacAddrEnd,omitempty"`
}