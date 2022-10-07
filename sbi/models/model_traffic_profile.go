/*
Nsmf_PDUSession

SMF PDU Session Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type TrafficProfile string

// List of TrafficProfileAnyOf
const (
	TRAFFICPROFILEANYOF_SINGLE_TRANS_UL     TrafficProfile = "SINGLE_TRANS_UL"
	TRAFFICPROFILEANYOF_SINGLE_TRANS_DL     TrafficProfile = "SINGLE_TRANS_DL"
	TRAFFICPROFILEANYOF_DUAL_TRANS_UL_FIRST TrafficProfile = "DUAL_TRANS_UL_FIRST"
	TRAFFICPROFILEANYOF_DUAL_TRANS_DL_FIRST TrafficProfile = "DUAL_TRANS_DL_FIRST"
	TRAFFICPROFILEANYOF_MULTI_TRANS         TrafficProfile = "MULTI_TRANS"
)
