/*
Nsmf_PDUSession

SMF PDU Session Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type PsaIndication string

// List of PsaIndication
const (
	PSAINDICATIONANYOF_INSERTED      PsaIndication = "PSA_INSERTED"
	PSAINDICATIONANYOF_REMOVED       PsaIndication = "PSA_REMOVED"
	PSAINDICATIONANYOF_INSERTED_ONLY PsaIndication = "PSA_INSERTED_ONLY"
	PSAINDICATIONANYOF_REMOVED_ONLY  PsaIndication = "PSA_REMOVED_ONLY"
)
