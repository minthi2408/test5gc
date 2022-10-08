/*
Nsmf_PDUSession

SMF PDU Session Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type EpsInterworkingIndication string

// List of EpsInterworkingIndication
const (
	EPSINTERWORKINGINDICATION_NONE          EpsInterworkingIndication = "NONE"
	EPSINTERWORKINGINDICATION_WITH_N26      EpsInterworkingIndication = "WITH_N26"
	EPSINTERWORKINGINDICATION_WITHOUT_N26   EpsInterworkingIndication = "WITHOUT_N26"
	EPSINTERWORKINGINDICATION_IWK_NON_3_GPP EpsInterworkingIndication = "IWK_NON_3GPP"
)
