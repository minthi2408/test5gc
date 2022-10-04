/*
Nsmf_PDUSession

SMF PDU Session Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type TriggerTypeAnyOf string

// List of TriggerTypeAnyOf
const (
	TRIGGERTYPEANYOF_QUOTA_THRESHOLD                                  TriggerTypeAnyOf = "QUOTA_THRESHOLD"
	TRIGGERTYPEANYOF_QHT                                              TriggerTypeAnyOf = "QHT"
	TRIGGERTYPEANYOF_FINAL                                            TriggerTypeAnyOf = "FINAL"
	TRIGGERTYPEANYOF_QUOTA_EXHAUSTED                                  TriggerTypeAnyOf = "QUOTA_EXHAUSTED"
	TRIGGERTYPEANYOF_VALIDITY_TIME                                    TriggerTypeAnyOf = "VALIDITY_TIME"
	TRIGGERTYPEANYOF_OTHER_QUOTA_TYPE                                 TriggerTypeAnyOf = "OTHER_QUOTA_TYPE"
	TRIGGERTYPEANYOF_FORCED_REAUTHORISATION                           TriggerTypeAnyOf = "FORCED_REAUTHORISATION"
	TRIGGERTYPEANYOF_UNUSED_QUOTA_TIMER                               TriggerTypeAnyOf = "UNUSED_QUOTA_TIMER"
	TRIGGERTYPEANYOF_UNIT_COUNT_INACTIVITY_TIMER                      TriggerTypeAnyOf = "UNIT_COUNT_INACTIVITY_TIMER"
	TRIGGERTYPEANYOF_ABNORMAL_RELEASE                                 TriggerTypeAnyOf = "ABNORMAL_RELEASE"
	TRIGGERTYPEANYOF_QOS_CHANGE                                       TriggerTypeAnyOf = "QOS_CHANGE"
	TRIGGERTYPEANYOF_VOLUME_LIMIT                                     TriggerTypeAnyOf = "VOLUME_LIMIT"
	TRIGGERTYPEANYOF_TIME_LIMIT                                       TriggerTypeAnyOf = "TIME_LIMIT"
	TRIGGERTYPEANYOF_EVENT_LIMIT                                      TriggerTypeAnyOf = "EVENT_LIMIT"
	TRIGGERTYPEANYOF_PLMN_CHANGE                                      TriggerTypeAnyOf = "PLMN_CHANGE"
	TRIGGERTYPEANYOF_USER_LOCATION_CHANGE                             TriggerTypeAnyOf = "USER_LOCATION_CHANGE"
	TRIGGERTYPEANYOF_RAT_CHANGE                                       TriggerTypeAnyOf = "RAT_CHANGE"
	TRIGGERTYPEANYOF_SESSION_AMBR_CHANGE                              TriggerTypeAnyOf = "SESSION_AMBR_CHANGE"
	TRIGGERTYPEANYOF_UE_TIMEZONE_CHANGE                               TriggerTypeAnyOf = "UE_TIMEZONE_CHANGE"
	TRIGGERTYPEANYOF_TARIFF_TIME_CHANGE                               TriggerTypeAnyOf = "TARIFF_TIME_CHANGE"
	TRIGGERTYPEANYOF_MAX_NUMBER_OF_CHANGES_IN_CHARGING_CONDITIONS     TriggerTypeAnyOf = "MAX_NUMBER_OF_CHANGES_IN_CHARGING_CONDITIONS"
	TRIGGERTYPEANYOF_MANAGEMENT_INTERVENTION                          TriggerTypeAnyOf = "MANAGEMENT_INTERVENTION"
	TRIGGERTYPEANYOF_CHANGE_OF_UE_PRESENCE_IN_PRESENCE_REPORTING_AREA TriggerTypeAnyOf = "CHANGE_OF_UE_PRESENCE_IN_PRESENCE_REPORTING_AREA"
	TRIGGERTYPEANYOF_CHANGE_OF_3_GPP_PS_DATA_OFF_STATUS               TriggerTypeAnyOf = "CHANGE_OF_3GPP_PS_DATA_OFF_STATUS"
	TRIGGERTYPEANYOF_SERVING_NODE_CHANGE                              TriggerTypeAnyOf = "SERVING_NODE_CHANGE"
	TRIGGERTYPEANYOF_REMOVAL_OF_UPF                                   TriggerTypeAnyOf = "REMOVAL_OF_UPF"
	TRIGGERTYPEANYOF_ADDITION_OF_UPF                                  TriggerTypeAnyOf = "ADDITION_OF_UPF"
	TRIGGERTYPEANYOF_INSERTION_OF_ISMF                                TriggerTypeAnyOf = "INSERTION_OF_ISMF"
	TRIGGERTYPEANYOF_REMOVAL_OF_ISMF                                  TriggerTypeAnyOf = "REMOVAL_OF_ISMF"
	TRIGGERTYPEANYOF_CHANGE_OF_ISMF                                   TriggerTypeAnyOf = "CHANGE_OF_ISMF"
	TRIGGERTYPEANYOF_START_OF_SERVICE_DATA_FLOW                       TriggerTypeAnyOf = "START_OF_SERVICE_DATA_FLOW"
	TRIGGERTYPEANYOF_ECGI_CHANGE                                      TriggerTypeAnyOf = "ECGI_CHANGE"
	TRIGGERTYPEANYOF_TAI_CHANGE                                       TriggerTypeAnyOf = "TAI_CHANGE"
	TRIGGERTYPEANYOF_HANDOVER_CANCEL                                  TriggerTypeAnyOf = "HANDOVER_CANCEL"
	TRIGGERTYPEANYOF_HANDOVER_START                                   TriggerTypeAnyOf = "HANDOVER_START"
	TRIGGERTYPEANYOF_HANDOVER_COMPLETE                                TriggerTypeAnyOf = "HANDOVER_COMPLETE"
	TRIGGERTYPEANYOF_GFBR_GUARANTEED_STATUS_CHANGE                    TriggerTypeAnyOf = "GFBR_GUARANTEED_STATUS_CHANGE"
	TRIGGERTYPEANYOF_ADDITION_OF_ACCESS                               TriggerTypeAnyOf = "ADDITION_OF_ACCESS"
	TRIGGERTYPEANYOF_REMOVAL_OF_ACCESS                                TriggerTypeAnyOf = "REMOVAL_OF_ACCESS"
	TRIGGERTYPEANYOF_START_OF_SDF_ADDITIONAL_ACCESS                   TriggerTypeAnyOf = "START_OF_SDF_ADDITIONAL_ACCESS"
)
