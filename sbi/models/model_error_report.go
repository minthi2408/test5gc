/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type ErrorReport struct {

	Error ProblemDetails `json:"error,omitempty"`

	// Used to report the PCC rule failure.
	RuleReports []RuleReport `json:"ruleReports,omitempty"`

	// Used to report the session rule failure.
	SessRuleReports []SessionRuleReport `json:"sessRuleReports,omitempty"`

	// Used to report failure of the policy decision and/or condition data.
	PolDecFailureReports []PolicyDecisionFailureCode `json:"polDecFailureReports,omitempty"`

	AltQosParamId string `json:"altQosParamId,omitempty"`
}