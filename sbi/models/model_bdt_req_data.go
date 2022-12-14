/*
Npcf_BDTPolicyControl Service API

PCF BDT Policy Control Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

// BdtReqData - Contains service requirements for creation a new Individual BDT policy resource.
type BdtReqData struct {

	// Contains an identity of an application service provider.
	AspId string `json:"aspId"`

	DesTimeInt TimeWindow `json:"desTimeInt"`

	Dnn string `json:"dnn,omitempty"`

	InterGroupId string `json:"interGroupId,omitempty"`

	NotifUri string `json:"notifUri,omitempty"`

	NwAreaInfo NetworkAreaInfo `json:"nwAreaInfo,omitempty"`

	// Indicates a number of UEs.
	NumOfUes int32 `json:"numOfUes"`

	VolPerUe UsageThreshold `json:"volPerUe"`

	Snssai Snssai `json:"snssai,omitempty"`

	SuppFeat string `json:"suppFeat,omitempty"`

	// Identify a traffic descriptor as defined in Figure 5.2.2 of 3GPP TS 24.526, octets v+5 to w.
	TrafficDes string `json:"trafficDes,omitempty"`

	// Indicates whether the BDT warning notification is enabled or disabled.
	WarnNotifReq bool `json:"warnNotifReq,omitempty"`
}
