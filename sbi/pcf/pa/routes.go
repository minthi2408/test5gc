/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package pa

import (
	"etrib5gc/sbi"
	"net/http"
)

var Routes = sbi.SbiRoutes{

	{
		Label:   "PostAppSessions",
		Method:  http.MethodPost,
		Path:    "/npcf-policyauthorization/v1/app-sessions",
		Handler: OnPostAppSessions,
	},
	{
		Label:   "DeleteEventsSubsc",
		Method:  http.MethodDelete,
		Path:    "/npcf-policyauthorization/v1/app-sessions/{appSessionId}/events-subscription",
		Handler: OnDeleteEventsSubsc,
	},
	{
		Label:   "UpdateEventsSubsc",
		Method:  http.MethodPut,
		Path:    "/npcf-policyauthorization/v1/app-sessions/{appSessionId}/events-subscription",
		Handler: OnUpdateEventsSubsc,
	},
	{
		Label:   "DeleteAppSession",
		Method:  http.MethodPost,
		Path:    "/npcf-policyauthorization/v1/app-sessions/{appSessionId}/delete",
		Handler: OnDeleteAppSession,
	},
	{
		Label:   "GetAppSession",
		Method:  http.MethodGet,
		Path:    "/npcf-policyauthorization/v1/app-sessions/{appSessionId}",
		Handler: OnGetAppSession,
	},
	{
		Label:   "ModAppSession",
		Method:  http.MethodPatch,
		Path:    "/npcf-policyauthorization/v1/app-sessions/{appSessionId}",
		Handler: OnModAppSession,
	},
	{
		Label:   "PcscfRestoration",
		Method:  http.MethodPost,
		Path:    "/npcf-policyauthorization/v1/app-sessions/pcscf-restoration",
		Handler: OnPcscfRestoration,
	},
}