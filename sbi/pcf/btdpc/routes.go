/*
Npcf_BDTPolicyControl Service API

PCF BDT Policy Control Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package btdpc

import (
	"etrib5gc/sbi"
	"net/http"
)

var Routes = sbi.SbiRoutes{

	{
		Label:   "CreateBDTPolicy",
		Method:  http.MethodPost,
		Path:    "/npcf-bdtpolicycontrol/v1/bdtpolicies",
		Handler: OnCreateBDTPolicy,
	},
	{
		Label:   "GetBDTPolicy",
		Method:  http.MethodGet,
		Path:    "/npcf-bdtpolicycontrol/v1/bdtpolicies/{bdtPolicyId}",
		Handler: OnGetBDTPolicy,
	},
	{
		Label:   "UpdateBDTPolicy",
		Method:  http.MethodPatch,
		Path:    "/npcf-bdtpolicycontrol/v1/bdtpolicies/{bdtPolicyId}",
		Handler: OnUpdateBDTPolicy,
	},
}
