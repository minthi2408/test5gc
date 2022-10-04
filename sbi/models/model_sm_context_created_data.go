/*
Nsmf_PDUSession

SMF PDU Session Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

import (
	"time"
)

type SmContextCreatedData struct {
	HSmfUri string `json:"hSmfUri,omitempty"`

	SmfUri string `json:"smfUri,omitempty"`

	PduSessionId int32 `json:"pduSessionId,omitempty"`

	SNssai Snssai `json:"sNssai,omitempty"`

	UpCnxState UpCnxState `json:"upCnxState,omitempty"`

	N2SmInfo RefToBinaryData `json:"n2SmInfo,omitempty"`

	N2SmInfoType N2SmInfoType `json:"n2SmInfoType,omitempty"`

	AllocatedEbiList []EbiArpMapping `json:"allocatedEbiList,omitempty"`

	HoState HoState `json:"hoState,omitempty"`

	Gpsi string `json:"gpsi,omitempty"`

	SmfServiceInstanceId string `json:"smfServiceInstanceId,omitempty"`

	RecoveryTime time.Time `json:"recoveryTime,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	SelectedSmfId string `json:"selectedSmfId,omitempty"`

	SelectedOldSmfId string `json:"selectedOldSmfId,omitempty"`
}
