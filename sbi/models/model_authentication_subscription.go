/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type AuthenticationSubscription struct {

	AuthenticationMethod AuthMethod `json:"authenticationMethod"`

	EncPermanentKey string `json:"encPermanentKey,omitempty"`

	ProtectionParameterId string `json:"protectionParameterId,omitempty"`

	SequenceNumber SequenceNumber `json:"sequenceNumber,omitempty"`

	AuthenticationManagementField string `json:"authenticationManagementField,omitempty"`

	AlgorithmId string `json:"algorithmId,omitempty"`

	EncOpcKey string `json:"encOpcKey,omitempty"`

	EncTopcKey string `json:"encTopcKey,omitempty"`

	VectorGenerationInHss bool `json:"vectorGenerationInHss,omitempty"`

	N5gcAuthMethod AuthMethod `json:"n5gcAuthMethod,omitempty"`

	RgAuthenticationInd bool `json:"rgAuthenticationInd,omitempty"`

	Supi string `json:"supi,omitempty"`
}
