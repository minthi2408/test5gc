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

type SmContextUpdateError struct {
	Error *ExtProblemDetails `json:"error"`

	N1SmMsg RefToBinaryData `json:"n1SmMsg,omitempty"`

	N2SmInfo RefToBinaryData `json:"n2SmInfo,omitempty"`

	N2SmInfoType N2SmInfoType `json:"n2SmInfoType,omitempty"`

	UpCnxState UpCnxState `json:"upCnxState,omitempty"`

	RecoveryTime time.Time `json:"recoveryTime,omitempty"`
}
