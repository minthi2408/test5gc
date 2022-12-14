/*
Npcf_UEPolicyControl

UE Policy Control Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type N1N2MessageTransferCause string

// List of N1N2MessageTransferCause
const (
	N1N2MESSAGETRANSFERCAUSE_ATTEMPTING_TO_REACH_UE                N1N2MessageTransferCause = "ATTEMPTING_TO_REACH_UE"
	N1N2MESSAGETRANSFERCAUSE_N1_N2_TRANSFER_INITIATED              N1N2MessageTransferCause = "N1_N2_TRANSFER_INITIATED"
	N1N2MESSAGETRANSFERCAUSE_WAITING_FOR_ASYNCHRONOUS_TRANSFER     N1N2MessageTransferCause = "WAITING_FOR_ASYNCHRONOUS_TRANSFER"
	N1N2MESSAGETRANSFERCAUSE_UE_NOT_RESPONDING                     N1N2MessageTransferCause = "UE_NOT_RESPONDING"
	N1N2MESSAGETRANSFERCAUSE_N1_MSG_NOT_TRANSFERRED                N1N2MessageTransferCause = "N1_MSG_NOT_TRANSFERRED"
	N1N2MESSAGETRANSFERCAUSE_UE_NOT_REACHABLE_FOR_SESSION          N1N2MessageTransferCause = "UE_NOT_REACHABLE_FOR_SESSION"
	N1N2MESSAGETRANSFERCAUSE_TEMPORARY_REJECT_REGISTRATION_ONGOING N1N2MessageTransferCause = "TEMPORARY_REJECT_REGISTRATION_ONGOING"
	N1N2MESSAGETRANSFERCAUSE_TEMPORARY_REJECT_HANDOVER_ONGOING     N1N2MessageTransferCause = "TEMPORARY_REJECT_HANDOVER_ONGOING"
)
