/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type LocationArea1 struct {

	// Identifies a list of geographic area of the user where the UE is located.
	GeographicAreas []GeographicArea `json:"geographicAreas,omitempty"`

	// Identifies a list of civic addresses of the user where the UE is located.
	CivicAddresses []CivicAddress `json:"civicAddresses,omitempty"`

	NwAreaInfo NetworkAreaInfo1 `json:"nwAreaInfo,omitempty"`
}
