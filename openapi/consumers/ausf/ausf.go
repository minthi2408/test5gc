package ausf

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
)

type AusfConsumer interface {
	//ue-authentications
	EapAuthMethod(string, *models.EapSession) (models.EapSession, error)
	UeAuthPost(models.AuthenticationInfo) (models.UeAuthenticationCtx,  error)
	UeAuthAuthCtxId5gAkaConfirmationPut(string, string) (models.ConfirmationDataResponse, error)

	//ue-sor
	SupiUeSorPost(string, models.SorInfo) (models.SorSecurityInfo, error)

	//ue-upu
	SupiUeUpuPost(string, models.UpuInfo) (models.UpuSecurityInfo, error)
}

type AusfCallback interface {
}

type ausfConsumerImpl struct {
	client openapi.Consumer
}

func New(c openapi.Consumer) AusfConsumer {
	return &ausfConsumerImpl{
		client: c,
	}
}

type ausfCallbackImpl struct {
}

func Callback() AusfCallback {
	return &ausfCallbackImpl{}
}
