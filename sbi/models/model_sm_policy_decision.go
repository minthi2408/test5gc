/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

import (
	"time"
)

type SmPolicyDecision struct {

	// A map of Sessionrules with the content being the SessionRule as described in subclause 5.6.2.7.
	SessRules map[string]SessionRule `json:"sessRules,omitempty"`

	// A map of PCC rules with the content being the PCCRule as described in subclause 5.6.2.6.
	PccRules *map[string]PccRule `json:"pccRules,omitempty"`

	// If it is included and set to true, it indicates the P-CSCF Restoration is requested.
	PcscfRestIndication bool `json:"pcscfRestIndication,omitempty"`

	// Map of QoS data policy decisions.
	QosDecs map[string]QosData `json:"qosDecs,omitempty"`

	// Map of Charging data policy decisions.
	ChgDecs *map[string]ChargingData `json:"chgDecs,omitempty"`

	ChargingInfo ChargingInformation `json:"chargingInfo,omitempty"`

	// Map of Traffic Control data policy decisions.
	TraffContDecs map[string]TrafficControlData `json:"traffContDecs,omitempty"`

	// Map of Usage Monitoring data policy decisions.
	UmDecs *map[string]UsageMonitoringData `json:"umDecs,omitempty"`

	// Map of QoS characteristics for non standard 5QIs. This map uses the 5QI values as keys.
	QosChars map[string]QosCharacteristics `json:"qosChars,omitempty"`

	// Map of QoS Monitoring data policy decisions.
	QosMonDecs *map[string]QosMonitoringData `json:"qosMonDecs,omitempty"`

	ReflectiveQoSTimer int32 `json:"reflectiveQoSTimer,omitempty"`

	// A map of condition data with the content being as described in subclause 5.6.2.9.
	Conds *map[string]ConditionData `json:"conds,omitempty"`

	RevalidationTime time.Time `json:"revalidationTime,omitempty"`

	// Indicates the offline charging is applicable to the PDU session when it is included and set to true.
	Offline bool `json:"offline,omitempty"`

	// Indicates the online charging is applicable to the PDU session when it is included and set to true.
	Online bool `json:"online,omitempty"`

	// Defines the policy control request triggers subscribed by the PCF.
	PolicyCtrlReqTriggers *[]PolicyControlRequestTrigger `json:"policyCtrlReqTriggers,omitempty"`

	// Defines the last list of rule control data requested by the PCF.
	LastReqRuleData []RequestedRuleData `json:"lastReqRuleData,omitempty"`

	LastReqUsageData RequestedUsageData `json:"lastReqUsageData,omitempty"`

	// Map of PRA information.
	PraInfos *map[string]PresenceInfoRm `json:"praInfos,omitempty"`

	Ipv4Index int32 `json:"ipv4Index,omitempty"`

	Ipv6Index int32 `json:"ipv6Index,omitempty"`

	QosFlowUsage QosFlowUsage `json:"qosFlowUsage,omitempty"`

	RelCause SmPolicyAssociationReleaseCause `json:"relCause,omitempty"`

	SuppFeat string `json:"suppFeat,omitempty"`

	TsnBridgeManCont BridgeManagementContainer `json:"tsnBridgeManCont,omitempty"`

	TsnPortManContDstt PortManagementContainer `json:"tsnPortManContDstt,omitempty"`

	TsnPortManContNwtts []PortManagementContainer `json:"tsnPortManContNwtts,omitempty"`

	// Indicates whether the PDU session is a redundant PDU session. If absent it means the PDU session is not a redundant PDU session.
	RedSessIndication bool `json:"redSessIndication,omitempty"`
}