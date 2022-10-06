package context

import (
	"etrib5gc/sbi/models"
)

type amfCallback struct {
	amf *AmfContext
}

func (c *amfCallback) SendAmfStatusChangeNotify(amfStatus string, guamiList []models.Guami) {
	/*
		subs := c.amf.GetStatusSubList()
		subs.Range(func(key, value interface{}) bool {
			subscriptionData := value.(models.SubscriptionData)
			client := callback_client()
			amfStatusNotification := models.AmfStatusChangeNotification{}
			amfStatusInfo := models.AmfStatusInfo{}

			for _, guami := range guamiList {
				for _, subGumi := range subscriptionData.GuamiList {
					if reflect.DeepEqual(guami, subGumi) {
						// AMF status is available
						amfStatusInfo.GuamiList = append(amfStatusInfo.GuamiList, guami)
					}
				}
			}

			amfStatusInfo = models.AmfStatusInfo{
				StatusChange:     (models.StatusChange)(amfStatus),
				TargetAmfRemoval: "",
				TargetAmfFailure: "",
			}

			amfStatusNotification.AmfStatusInfoList = append(amfStatusNotification.AmfStatusInfoList, amfStatusInfo)
			uri := subscriptionData.AmfStatusUri

			//logger.ProducerLog.Infof("[AMF] Send Amf Status Change Notify to %s", uri)
			httpResponse, err := client.AmfStatusChangeCallbackDocumentApiServiceCallbackDocumentApi.
				AmfStatusChangeNotify(org_context.Background(), uri, amfStatusNotification)
			if err != nil {
				if httpResponse == nil {
					//HttpLog.Errorln(err.Error())
				} else if err.Error() != httpResponse.Status {
					//HttpLog.Errorln(err.Error())
				}
			}
			return true
		})
	*/
}
