/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

// MediaSubComponentRm - This data type is defined in the same way as the MediaSubComponent data type, but with the OpenAPI nullable property set to true. Removable attributes marBwDl and marBwUl are defined with the corresponding removable data type.
type MediaSubComponentRm struct {

	AfSigProtocol AfSigProtocol `json:"afSigProtocol,omitempty"`

	EthfDescs *[]EthFlowDescription `json:"ethfDescs,omitempty"`

	FNum int32 `json:"fNum"`

	FDescs *[]string `json:"fDescs,omitempty"`

	FStatus FlowStatus `json:"fStatus,omitempty"`

	MarBwDl *string `json:"marBwDl,omitempty"`

	MarBwUl *string `json:"marBwUl,omitempty"`

	// this data type is defined in the same way as the TosTrafficClass data type, but with the OpenAPI nullable property set to true
	TosTrCl *string `json:"tosTrCl,omitempty"`

	FlowUsage FlowUsage `json:"flowUsage,omitempty"`
}
