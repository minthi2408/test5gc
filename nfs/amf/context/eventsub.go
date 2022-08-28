package context

import (
	"etri5gc/openapi/models"
	"net/http"
	"strconv"
	"time"
	//	"github.com/free5gc/util/idgenerator"
)

type AmfContextEventSubscription struct {
	IsAnyUe           bool
	IsGroupUe         bool
	UeSupiList        []string
	Expiry            *time.Time
	EventSubscription models.AmfEventSubscription
}

type AmfUeEventSubscription struct {
	Timestamp         time.Time
	AnyUe             bool
	RemainReports     *int32
	EventSubscription *models.AmfEventSubscription
}

//TungTQ: move event subscription handling methods from the `producer` to `context`

// TODO (free5gc): handle event filter
func (amf *AmfContext) CreateAMFEventSub(reqsub models.AmfCreateEventSubscription) (*models.AmfCreatedEventSubscription, *models.ProblemDetails) {
	/*
		TODO: tungtq - it seems current implementation is a bit messy, let clean it up later

		ressub := &models.AmfCreatedEventSubscription{}

		subscription := reqsub.Subscription

		eventsub := &AmfContextEventSubscription{
			EventSubscription: *subscription, //copy (not a pointer)
		}

		var isImmediate bool
		var immediateFlags []bool
		var reportlist []models.AmfEventReport

		// generate even id
		id, err := amf.eventsubIdGen.Allocate()
		if err != nil {
			problemDetails := &models.ProblemDetails{
				Status: http.StatusInternalServerError,
				Cause:  "UNSPECIFIED_NF_FAILURE",
			}
			return nil, problemDetails
		}
		subId := strconv.Itoa(int(id))

		// store subscription in context
		ueEventSubscription := AmfUeEventSubscription{
			EventSubscriptioni:  &eventsub.EventSubscription,
			Timestamp:  time.Now().UTC(),
		}

		if subscription.Options != nil && subscription.Options.Trigger == models.AmfEventTrigger_CONTINUOUS {
			ueEventSubscription.RemainReports = new(int32)
			*ueEventSubscription.RemainReports = subscription.Options.MaxReports
		}
		for _, events := range *subscription.EventList {
			immediateFlags = append(immediateFlags, events.ImmediateFlag)
			if events.ImmediateFlag {
				isImmediate = true
			}
		}

		if subscription.AnyUE {
			eventsub.IsAnyUe = true
			ueEventSubscription.AnyUe = true
			amf.uepool.Range(func(key, value interface{}) bool {
				ue := value.(*AmfUe)
				ue.EventSubscriptionsInfo[subID] = new(AmfUeEventSubscription)
				*ue.EventSubscriptionsInfo[subId] = ueEventSubscription
				eventsub.UeSupiList = append(eventsub.UeSupiList, ue.Supi)
				return true
			})
		} else if subscription.GroupId != "" {
			contextEventSubscription.IsGroupUe = true
			ueEventSubscription.AnyUe = true
			amfSelf.UePool.Range(func(key, value interface{}) bool {
				ue := value.(*context.AmfUe)
				if ue.GroupID == subscription.GroupId {
					ue.EventSubscriptionsInfo[newSubscriptionID] = new(context.AmfUeEventSubscription)
					*ue.EventSubscriptionsInfo[newSubscriptionID] = ueEventSubscription
					contextEventSubscription.UeSupiList = append(contextEventSubscription.UeSupiList, ue.Supi)
				}
				return true
			})
		} else {
			if ue, ok := amfSelf.AmfUeFindBySupi(subscription.Supi); !ok {
				problemDetails := &models.ProblemDetails{
					Status: http.StatusForbidden,
					Cause:  "UE_NOT_SERVED_BY_AMF",
				}
				return nil, problemDetails
			} else {
				ue.EventSubscriptionsInfo[newSubscriptionID] = new(context.AmfUeEventSubscription)
				*ue.EventSubscriptionsInfo[newSubscriptionID] = ueEventSubscription
				contextEventSubscription.UeSupiList = append(contextEventSubscription.UeSupiList, ue.Supi)
			}
		}

		// delete subscription
		if subscription.Options != nil {
			eventsub.Expiry = subscription.Options.Expiry
		}

		//add to event pool
		amf.eventsubpool.Store(subId, subscription)

		// build response

		ressub.Subscription = subscription
		ressub.SubscriptionId = subId

		// for immediate use
		if subscription.AnyUE {
			amf.UePool.Range(func(key, value interface{}) bool {
				ue := value.(*context.AmfUe)
				if isImmediate {
					subReports(ue, newSubscriptionID)
				}
				for i, flag := range immediateFlags {
					if flag {
						report, ok := NewAmfEventReport(ue, (*subscription.EventList)[i].Type, subId)
						if ok {
							reportlist = append(reportlist, report)
						}
					}
				}
				// delete subscription
				if reportlistLen := len(reportlist); reportlistLen > 0 && (!reportlist[reportlistLen-1].State.Active) {
					delete(ue.EventSubscriptionsInfo, newSubscriptionID)
				}
				return true
			})
		} else if subscription.GroupId != "" {
			amf.UePool.Range(func(key, value interface{}) bool {
				ue := value.(*context.AmfUe)
				if isImmediate {
					subReports(ue, newSubscriptionID)
				}
				if ue.GroupID == subscription.GroupId {
					for i, flag := range immediateFlags {
						if flag {
							report, ok := NewAmfEventReport(ue, (*subscription.EventList)[i].Type, newSubscriptionID)
							if ok {
								reportlist = append(reportlist, report)
							}
						}
					}
					// delete subscription
					if reportlistLen := len(reportlist); reportlistLen > 0 && (!reportlist[reportlistLen-1].State.Active) {
						delete(ue.EventSubscriptionsInfo, newSubscriptionID)
					}
				}
				return true
			})
		} else {
			ue, _ := amfSelf.AmfUeFindBySupi(subscription.Supi)
			if isImmediate {
				subReports(ue, newSubscriptionID)
			}
			for i, flag := range immediateFlags {
				if flag {
					report, ok := NewAmfEventReport(ue, (*subscription.EventList)[i].Type, newSubscriptionID)
					if ok {
						reportlist = append(reportlist, report)
					}
				}
			}
			// delete subscription
			if reportlistLen := len(reportlist); reportlistLen > 0 && (!reportlist[reportlistLen-1].State.Active) {
				delete(ue.EventSubscriptionsInfo, newSubscriptionID)
			}
		}
		if len(reportlist) > 0 {
			ressub.ReportList = reportlist
			// delete subscription
			if !reportlist[0].State.Active {
				amf.DeleteEventSubscription(newSubscriptionID)
			}
		}

		return ressub, nil
	*/
	return nil, nil
}

func (amf *AmfContext) DeleteAMFEventSub(subId string) *models.ProblemDetails {
	if value, ok := amf.eventsubpool.Load(subId); !ok {
		return &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "SUBSCRIPTION_NOT_FOUND",
		}

	} else {
		sub, _ := value.(*AmfContextEventSubscription)
		for _, supi := range sub.UeSupiList {
			//find AmfUe to delete subscription
			if uev, ok := amf.uepool.Load(supi); ok {
				ue, _ := uev.(*AmfUe)
				ue.DeleteEventSub(subId)
			}
		}
		amf.eventsubpool.Delete(subId)
		//free Id
		if id, err := strconv.ParseInt(subId, 10, 32); err != nil {
			//logger.ContextLog.Error(err)
		} else {
			amf.eventsubIdGen.FreeID(id)
		}

		return nil
	}
}

func (amf *AmfContext) ModifyAMFEventSub(subId string, modreq models.ModifySubscriptionRequest) (*models.AmfUpdatedEventSubscription, *models.ProblemDetails) {

	if value, ok := amf.eventsubpool.Load(subId); !ok {
		return nil, &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "SUBSCRIPTION_NOT_FOUND",
		}
	} else {
		cursub, _ := value.(*AmfContextEventSubscription)

		if modreq.OptionItem != nil {
			cursub.Expiry = modreq.OptionItem.Value
		} else if modreq.SubscriptionItemInner != nil {
			subscription := &cursub.EventSubscription
			if !cursub.IsAnyUe && !cursub.IsGroupUe {
				if _, ok := amf.uepool.Load(subscription.Supi); !ok {
					//ue is not found
					return nil, &models.ProblemDetails{
						Status: http.StatusForbidden,
						Cause:  "UE_NOT_SERVED_BY_AMF",
					}
				}
			}

			op := modreq.SubscriptionItemInner.Op
			index, err := strconv.Atoi(modreq.SubscriptionItemInner.Path[11:])
			if err != nil {
				return nil, &models.ProblemDetails{
					Status: http.StatusInternalServerError,
					Cause:  "UNSPECIFIED_NF_FAILURE",
				}
			}
			lists := (*subscription.EventList)
			eventlistLen := len(*subscription.EventList)
			switch op {
			case "replace":
				event := *modreq.SubscriptionItemInner.Value
				if index < eventlistLen {
					(*subscription.EventList)[index] = event
				}
			case "remove":
				if index < eventlistLen {
					*subscription.EventList = append(lists[:index], lists[index+1:]...)
				}
			case "add":
				event := *modreq.SubscriptionItemInner.Value
				*subscription.EventList = append(lists, event)
			}
		}

		upsub := &models.AmfUpdatedEventSubscription{
			Subscription: &cursub.EventSubscription,
		}
		return upsub, nil
	}
}

func (context *AmfContext) FindEventSubscription(subscriptionID string) (*AmfContextEventSubscription, bool) {
	if value, ok := context.eventsubpool.Load(subscriptionID); ok {
		return value.(*AmfContextEventSubscription), ok
	} else {
		return nil, false
	}
}

///TungTQ: AfmUe methods for handling event subscription

func (ue *AmfUe) DeleteEventSub(subId string) {
	delete(ue.EventSubscriptionsInfo, subId)
}

/*
func subReports(ue *context.AmfUe, subscriptionId string) {
	remainReport := ue.EventSubscriptionsInfo[subscriptionId].RemainReports
	if remainReport == nil {
		return
	}
	*remainReport--
}

// DO NOT handle AmfEventType_PRESENCE_IN_AOI_REPORT and AmfEventType_UES_IN_AREA_REPORT(about area)
func NewAmfEventReport(ue *context.AmfUe, Type models.AmfEventType, subscriptionId string) (
	report models.AmfEventReport, ok bool) {
	ueSubscription, ok := ue.EventSubscriptionsInfo[subscriptionId]
	if !ok {
		return report, ok
	}

	report.AnyUe = ueSubscription.AnyUe
	report.Supi = ue.Supi
	report.Type = Type
	report.TimeStamp = &ueSubscription.Timestamp
	report.State = new(models.AmfEventState)
	mode := ueSubscription.EventSubscription.Options
	if mode == nil {
		report.State.Active = true
	} else if mode.Trigger == models.AmfEventTrigger_ONE_TIME {
		report.State.Active = false
	} else if *ueSubscription.RemainReports <= 0 {
		report.State.Active = false
	} else {
		report.State.Active = getDuration(mode.Expiry, &report.State.RemainDuration)
		if report.State.Active {
			report.State.RemainReports = *ueSubscription.RemainReports
		}
	}

	switch Type {
	case models.AmfEventType_LOCATION_REPORT:
		report.Location = &ue.Location
	// case models.AmfEventType_PRESENCE_IN_AOI_REPORT:
	// report.AreaList = (*subscription.EventList)[eventIndex].AreaList
	case models.AmfEventType_TIMEZONE_REPORT:
		report.Timezone = ue.TimeZone
	case models.AmfEventType_ACCESS_TYPE_REPORT:
		for accessType, state := range ue.State {
			if state.Is(context.Registered) {
				report.AccessTypeList = append(report.AccessTypeList, accessType)
			}
		}
	case models.AmfEventType_REGISTRATION_STATE_REPORT:
		var rmInfos []models.RmInfo
		for accessType, state := range ue.State {
			rmInfo := models.RmInfo{
				RmState:    models.RmState_DEREGISTERED,
				AccessType: accessType,
			}
			if state.Is(context.Registered) {
				rmInfo.RmState = models.RmState_REGISTERED
			}
			rmInfos = append(rmInfos, rmInfo)
		}
		report.RmInfoList = rmInfos
	case models.AmfEventType_CONNECTIVITY_STATE_REPORT:
		report.CmInfoList = ue.GetCmInfo()
	case models.AmfEventType_REACHABILITY_REPORT:
		report.Reachability = ue.Reachability
	case models.AmfEventType_SUBSCRIBED_DATA_REPORT:
		report.SubscribedData = &ue.SubscribedData
	case models.AmfEventType_COMMUNICATION_FAILURE_REPORT:
		// TODO : report.CommFailure
	case models.AmfEventType_SUBSCRIPTION_ID_CHANGE:
		report.SubscriptionId = subscriptionId
	case models.AmfEventType_SUBSCRIPTION_ID_ADDITION:
		report.SubscriptionId = subscriptionId
	}
	return report, ok
}

*/
