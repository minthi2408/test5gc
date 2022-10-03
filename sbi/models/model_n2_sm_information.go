/*
Namf_Communication

AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type N2SmInformation struct {

	PduSessionId int32 `json:"pduSessionId"`

	N2InfoContent N2InfoContent `json:"n2InfoContent,omitempty"`

	SNssai Snssai `json:"sNssai,omitempty"`

	HomePlmnSnssai Snssai `json:"homePlmnSnssai,omitempty"`

	IwkSnssai Snssai `json:"iwkSnssai,omitempty"`

	SubjectToHo bool `json:"subjectToHo,omitempty"`
}
