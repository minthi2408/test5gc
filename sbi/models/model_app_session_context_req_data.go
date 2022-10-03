/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

// AppSessionContextReqData - Identifies the service requirements of an Individual Application Session Context.
type AppSessionContextReqData struct {

	// Contains an AF application identifier.
	AfAppId string `json:"afAppId,omitempty"`

	AfChargId string `json:"afChargId,omitempty"`

	AfReqData AfRequestedData `json:"afReqData,omitempty"`

	AfRoutReq AfRoutingRequirement `json:"afRoutReq,omitempty"`

	// Contains an identity of an application service provider.
	AspId string `json:"aspId,omitempty"`

	// string identifying a BDT Reference ID as defined in subclause 5.3.3 of 3GPP TS 29.154.
	BdtRefId string `json:"bdtRefId,omitempty"`

	Dnn string `json:"dnn,omitempty"`

	EvSubsc EventsSubscReqData `json:"evSubsc,omitempty"`

	// indication of MCPTT service request
	McpttId string `json:"mcpttId,omitempty"`

	// indication of MCVideo service request
	McVideoId string `json:"mcVideoId,omitempty"`

	MedComponents map[string]MediaComponent `json:"medComponents,omitempty"`

	IpDomain string `json:"ipDomain,omitempty"`

	// indication of MPS service request
	MpsId string `json:"mpsId,omitempty"`

	// indication of MCS service request
	McsId string `json:"mcsId,omitempty"`

	PreemptControlInfo PreemptionControlInformation `json:"preemptControlInfo,omitempty"`

	ResPrio ReservPriority `json:"resPrio,omitempty"`

	ServInfStatus ServiceInfoStatus `json:"servInfStatus,omitempty"`

	NotifUri string `json:"notifUri"`

	// Contains values of the service URN and may include subservices.
	ServUrn string `json:"servUrn,omitempty"`

	SliceInfo Snssai `json:"sliceInfo,omitempty"`

	// Contains an identity of a sponsor.
	SponId string `json:"sponId,omitempty"`

	SponStatus SponsoringStatus `json:"sponStatus,omitempty"`

	Supi string `json:"supi,omitempty"`

	Gpsi string `json:"gpsi,omitempty"`

	SuppFeat string `json:"suppFeat"`

	UeIpv4 string `json:"ueIpv4,omitempty"`

	UeIpv6 Ipv6Addr `json:"ueIpv6,omitempty"`

	UeMac string `json:"ueMac,omitempty"`

	TsnBridgeManCont BridgeManagementContainer `json:"tsnBridgeManCont,omitempty"`

	TsnPortManContDstt PortManagementContainer `json:"tsnPortManContDstt,omitempty"`

	TsnPortManContNwtts []PortManagementContainer `json:"tsnPortManContNwtts,omitempty"`
}
