/*
Nsmf_PDUSession

SMF PDU Session Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

import (
	"time"
)

type SmContext struct {

	PduSessionId int32 `json:"pduSessionId"`

	Dnn string `json:"dnn"`

	SelectedDnn string `json:"selectedDnn,omitempty"`

	SNssai Snssai `json:"sNssai"`

	HplmnSnssai Snssai `json:"hplmnSnssai,omitempty"`

	PduSessionType PduSessionType `json:"pduSessionType"`

	Gpsi string `json:"gpsi,omitempty"`

	HSmfUri string `json:"hSmfUri,omitempty"`

	SmfUri string `json:"smfUri,omitempty"`

	PduSessionRef string `json:"pduSessionRef,omitempty"`

	PcfId string `json:"pcfId,omitempty"`

	PcfGroupId string `json:"pcfGroupId,omitempty"`

	PcfSetId string `json:"pcfSetId,omitempty"`

	SelMode DnnSelectionMode `json:"selMode,omitempty"`

	UdmGroupId string `json:"udmGroupId,omitempty"`

	RoutingIndicator string `json:"routingIndicator,omitempty"`

	SessionAmbr Ambr `json:"sessionAmbr"`

	QosFlowsList []QosFlowSetupItem `json:"qosFlowsList"`

	HSmfInstanceId string `json:"hSmfInstanceId,omitempty"`

	SmfInstanceId string `json:"smfInstanceId,omitempty"`

	PduSessionSmfSetId string `json:"pduSessionSmfSetId,omitempty"`

	PduSessionSmfServiceSetId string `json:"pduSessionSmfServiceSetId,omitempty"`

	PduSessionSmfBinding SbiBindingLevel `json:"pduSessionSmfBinding,omitempty"`

	EnablePauseCharging bool `json:"enablePauseCharging,omitempty"`

	UeIpv4Address string `json:"ueIpv4Address,omitempty"`

	UeIpv6Prefix Ipv6Prefix `json:"ueIpv6Prefix,omitempty"`

	EpsPdnCnxInfo EpsPdnCnxInfo `json:"epsPdnCnxInfo,omitempty"`

	EpsBearerInfo []EpsBearerInfo `json:"epsBearerInfo,omitempty"`

	MaxIntegrityProtectedDataRate MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRate,omitempty"`

	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateDl,omitempty"`

	AlwaysOnGranted bool `json:"alwaysOnGranted,omitempty"`

	UpSecurity UpSecurity `json:"upSecurity,omitempty"`

	HSmfServiceInstanceId string `json:"hSmfServiceInstanceId,omitempty"`

	SmfServiceInstanceId string `json:"smfServiceInstanceId,omitempty"`

	RecoveryTime time.Time `json:"recoveryTime,omitempty"`

	ForwardingInd bool `json:"forwardingInd,omitempty"`

	PsaTunnelInfo TunnelInfo `json:"psaTunnelInfo,omitempty"`

	ChargingId string `json:"chargingId,omitempty"`

	ChargingInfo ChargingInformation `json:"chargingInfo,omitempty"`

	RoamingChargingProfile RoamingChargingProfile `json:"roamingChargingProfile,omitempty"`

	NefExtBufSupportInd bool `json:"nefExtBufSupportInd,omitempty"`

	Ipv6Index int32 `json:"ipv6Index,omitempty"`

	DnAaaAddress IpAddress `json:"dnAaaAddress,omitempty"`

	RedundantPduSessionInfo RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`

	RanTunnelInfo QosFlowTunnel `json:"ranTunnelInfo,omitempty"`

	AddRanTunnelInfo []QosFlowTunnel `json:"addRanTunnelInfo,omitempty"`

	RedRanTunnelInfo QosFlowTunnel `json:"redRanTunnelInfo,omitempty"`

	AddRedRanTunnelInfo []QosFlowTunnel `json:"addRedRanTunnelInfo,omitempty"`

	DlsetSupportInd bool `json:"dlsetSupportInd,omitempty"`
}
