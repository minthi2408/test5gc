package consumer

import (
	org_context "context"
	"etri5gc/nfs/amf/context"
	"github.com/antihax/optional"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/Nudm_SubscriberDataManagement"
	"github.com/free5gc/openapi/Nudm_UEContextManagement"
)

type UdmConsumer interface {
}

type udmConsumer struct {
	amf		*context.AMFContext
}


func newUdmConsumer(amf *context.AMFContext) UdmConsumer {
	return &udmConsumer{amf: amf}
}

func udm_sdm_client(ue *context.AmfUe) (*Nudm_SubscriberDataManagement.APIClient) {
	conf := Nudm_SubscriberDataManagement.NewConfiguration()
	conf.SetBasePath(ue.NudmSDMUri)
	return Nudm_SubscriberDataManagement.NewAPIClient(conf)
}

func udm_ucm_client(ue *context.AmfUe) (*Nudm_UEContextManagement.APIClient) {
	conf := Nudm_UEContextManagement.NewConfiguration()
	conf.SetBasePath(ue.NudmSDMUri)
	return Nudm_UEContextManagement.NewAPIClient(conf)
}


//subcriber data management
func (c *udmConsumer) PutUpuAck(ue *context.AmfUe, upuMacIue string) error {
	client := udm_sdm_client(ue)

	ackInfo := models.AcknowledgeInfo{
		UpuMacIue: upuMacIue,
	}
	upuOpt := Nudm_SubscriberDataManagement.PutUpuAckParamOpts{
		AcknowledgeInfo: optional.NewInterface(ackInfo),
	}
	_, err := client.ProvidingAcknowledgementOfUEParametersUpdateApi.PutUpuAck(org_context.Background(), ue.Supi, &upuOpt)
	return err
}

func (c *udmConsumer) SDMGetAmData(ue *context.AmfUe) (problemDetails *models.ProblemDetails, err error) {
	client := udm_sdm_client(ue)

	getAmDataParamOpt := Nudm_SubscriberDataManagement.GetAmDataParamOpts{
		PlmnId: optional.NewInterface(openapi.MarshToJsonString(ue.PlmnId)),
	}

	data, httpResp, localErr := client.AccessAndMobilitySubscriptionDataRetrievalApi.GetAmData(
		org_context.Background(), ue.Supi, &getAmDataParamOpt)
	if localErr == nil {
		ue.AccessAndMobilitySubscriptionData = &data
		ue.Gpsi = data.Gpsis[0] // TODO: select GPSI
	} else if httpResp != nil {
		if httpResp.Status != localErr.Error() {
			err = localErr
			return
		}
		problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
		problemDetails = &problem
	} else {
		err = openapi.ReportError("server no response")
	}
	return
}

func (c *udmConsumer) SDMGetSmfSelectData(ue *context.AmfUe) (problemDetails *models.ProblemDetails, err error) {

	client := udm_sdm_client(ue)

	paramOpt := Nudm_SubscriberDataManagement.GetSmfSelectDataParamOpts{
		PlmnId: optional.NewInterface(openapi.MarshToJsonString(ue.PlmnId)),
	}
	data, httpResp, localErr :=
		client.SMFSelectionSubscriptionDataRetrievalApi.GetSmfSelectData(org_context.Background(), ue.Supi, &paramOpt)
	if localErr == nil {
		ue.SmfSelectionData = &data
	} else if httpResp != nil {
		if httpResp.Status != localErr.Error() {
			err = localErr
			return
		}
		problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
		problemDetails = &problem
	} else {
		err = openapi.ReportError("server no response")
	}

	return
}

func (c *udmConsumer) SDMGetUeContextInSmfData(ue *context.AmfUe) (problemDetails *models.ProblemDetails, err error) {

	client := udm_sdm_client(ue)

	data, httpResp, localErr :=
		client.UEContextInSMFDataRetrievalApi.GetUeContextInSmfData(org_context.Background(), ue.Supi, nil)
	if localErr == nil {
		ue.UeContextInSmfData = &data
	} else if httpResp != nil {
		if httpResp.Status != localErr.Error() {
			err = localErr
			return
		}
		problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
		problemDetails = &problem
	} else {
		err = openapi.ReportError("server no response")
	}

	return
}

func (c *udmConsumer) SDMSubscribe(ue *context.AmfUe) (problemDetails *models.ProblemDetails, err error) {

	client := udm_sdm_client(ue)

	sdmSubscription := models.SdmSubscription{
		NfInstanceId: c.amf.NfId(),
		PlmnId:       &ue.PlmnId,
	}

	resSubscription, httpResp, localErr := client.SubscriptionCreationApi.Subscribe(
		org_context.Background(), ue.Supi, sdmSubscription)
	if localErr == nil {
		ue.SdmSubscriptionId = resSubscription.SubscriptionId
		return
	} else if httpResp != nil {
		if httpResp.Status != localErr.Error() {
			err = localErr
			return
		}
		problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
		problemDetails = &problem
	} else {
		err = openapi.ReportError("server no response")
	}
	return
}

func (c *udmConsumer) SDMGetSliceSelectionSubscriptionData(ue *context.AmfUe) (problemDetails *models.ProblemDetails, err error) {

	client := udm_sdm_client(ue)

	paramOpt := Nudm_SubscriberDataManagement.GetNssaiParamOpts{
		PlmnId: optional.NewInterface(openapi.MarshToJsonString(ue.PlmnId)),
	}
	nssai, httpResp, localErr :=
		client.SliceSelectionSubscriptionDataRetrievalApi.GetNssai(org_context.Background(), ue.Supi, &paramOpt)
	if localErr == nil {
		for _, defaultSnssai := range nssai.DefaultSingleNssais {
			subscribedSnssai := models.SubscribedSnssai{
				SubscribedSnssai: &models.Snssai{
					Sst: defaultSnssai.Sst,
					Sd:  defaultSnssai.Sd,
				},
				DefaultIndication: true,
			}
			ue.SubscribedNssai = append(ue.SubscribedNssai, subscribedSnssai)
		}
		for _, snssai := range nssai.SingleNssais {
			subscribedSnssai := models.SubscribedSnssai{
				SubscribedSnssai: &models.Snssai{
					Sst: snssai.Sst,
					Sd:  snssai.Sd,
				},
				DefaultIndication: false,
			}
			ue.SubscribedNssai = append(ue.SubscribedNssai, subscribedSnssai)
		}
	} else if httpResp != nil {
		if httpResp.Status != localErr.Error() {
			err = localErr
			return
		}
		problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
		problemDetails = &problem
	} else {
		err = openapi.ReportError("server no response")
	}
	return problemDetails, err
}

func (c *udmConsumer) SDMUnsubscribe(ue *context.AmfUe) (problemDetails *models.ProblemDetails, err error) {
	client := udm_sdm_client(ue)

	httpResp, localErr := client.SubscriptionDeletionApi.Unsubscribe(org_context.Background(), ue.Supi, ue.SdmSubscriptionId)
	if localErr == nil {
		return
	} else if httpResp != nil {
		if httpResp.Status != localErr.Error() {
			err = localErr
			return
		}
		problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
		problemDetails = &problem
	} else {
		err = openapi.ReportError("server no response")
	}
	return
}

//uecontext management

func (c *udmConsumer) UeCmRegistration(ue *context.AmfUe, accessType models.AccessType, initialRegistrationInd bool) (
	*models.ProblemDetails, error) {

	client := udm_ucm_client(ue)

	switch accessType {
	case models.AccessType__3_GPP_ACCESS:
		registrationData := models.Amf3GppAccessRegistration{
			AmfInstanceId:          c.amf.NfId(),
			InitialRegistrationInd: initialRegistrationInd,
			Guami:                  c.amf.ServedGuami(),
			RatType:                ue.RatType,
			// TODO: not support Homogenous Support of IMS Voice over PS Sessions this stage
			ImsVoPs: models.ImsVoPs_HOMOGENEOUS_NON_SUPPORT,
		}

		_, httpResp, localErr := client.AMFRegistrationFor3GPPAccessApi.Registration(org_context.Background(),
			ue.Supi, registrationData)
		if localErr == nil {
			ue.UeCmRegistered = true
			return nil, nil
		} else if httpResp != nil {
			if httpResp.Status != localErr.Error() {
				return nil, localErr
			}
			problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
			return &problem, nil
		} else {
			return nil, openapi.ReportError("server no response")
		}
	case models.AccessType_NON_3_GPP_ACCESS:
		registrationData := models.AmfNon3GppAccessRegistration{
			AmfInstanceId: c.amf.NfId(),
			Guami:         c.amf.ServedGuami(),
			RatType:       ue.RatType,
		}

		_, httpResp, localErr :=
			client.AMFRegistrationForNon3GPPAccessApi.Register(org_context.Background(), ue.Supi, registrationData)
		if localErr == nil {
			ue.UeCmRegistered = true
			return nil, nil
		} else if httpResp != nil {
			if httpResp.Status != localErr.Error() {
				return nil, localErr
			}
			problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
			return &problem, nil
		} else {
			return nil, openapi.ReportError("server no response")
		}
	}

	return nil, nil
}

func (c *udmConsumer) UeCmDeregistration(ue *context.AmfUe, accessType models.AccessType) (
	*models.ProblemDetails, error) {
	client := udm_ucm_client(ue)

	switch accessType {
	case models.AccessType__3_GPP_ACCESS:
		modificationData := models.Amf3GppAccessRegistrationModification{
			Guami:     c.amf.ServedGuami(),
			PurgeFlag: true,
		}

		httpResp, localErr := client.ParameterUpdateInTheAMFRegistrationFor3GPPAccessApi.Update(org_context.Background(),
			ue.Supi, modificationData)
		if localErr == nil {
			return nil, nil
		} else if httpResp != nil {
			if httpResp.Status != localErr.Error() {
				return nil, localErr
			}
			problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
			return &problem, nil
		} else {
			return nil, openapi.ReportError("server no response")
		}
	case models.AccessType_NON_3_GPP_ACCESS:
		modificationData := models.AmfNon3GppAccessRegistrationModification{
			Guami:     c.amf.ServedGuami(),
			PurgeFlag: true,
		}

		httpResp, localErr :=
			client.ParameterUpdateInTheAMFRegistrationForNon3GPPAccessApi.UpdateAmfNon3gppAccess(
				org_context.Background(), ue.Supi, modificationData)
		if localErr == nil {
			return nil, nil
		} else if httpResp != nil {
			if httpResp.Status != localErr.Error() {
				return nil, localErr
			}
			problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
			return &problem, nil
		} else {
			return nil, openapi.ReportError("server no response")
		}
	}

	return nil, nil
}
