package context
import (
	"net/http"
//	"time"
	"strconv"
	"reflect"

	"github.com/free5gc/openapi/models"
)


func (amf *AMFContext) AMFStatusChangeSubscribe(datreq models.SubscriptionData) (
	datres models.SubscriptionData, locheader string, prob *models.ProblemDetails) {

	for _, guami := range datreq.GuamiList {
		for _, servedGumi := range amf.cfg.Configuration.ServedGuamiList {
			if reflect.DeepEqual(guami, servedGumi) {
				// AMF status is available
				datres.GuamiList = append(datres.GuamiList, guami)
			}
		}
	}

	if datres.GuamiList != nil {
		subId := amf.NewAMFStatusSubscription(datreq)
		locheader = datreq.AmfStatusUri + "/" + subId
		//logger.CommLog.Infof("new AMF Status Subscription[%s]", newSubscriptionID)
	} else {
		prob = &models.ProblemDetails{
			Status: http.StatusForbidden,
			Cause:  "UNSPECIFIED",
		}
	}
	return
}



func (amf *AMFContext) AMFStatusChangeUnSubscribe(subId string) *models.ProblemDetails {

	if _, ok := amf.statussubpool.Load(subId); !ok {
		return &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "SUBSCRIPTION_NOT_FOUND",
		}

	} else {
		amf.statussubpool.Delete(subId)
		if id, err := strconv.ParseInt(subId, 10, 64); err == nil {
			amf.statussubIdGen.FreeID(id)
		}
		return nil
	}
}

func (amf *AMFContext) AMFStatusChangeSubscribeModify(subId string, dat models.SubscriptionData) (*models.SubscriptionData, *models.ProblemDetails) {

	if value, ok := amf.statussubpool.Load(subId); ok {
		curdat := value.(models.SubscriptionData)
		curdat.GuamiList = curdat.GuamiList[:0]
		curdat.GuamiList = append(curdat.GuamiList, dat.GuamiList...)
		curdat.AmfStatusUri = dat.AmfStatusUri

		amf.statussubpool.Store(subId, curdat)
		return &curdat, nil

	} else {
		return nil, &models.ProblemDetails{
			Status: http.StatusForbidden,
			Cause:  "Forbidden",
		}
	}
}
