/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

import (
	"time"
)

type TrafficInfluData struct {

	// Contains the Notification Correlation Id allocated by the NEF for the UP path change notification.
	UpPathChgNotifCorreId string `json:"upPathChgNotifCorreId,omitempty"`

	// Identifies whether an application can be relocated once a location of the application has been selected.
	AppReloInd bool `json:"appReloInd,omitempty"`

	// Identifies an application.
	AfAppId string `json:"afAppId,omitempty"`

	Dnn string `json:"dnn,omitempty"`

	// Identifies Ethernet packet filters. Either \"trafficFilters\" or \"ethTrafficFilters\" shall be included if applicable.
	EthTrafficFilters []EthFlowDescription `json:"ethTrafficFilters,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	InterGroupId string `json:"interGroupId,omitempty"`

	Supi string `json:"supi,omitempty"`

	// Identifies IP packet filters. Either \"trafficFilters\" or \"ethTrafficFilters\" shall be included if applicable.
	TrafficFilters []FlowInfo `json:"trafficFilters,omitempty"`

	// Identifies the N6 traffic routing requirement.
	TrafficRoutes []RouteToLocation `json:"trafficRoutes,omitempty"`

	TraffCorreInd bool `json:"traffCorreInd,omitempty"`

	ValidStartTime time.Time `json:"validStartTime,omitempty"`

	ValidEndTime time.Time `json:"validEndTime,omitempty"`

	// Identifies the temporal validities for the N6 traffic routing requirement.
	TempValidities []TemporalValidity `json:"tempValidities,omitempty"`

	NwAreaInfo NetworkAreaInfo1 `json:"nwAreaInfo,omitempty"`

	UpPathChgNotifUri string `json:"upPathChgNotifUri,omitempty"`

	Headers []string `json:"headers,omitempty"`

	SubscribedEvents []SubscribedEvent `json:"subscribedEvents,omitempty"`

	DnaiChgType DnaiChangeType `json:"dnaiChgType,omitempty"`

	AfAckInd bool `json:"afAckInd,omitempty"`

	AddrPreserInd bool `json:"addrPreserInd,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	ResUri string `json:"resUri,omitempty"`
}
