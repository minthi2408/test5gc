/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type ContextDataSets struct {

	Amf3Gpp Amf3GppAccessRegistration `json:"amf3Gpp,omitempty"`

	AmfNon3Gpp AmfNon3GppAccessRegistration `json:"amfNon3Gpp,omitempty"`

	SdmSubscriptions []SdmSubscription `json:"sdmSubscriptions,omitempty"`

	EeSubscriptions []EeSubscription `json:"eeSubscriptions,omitempty"`

	Smsf3GppAccess SmsfRegistration `json:"smsf3GppAccess,omitempty"`

	SmsfNon3GppAccess SmsfRegistration `json:"smsfNon3GppAccess,omitempty"`

	SubscriptionDataSubscriptions []SubscriptionDataSubscriptions `json:"subscriptionDataSubscriptions,omitempty"`

	SmfRegistrations []SmfRegistration `json:"smfRegistrations,omitempty"`

	IpSmGw IpSmGwRegistration `json:"ipSmGw,omitempty"`
}
