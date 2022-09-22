package openapi

// ausf
const (
	AUSF_AUTHENTICATIONS = "/nausf-auth/v1"
	AUSF_SORPROTECTION   = "/nausf-sorprotection/v1"
	AUSF_UPUPROTECTION   = "/nausf-upuprotection/v1"

	AUSF_UE_AUTHENTICATIONS = "ue-authentications"
	AUSF_UE_SORPROTECTION   = "ue-sor"
	AUSF_UE_UPUPROTECTION   = "ue-upu"
)

// amf
const (
	AMF_COMMUNICATION                   = "/namf-comm/v1"
	AMF_COMMUNICATION_SUBSCRIPTIONS     = "subscriptions"
	AMF_COMMUNICATION_UE_CONTEXTS       = "ue-contexts"
	AMF_COMMUNICATION_NONUE_N2_MESSAGES = "non-ue-n2-messages"

	AMF_LOCATION = "/namf-loc/v1"

	AMF_MT             = "/namf-mt/v1"
	AMF_MT_UE_CONTEXTS = "ue-contexts"

	AMF_EVENTEXPOSURE = "/namf-eventexposure/v1"

	AMF_CALLBACK = "/namf-callback/v1"

	AMF_OAM = "/namf-oam/v1"
)

//udm
const (
	UDM_EVENTEXPOSURE    = "/nudm-evenrexposure/v1"
	UDM_EE_SUBSCRIPTIONS = "ee-subscriptions"

	UDM_PARAMETERPROVISION = "/nudm-parameterprovision/v1"
	UDM_PP_DATA            = "pp-data"

	UDM_SUBDATAMANAGEMENT   = "/nudm-subsriber-data-management/v1"
	UDM_SDM_AMDATA          = "am-data"
	UDM_SDM_ID_TRANS_RESULT = "id-translation-result"
	UDM_SDM_SHARED_DATA     = "shared-data"
	UDM_SDM_SM_DATA         = "sm-data"
	UDM_SDM_NSSAI           = "nssai"
	UDM_SDM_SMS_MNG_DATA    = "sms-mng-data"
	UDM_SDM_SMS_DATA        = "sms-data"
	UDM_SDM_SUBSCRIPTIONS   = "sdm-subscriptions"
	UDM_SDM_SHARED_DATA_SUB = "shared-data-subscriptions"
	UDM_SDM_TRACE_DATA      = "trace-data"
	UDM_SDM_UECONTEXT_SMF   = "ue-context-in-smf-data"
	UDM_SDM_UECONTEXT_SMSF  = "ue-context-in-smsf-data"

	UDM_UEAUTHENTICATION     = "/nudm-ueau/v1"
	UDM_UEAUTH_EVENTS        = "auth-events"
	UDM_UEAUTH_GEN_AUTH_DATA = "security-information/generate-auth-data"

	UDM_UECMAN                 = "/nudm-uecm/v1"
	UDM_UECMAN_REGISTRATIONS = "registrations"
	UDM_UECMAN_AMF_3GPP_ACCESS = "amf-3gpp-access"
	UDM_UECMAN_AMF_NON3GPP_ACCESS = "amf-non-3gpp-access"
)
