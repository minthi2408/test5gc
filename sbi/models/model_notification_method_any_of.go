/*
Npcf_EventExposure

PCF Event Exposure Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type NotificationMethodAnyOf string

// List of NotificationMethodAnyOf
const (
	NOTIFICATIONMETHODANYOF_PERIODIC NotificationMethodAnyOf = "PERIODIC"
	NOTIFICATIONMETHODANYOF_ONE_TIME NotificationMethodAnyOf = "ONE_TIME"
	NOTIFICATIONMETHODANYOF_ON_EVENT_DETECTION NotificationMethodAnyOf = "ON_EVENT_DETECTION"
)