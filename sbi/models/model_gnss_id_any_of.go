/*
Namf_Location

AMF Location Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type GnssIdAnyOf string

// List of GnssIdAnyOf
const (
	GNSSIDANYOF_GPS GnssIdAnyOf = "GPS"
	GNSSIDANYOF_GALILEO GnssIdAnyOf = "GALILEO"
	GNSSIDANYOF_SBAS GnssIdAnyOf = "SBAS"
	GNSSIDANYOF_MODERNIZED_GPS GnssIdAnyOf = "MODERNIZED_GPS"
	GNSSIDANYOF_QZSS GnssIdAnyOf = "QZSS"
	GNSSIDANYOF_GLONASS GnssIdAnyOf = "GLONASS"
	GNSSIDANYOF_BDS GnssIdAnyOf = "BDS"
	GNSSIDANYOF_NAVIC GnssIdAnyOf = "NAVIC"
)