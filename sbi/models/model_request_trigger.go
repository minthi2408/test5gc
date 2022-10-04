/*
Npcf_UEPolicyControl

UE Policy Control Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

// RequestTrigger - Possible values are - LOC_CH: Location change (tracking area). The tracking area of the UE has changed. - PRA_CH: Change of UE presence in PRA. The AMF reports the current presence status of the UE in a Presence Reporting Area, and notifies that the UE enters/leaves the Presence Reporting Area. - UE_POLICY: A MANAGE UE POLICY COMPLETE message or a MANAGE UE POLICY COMMAND REJECT message, as defined in Annex D.5 of 3GPP TS 24.501 or a \"UE POLICY PROVISIONING REQUEST\" message, as defined in subclause 7.2.1.1 of 3GPP TS 24.587 , has been received by the AMF and is being forwarded. - PLMN_CH: PLMN change. the serving PLMN of UE has changed.  - CON_STATE_CH: Connectivity state change: the connectivity state of UE has changed.  - GROUP_ID_LIST_CHG: UE Internal Group Identifier(s) has changed. This event does not require a subscription
type RequestTrigger struct {
}
