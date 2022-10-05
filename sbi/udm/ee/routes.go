/*
Nudm_EE

Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package ee

import (
	"etri5gc/sbi"
	"net/http"
)

var Routes = sbi.SbiRoutes{

	{
		Label:   "CreateEeSubscription",
		Method:  http.MethodPost,
		Path:    "/nudm-ee/v1/{ueIdentity}/ee-subscriptions",
		Handler: OnCreateEeSubscription,
	},
	{
		Label:   "DeleteEeSubscription",
		Method:  http.MethodDelete,
		Path:    "/nudm-ee/v1/{ueIdentity}/ee-subscriptions/{subscriptionId}",
		Handler: OnDeleteEeSubscription,
	},
	{
		Label:   "UpdateEeSubscription",
		Method:  http.MethodPatch,
		Path:    "/nudm-ee/v1/{ueIdentity}/ee-subscriptions/{subscriptionId}",
		Handler: OnUpdateEeSubscription,
	},
}
