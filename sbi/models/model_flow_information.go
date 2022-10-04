/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type FlowInformation struct {

	// Defines a packet filter for an IP flow.
	FlowDescription string `json:"flowDescription,omitempty"`

	EthFlowDescription EthFlowDescription `json:"ethFlowDescription,omitempty"`

	// An identifier of packet filter.
	PackFiltId string `json:"packFiltId,omitempty"`

	// The packet shall be sent to the UE.
	PacketFilterUsage bool `json:"packetFilterUsage,omitempty"`

	// Contains the Ipv4 Type-of-Service and mask field or the Ipv6 Traffic-Class field and mask field.
	TosTrafficClass *string `json:"tosTrafficClass,omitempty"`

	// the security parameter index of the IPSec packet.
	Spi *string `json:"spi,omitempty"`

	// the Ipv6 flow label header field.
	FlowLabel *string `json:"flowLabel,omitempty"`

	FlowDirection FlowDirectionRm `json:"flowDirection,omitempty"`
}
