/*
Nudm_UEAU

UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type AvEapAkaPrime struct {

	AvType AvType `json:"avType"`

	Rand string `json:"rand"`

	Xres string `json:"xres"`

	Autn string `json:"autn"`

	CkPrime string `json:"ckPrime"`

	IkPrime string `json:"ikPrime"`
}
