/*
Nsmf_PDUSession

SMF PDU Session Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type VsmfUpdatedData struct {
	QosFlowsAddModList []QosFlowItem `json:"qosFlowsAddModList,omitempty"`

	QosFlowsRelList []QosFlowItem `json:"qosFlowsRelList,omitempty"`

	QosFlowsFailedtoAddModList []QosFlowItem `json:"qosFlowsFailedtoAddModList,omitempty"`

	QosFlowsFailedtoRelList []QosFlowItem `json:"qosFlowsFailedtoRelList,omitempty"`

	N1SmInfoFromUe RefToBinaryData `json:"n1SmInfoFromUe,omitempty"`

	UnknownN1SmInfo RefToBinaryData `json:"unknownN1SmInfo,omitempty"`

	UeLocation UserLocation `json:"ueLocation,omitempty"`

	UeTimeZone string `json:"ueTimeZone,omitempty"`

	AddUeLocation UserLocation `json:"addUeLocation,omitempty"`

	AssignedEbiList []EbiArpMapping `json:"assignedEbiList,omitempty"`

	FailedToAssignEbiList []Arp `json:"failedToAssignEbiList,omitempty"`

	ReleasedEbiList []int32 `json:"releasedEbiList,omitempty"`

	SecondaryRatUsageReport []SecondaryRatUsageReport `json:"secondaryRatUsageReport,omitempty"`

	SecondaryRatUsageInfo []SecondaryRatUsageInfo `json:"secondaryRatUsageInfo,omitempty"`

	N4Info N4Information `json:"n4Info,omitempty"`

	N4InfoExt1 N4Information `json:"n4InfoExt1,omitempty"`

	N4InfoExt2 N4Information `json:"n4InfoExt2,omitempty"`
}
