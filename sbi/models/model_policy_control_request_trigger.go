/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type PolicyControlRequestTrigger string

// List of PolicyControlRequestTrigger
const (
	POLICYCONTROLREQUESTTRIGGER_PLMN_CH                          PolicyControlRequestTrigger = "PLMN_CH"
	POLICYCONTROLREQUESTTRIGGER_RES_MO_RE                        PolicyControlRequestTrigger = "RES_MO_RE"
	POLICYCONTROLREQUESTTRIGGER_AC_TY_CH                         PolicyControlRequestTrigger = "AC_TY_CH"
	POLICYCONTROLREQUESTTRIGGER_UE_IP_CH                         PolicyControlRequestTrigger = "UE_IP_CH"
	POLICYCONTROLREQUESTTRIGGER_UE_MAC_CH                        PolicyControlRequestTrigger = "UE_MAC_CH"
	POLICYCONTROLREQUESTTRIGGER_AN_CH_COR                        PolicyControlRequestTrigger = "AN_CH_COR"
	POLICYCONTROLREQUESTTRIGGER_US_RE                            PolicyControlRequestTrigger = "US_RE"
	POLICYCONTROLREQUESTTRIGGER_APP_STA                          PolicyControlRequestTrigger = "APP_STA"
	POLICYCONTROLREQUESTTRIGGER_APP_STO                          PolicyControlRequestTrigger = "APP_STO"
	POLICYCONTROLREQUESTTRIGGER_AN_INFO                          PolicyControlRequestTrigger = "AN_INFO"
	POLICYCONTROLREQUESTTRIGGER_CM_SES_FAIL                      PolicyControlRequestTrigger = "CM_SES_FAIL"
	POLICYCONTROLREQUESTTRIGGER_PS_DA_OFF                        PolicyControlRequestTrigger = "PS_DA_OFF"
	POLICYCONTROLREQUESTTRIGGER_DEF_QOS_CH                       PolicyControlRequestTrigger = "DEF_QOS_CH"
	POLICYCONTROLREQUESTTRIGGER_SE_AMBR_CH                       PolicyControlRequestTrigger = "SE_AMBR_CH"
	POLICYCONTROLREQUESTTRIGGER_QOS_NOTIF                        PolicyControlRequestTrigger = "QOS_NOTIF"
	POLICYCONTROLREQUESTTRIGGER_NO_CREDIT                        PolicyControlRequestTrigger = "NO_CREDIT"
	POLICYCONTROLREQUESTTRIGGER_REALLO_OF_CREDIT                 PolicyControlRequestTrigger = "REALLO_OF_CREDIT"
	POLICYCONTROLREQUESTTRIGGER_PRA_CH                           PolicyControlRequestTrigger = "PRA_CH"
	POLICYCONTROLREQUESTTRIGGER_SAREA_CH                         PolicyControlRequestTrigger = "SAREA_CH"
	POLICYCONTROLREQUESTTRIGGER_SCNN_CH                          PolicyControlRequestTrigger = "SCNN_CH"
	POLICYCONTROLREQUESTTRIGGER_RE_TIMEOUT                       PolicyControlRequestTrigger = "RE_TIMEOUT"
	POLICYCONTROLREQUESTTRIGGER_RES_RELEASE                      PolicyControlRequestTrigger = "RES_RELEASE"
	POLICYCONTROLREQUESTTRIGGER_SUCC_RES_ALLO                    PolicyControlRequestTrigger = "SUCC_RES_ALLO"
	POLICYCONTROLREQUESTTRIGGER_RAT_TY_CH                        PolicyControlRequestTrigger = "RAT_TY_CH"
	POLICYCONTROLREQUESTTRIGGER_REF_QOS_IND_CH                   PolicyControlRequestTrigger = "REF_QOS_IND_CH"
	POLICYCONTROLREQUESTTRIGGER_NUM_OF_PACKET_FILTER             PolicyControlRequestTrigger = "NUM_OF_PACKET_FILTER"
	POLICYCONTROLREQUESTTRIGGER_UE_STATUS_RESUME                 PolicyControlRequestTrigger = "UE_STATUS_RESUME"
	POLICYCONTROLREQUESTTRIGGER_UE_TZ_CH                         PolicyControlRequestTrigger = "UE_TZ_CH"
	POLICYCONTROLREQUESTTRIGGER_AUTH_PROF_CH                     PolicyControlRequestTrigger = "AUTH_PROF_CH"
	POLICYCONTROLREQUESTTRIGGER_QOS_MONITORING                   PolicyControlRequestTrigger = "QOS_MONITORING"
	POLICYCONTROLREQUESTTRIGGER_SCELL_CH                         PolicyControlRequestTrigger = "SCELL_CH"
	POLICYCONTROLREQUESTTRIGGER_EPS_FALLBACK                     PolicyControlRequestTrigger = "EPS_FALLBACK"
	POLICYCONTROLREQUESTTRIGGER_MA_PDU                           PolicyControlRequestTrigger = "MA_PDU"
	POLICYCONTROLREQUESTTRIGGER_TSN_BRIDGE_INFO                  PolicyControlRequestTrigger = "TSN_BRIDGE_INFO"
	POLICYCONTROLREQUESTTRIGGER__5_G_RG_JOIN                     PolicyControlRequestTrigger = "5G_RG_JOIN"
	POLICYCONTROLREQUESTTRIGGER__5_G_RG_LEAVE                    PolicyControlRequestTrigger = "5G_RG_LEAVE"
	POLICYCONTROLREQUESTTRIGGER_DDN_FAILURE                      PolicyControlRequestTrigger = "DDN_FAILURE"
	POLICYCONTROLREQUESTTRIGGER_DDN_DELIVERY_STATUS              PolicyControlRequestTrigger = "DDN_DELIVERY_STATUS"
	POLICYCONTROLREQUESTTRIGGER_GROUP_ID_LIST_CHG                PolicyControlRequestTrigger = "GROUP_ID_LIST_CHG"
	POLICYCONTROLREQUESTTRIGGER_DDN_FAILURE_CANCELLATION         PolicyControlRequestTrigger = "DDN_FAILURE_CANCELLATION"
	POLICYCONTROLREQUESTTRIGGER_DDN_DELIVERY_STATUS_CANCELLATION PolicyControlRequestTrigger = "DDN_DELIVERY_STATUS_CANCELLATION"
	POLICYCONTROLREQUESTTRIGGER_VPLMN_QOS_CH                     PolicyControlRequestTrigger = "VPLMN_QOS_CH"
)
