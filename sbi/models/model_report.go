/*
Nudm_EE

Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type Report struct {

	NewPei string `json:"newPei"`

	Roaming bool `json:"roaming"`

	NewServingPlmn PlmnId `json:"newServingPlmn"`

	NewCnType CnType `json:"newCnType"`

	OldCnType CnType `json:"oldCnType,omitempty"`

	OldCmInfoList []CmInfo `json:"oldCmInfoList,omitempty"`

	NewCmInfoList []CmInfo `json:"newCmInfoList"`

	LossOfConnectReason LossOfConnectivityReason `json:"lossOfConnectReason"`

	Location UserLocation `json:"location"`

	PdnConnStat PdnConnectivityStatus `json:"pdnConnStat"`

	Dnn string `json:"dnn,omitempty"`

	PduSeId int32 `json:"pduSeId,omitempty"`

	Ipv4Addr string `json:"ipv4Addr,omitempty"`

	Ipv6Prefixes []Ipv6Prefix `json:"ipv6Prefixes,omitempty"`

	Ipv6Addrs []Ipv6Addr `json:"ipv6Addrs,omitempty"`

	PduSessType PduSessionType `json:"pduSessType,omitempty"`
}
