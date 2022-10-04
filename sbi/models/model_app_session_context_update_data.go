/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

// AppSessionContextUpdateData - Identifies the modifications to the \"ascReqData\" property of an Individual Application Session Context which may include the modifications to the sub-resource Events Subscription.
type AppSessionContextUpdateData struct {

	// Contains an AF application identifier.
	AfAppId string `json:"afAppId,omitempty"`

	AfRoutReq *AfRoutingRequirementRm `json:"afRoutReq,omitempty"`

	// Contains an identity of an application service provider.
	AspId string `json:"aspId,omitempty"`

	// string identifying a BDT Reference ID as defined in subclause 5.3.3 of 3GPP TS 29.154.
	BdtRefId string `json:"bdtRefId,omitempty"`

	EvSubsc *EventsSubscReqDataRm `json:"evSubsc,omitempty"`

	// indication of MCPTT service request
	McpttId string `json:"mcpttId,omitempty"`

	// indication of modification of MCVideo service
	McVideoId string `json:"mcVideoId,omitempty"`

	MedComponents map[string]MediaComponentRm `json:"medComponents,omitempty"`

	// indication of MPS service request
	MpsId string `json:"mpsId,omitempty"`

	// indication of MCS service request
	McsId string `json:"mcsId,omitempty"`

	PreemptControlInfo PreemptionControlInformationRm `json:"preemptControlInfo,omitempty"`

	ResPrio ReservPriority `json:"resPrio,omitempty"`

	ServInfStatus ServiceInfoStatus `json:"servInfStatus,omitempty"`

	SipForkInd SipForkingIndication `json:"sipForkInd,omitempty"`

	// Contains an identity of a sponsor.
	SponId string `json:"sponId,omitempty"`

	SponStatus SponsoringStatus `json:"sponStatus,omitempty"`

	TsnBridgeManCont BridgeManagementContainer `json:"tsnBridgeManCont,omitempty"`

	TsnPortManContDstt PortManagementContainer `json:"tsnPortManContDstt,omitempty"`

	TsnPortManContNwtts []PortManagementContainer `json:"tsnPortManContNwtts,omitempty"`
}