/*
Nudm_MT

UDM MT Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type LocationInfoResult struct {

	VPlmnId PlmnId `json:"vPlmnId,omitempty"`

	AmfInstanceId string `json:"amfInstanceId,omitempty"`

	SmsfInstanceId string `json:"smsfInstanceId,omitempty"`

	Ncgi Ncgi `json:"ncgi,omitempty"`

	Ecgi Ecgi `json:"ecgi,omitempty"`

	Tai Tai `json:"tai,omitempty"`

	CurrentLoc bool `json:"currentLoc,omitempty"`

	GeoInfo GeographicArea `json:"geoInfo,omitempty"`

	LocationAge int32 `json:"locationAge,omitempty"`

	RatType RatType `json:"ratType,omitempty"`

	Timezone string `json:"timezone,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}
