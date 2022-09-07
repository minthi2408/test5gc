package producer

import (
	"etri5gc/fabric/httpdp"
	"etri5gc/openapi"
	amfprod "etri5gc/openapi/producers/amf"
	"fmt"
	"strings"

	"etri5gc/openapi/models"
	"github.com/free5gc/util/httpwrapper"
	"net/http"
)

func mtService(p *Producer) (service httpdp.HttpService) {
	fn := ginHandler 
	uecontexts := openapi.AMF_MT_UE_CONTEXTS
	service.Routes = httpdp.HttpRoutes{
		{
			"Index",
			"GET",
			"/",
			httpdp.HttpIndexHandler,
		},

		{
			"ProvideDomainSelectionInfo",
			strings.ToUpper("Get"),
			fmt.Sprintf("/%s/:ueContextId", uecontexts),
			fn(amfprod.ProvideDomainSelectionInfo, p),
		},

		{
			"EnableUeReachability",
			strings.ToUpper("Post"),
			fmt.Sprintf("/%s/:ueContextId/ue-reachind", uecontexts),
			fn(amfprod.EnableUeReachability, p),
		},
	}
	service.Group = openapi.AMF_MT
	return
}

func (p *Producer) HandleProvideDomainSelectionInfo(ueContextId, infoclass, supfeatures string) (result models.UeContextInfo, prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleEnableUeReachability(ueContextId string, input *models.EnableUeReachabilityReqData) (result models.EnableUeReachabilityRspData, prob *models.ProblemDetails) {
	return
}

//Old

func (h *Producer) HandleProvideDomainSelectionInfoRequest(request *httpwrapper.Request) *httpwrapper.Response {
	log.Info("Handle Provide Domain Selection Info Request")

	ueContextID := request.Params["ueContextId"]
	infoClassQuery := request.Query.Get("info-class")
	supportedFeaturesQuery := request.Query.Get("supported-features")

	if ueContextInfo, prob := h.doProvideDomainSelectionInfo(ueContextID,
		infoClassQuery, supportedFeaturesQuery); prob != nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else {
		return httpwrapper.NewResponse(http.StatusOK, nil, ueContextInfo)
	}
}

func (h *Producer) doProvideDomainSelectionInfo(ueContextID string, infoClassQuery string, supportedFeaturesQuery string) (
	*models.UeContextInfo, *models.ProblemDetails) {
	amf := h.amf()

	if ue, ok := amf.AmfUeFindByUeContextID(ueContextID); !ok {
		return nil, &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "CONTEXT_NOT_FOUND",
		}
	} else {
		return ue.GetContextInfo(infoClassQuery, supportedFeaturesQuery), nil
	}
}
