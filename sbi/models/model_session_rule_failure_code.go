/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

// SessionRuleFailureCode - Possible values are   - NF_MAL: Indicate that the PCC rule could not be successfully installed (for those provisioned from the PCF) or activated (for those pre-defined in SMF) or enforced (for those already successfully installed) due to SMF/UPF malfunction.   - RES_LIM: Indicate that the PCC rule could not be successfully installed (for those provisioned from PCF) or activated (for those pre-defined in SMF) or enforced (for those already successfully installed) due to a limitation of resources at the SMF/UPF.   - UNSUCC_QOS_VAL: indicate that the QoS validation has failed.   - UE_STA_SUSP: Indicates that the UE is in suspend state.
type SessionRuleFailureCode struct {
}
