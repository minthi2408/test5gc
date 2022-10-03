/*
Nsmf_PDUSession

SMF PDU Session Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type ModifyPduSession400Response struct {

	JsonData VsmfUpdateError `json:"jsonData,omitempty"`

	BinaryDataN1SmInfoFromUe []byte `json:"binaryDataN1SmInfoFromUe,omitempty"`

	BinaryDataUnknownN1SmInfo []byte `json:"binaryDataUnknownN1SmInfo,omitempty"`

	BinaryDataN4Information []byte `json:"binaryDataN4Information,omitempty"`

	BinaryDataN4InformationExt1 []byte `json:"binaryDataN4InformationExt1,omitempty"`

	BinaryDataN4InformationExt2 []byte `json:"binaryDataN4InformationExt2,omitempty"`
}
