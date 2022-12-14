/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

import (
	"time"
)

// EventsNotification - describes the notification of a matched event
type EventsNotification struct {
	AccessType AccessType `json:"accessType,omitempty"`

	AddAccessInfo AdditionalAccessInfo `json:"addAccessInfo,omitempty"`

	RelAccessInfo AdditionalAccessInfo `json:"relAccessInfo,omitempty"`

	AnChargAddr AccNetChargingAddress `json:"anChargAddr,omitempty"`

	AnChargIds []AccessNetChargingIdentifier `json:"anChargIds,omitempty"`

	AnGwAddr AnGwAddress `json:"anGwAddr,omitempty"`

	EvSubsUri string `json:"evSubsUri"`

	EvNotifs []AfEventNotification `json:"evNotifs"`

	FailedResourcAllocReports []ResourcesAllocationInfo `json:"failedResourcAllocReports,omitempty"`

	SuccResourcAllocReports []ResourcesAllocationInfo `json:"succResourcAllocReports,omitempty"`

	NoNetLocSupp NetLocAccessSupport `json:"noNetLocSupp,omitempty"`

	OutOfCredReports []OutOfCreditInformation `json:"outOfCredReports,omitempty"`

	PlmnId PlmnIdNid `json:"plmnId,omitempty"`

	QncReports []QosNotificationControlInfo `json:"qncReports,omitempty"`

	QosMonReports []QosMonitoringReport `json:"qosMonReports,omitempty"`

	// Contains the RAN and/or NAS release cause.
	RanNasRelCauses []RanNasRelCause `json:"ranNasRelCauses,omitempty"`

	RatType RatType `json:"ratType,omitempty"`

	UeLoc UserLocation `json:"ueLoc,omitempty"`

	UeLocTime time.Time `json:"ueLocTime,omitempty"`

	UeTimeZone string `json:"ueTimeZone,omitempty"`

	UsgRep AccumulatedUsage `json:"usgRep,omitempty"`

	TsnBridgeManCont BridgeManagementContainer `json:"tsnBridgeManCont,omitempty"`

	TsnPortManContDstt PortManagementContainer `json:"tsnPortManContDstt,omitempty"`

	TsnPortManContNwtts []PortManagementContainer `json:"tsnPortManContNwtts,omitempty"`
}
