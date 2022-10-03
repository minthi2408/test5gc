/*
Namf_Communication

AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type N2InformationTransferReqData struct {

	TaiList []Tai `json:"taiList,omitempty"`

	RatSelector RatSelector `json:"ratSelector,omitempty"`

	GlobalRanNodeList []GlobalRanNodeId `json:"globalRanNodeList,omitempty"`

	N2Information N2InfoContainer `json:"n2Information"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}
