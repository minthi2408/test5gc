/*
Npcf_EventExposure

PCF Event Exposure Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type NotificationMethod string

// List of NotificationMethod
const (
	NOTIFICATIONMETHOD_PERIODIC           NotificationMethod = "PERIODIC"
	NOTIFICATIONMETHOD_ONE_TIME           NotificationMethod = "ONE_TIME"
	NOTIFICATIONMETHOD_ON_EVENT_DETECTION NotificationMethod = "ON_EVENT_DETECTION"
)
