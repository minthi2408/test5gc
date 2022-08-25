package ausf
import (
    "net/http"
	"etri5gc/openapi"
	"etri5gc/openapi/models"

//	"github.com/antihax/optional"
)

type AusfConsumer interface {
    //ue-authentications
    EapAuthMethod(string, *models.EapSession) (models.EapSession, *http.Response, error)
    UeAuthPost(models.AuthenticationInfo) (models.UeAuthenticationCtx, *http.Response, error)
    UeAuthAuthCtxId5gAkaConfirmationPut(string, string) (models.ConfirmationDataResponse, *http.Response, error)

}


type ausfImpl struct {
    client  openapi.Consumer
}

func New(c openapi.Consumer) AusfConsumer {
    return &ausfImpl{
        client: c,
    }
}


