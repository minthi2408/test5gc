/*
Npcf_AMPolicyControl

Access and Mobility Policy Control Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package ampc

import (
	"etri5gc/sbi"
	"net/http"
)

var Routes = sbi.SbiRoutes{

	{
		Label:   "CreateIndividualAMPolicyAssociation",
		Method:  http.MethodPost,
		Path:    "/npcf-am-policy-control/v1/policies",
		Handler: OnCreateIndividualAMPolicyAssociation,
	},
	{
		Label:   "DeleteIndividualAMPolicyAssociation",
		Method:  http.MethodDelete,
		Path:    "/npcf-am-policy-control/v1/policies/{polAssoId}",
		Handler: OnDeleteIndividualAMPolicyAssociation,
	},
	{
		Label:   "ReadIndividualAMPolicyAssociation",
		Method:  http.MethodGet,
		Path:    "/npcf-am-policy-control/v1/policies/{polAssoId}",
		Handler: OnReadIndividualAMPolicyAssociation,
	},
	{
		Label:   "ReportObservedEventTriggersForIndividualAMPolicyAssociation",
		Method:  http.MethodPost,
		Path:    "/npcf-am-policy-control/v1/policies/{polAssoId}/update",
		Handler: OnReportObservedEventTriggersForIndividualAMPolicyAssociation,
	},
}
