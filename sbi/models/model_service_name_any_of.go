/*
Npcf_UEPolicyControl

UE Policy Control Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type ServiceNameAnyOf string

// List of ServiceNameAnyOf
const (
	SERVICENAMEANYOF_NNRF_NFM                          ServiceNameAnyOf = "nnrf-nfm"
	SERVICENAMEANYOF_NNRF_DISC                         ServiceNameAnyOf = "nnrf-disc"
	SERVICENAMEANYOF_NNRF_OAUTH2                       ServiceNameAnyOf = "nnrf-oauth2"
	SERVICENAMEANYOF_NUDM_SDM                          ServiceNameAnyOf = "nudm-sdm"
	SERVICENAMEANYOF_NUDM_UECM                         ServiceNameAnyOf = "nudm-uecm"
	SERVICENAMEANYOF_NUDM_UEAU                         ServiceNameAnyOf = "nudm-ueau"
	SERVICENAMEANYOF_NUDM_EE                           ServiceNameAnyOf = "nudm-ee"
	SERVICENAMEANYOF_NUDM_PP                           ServiceNameAnyOf = "nudm-pp"
	SERVICENAMEANYOF_NUDM_NIDDAU                       ServiceNameAnyOf = "nudm-niddau"
	SERVICENAMEANYOF_NUDM_MT                           ServiceNameAnyOf = "nudm-mt"
	SERVICENAMEANYOF_NAMF_COMM                         ServiceNameAnyOf = "namf-comm"
	SERVICENAMEANYOF_NAMF_EVTS                         ServiceNameAnyOf = "namf-evts"
	SERVICENAMEANYOF_NAMF_MT                           ServiceNameAnyOf = "namf-mt"
	SERVICENAMEANYOF_NAMF_LOC                          ServiceNameAnyOf = "namf-loc"
	SERVICENAMEANYOF_NSMF_PDUSESSION                   ServiceNameAnyOf = "nsmf-pdusession"
	SERVICENAMEANYOF_NSMF_EVENT_EXPOSURE               ServiceNameAnyOf = "nsmf-event-exposure"
	SERVICENAMEANYOF_NSMF_NIDD                         ServiceNameAnyOf = "nsmf-nidd"
	SERVICENAMEANYOF_NAUSF_AUTH                        ServiceNameAnyOf = "nausf-auth"
	SERVICENAMEANYOF_NAUSF_SORPROTECTION               ServiceNameAnyOf = "nausf-sorprotection"
	SERVICENAMEANYOF_NAUSF_UPUPROTECTION               ServiceNameAnyOf = "nausf-upuprotection"
	SERVICENAMEANYOF_NNEF_PFDMANAGEMENT                ServiceNameAnyOf = "nnef-pfdmanagement"
	SERVICENAMEANYOF_NNEF_SMCONTEXT                    ServiceNameAnyOf = "nnef-smcontext"
	SERVICENAMEANYOF_NNEF_EVENTEXPOSURE                ServiceNameAnyOf = "nnef-eventexposure"
	SERVICENAMEANYOF__3GPP_CP_PARAMETER_PROVISIONING   ServiceNameAnyOf = "3gpp-cp-parameter-provisioning"
	SERVICENAMEANYOF__3GPP_DEVICE_TRIGGERING           ServiceNameAnyOf = "3gpp-device-triggering"
	SERVICENAMEANYOF__3GPP_BDT                         ServiceNameAnyOf = "3gpp-bdt"
	SERVICENAMEANYOF__3GPP_TRAFFIC_INFLUENCE           ServiceNameAnyOf = "3gpp-traffic-influence"
	SERVICENAMEANYOF__3GPP_CHARGEABLE_PARTY            ServiceNameAnyOf = "3gpp-chargeable-party"
	SERVICENAMEANYOF__3GPP_AS_SESSION_WITH_QOS         ServiceNameAnyOf = "3gpp-as-session-with-qos"
	SERVICENAMEANYOF__3GPP_MSISDN_LESS_MO_SMS          ServiceNameAnyOf = "3gpp-msisdn-less-mo-sms"
	SERVICENAMEANYOF__3GPP_SERVICE_PARAMETER           ServiceNameAnyOf = "3gpp-service-parameter"
	SERVICENAMEANYOF__3GPP_MONITORING_EVENT            ServiceNameAnyOf = "3gpp-monitoring-event"
	SERVICENAMEANYOF__3GPP_NIDD_CONFIGURATION_TRIGGER  ServiceNameAnyOf = "3gpp-nidd-configuration-trigger"
	SERVICENAMEANYOF__3GPP_NIDD                        ServiceNameAnyOf = "3gpp-nidd"
	SERVICENAMEANYOF__3GPP_ANALYTICSEXPOSURE           ServiceNameAnyOf = "3gpp-analyticsexposure"
	SERVICENAMEANYOF__3GPP_RACS_PARAMETER_PROVISIONING ServiceNameAnyOf = "3gpp-racs-parameter-provisioning"
	SERVICENAMEANYOF__3GPP_ECR_CONTROL                 ServiceNameAnyOf = "3gpp-ecr-control"
	SERVICENAMEANYOF__3GPP_APPLYING_BDT_POLICY         ServiceNameAnyOf = "3gpp-applying-bdt-policy"
	SERVICENAMEANYOF__3GPP_MO_LCS_NOTIFY               ServiceNameAnyOf = "3gpp-mo-lcs-notify"
	SERVICENAMEANYOF_NPCF_AM_POLICY_CONTROL            ServiceNameAnyOf = "npcf-am-policy-control"
	SERVICENAMEANYOF_NPCF_SMPOLICYCONTROL              ServiceNameAnyOf = "npcf-smpolicycontrol"
	SERVICENAMEANYOF_NPCF_POLICYAUTHORIZATION          ServiceNameAnyOf = "npcf-policyauthorization"
	SERVICENAMEANYOF_NPCF_BDTPOLICYCONTROL             ServiceNameAnyOf = "npcf-bdtpolicycontrol"
	SERVICENAMEANYOF_NPCF_EVENTEXPOSURE                ServiceNameAnyOf = "npcf-eventexposure"
	SERVICENAMEANYOF_NPCF_UE_POLICY_CONTROL            ServiceNameAnyOf = "npcf-ue-policy-control"
	SERVICENAMEANYOF_NSMSF_SMS                         ServiceNameAnyOf = "nsmsf-sms"
	SERVICENAMEANYOF_NNSSF_NSSELECTION                 ServiceNameAnyOf = "nnssf-nsselection"
	SERVICENAMEANYOF_NNSSF_NSSAIAVAILABILITY           ServiceNameAnyOf = "nnssf-nssaiavailability"
	SERVICENAMEANYOF_NUDR_DR                           ServiceNameAnyOf = "nudr-dr"
	SERVICENAMEANYOF_NUDR_GROUP_ID_MAP                 ServiceNameAnyOf = "nudr-group-id-map"
	SERVICENAMEANYOF_NLMF_LOC                          ServiceNameAnyOf = "nlmf-loc"
	SERVICENAMEANYOF_N5G_EIR_EIC                       ServiceNameAnyOf = "n5g-eir-eic"
	SERVICENAMEANYOF_NBSF_MANAGEMENT                   ServiceNameAnyOf = "nbsf-management"
	SERVICENAMEANYOF_NCHF_SPENDINGLIMITCONTROL         ServiceNameAnyOf = "nchf-spendinglimitcontrol"
	SERVICENAMEANYOF_NCHF_CONVERGEDCHARGING            ServiceNameAnyOf = "nchf-convergedcharging"
	SERVICENAMEANYOF_NCHF_OFFLINEONLYCHARGING          ServiceNameAnyOf = "nchf-offlineonlycharging"
	SERVICENAMEANYOF_NNWDAF_EVENTSSUBSCRIPTION         ServiceNameAnyOf = "nnwdaf-eventssubscription"
	SERVICENAMEANYOF_NNWDAF_ANALYTICSINFO              ServiceNameAnyOf = "nnwdaf-analyticsinfo"
	SERVICENAMEANYOF_NGMLC_LOC                         ServiceNameAnyOf = "ngmlc-loc"
	SERVICENAMEANYOF_NUCMF_PROVISIONING                ServiceNameAnyOf = "nucmf-provisioning"
	SERVICENAMEANYOF_NUCMF_UECAPABILITYMANAGEMENT      ServiceNameAnyOf = "nucmf-uecapabilitymanagement"
	SERVICENAMEANYOF_NHSS_SDM                          ServiceNameAnyOf = "nhss-sdm"
	SERVICENAMEANYOF_NHSS_UECM                         ServiceNameAnyOf = "nhss-uecm"
	SERVICENAMEANYOF_NHSS_UEAU                         ServiceNameAnyOf = "nhss-ueau"
	SERVICENAMEANYOF_NHSS_EE                           ServiceNameAnyOf = "nhss-ee"
	SERVICENAMEANYOF_NHSS_IMS_SDM                      ServiceNameAnyOf = "nhss-ims-sdm"
	SERVICENAMEANYOF_NHSS_IMS_UECM                     ServiceNameAnyOf = "nhss-ims-uecm"
	SERVICENAMEANYOF_NHSS_IMS_UEAU                     ServiceNameAnyOf = "nhss-ims-ueau"
	SERVICENAMEANYOF_NSEPP_TELESCOPIC                  ServiceNameAnyOf = "nsepp-telescopic"
	SERVICENAMEANYOF_NSORAF_SOR                        ServiceNameAnyOf = "nsoraf-sor"
	SERVICENAMEANYOF_NSPAF_SECURED_PACKET              ServiceNameAnyOf = "nspaf-secured-packet"
	SERVICENAMEANYOF_NUDSF_DR                          ServiceNameAnyOf = "nudsf-dr"
	SERVICENAMEANYOF_NNSSAAF_NSSAA                     ServiceNameAnyOf = "nnssaaf-nssaa"
)
