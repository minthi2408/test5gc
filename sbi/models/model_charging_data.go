/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type ChargingData struct {

	// Univocally identifies the charging control policy data within a PDU session.
	ChgId string `json:"chgId"`

	MeteringMethod MeteringMethod `json:"meteringMethod,omitempty"`

	// Indicates the offline charging is applicable to the PCC rule when it is included and set to true.
	Offline bool `json:"offline,omitempty"`

	// Indicates the online charging is applicable to the PCC rule when it is included and set to true.
	Online bool `json:"online,omitempty"`

	// Indicates whether the service data flow is allowed to start while the SMF is waiting for the response to the credit request.
	SdfHandl bool `json:"sdfHandl,omitempty"`

	RatingGroup int32 `json:"ratingGroup,omitempty"`

	ReportingLevel ReportingLevel `json:"reportingLevel,omitempty"`

	ServiceId int32 `json:"serviceId,omitempty"`

	// Indicates the sponsor identity.
	SponsorId string `json:"sponsorId,omitempty"`

	// Indicates the application service provider identity.
	AppSvcProvId string `json:"appSvcProvId,omitempty"`

	AfChargingIdentifier int32 `json:"afChargingIdentifier,omitempty"`

	AfChargId string `json:"afChargId,omitempty"`
}
