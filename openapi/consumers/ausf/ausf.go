package ausf

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	"net/http"
	//	"github.com/antihax/optional"
)

type AusfConsumer interface {
	//ue-authentications
	EapAuthMethod(string, *models.EapSession) (models.EapSession, *http.Response, error)
	UeAuthPost(models.AuthenticationInfo) (models.UeAuthenticationCtx, *http.Response, error)
	UeAuthAuthCtxId5gAkaConfirmationPut(string, string) (models.ConfirmationDataResponse, *http.Response, error)

	//ue-sor
	SupiUeSorPost(string, models.SorInfo) (models.SorSecurityInfo, *http.Response, error)

	//ue-upu
	SupiUeUpuPost(string, models.UpuInfo) (models.UpuSecurityInfo, *http.Response, error)
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
