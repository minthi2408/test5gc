/*
Namf_Communication

AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type N2InfoContainer struct {

	N2InformationClass N2InformationClass `json:"n2InformationClass"`

	SmInfo N2SmInformation `json:"smInfo,omitempty"`

	RanInfo N2RanInformation `json:"ranInfo,omitempty"`

	NrppaInfo NrppaInformation `json:"nrppaInfo,omitempty"`

	PwsInfo PwsInformation `json:"pwsInfo,omitempty"`

	V2xInfo V2xInformation `json:"v2xInfo,omitempty"`
}
