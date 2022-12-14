/*
Nudm_UECM

Nudm Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type Amf3GppAccessRegistrationModification struct {
	Guami Guami `json:"guami"`

	PurgeFlag bool `json:"purgeFlag,omitempty"`

	Pei string `json:"pei,omitempty"`

	ImsVoPs ImsVoPs `json:"imsVoPs,omitempty"`

	BackupAmfInfo []BackupAmfInfo `json:"backupAmfInfo,omitempty"`

	EpsInterworkingInfo EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`

	UeSrvccCapability *bool `json:"ueSrvccCapability,omitempty"`
}
