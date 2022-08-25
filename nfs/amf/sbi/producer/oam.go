package producer

import (
	"net/http"
	"strconv"

	"etri5gc/nfs/amf/context"
	"etri5gc/openapi/models"

	"github.com/free5gc/util/httpwrapper"
)

type PduSession struct {
	PduSessionId string
	SmContextRef string
	Sst          string
	Sd           string
	Dnn          string
}

type UEContext struct {
	AccessType models.AccessType
	Supi       string
	Guti       string
	/* Tai */
	Mcc string
	Mnc string
	Tac string
	/* PDU sessions */
	PduSessions []PduSession
	/*Connection state */
	CmState models.CmState
}

type UEContexts []UEContext

func (h *Producer) HandleOAMRegisteredUEContext(request *httpwrapper.Request) *httpwrapper.Response {
	log.Infof("[OAM] Handle Registered UE Context")

	supi := request.Params["supi"]

	if ueContexts, prob := h.doOAMRegisteredUEContext(supi); prob != nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else {
		return httpwrapper.NewResponse(http.StatusOK, nil, ueContexts)
	}
}

func (h *Producer) doOAMRegisteredUEContext(supi string) (UEContexts, *models.ProblemDetails) {
	var ueContexts UEContexts
	amf := h.amf()

	if supi != "" {
		if ue, ok := amf.AmfUeFindBySupi(supi); ok {
			ueContext := buildUEContext(ue, models.AccessType__3_GPP_ACCESS)
			if ueContext != nil {
				ueContexts = append(ueContexts, *ueContext)
			}
			ueContext = buildUEContext(ue, models.AccessType_NON_3_GPP_ACCESS)
			if ueContext != nil {
				ueContexts = append(ueContexts, *ueContext)
			}
		} else {
			return nil, &models.ProblemDetails{
				Status: http.StatusNotFound,
				Cause:  "CONTEXT_NOT_FOUND",
			}
		}
	} else {
		uepool := amf.UePool()
		uepool.Range(func(key, value interface{}) bool {
			ue := value.(*context.AmfUe)
			ueContext := buildUEContext(ue, models.AccessType__3_GPP_ACCESS)
			if ueContext != nil {
				ueContexts = append(ueContexts, *ueContext)
			}
			ueContext = buildUEContext(ue, models.AccessType_NON_3_GPP_ACCESS)
			if ueContext != nil {
				ueContexts = append(ueContexts, *ueContext)
			}
			return true
		})
	}

	return ueContexts, nil
}

//NOTE: (by TungTQ) should this method be implemented at AmfUe class?
// may be not because UEContext is defined here
func buildUEContext(ue *context.AmfUe, accessType models.AccessType) *UEContext {
	if ue.State[accessType].Is(context.Registered) {
		locinfo := ue.GetLocInfo()
		ueContext := &UEContext{
			AccessType: models.AccessType__3_GPP_ACCESS,
			Supi:       ue.Supi,
			Guti:       ue.Guti,
			Mcc:        locinfo.Tai.PlmnId.Mcc,
			Mnc:        locinfo.Tai.PlmnId.Mnc,
			Tac:        locinfo.Tai.Tac,
		}

		ue.SmContextList.Range(func(key, value interface{}) bool {
			smContext := value.(*context.SmContext)
			if smContext.AccessType() == accessType {
				pduSession := PduSession{
					PduSessionId: strconv.Itoa(int(smContext.PduSessionID())),
					SmContextRef: smContext.SmContextRef(),
					Sst:          strconv.Itoa(int(smContext.Snssai().Sst)),
					Sd:           smContext.Snssai().Sd,
					Dnn:          smContext.Dnn(),
				}
				ueContext.PduSessions = append(ueContext.PduSessions, pduSession)
			}
			return true
		})

		if ue.CmConnect(accessType) {
			ueContext.CmState = models.CmState_CONNECTED
		} else {
			ueContext.CmState = models.CmState_IDLE
		}
		return ueContext
	}
	return nil
}
