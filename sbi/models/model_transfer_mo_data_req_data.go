/*
Nsmf_PDUSession

SMF PDU Session Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type TransferMoDataReqData struct {

	MoData RefToBinaryData `json:"moData"`

	MoExpDataCounter MoExpDataCounter `json:"moExpDataCounter,omitempty"`

	UeLocation UserLocation `json:"ueLocation,omitempty"`
}