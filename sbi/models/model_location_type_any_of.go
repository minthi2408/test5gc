/*
Namf_Location

AMF Location Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type LocationTypeAnyOf string

// List of LocationTypeAnyOf
const (
	LOCATIONTYPEANYOF_CURRENT_LOCATION LocationTypeAnyOf = "CURRENT_LOCATION"
	LOCATIONTYPEANYOF_CURRENT_OR_LAST_KNOWN_LOCATION LocationTypeAnyOf = "CURRENT_OR_LAST_KNOWN_LOCATION"
	LOCATIONTYPEANYOF_NOTIFICATION_VERIFICATION_ONLY LocationTypeAnyOf = "NOTIFICATION_VERIFICATION_ONLY"
	LOCATIONTYPEANYOF_DEFERRED_LOCATION LocationTypeAnyOf = "DEFERRED_LOCATION"
)
