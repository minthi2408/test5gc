/*
Namf_Communication

AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type PduSessionContext struct {
	PduSessionId int32 `json:"pduSessionId"`

	SmContextRef string `json:"smContextRef"`

	SNssai Snssai `json:"sNssai"`

	Dnn string `json:"dnn"`

	SelectedDnn string `json:"selectedDnn,omitempty"`

	AccessType AccessType `json:"accessType"`

	AdditionalAccessType AccessType `json:"additionalAccessType,omitempty"`

	AllocatedEbiList []EbiArpMapping `json:"allocatedEbiList,omitempty"`

	HsmfId string `json:"hsmfId,omitempty"`

	HsmfSetId string `json:"hsmfSetId,omitempty"`

	HsmfServiceSetId string `json:"hsmfServiceSetId,omitempty"`

	SmfBinding SbiBindingLevel `json:"smfBinding,omitempty"`

	VsmfId string `json:"vsmfId,omitempty"`

	VsmfSetId string `json:"vsmfSetId,omitempty"`

	VsmfServiceSetId string `json:"vsmfServiceSetId,omitempty"`

	VsmfBinding SbiBindingLevel `json:"vsmfBinding,omitempty"`

	IsmfId string `json:"ismfId,omitempty"`

	IsmfSetId string `json:"ismfSetId,omitempty"`

	IsmfServiceSetId string `json:"ismfServiceSetId,omitempty"`

	IsmfBinding SbiBindingLevel `json:"ismfBinding,omitempty"`

	NsInstance string `json:"nsInstance,omitempty"`

	SmfServiceInstanceId string `json:"smfServiceInstanceId,omitempty"`

	MaPduSession bool `json:"maPduSession,omitempty"`

	CnAssistedRanPara CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
}
