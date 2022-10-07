/*
Npcf_UEPolicyControl

UE Policy Control Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type ServiceName string

// List of ServiceName
const (
	SERVICENAMEANYOF_NNRF_NFM                          ServiceName = "nnrf-nfm"
	SERVICENAMEANYOF_NNRF_DISC                         ServiceName = "nnrf-disc"
	SERVICENAMEANYOF_NNRF_OAUTH2                       ServiceName = "nnrf-oauth2"
	SERVICENAMEANYOF_NUDM_SDM                          ServiceName = "nudm-sdm"
	SERVICENAMEANYOF_NUDM_UECM                         ServiceName = "nudm-uecm"
	SERVICENAMEANYOF_NUDM_UEAU                         ServiceName = "nudm-ueau"
	SERVICENAMEANYOF_NUDM_EE                           ServiceName = "nudm-ee"
	SERVICENAMEANYOF_NUDM_PP                           ServiceName = "nudm-pp"
	SERVICENAMEANYOF_NUDM_NIDDAU                       ServiceName = "nudm-niddau"
	SERVICENAMEANYOF_NUDM_MT                           ServiceName = "nudm-mt"
	SERVICENAMEANYOF_NAMF_COMM                         ServiceName = "namf-comm"
	SERVICENAMEANYOF_NAMF_EVTS                         ServiceName = "namf-evts"
	SERVICENAMEANYOF_NAMF_MT                           ServiceName = "namf-mt"
	SERVICENAMEANYOF_NAMF_LOC                          ServiceName = "namf-loc"
	SERVICENAMEANYOF_NSMF_PDUSESSION                   ServiceName = "nsmf-pdusession"
	SERVICENAMEANYOF_NSMF_EVENT_EXPOSURE               ServiceName = "nsmf-event-exposure"
	SERVICENAMEANYOF_NSMF_NIDD                         ServiceName = "nsmf-nidd"
	SERVICENAMEANYOF_NAUSF_AUTH                        ServiceName = "nausf-auth"
	SERVICENAMEANYOF_NAUSF_SORPROTECTION               ServiceName = "nausf-sorprotection"
	SERVICENAMEANYOF_NAUSF_UPUPROTECTION               ServiceName = "nausf-upuprotection"
	SERVICENAMEANYOF_NNEF_PFDMANAGEMENT                ServiceName = "nnef-pfdmanagement"
	SERVICENAMEANYOF_NNEF_SMCONTEXT                    ServiceName = "nnef-smcontext"
	SERVICENAMEANYOF_NNEF_EVENTEXPOSURE                ServiceName = "nnef-eventexposure"
	SERVICENAMEANYOF__3GPP_CP_PARAMETER_PROVISIONING   ServiceName = "3gpp-cp-parameter-provisioning"
	SERVICENAMEANYOF__3GPP_DEVICE_TRIGGERING           ServiceName = "3gpp-device-triggering"
	SERVICENAMEANYOF__3GPP_BDT                         ServiceName = "3gpp-bdt"
	SERVICENAMEANYOF__3GPP_TRAFFIC_INFLUENCE           ServiceName = "3gpp-traffic-influence"
	SERVICENAMEANYOF__3GPP_CHARGEABLE_PARTY            ServiceName = "3gpp-chargeable-party"
	SERVICENAMEANYOF__3GPP_AS_SESSION_WITH_QOS         ServiceName = "3gpp-as-session-with-qos"
	SERVICENAMEANYOF__3GPP_MSISDN_LESS_MO_SMS          ServiceName = "3gpp-msisdn-less-mo-sms"
	SERVICENAMEANYOF__3GPP_SERVICE_PARAMETER           ServiceName = "3gpp-service-parameter"
	SERVICENAMEANYOF__3GPP_MONITORING_EVENT            ServiceName = "3gpp-monitoring-event"
	SERVICENAMEANYOF__3GPP_NIDD_CONFIGURATION_TRIGGER  ServiceName = "3gpp-nidd-configuration-trigger"
	SERVICENAMEANYOF__3GPP_NIDD                        ServiceName = "3gpp-nidd"
	SERVICENAMEANYOF__3GPP_ANALYTICSEXPOSURE           ServiceName = "3gpp-analyticsexposure"
	SERVICENAMEANYOF__3GPP_RACS_PARAMETER_PROVISIONING ServiceName = "3gpp-racs-parameter-provisioning"
	SERVICENAMEANYOF__3GPP_ECR_CONTROL                 ServiceName = "3gpp-ecr-control"
	SERVICENAMEANYOF__3GPP_APPLYING_BDT_POLICY         ServiceName = "3gpp-applying-bdt-policy"
	SERVICENAMEANYOF__3GPP_MO_LCS_NOTIFY               ServiceName = "3gpp-mo-lcs-notify"
	SERVICENAMEANYOF_NPCF_AM_POLICY_CONTROL            ServiceName = "npcf-am-policy-control"
	SERVICENAMEANYOF_NPCF_SMPOLICYCONTROL              ServiceName = "npcf-smpolicycontrol"
	SERVICENAMEANYOF_NPCF_POLICYAUTHORIZATION          ServiceName = "npcf-policyauthorization"
	SERVICENAMEANYOF_NPCF_BDTPOLICYCONTROL             ServiceName = "npcf-bdtpolicycontrol"
	SERVICENAMEANYOF_NPCF_EVENTEXPOSURE                ServiceName = "npcf-eventexposure"
	SERVICENAMEANYOF_NPCF_UE_POLICY_CONTROL            ServiceName = "npcf-ue-policy-control"
	SERVICENAMEANYOF_NSMSF_SMS                         ServiceName = "nsmsf-sms"
	SERVICENAMEANYOF_NNSSF_NSSELECTION                 ServiceName = "nnssf-nsselection"
	SERVICENAMEANYOF_NNSSF_NSSAIAVAILABILITY           ServiceName = "nnssf-nssaiavailability"
	SERVICENAMEANYOF_NUDR_DR                           ServiceName = "nudr-dr"
	SERVICENAMEANYOF_NUDR_GROUP_ID_MAP                 ServiceName = "nudr-group-id-map"
	SERVICENAMEANYOF_NLMF_LOC                          ServiceName = "nlmf-loc"
	SERVICENAMEANYOF_N5G_EIR_EIC                       ServiceName = "n5g-eir-eic"
	SERVICENAMEANYOF_NBSF_MANAGEMENT                   ServiceName = "nbsf-management"
	SERVICENAMEANYOF_NCHF_SPENDINGLIMITCONTROL         ServiceName = "nchf-spendinglimitcontrol"
	SERVICENAMEANYOF_NCHF_CONVERGEDCHARGING            ServiceName = "nchf-convergedcharging"
	SERVICENAMEANYOF_NCHF_OFFLINEONLYCHARGING          ServiceName = "nchf-offlineonlycharging"
	SERVICENAMEANYOF_NNWDAF_EVENTSSUBSCRIPTION         ServiceName = "nnwdaf-eventssubscription"
	SERVICENAMEANYOF_NNWDAF_ANALYTICSINFO              ServiceName = "nnwdaf-analyticsinfo"
	SERVICENAMEANYOF_NGMLC_LOC                         ServiceName = "ngmlc-loc"
	SERVICENAMEANYOF_NUCMF_PROVISIONING                ServiceName = "nucmf-provisioning"
	SERVICENAMEANYOF_NUCMF_UECAPABILITYMANAGEMENT      ServiceName = "nucmf-uecapabilitymanagement"
	SERVICENAMEANYOF_NHSS_SDM                          ServiceName = "nhss-sdm"
	SERVICENAMEANYOF_NHSS_UECM                         ServiceName = "nhss-uecm"
	SERVICENAMEANYOF_NHSS_UEAU                         ServiceName = "nhss-ueau"
	SERVICENAMEANYOF_NHSS_EE                           ServiceName = "nhss-ee"
	SERVICENAMEANYOF_NHSS_IMS_SDM                      ServiceName = "nhss-ims-sdm"
	SERVICENAMEANYOF_NHSS_IMS_UECM                     ServiceName = "nhss-ims-uecm"
	SERVICENAMEANYOF_NHSS_IMS_UEAU                     ServiceName = "nhss-ims-ueau"
	SERVICENAMEANYOF_NSEPP_TELESCOPIC                  ServiceName = "nsepp-telescopic"
	SERVICENAMEANYOF_NSORAF_SOR                        ServiceName = "nsoraf-sor"
	SERVICENAMEANYOF_NSPAF_SECURED_PACKET              ServiceName = "nspaf-secured-packet"
	SERVICENAMEANYOF_NUDSF_DR                          ServiceName = "nudsf-dr"
	SERVICENAMEANYOF_NNSSAAF_NSSAA                     ServiceName = "nnssaaf-nssaa"
)
