/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

import (
	"time"
)

type Amf3GppAccessRegistration struct {
	AmfInstanceId string `json:"amfInstanceId"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	PurgeFlag bool `json:"purgeFlag,omitempty"`

	Pei string `json:"pei,omitempty"`

	ImsVoPs ImsVoPs `json:"imsVoPs,omitempty"`

	DeregCallbackUri string `json:"deregCallbackUri"`

	AmfServiceNameDereg ServiceName `json:"amfServiceNameDereg,omitempty"`

	PcscfRestorationCallbackUri string `json:"pcscfRestorationCallbackUri,omitempty"`

	AmfServiceNamePcscfRest ServiceName `json:"amfServiceNamePcscfRest,omitempty"`

	InitialRegistrationInd bool `json:"initialRegistrationInd,omitempty"`

	Guami Guami `json:"guami"`

	BackupAmfInfo []BackupAmfInfo `json:"backupAmfInfo,omitempty"`

	DrFlag bool `json:"drFlag,omitempty"`

	RatType RatType `json:"ratType"`

	UrrpIndicator bool `json:"urrpIndicator,omitempty"`

	AmfEeSubscriptionId string `json:"amfEeSubscriptionId,omitempty"`

	EpsInterworkingInfo EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`

	UeSrvccCapability bool `json:"ueSrvccCapability,omitempty"`

	RegistrationTime time.Time `json:"registrationTime,omitempty"`

	VgmlcAddress VgmlcAddress `json:"vgmlcAddress,omitempty"`

	ContextInfo ContextInfo `json:"contextInfo,omitempty"`

	NoEeSubscriptionInd bool `json:"noEeSubscriptionInd,omitempty"`

	Supi string `json:"supi,omitempty"`
}
