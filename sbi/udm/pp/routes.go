/*
Nudm_PP

Nudm Parameter Provision Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package pp

import (
	"etrib5gc/sbi"
	"net/http"
)

var Routes = sbi.SbiRoutes{

	{
		Label:   "Create5GVNGroup",
		Method:  http.MethodPut,
		Path:    "/nudm-pp/v1/5g-vn-groups/{extGroupId}",
		Handler: OnCreate5GVNGroup,
	},
	{
		Label:   "Delete5GVNGroup",
		Method:  http.MethodDelete,
		Path:    "/nudm-pp/v1/5g-vn-groups/{extGroupId}",
		Handler: OnDelete5GVNGroup,
	},
	{
		Label:   "Get5GVNGroup",
		Method:  http.MethodGet,
		Path:    "/nudm-pp/v1/5g-vn-groups/{extGroupId}",
		Handler: OnGet5GVNGroup,
	},
	{
		Label:   "Modify5GVNGroup",
		Method:  http.MethodPatch,
		Path:    "/nudm-pp/v1/5g-vn-groups/{extGroupId}",
		Handler: OnModify5GVNGroup,
	},
	{
		Label:   "Update",
		Method:  http.MethodPatch,
		Path:    "/nudm-pp/v1/{ueId}/pp-data",
		Handler: OnUpdate,
	},
}
