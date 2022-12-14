/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

// MediaComponent - Identifies a media component.
type MediaComponent struct {

	// Contains an AF application identifier.
	AfAppId string `json:"afAppId,omitempty"`

	AfRoutReq AfRoutingRequirement `json:"afRoutReq,omitempty"`

	QosReference string `json:"qosReference,omitempty"`

	DisUeNotif bool `json:"disUeNotif,omitempty"`

	AltSerReqs []string `json:"altSerReqs,omitempty"`

	// Represents the content version of some content.
	ContVer int32 `json:"contVer,omitempty"`

	Codecs []string `json:"codecs,omitempty"`

	DesMaxLatency float32 `json:"desMaxLatency,omitempty"`

	DesMaxLoss float32 `json:"desMaxLoss,omitempty"`

	FlusId string `json:"flusId,omitempty"`

	FStatus FlowStatus `json:"fStatus,omitempty"`

	MarBwDl string `json:"marBwDl,omitempty"`

	MarBwUl string `json:"marBwUl,omitempty"`

	MaxPacketLossRateDl *int32 `json:"maxPacketLossRateDl,omitempty"`

	MaxPacketLossRateUl *int32 `json:"maxPacketLossRateUl,omitempty"`

	MaxSuppBwDl string `json:"maxSuppBwDl,omitempty"`

	MaxSuppBwUl string `json:"maxSuppBwUl,omitempty"`

	MedCompN int32 `json:"medCompN"`

	MedSubComps map[string]MediaSubComponent `json:"medSubComps,omitempty"`

	MedType MediaType `json:"medType,omitempty"`

	MinDesBwDl string `json:"minDesBwDl,omitempty"`

	MinDesBwUl string `json:"minDesBwUl,omitempty"`

	MirBwDl string `json:"mirBwDl,omitempty"`

	MirBwUl string `json:"mirBwUl,omitempty"`

	PreemptCap PreemptionCapability `json:"preemptCap,omitempty"`

	PreemptVuln PreemptionVulnerability `json:"preemptVuln,omitempty"`

	PrioSharingInd PrioritySharingIndicator `json:"prioSharingInd,omitempty"`

	ResPrio ReservPriority `json:"resPrio,omitempty"`

	RrBw string `json:"rrBw,omitempty"`

	RsBw string `json:"rsBw,omitempty"`

	SharingKeyDl int32 `json:"sharingKeyDl,omitempty"`

	SharingKeyUl int32 `json:"sharingKeyUl,omitempty"`

	TsnQos TsnQosContainer `json:"tsnQos,omitempty"`

	TscaiInputDl *TscaiInputContainer `json:"tscaiInputDl,omitempty"`

	TscaiInputUl *TscaiInputContainer `json:"tscaiInputUl,omitempty"`
}
