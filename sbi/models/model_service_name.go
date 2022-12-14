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
	SERVICENAME_NNRF_NFM                          ServiceName = "nnrf-nfm"
	SERVICENAME_NNRF_DISC                         ServiceName = "nnrf-disc"
	SERVICENAME_NNRF_OAUTH2                       ServiceName = "nnrf-oauth2"
	SERVICENAME_NUDM_SDM                          ServiceName = "nudm-sdm"
	SERVICENAME_NUDM_UECM                         ServiceName = "nudm-uecm"
	SERVICENAME_NUDM_UEAU                         ServiceName = "nudm-ueau"
	SERVICENAME_NUDM_EE                           ServiceName = "nudm-ee"
	SERVICENAME_NUDM_PP                           ServiceName = "nudm-pp"
	SERVICENAME_NUDM_NIDDAU                       ServiceName = "nudm-niddau"
	SERVICENAME_NUDM_MT                           ServiceName = "nudm-mt"
	SERVICENAME_NAMF_COMM                         ServiceName = "namf-comm"
	SERVICENAME_NAMF_EVTS                         ServiceName = "namf-evts"
	SERVICENAME_NAMF_MT                           ServiceName = "namf-mt"
	SERVICENAME_NAMF_LOC                          ServiceName = "namf-loc"
	SERVICENAME_NSMF_PDUSESSION                   ServiceName = "nsmf-pdusession"
	SERVICENAME_NSMF_EVENT_EXPOSURE               ServiceName = "nsmf-event-exposure"
	SERVICENAME_NSMF_NIDD                         ServiceName = "nsmf-nidd"
	SERVICENAME_NAUSF_AUTH                        ServiceName = "nausf-auth"
	SERVICENAME_NAUSF_SORPROTECTION               ServiceName = "nausf-sorprotection"
	SERVICENAME_NAUSF_UPUPROTECTION               ServiceName = "nausf-upuprotection"
	SERVICENAME_NNEF_PFDMANAGEMENT                ServiceName = "nnef-pfdmanagement"
	SERVICENAME_NNEF_SMCONTEXT                    ServiceName = "nnef-smcontext"
	SERVICENAME_NNEF_EVENTEXPOSURE                ServiceName = "nnef-eventexposure"
	SERVICENAME__3GPP_CP_PARAMETER_PROVISIONING   ServiceName = "3gpp-cp-parameter-provisioning"
	SERVICENAME__3GPP_DEVICE_TRIGGERING           ServiceName = "3gpp-device-triggering"
	SERVICENAME__3GPP_BDT                         ServiceName = "3gpp-bdt"
	SERVICENAME__3GPP_TRAFFIC_INFLUENCE           ServiceName = "3gpp-traffic-influence"
	SERVICENAME__3GPP_CHARGEABLE_PARTY            ServiceName = "3gpp-chargeable-party"
	SERVICENAME__3GPP_AS_SESSION_WITH_QOS         ServiceName = "3gpp-as-session-with-qos"
	SERVICENAME__3GPP_MSISDN_LESS_MO_SMS          ServiceName = "3gpp-msisdn-less-mo-sms"
	SERVICENAME__3GPP_SERVICE_PARAMETER           ServiceName = "3gpp-service-parameter"
	SERVICENAME__3GPP_MONITORING_EVENT            ServiceName = "3gpp-monitoring-event"
	SERVICENAME__3GPP_NIDD_CONFIGURATION_TRIGGER  ServiceName = "3gpp-nidd-configuration-trigger"
	SERVICENAME__3GPP_NIDD                        ServiceName = "3gpp-nidd"
	SERVICENAME__3GPP_ANALYTICSEXPOSURE           ServiceName = "3gpp-analyticsexposure"
	SERVICENAME__3GPP_RACS_PARAMETER_PROVISIONING ServiceName = "3gpp-racs-parameter-provisioning"
	SERVICENAME__3GPP_ECR_CONTROL                 ServiceName = "3gpp-ecr-control"
	SERVICENAME__3GPP_APPLYING_BDT_POLICY         ServiceName = "3gpp-applying-bdt-policy"
	SERVICENAME__3GPP_MO_LCS_NOTIFY               ServiceName = "3gpp-mo-lcs-notify"
	SERVICENAME_NPCF_AM_POLICY_CONTROL            ServiceName = "npcf-am-policy-control"
	SERVICENAME_NPCF_SMPOLICYCONTROL              ServiceName = "npcf-smpolicycontrol"
	SERVICENAME_NPCF_POLICYAUTHORIZATION          ServiceName = "npcf-policyauthorization"
	SERVICENAME_NPCF_BDTPOLICYCONTROL             ServiceName = "npcf-bdtpolicycontrol"
	SERVICENAME_NPCF_EVENTEXPOSURE                ServiceName = "npcf-eventexposure"
	SERVICENAME_NPCF_UE_POLICY_CONTROL            ServiceName = "npcf-ue-policy-control"
	SERVICENAME_NSMSF_SMS                         ServiceName = "nsmsf-sms"
	SERVICENAME_NNSSF_NSSELECTION                 ServiceName = "nnssf-nsselection"
	SERVICENAME_NNSSF_NSSAIAVAILABILITY           ServiceName = "nnssf-nssaiavailability"
	SERVICENAME_NUDR_DR                           ServiceName = "nudr-dr"
	SERVICENAME_NUDR_GROUP_ID_MAP                 ServiceName = "nudr-group-id-map"
	SERVICENAME_NLMF_LOC                          ServiceName = "nlmf-loc"
	SERVICENAME_N5G_EIR_EIC                       ServiceName = "n5g-eir-eic"
	SERVICENAME_NBSF_MANAGEMENT                   ServiceName = "nbsf-management"
	SERVICENAME_NCHF_SPENDINGLIMITCONTROL         ServiceName = "nchf-spendinglimitcontrol"
	SERVICENAME_NCHF_CONVERGEDCHARGING            ServiceName = "nchf-convergedcharging"
	SERVICENAME_NCHF_OFFLINEONLYCHARGING          ServiceName = "nchf-offlineonlycharging"
	SERVICENAME_NNWDAF_EVENTSSUBSCRIPTION         ServiceName = "nnwdaf-eventssubscription"
	SERVICENAME_NNWDAF_ANALYTICSINFO              ServiceName = "nnwdaf-analyticsinfo"
	SERVICENAME_NGMLC_LOC                         ServiceName = "ngmlc-loc"
	SERVICENAME_NUCMF_PROVISIONING                ServiceName = "nucmf-provisioning"
	SERVICENAME_NUCMF_UECAPABILITYMANAGEMENT      ServiceName = "nucmf-uecapabilitymanagement"
	SERVICENAME_NHSS_SDM                          ServiceName = "nhss-sdm"
	SERVICENAME_NHSS_UECM                         ServiceName = "nhss-uecm"
	SERVICENAME_NHSS_UEAU                         ServiceName = "nhss-ueau"
	SERVICENAME_NHSS_EE                           ServiceName = "nhss-ee"
	SERVICENAME_NHSS_IMS_SDM                      ServiceName = "nhss-ims-sdm"
	SERVICENAME_NHSS_IMS_UECM                     ServiceName = "nhss-ims-uecm"
	SERVICENAME_NHSS_IMS_UEAU                     ServiceName = "nhss-ims-ueau"
	SERVICENAME_NSEPP_TELESCOPIC                  ServiceName = "nsepp-telescopic"
	SERVICENAME_NSORAF_SOR                        ServiceName = "nsoraf-sor"
	SERVICENAME_NSPAF_SECURED_PACKET              ServiceName = "nspaf-secured-packet"
	SERVICENAME_NUDSF_DR                          ServiceName = "nudsf-dr"
	SERVICENAME_NNSSAAF_NSSAA                     ServiceName = "nnssaaf-nssaa"
)
