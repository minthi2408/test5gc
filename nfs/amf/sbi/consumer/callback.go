package consumer

import (
	"fmt"
//	"bytes"
	"reflect"
	"strconv"
	org_context	"context"
	"etri5gc/nfs/amf/context"
//	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/openapi/models"
//	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/Namf_Communication"
)


type Callback interface {
	SendN2InfoNotifyN2Handover(*context.AmfUe, []int32) error
	SendAmfStatusChangeNotify(string, []models.Guami) 
    SendN1N2TransferFailureNotification(*context.AmfUe, models.N1N2MessageTransferCause)
	SendN1MessageNotify(*context.AmfUe, models.N1MessageClass, []byte, *models.RegistrationContextContainer)
}


type callback struct {
	amf *context.AMFContext
}

func newCallback(c *context.AMFContext) Callback {
	return &callback {
		amf: c,
	}
}

func callback_client() *Namf_Communication.APIClient {
	conf := Namf_Communication.NewConfiguration()
	return Namf_Communication.NewAPIClient(conf)
}

func (c *callback) SendN2InfoNotifyN2Handover(ue *context.AmfUe, releaseList []int32) error {
	if ue.HandoverNotifyUri == "" {
		return fmt.Errorf("N2 Info Notify N2Handover failed(uri dose not exist)")
	}
	client := callback_client()

	n2InformationNotification := models.N2InformationNotification{
		N2NotifySubscriptionId: ue.Supi,
		ToReleaseSessionList:   releaseList,
		NotifyReason:           models.N2InfoNotifyReason_HANDOVER_COMPLETED,
	}

	_, httpResponse, err := client.N2MessageNotifyCallbackDocumentApiServiceCallbackDocumentApi.
		N2InfoNotify(org_context.Background(), ue.HandoverNotifyUri, n2InformationNotification)

	if err == nil {
		// TODO: handle Msg
	} else {
		if httpResponse == nil {
//			HttpLog.Errorln(err.Error())
		} else if err.Error() != httpResponse.Status {
//			HttpLog.Errorln(err.Error())
		}
	}
	return nil
}


func (c callback) SendAmfStatusChangeNotify(amfStatus string, guamiList []models.Guami) {
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
}

func (c *callback) SendN1N2TransferFailureNotification(ue *context.AmfUe, cause models.N1N2MessageTransferCause) {
	if ue.N1N2Message == nil {
		return
	}
	n1n2Message := ue.N1N2Message
	uri := n1n2Message.Request.JsonData.N1n2FailureTxfNotifURI
	if n1n2Message.Status == models.N1N2MessageTransferCause_ATTEMPTING_TO_REACH_UE && uri != "" {
		client := callback_client()
		n1N2MsgTxfrFailureNotification := models.N1N2MsgTxfrFailureNotification{
			Cause:          cause,
			N1n2MsgDataUri: n1n2Message.ResourceUri,
		}

		httpResponse, err := client.N1N2MessageTransferStatusNotificationCallbackDocumentApi.
			N1N2TransferFailureNotification(org_context.Background(), uri, n1N2MsgTxfrFailureNotification)

		if err != nil {
			if httpResponse == nil {
				//HttpLog.Errorln(err.Error())
			} else if err.Error() != httpResponse.Status {
				//HttpLog.Errorln(err.Error())
			}
		} else {
			ue.N1N2Message = nil
		}
	}
}

func (c *callback) SendN1MessageNotify(ue *context.AmfUe, n1class models.N1MessageClass, n1Msg []byte,
	registerContext *models.RegistrationContextContainer) {
	ue.N1N2MessageSubscription.Range(func(key, value interface{}) bool {
		subscriptionID := key.(int64)
		subscription := value.(models.UeN1N2InfoSubscriptionCreateData)

		if subscription.N1NotifyCallbackUri != "" && subscription.N1MessageClass == n1class {
			client := callback_client()

			n1MessageNotify := models.N1MessageNotify{
				JsonData: &models.N1MessageNotification{
					N1NotifySubscriptionId: strconv.Itoa(int(subscriptionID)),
					N1MessageContainer: &models.N1MessageContainer{
						N1MessageClass: subscription.N1MessageClass,
						N1MessageContent: &models.RefToBinaryData{
							ContentId: "n1Msg",
						},
					},
					RegistrationCtxtContainer: registerContext,
				},
				BinaryDataN1Message: n1Msg,
			}
			httpResponse, err := client.N1MessageNotifyCallbackDocumentApiServiceCallbackDocumentApi.
				N1MessageNotify(org_context.Background(), subscription.N1NotifyCallbackUri, n1MessageNotify)
			if err != nil {
				if httpResponse == nil {
					//HttpLog.Errorln(err.Error())
				} else if err.Error() != httpResponse.Status {
					//HttpLog.Errorln(err.Error())
				}
			}
		}
		return true
	})
}

// TS 29.518 5.2.2.3.5.2
func (c *callback) SendN1MessageNotifyAtAMFReAllocation(
	ue *context.AmfUe, n1Msg []byte, registerContext *models.RegistrationContextContainer) {
	client := callback_client()

	n1MessageNotify := models.N1MessageNotify{
		JsonData: &models.N1MessageNotification{
			N1MessageContainer: &models.N1MessageContainer{
				N1MessageClass: models.N1MessageClass__5_GMM,
				N1MessageContent: &models.RefToBinaryData{
					ContentId: "n1Msg",
				},
			},
			RegistrationCtxtContainer: registerContext,
		},
		BinaryDataN1Message: n1Msg,
	}

	var callbackUri string
	for _, subscription := range ue.TargetAmfProfile.DefaultNotificationSubscriptions {
		if subscription.NotificationType == models.NotificationType_N1_MESSAGES &&
			subscription.N1MessageClass == models.N1MessageClass__5_GMM {
			callbackUri = subscription.CallbackUri
			break
		}
	}

	httpResp, err := client.N1MessageNotifyCallbackDocumentApiServiceCallbackDocumentApi.
		N1MessageNotify(org_context.Background(), callbackUri, n1MessageNotify)
	if err != nil {
		if httpResp == nil {
			//HttpLog.Errorln(err.Error())
		} else if err.Error() != httpResp.Status {
			//HttpLog.Errorln(err.Error())
		}
	}
}

func (c *callback) SendN2InfoNotify(ue *context.AmfUe, n2class models.N2InformationClass, n1Msg, n2Msg []byte) {
	ue.N1N2MessageSubscription.Range(func(key, value interface{}) bool {
		subscriptionID := key.(int64)
		subscription := value.(models.UeN1N2InfoSubscriptionCreateData)

		if subscription.N2NotifyCallbackUri != "" && subscription.N2InformationClass == n2class {
			client := callback_client()

			n2InformationNotify := models.N2InfoNotifyRequest{
				JsonData: &models.N2InformationNotification{
					N2NotifySubscriptionId: strconv.Itoa(int(subscriptionID)),
					N2InfoContainer: &models.N2InfoContainer{
						N2InformationClass: n2class,
					},
				},
				BinaryDataN1Message:     n1Msg,
				BinaryDataN2Information: n2Msg,
			}
			if n2Msg == nil {
				//HttpLog.Errorln("Send N2 Info Notify Error(N2 Info does not exist)")
			}
			switch n2class {
			case models.N2InformationClass_SM:
				n2InformationNotify.JsonData.N2InfoContainer.SmInfo = &models.N2SmInformation{
					N2InfoContent: &models.N2InfoContent{
						NgapData: &models.RefToBinaryData{
							ContentId: "n2Info",
						},
					},
				}
			case models.N2InformationClass_NRP_PA:
				n2InformationNotify.JsonData.N2InfoContainer.NrppaInfo = &models.NrppaInformation{
					NrppaPdu: &models.N2InfoContent{
						NgapData: &models.RefToBinaryData{
							ContentId: "n2Info",
						},
					},
				}
			case models.N2InformationClass_PWS, models.N2InformationClass_PWS_BCAL, models.N2InformationClass_PWS_RF:
				n2InformationNotify.JsonData.N2InfoContainer.PwsInfo = &models.PwsInformation{
					PwsContainer: &models.N2InfoContent{
						NgapData: &models.RefToBinaryData{
							ContentId: "n2Info",
						},
					},
				}
			case models.N2InformationClass_RAN:
				n2InformationNotify.JsonData.N2InfoContainer.RanInfo = &models.N2RanInformation{
					N2InfoContent: &models.N2InfoContent{
						NgapData: &models.RefToBinaryData{
							ContentId: "n2Info",
						},
					},
				}
			}

			httpResponse, err := client.N2InfoNotifyCallbackDocumentApiServiceCallbackDocumentApi.
				N2InfoNotify(org_context.Background(), subscription.N2NotifyCallbackUri, n2InformationNotify)
			if err != nil {
				if httpResponse == nil {
					//HttpLog.Errorln(err.Error())
				} else if err.Error() != httpResponse.Status {
					//HttpLog.Errorln(err.Error())
				}
			}
		}
		return true
	})
}
