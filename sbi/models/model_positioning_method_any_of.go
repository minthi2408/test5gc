/*
Namf_Location

AMF Location Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type PositioningMethodAnyOf string

// List of PositioningMethodAnyOf
const (
	POSITIONINGMETHODANYOF_CELLID PositioningMethodAnyOf = "CELLID"
	POSITIONINGMETHODANYOF_ECID PositioningMethodAnyOf = "ECID"
	POSITIONINGMETHODANYOF_OTDOA PositioningMethodAnyOf = "OTDOA"
	POSITIONINGMETHODANYOF_BAROMETRIC_PRESSURE PositioningMethodAnyOf = "BAROMETRIC_PRESSURE"
	POSITIONINGMETHODANYOF_WLAN PositioningMethodAnyOf = "WLAN"
	POSITIONINGMETHODANYOF_BLUETOOTH PositioningMethodAnyOf = "BLUETOOTH"
	POSITIONINGMETHODANYOF_MBS PositioningMethodAnyOf = "MBS"
	POSITIONINGMETHODANYOF_MOTION_SENSOR PositioningMethodAnyOf = "MOTION_SENSOR"
	POSITIONINGMETHODANYOF_DL_TDOA PositioningMethodAnyOf = "DL_TDOA"
	POSITIONINGMETHODANYOF_DL_AOD PositioningMethodAnyOf = "DL_AOD"
	POSITIONINGMETHODANYOF_MULTI_RTT PositioningMethodAnyOf = "MULTI-RTT"
	POSITIONINGMETHODANYOF_NR_ECID PositioningMethodAnyOf = "NR_ECID"
	POSITIONINGMETHODANYOF_UL_TDOA PositioningMethodAnyOf = "UL_TDOA"
	POSITIONINGMETHODANYOF_UL_AOA PositioningMethodAnyOf = "UL_AOA"
	POSITIONINGMETHODANYOF_NETWORK_SPECIFIC PositioningMethodAnyOf = "NETWORK_SPECIFIC"
)