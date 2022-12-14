/*
Nsmf_PDUSession

SMF PDU Session Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type SmContextUpdatedData struct {
	UpCnxState UpCnxState `json:"upCnxState,omitempty"`

	HoState HoState `json:"hoState,omitempty"`

	ReleaseEbiList []int32 `json:"releaseEbiList,omitempty"`

	AllocatedEbiList []EbiArpMapping `json:"allocatedEbiList,omitempty"`

	ModifiedEbiList []EbiArpMapping `json:"modifiedEbiList,omitempty"`

	N1SmMsg RefToBinaryData `json:"n1SmMsg,omitempty"`

	N2SmInfo RefToBinaryData `json:"n2SmInfo,omitempty"`

	N2SmInfoType N2SmInfoType `json:"n2SmInfoType,omitempty"`

	EpsBearerSetup []string `json:"epsBearerSetup,omitempty"`

	DataForwarding bool `json:"dataForwarding,omitempty"`

	N3DlForwardingTnlList []IndirectDataForwardingTunnelInfo `json:"n3DlForwardingTnlList,omitempty"`

	N3UlForwardingTnlList []IndirectDataForwardingTunnelInfo `json:"n3UlForwardingTnlList,omitempty"`

	Cause Cause `json:"cause,omitempty"`

	MaAcceptedInd bool `json:"maAcceptedInd,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	ForwardingFTeid string `json:"forwardingFTeid,omitempty"`

	ForwardingBearerContexts []string `json:"forwardingBearerContexts,omitempty"`

	SelectedSmfId string `json:"selectedSmfId,omitempty"`

	SelectedOldSmfId string `json:"selectedOldSmfId,omitempty"`

	AnchorSmfFeatures AnchorSmfFeatures `json:"anchorSmfFeatures,omitempty"`
}
