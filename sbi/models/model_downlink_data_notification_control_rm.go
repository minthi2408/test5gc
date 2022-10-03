/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

// DownlinkDataNotificationControlRm - this data type is defined in the same way as the DownlinkDataNotificationControl data type, but with the OpenAPI nullable property set to true.
type DownlinkDataNotificationControlRm struct {

	NotifCtrlInds *[]NotificationControlIndication `json:"notifCtrlInds,omitempty"`

	TypesOfNotif *[]DlDataDeliveryStatus `json:"typesOfNotif,omitempty"`
}
