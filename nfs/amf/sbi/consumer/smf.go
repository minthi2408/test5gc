package consumer

import (
	org_context "context"
	"etri5gc/nfs/amf/context"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/Nsmf_PDUSession"
)

type smfConsumer struct {
}


func newSmfConsumer(context *context.AMFContext) *smfConsumer {
	return &smfConsumer{}
}

func pdusession_client(sm *context.SmContext) (*Nsmf_PDUSession.APIClient) {
	conf := Nsmf_PDUSession.NewConfiguration()
	conf.SetBasePath(sm.SmfUri())
	return Nsmf_PDUSession.NewAPIClient(conf)
}

func (c *smfConsumer) SendCreateSmContextRequest(ue *context.AmfUe, sm *context.SmContext, rtype *models.RequestType, naspdu []byte) (res *models.PostSmContextsResponse, ref string, errRes *models.PostSmContextsErrorResponse, prob *models.ProblemDetails, err1 error) {

	dat := ue.BuildCreateSmContextData(sm, nil)

	postSmContextsRequest := models.PostSmContextsRequest{
		JsonData:              &dat,
		BinaryDataN1SmMessage: naspdu,
	}
	
	client := pdusession_client(sm)

	postSmContextReponse, httpResponse, err :=
		client.SMContextsCollectionApi.PostSmContexts(org_context.Background(), postSmContextsRequest)

	if err == nil {
		res = &postSmContextReponse
		ref = httpResponse.Header.Get("Location")
	} else if httpResponse != nil {
		if httpResponse.Status != err.Error() {
			err1 = err
			return
		}
		switch httpResponse.StatusCode {
		case 400, 403, 404, 500, 503, 504:
			errResponse := err.(openapi.GenericOpenAPIError).Model().(models.PostSmContextsErrorResponse)
			errRes = &errResponse
		case 411, 413, 415, 429:
			problem := err.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
			prob = &problem
		}
	} else {
		err1 = openapi.ReportError("server no response")
	}
	return 
}
// Upadate SmContext Request
// servingNfId, smContextStatusUri, guami, servingNetwork -> amf change
// anType -> anType change
// ratType -> ratType change
// presenceInLadn -> Service Request , Xn handover, N2 handover and dnn is a ladn
// ueLocation -> the user location has changed or the user plane of the PDU session is deactivated
// upCnxState -> request the activation or the deactivation of the user plane connection of the PDU session
// hoState -> the preparation, execution or cancellation of a handover of the PDU session
// toBeSwitch -> Xn Handover to request to switch the PDU session to a new downlink N3 tunnel endpoint
// failedToBeSwitch -> indicate that the PDU session failed to be setup in the target RAN
// targetId, targetServingNfId(preparation with AMF change) -> N2 handover
// release -> duplicated PDU Session Id in subclause 5.2.2.3.11, slice not available in subclause 5.2.2.3.12
// ngApCause -> e.g. the NGAP cause for requesting to deactivate the user plane connection of the PDU session.
// 5gMmCauseValue -> AMF received a 5GMM cause code from the UE e.g 5GMM Status message in response to
// a Downlink NAS Transport message carrying 5GSM payload
// anTypeCanBeChanged

func (c *smfConsumer) SendUpdateSmContextActivateUpCnxState(
	ue *context.AmfUe, sm *context.SmContext, accessType models.AccessType) (
	*models.UpdateSmContextResponse, *models.UpdateSmContextErrorResponse, *models.ProblemDetails, error) {
	updateData := models.SmContextUpdateData{}
	updateData.UpCnxState = models.UpCnxState_ACTIVATING
	if !context.CompareUserLocation(ue.Location, sm.UserLocation()) {
		updateData.UeLocation = &ue.Location
	}
	if sm.AccessType() != accessType {
		updateData.AnType = sm.AccessType()
	}
	if ladn, ok := ue.ServingAMF().Ladn(sm.Dnn()); ok {
		if context.InTaiList(ue.Tai, ladn.TaiLists) {
			updateData.PresenceInLadn = models.PresenceState_IN_AREA
		}
	}
	return c.SendUpdateSmContextRequest(sm, updateData, nil, nil)
}

func (c *smfConsumer) SendUpdateSmContextDeactivateUpCnxState(ue *context.AmfUe,
	sm *context.SmContext, cause context.CauseAll) (
	*models.UpdateSmContextResponse, *models.UpdateSmContextErrorResponse, *models.ProblemDetails, error) {
	updateData := models.SmContextUpdateData{}
	updateData.UpCnxState = models.UpCnxState_DEACTIVATED
	updateData.UeLocation = &ue.Location
	if cause.Cause != nil {
		updateData.Cause = *cause.Cause
	}
	if cause.NgapCause != nil {
		updateData.NgApCause = cause.NgapCause
	}
	if cause.Var5GmmCause != nil {
		updateData.Var5gMmCauseValue = *cause.Var5GmmCause
	}
	return c.SendUpdateSmContextRequest(sm, updateData, nil, nil)
}

func (c *smfConsumer) SendUpdateSmContextChangeAccessType(ue *context.AmfUe,
	sm *context.SmContext, anTypeCanBeChanged bool) (
	*models.UpdateSmContextResponse, *models.UpdateSmContextErrorResponse, *models.ProblemDetails, error) {
	updateData := models.SmContextUpdateData{}
	updateData.AnTypeCanBeChanged = anTypeCanBeChanged
	return c.SendUpdateSmContextRequest(sm, updateData, nil, nil)
}

func (c *smfConsumer) SendUpdateSmContextN2Info(
	ue *context.AmfUe, sm *context.SmContext, n2SmType models.N2SmInfoType, N2SmInfo []byte) (
	*models.UpdateSmContextResponse, *models.UpdateSmContextErrorResponse, *models.ProblemDetails, error) {
	updateData := models.SmContextUpdateData{}
	updateData.N2SmInfoType = n2SmType
	updateData.N2SmInfo = new(models.RefToBinaryData)
	updateData.N2SmInfo.ContentId = "N2SmInfo"
	updateData.UeLocation = &ue.Location
	return c.SendUpdateSmContextRequest(sm, updateData, nil, N2SmInfo)
}

func (c *smfConsumer) SendUpdateSmContextXnHandover(
	ue *context.AmfUe, sm *context.SmContext, n2SmType models.N2SmInfoType, N2SmInfo []byte) (
	*models.UpdateSmContextResponse, *models.UpdateSmContextErrorResponse, *models.ProblemDetails, error) {
	updateData := models.SmContextUpdateData{}
	if n2SmType != "" {
		updateData.N2SmInfoType = n2SmType
		updateData.N2SmInfo = new(models.RefToBinaryData)
		updateData.N2SmInfo.ContentId = "N2SmInfo"
	}
	updateData.ToBeSwitched = true
	updateData.UeLocation = &ue.Location
	if ladn, ok := ue.ServingAMF().Ladn(sm.Dnn()); ok {
		if context.InTaiList(ue.Tai, ladn.TaiLists) {
			updateData.PresenceInLadn = models.PresenceState_IN_AREA
		} else {
			updateData.PresenceInLadn = models.PresenceState_OUT_OF_AREA
		}
	}
	return c.SendUpdateSmContextRequest(sm, updateData, nil, N2SmInfo)
}

func (c *smfConsumer) SendUpdateSmContextXnHandoverFailed(
	ue *context.AmfUe, sm *context.SmContext, n2SmType models.N2SmInfoType, N2SmInfo []byte) (
	*models.UpdateSmContextResponse, *models.UpdateSmContextErrorResponse, *models.ProblemDetails, error) {
	updateData := models.SmContextUpdateData{}
	if n2SmType != "" {
		updateData.N2SmInfoType = n2SmType
		updateData.N2SmInfo = new(models.RefToBinaryData)
		updateData.N2SmInfo.ContentId = "N2SmInfo"
	}
	updateData.FailedToBeSwitched = true
	return c.SendUpdateSmContextRequest(sm, updateData, nil, N2SmInfo)
}

func (c *smfConsumer) SendUpdateSmContextN2HandoverPreparing(
	ue *context.AmfUe,
	sm *context.SmContext,
	n2SmType models.N2SmInfoType,
	N2SmInfo []byte, amfid string, targetId *models.NgRanTargetId) (
	*models.UpdateSmContextResponse, *models.UpdateSmContextErrorResponse, *models.ProblemDetails, error) {
	updateData := models.SmContextUpdateData{}
	if n2SmType != "" {
		updateData.N2SmInfoType = n2SmType
		updateData.N2SmInfo = new(models.RefToBinaryData)
		updateData.N2SmInfo.ContentId = "N2SmInfo"
	}
	updateData.HoState = models.HoState_PREPARING
	updateData.TargetId = targetId
	// amf changed in same plmn
	if amfid != "" {
		updateData.TargetServingNfId = amfid
	}
	return c.SendUpdateSmContextRequest(sm, updateData, nil, N2SmInfo)
}

func (c *smfConsumer) SendUpdateSmContextN2HandoverPrepared(
	ue *context.AmfUe, sm *context.SmContext, n2SmType models.N2SmInfoType, N2SmInfo []byte) (
	*models.UpdateSmContextResponse, *models.UpdateSmContextErrorResponse, *models.ProblemDetails, error) {
	updateData := models.SmContextUpdateData{}
	if n2SmType != "" {
		updateData.N2SmInfoType = n2SmType
		updateData.N2SmInfo = new(models.RefToBinaryData)
		updateData.N2SmInfo.ContentId = "N2SmInfo"
	}
	updateData.HoState = models.HoState_PREPARED
	return c.SendUpdateSmContextRequest(sm, updateData, nil, N2SmInfo)
}

func (c *smfConsumer) SendUpdateSmContextN2HandoverComplete(
	ue *context.AmfUe, sm *context.SmContext, amfid string, guami *models.Guami) (
	*models.UpdateSmContextResponse, *models.UpdateSmContextErrorResponse, *models.ProblemDetails, error) {
	updateData := models.SmContextUpdateData{}
	updateData.HoState = models.HoState_COMPLETED
	if amfid != "" {
		updateData.ServingNfId = amfid
		updateData.ServingNetwork = guami.PlmnId
		updateData.Guami = guami
	}
	if ladn, ok := ue.ServingAMF().Ladn(sm.Dnn()); ok {
		if context.InTaiList(ue.Tai, ladn.TaiLists) {
			updateData.PresenceInLadn = models.PresenceState_IN_AREA
		} else {
			updateData.PresenceInLadn = models.PresenceState_OUT_OF_AREA
		}
	}
	return c.SendUpdateSmContextRequest(sm, updateData, nil, nil)
}

func (c *smfConsumer) SendUpdateSmContextN2HandoverCanceled(ue *context.AmfUe,
	sm *context.SmContext, cause context.CauseAll) (
	*models.UpdateSmContextResponse, *models.UpdateSmContextErrorResponse, *models.ProblemDetails, error) {
	updateData := models.SmContextUpdateData{}
	updateData.HoState = models.HoState_CANCELLED
	if cause.Cause != nil {
		updateData.Cause = *cause.Cause
	}
	if cause.NgapCause != nil {
		updateData.NgApCause = cause.NgapCause
	}
	if cause.Var5GmmCause != nil {
		updateData.Var5gMmCauseValue = *cause.Var5GmmCause
	}
	return c.SendUpdateSmContextRequest(sm, updateData, nil, nil)
}

func (c *smfConsumer) SendUpdateSmContextHandoverBetweenAccessType(
	ue *context.AmfUe, sm *context.SmContext, targetAccessType models.AccessType, N1SmMsg []byte) (
	*models.UpdateSmContextResponse, *models.UpdateSmContextErrorResponse, *models.ProblemDetails, error) {
	updateData := models.SmContextUpdateData{}
	updateData.AnType = targetAccessType
	if N1SmMsg != nil {
		updateData.N1SmMsg = new(models.RefToBinaryData)
		updateData.N1SmMsg.ContentId = "N1Msg"
	}
	return c.SendUpdateSmContextRequest(sm, updateData, N1SmMsg, nil)
}

func (c *smfConsumer) SendUpdateSmContextHandoverBetweenAMF(
	ue *context.AmfUe, sm *context.SmContext, amfid string, guami *models.Guami, activate bool) (
	*models.UpdateSmContextResponse, *models.UpdateSmContextErrorResponse, *models.ProblemDetails, error) {
	updateData := models.SmContextUpdateData{}
	updateData.ServingNfId = amfid
	updateData.ServingNetwork = guami.PlmnId
	updateData.Guami = guami
	if activate {
		updateData.UpCnxState = models.UpCnxState_ACTIVATING
		if !context.CompareUserLocation(ue.Location, sm.UserLocation()) {
			updateData.UeLocation = &ue.Location
		}
		if ladn, ok := ue.ServingAMF().Ladn(sm.Dnn()); ok {
			if context.InTaiList(ue.Tai, ladn.TaiLists) {
				updateData.PresenceInLadn = models.PresenceState_IN_AREA
			}
		}
	}
	return c.SendUpdateSmContextRequest(sm, updateData, nil, nil)
}

func (c *smfConsumer) SendUpdateSmContextRequest(sm *context.SmContext,
	dat models.SmContextUpdateData, n1Msg []byte, n2Info []byte) (
	response *models.UpdateSmContextResponse, errorResponse *models.UpdateSmContextErrorResponse,
	problemDetail *models.ProblemDetails, err1 error) {
	client := pdusession_client(sm) 

	updateSmContextRequest := models.UpdateSmContextRequest {
		JsonData:  &dat,
		BinaryDataN1SmMessage:  n1Msg,
		BinaryDataN2SmInformation: n2Info,
	}

	updateSmContextReponse, httpResponse, err :=
		client.IndividualSMContextApi.UpdateSmContext(org_context.Background(), sm.SmContextRef(), updateSmContextRequest)

	if err == nil {
		response = &updateSmContextReponse
	} else if httpResponse != nil {
		if httpResponse.Status != err.Error() {
			err1 = err
			return
		}
		switch httpResponse.StatusCode {
		case 400, 403, 404, 500, 503:
			errResponse := err.(openapi.GenericOpenAPIError).Model().(models.UpdateSmContextErrorResponse)
			errorResponse = &errResponse
		case 411, 413, 415, 429:
			problem := err.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
			problemDetail = &problem
		}
	} else {
		err1 = openapi.ReportError("server no response")
	}
	return response, errorResponse, problemDetail, err1
}

// Release SmContext Request

func (c *smfConsumer) SendReleaseSmContextRequest(ue *context.AmfUe, sm *context.SmContext,
	cause *context.CauseAll, n2SmInfoType models.N2SmInfoType,
	n2Info []byte) (detail *models.ProblemDetails, err error) {

	client := pdusession_client(sm)

	releaseData := ue.BuildReleaseSmContextData(cause, n2SmInfoType, n2Info)
	releaseSmContextRequest := models.ReleaseSmContextRequest{
		JsonData: &releaseData,
	}

	response, err1 := client.IndividualSMContextApi.ReleaseSmContext(
		org_context.Background(), sm.SmContextRef(), releaseSmContextRequest)

	if err1 == nil {
		ue.SmContextList.Delete(sm.PduSessionID())
	} else if response != nil && response.Status == err1.Error() {
		problem := err1.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
		detail = &problem
	} else {
		err = err1
	}
	return
}


