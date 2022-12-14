/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

// NotificationItem - Identifies a data change notification when the change occurs in a fragment (subset of resource data) of a given resource.
type NotificationItem struct {
	ResourceId string `json:"resourceId"`

	NotifItems []UpdatedItem `json:"notifItems"`
}
