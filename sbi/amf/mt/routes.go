/*
Namf_MT

AMF Mobile Terminated Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package mt

import (
	"net/http"

	"etrib5gc/sbi"
)

var Routes = sbi.SbiRoutes{

	{
		Label:   "ProvideDomainSelectionInfo",
		Method:  http.MethodGet,
		Path:    "/namf-mt/v1/ue-contexts/{ueContextId}",
		Handler: OnProvideDomainSelectionInfo,
	},
	{
		Label:   "EnableUeReachability",
		Method:  http.MethodPut,
		Path:    "/namf-mt/v1/ue-contexts/{ueContextId}/ue-reachind",
		Handler: OnEnableUeReachability,
	},
}
