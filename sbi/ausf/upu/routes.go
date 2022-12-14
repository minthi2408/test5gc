/*
Nausf_UPUProtection Service

AUSF UPU Protection Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package upu

import (
	"etrib5gc/sbi"
	"net/http"
)

var Routes = sbi.SbiRoutes{

	{
		Label:   "SupiUeUpuPost",
		Method:  http.MethodPost,
		Path:    "/nausf-upuprotection/v1/{supi}/ue-upu",
		Handler: OnSupiUeUpuPost,
	},
}
