/*
Namf_MT

AMF Mobile Terminated Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type ProblemDetailsEnableUeReachability struct {

	Type string `json:"type,omitempty"`

	Title string `json:"title,omitempty"`

	Status int32 `json:"status,omitempty"`

	Detail string `json:"detail,omitempty"`

	Instance string `json:"instance,omitempty"`

	Cause string `json:"cause,omitempty"`

	InvalidParams []InvalidParam `json:"invalidParams,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	AccessTokenError AccessTokenErr `json:"accessTokenError,omitempty"`

	AccessTokenRequest AccessTokenReq `json:"accessTokenRequest,omitempty"`

	NrfId string `json:"nrfId,omitempty"`

	MaxWaitingTime int32 `json:"maxWaitingTime,omitempty"`
}
