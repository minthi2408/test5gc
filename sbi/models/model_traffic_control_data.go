/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type TrafficControlData struct {

	// Univocally identifies the traffic control policy data within a PDU session.
	TcId string `json:"tcId"`

	FlowStatus FlowStatus `json:"flowStatus,omitempty"`

	RedirectInfo RedirectInformation `json:"redirectInfo,omitempty"`

	AddRedirectInfo []RedirectInformation `json:"addRedirectInfo,omitempty"`

	// Indicates whether applicat'on's start or stop notification is to be muted.
	MuteNotif bool `json:"muteNotif,omitempty"`

	// Reference to a pre-configured traffic steering policy for downlink traffic at the SMF.
	TrafficSteeringPolIdDl *string `json:"trafficSteeringPolIdDl,omitempty"`

	// Reference to a pre-configured traffic steering policy for uplink traffic at the SMF.
	TrafficSteeringPolIdUl *string `json:"trafficSteeringPolIdUl,omitempty"`

	// A list of location which the traffic shall be routed to for the AF request
	RouteToLocs []RouteToLocation `json:"routeToLocs,omitempty"`

	TraffCorreInd bool `json:"traffCorreInd,omitempty"`

	UpPathChgEvent *UpPathChgEvent `json:"upPathChgEvent,omitempty"`

	SteerFun SteeringFunctionality `json:"steerFun,omitempty"`

	SteerModeDl SteeringMode `json:"steerModeDl,omitempty"`

	SteerModeUl SteeringMode `json:"steerModeUl,omitempty"`

	MulAccCtrl MulticastAccessControl `json:"mulAccCtrl,omitempty"`
}
