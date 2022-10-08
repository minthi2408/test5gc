/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type AccessTech string

// List of AccessTech
const (
	ACCESSTECH_NR                                AccessTech = "NR"
	ACCESSTECH_EUTRAN_IN_WBS1_MODE_AND_NBS1_MODE AccessTech = "EUTRAN_IN_WBS1_MODE_AND_NBS1_MODE"
	ACCESSTECH_EUTRAN_IN_NBS1_MODE_ONLY          AccessTech = "EUTRAN_IN_NBS1_MODE_ONLY"
	ACCESSTECH_EUTRAN_IN_WBS1_MODE_ONLY          AccessTech = "EUTRAN_IN_WBS1_MODE_ONLY"
	ACCESSTECH_UTRAN                             AccessTech = "UTRAN"
	ACCESSTECH_GSM_AND_ECGSM_IO_T                AccessTech = "GSM_AND_ECGSM_IoT"
	ACCESSTECH_GSM_WITHOUT_ECGSM_IO_T            AccessTech = "GSM_WITHOUT_ECGSM_IoT"
	ACCESSTECH_ECGSM_IO_T_ONLY                   AccessTech = "ECGSM_IoT_ONLY"
	ACCESSTECH_CDMA_1X_RTT                       AccessTech = "CDMA_1xRTT"
	ACCESSTECH_CDMA_HRPD                         AccessTech = "CDMA_HRPD"
	ACCESSTECH_GSM_COMPACT                       AccessTech = "GSM_COMPACT"
)
