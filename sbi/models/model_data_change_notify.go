/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type DataChangeNotify struct {
	OriginalCallbackReference []string `json:"originalCallbackReference,omitempty"`

	UeId string `json:"ueId,omitempty"`

	NotifyItems []NotifyItem `json:"notifyItems,omitempty"`

	SdmSubscription SdmSubscription1 `json:"sdmSubscription,omitempty"`

	AdditionalSdmSubscriptions []SdmSubscription1 `json:"additionalSdmSubscriptions,omitempty"`

	SubscriptionDataSubscriptions []SubscriptionDataSubscriptions `json:"subscriptionDataSubscriptions,omitempty"`
}
