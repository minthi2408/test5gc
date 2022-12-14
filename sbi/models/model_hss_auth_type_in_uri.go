/*
Nudm_UEAU

UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type HssAuthTypeInUri string

// List of HssAuthTypeInUri
const (
	HSSAUTHTYPEINURI_EPS_AKA       HssAuthTypeInUri = "eps-aka"
	HSSAUTHTYPEINURI_EAP_AKA       HssAuthTypeInUri = "eap-aka"
	HSSAUTHTYPEINURI_EAP_AKA_PRIME HssAuthTypeInUri = "eap-aka-prime"
	HSSAUTHTYPEINURI_IMS_AKA       HssAuthTypeInUri = "ims-aka"
	HSSAUTHTYPEINURI_GBA_AKA       HssAuthTypeInUri = "gba-aka"
)
