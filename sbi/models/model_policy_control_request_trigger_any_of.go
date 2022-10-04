/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type PolicyControlRequestTriggerAnyOf string

// List of PolicyControlRequestTriggerAnyOf
const (
	POLICYCONTROLREQUESTTRIGGERANYOF_PLMN_CH                          PolicyControlRequestTriggerAnyOf = "PLMN_CH"
	POLICYCONTROLREQUESTTRIGGERANYOF_RES_MO_RE                        PolicyControlRequestTriggerAnyOf = "RES_MO_RE"
	POLICYCONTROLREQUESTTRIGGERANYOF_AC_TY_CH                         PolicyControlRequestTriggerAnyOf = "AC_TY_CH"
	POLICYCONTROLREQUESTTRIGGERANYOF_UE_IP_CH                         PolicyControlRequestTriggerAnyOf = "UE_IP_CH"
	POLICYCONTROLREQUESTTRIGGERANYOF_UE_MAC_CH                        PolicyControlRequestTriggerAnyOf = "UE_MAC_CH"
	POLICYCONTROLREQUESTTRIGGERANYOF_AN_CH_COR                        PolicyControlRequestTriggerAnyOf = "AN_CH_COR"
	POLICYCONTROLREQUESTTRIGGERANYOF_US_RE                            PolicyControlRequestTriggerAnyOf = "US_RE"
	POLICYCONTROLREQUESTTRIGGERANYOF_APP_STA                          PolicyControlRequestTriggerAnyOf = "APP_STA"
	POLICYCONTROLREQUESTTRIGGERANYOF_APP_STO                          PolicyControlRequestTriggerAnyOf = "APP_STO"
	POLICYCONTROLREQUESTTRIGGERANYOF_AN_INFO                          PolicyControlRequestTriggerAnyOf = "AN_INFO"
	POLICYCONTROLREQUESTTRIGGERANYOF_CM_SES_FAIL                      PolicyControlRequestTriggerAnyOf = "CM_SES_FAIL"
	POLICYCONTROLREQUESTTRIGGERANYOF_PS_DA_OFF                        PolicyControlRequestTriggerAnyOf = "PS_DA_OFF"
	POLICYCONTROLREQUESTTRIGGERANYOF_DEF_QOS_CH                       PolicyControlRequestTriggerAnyOf = "DEF_QOS_CH"
	POLICYCONTROLREQUESTTRIGGERANYOF_SE_AMBR_CH                       PolicyControlRequestTriggerAnyOf = "SE_AMBR_CH"
	POLICYCONTROLREQUESTTRIGGERANYOF_QOS_NOTIF                        PolicyControlRequestTriggerAnyOf = "QOS_NOTIF"
	POLICYCONTROLREQUESTTRIGGERANYOF_NO_CREDIT                        PolicyControlRequestTriggerAnyOf = "NO_CREDIT"
	POLICYCONTROLREQUESTTRIGGERANYOF_REALLO_OF_CREDIT                 PolicyControlRequestTriggerAnyOf = "REALLO_OF_CREDIT"
	POLICYCONTROLREQUESTTRIGGERANYOF_PRA_CH                           PolicyControlRequestTriggerAnyOf = "PRA_CH"
	POLICYCONTROLREQUESTTRIGGERANYOF_SAREA_CH                         PolicyControlRequestTriggerAnyOf = "SAREA_CH"
	POLICYCONTROLREQUESTTRIGGERANYOF_SCNN_CH                          PolicyControlRequestTriggerAnyOf = "SCNN_CH"
	POLICYCONTROLREQUESTTRIGGERANYOF_RE_TIMEOUT                       PolicyControlRequestTriggerAnyOf = "RE_TIMEOUT"
	POLICYCONTROLREQUESTTRIGGERANYOF_RES_RELEASE                      PolicyControlRequestTriggerAnyOf = "RES_RELEASE"
	POLICYCONTROLREQUESTTRIGGERANYOF_SUCC_RES_ALLO                    PolicyControlRequestTriggerAnyOf = "SUCC_RES_ALLO"
	POLICYCONTROLREQUESTTRIGGERANYOF_RAT_TY_CH                        PolicyControlRequestTriggerAnyOf = "RAT_TY_CH"
	POLICYCONTROLREQUESTTRIGGERANYOF_REF_QOS_IND_CH                   PolicyControlRequestTriggerAnyOf = "REF_QOS_IND_CH"
	POLICYCONTROLREQUESTTRIGGERANYOF_NUM_OF_PACKET_FILTER             PolicyControlRequestTriggerAnyOf = "NUM_OF_PACKET_FILTER"
	POLICYCONTROLREQUESTTRIGGERANYOF_UE_STATUS_RESUME                 PolicyControlRequestTriggerAnyOf = "UE_STATUS_RESUME"
	POLICYCONTROLREQUESTTRIGGERANYOF_UE_TZ_CH                         PolicyControlRequestTriggerAnyOf = "UE_TZ_CH"
	POLICYCONTROLREQUESTTRIGGERANYOF_AUTH_PROF_CH                     PolicyControlRequestTriggerAnyOf = "AUTH_PROF_CH"
	POLICYCONTROLREQUESTTRIGGERANYOF_QOS_MONITORING                   PolicyControlRequestTriggerAnyOf = "QOS_MONITORING"
	POLICYCONTROLREQUESTTRIGGERANYOF_SCELL_CH                         PolicyControlRequestTriggerAnyOf = "SCELL_CH"
	POLICYCONTROLREQUESTTRIGGERANYOF_EPS_FALLBACK                     PolicyControlRequestTriggerAnyOf = "EPS_FALLBACK"
	POLICYCONTROLREQUESTTRIGGERANYOF_MA_PDU                           PolicyControlRequestTriggerAnyOf = "MA_PDU"
	POLICYCONTROLREQUESTTRIGGERANYOF_TSN_BRIDGE_INFO                  PolicyControlRequestTriggerAnyOf = "TSN_BRIDGE_INFO"
	POLICYCONTROLREQUESTTRIGGERANYOF__5_G_RG_JOIN                     PolicyControlRequestTriggerAnyOf = "5G_RG_JOIN"
	POLICYCONTROLREQUESTTRIGGERANYOF__5_G_RG_LEAVE                    PolicyControlRequestTriggerAnyOf = "5G_RG_LEAVE"
	POLICYCONTROLREQUESTTRIGGERANYOF_DDN_FAILURE                      PolicyControlRequestTriggerAnyOf = "DDN_FAILURE"
	POLICYCONTROLREQUESTTRIGGERANYOF_DDN_DELIVERY_STATUS              PolicyControlRequestTriggerAnyOf = "DDN_DELIVERY_STATUS"
	POLICYCONTROLREQUESTTRIGGERANYOF_GROUP_ID_LIST_CHG                PolicyControlRequestTriggerAnyOf = "GROUP_ID_LIST_CHG"
	POLICYCONTROLREQUESTTRIGGERANYOF_DDN_FAILURE_CANCELLATION         PolicyControlRequestTriggerAnyOf = "DDN_FAILURE_CANCELLATION"
	POLICYCONTROLREQUESTTRIGGERANYOF_DDN_DELIVERY_STATUS_CANCELLATION PolicyControlRequestTriggerAnyOf = "DDN_DELIVERY_STATUS_CANCELLATION"
	POLICYCONTROLREQUESTTRIGGERANYOF_VPLMN_QOS_CH                     PolicyControlRequestTriggerAnyOf = "VPLMN_QOS_CH"
)
