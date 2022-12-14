/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type PartialSuccessReport struct {
	FailureCause FailureCause `json:"failureCause"`

	// Information about the PCC rules provisioned by the PCF not successfully installed/activated.
	RuleReports []RuleReport `json:"ruleReports,omitempty"`

	// Information about the session rules provisioned by the PCF not successfully installed.
	SessRuleReports []SessionRuleReport `json:"sessRuleReports,omitempty"`

	UeCampingRep UeCampingRep `json:"ueCampingRep,omitempty"`

	// Contains the type(s) of failed policy decision and/or condition data.
	PolicyDecFailureReports []PolicyDecisionFailureCode `json:"policyDecFailureReports,omitempty"`
}
