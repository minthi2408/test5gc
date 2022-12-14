/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type UpPathChgEvent struct {
	NotificationUri string `json:"notificationUri"`

	// It is used to set the value of Notification Correlation ID in the notification sent by the SMF.
	NotifCorreId string `json:"notifCorreId"`

	DnaiChgType DnaiChangeType `json:"dnaiChgType"`

	AfAckInd bool `json:"afAckInd,omitempty"`
}
