package context

import (
	"github.com/free5gc/openapi/models"
)


type udmClient struct {
	ue		*AmfUe
}


//subcriber data management
func (c *udmClient) PutUpuAck(upuMacIue string) error {
	/*
	client := udm_sdm_client(ue)

	ackInfo := models.AcknowledgeInfo{
		UpuMacIue: upuMacIue,
	}
	upuOpt := Nudm_SubscriberDataManagement.PutUpuAckParamOpts{
		AcknowledgeInfo: optional.NewInterface(ackInfo),
	}
	_, err := client.ProvidingAcknowledgementOfUEParametersUpdateApi.PutUpuAck(org_context.Background(), ue.Supi, &upuOpt)
	return err
	*/
	return nil
}

func (c *udmClient) SDMGetAmData() (problemDetails *models.ProblemDetails, err error) {
	/*
	client := udm_sdm_client(ue)

	getAmDataParamOpt := Nudm_SubscriberDataManagement.GetAmDataParamOpts{
		PlmnId: optional.NewInterface(openapi.MarshToJsonString(ue.PlmnId)),
	}

	data, httpResp, localErr := client.AccessAndMobilitySubscriptionDataRetrievalApi.GetAmData(
		org_context.Background(), ue.Supi, &getAmDataParamOpt)
	if localErr == nil {
		ue.GetUdmInfo().AccessAndMobilitySubscriptionData = &data
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
	*/
	return
}

func (c *udmClient) SDMGetSmfSelectData() (problemDetails *models.ProblemDetails, err error) {
	/*
	client := udm_sdm_client(ue)

	paramOpt := Nudm_SubscriberDataManagement.GetSmfSelectDataParamOpts{
		PlmnId: optional.NewInterface(openapi.MarshToJsonString(ue.PlmnId)),
	}
	data, httpResp, localErr :=
		client.SMFSelectionSubscriptionDataRetrievalApi.GetSmfSelectData(org_context.Background(), ue.Supi, &paramOpt)
	if localErr == nil {
		ue.GetUdmInfo().SmfSelectionData = &data
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
	*/
	return
}

func (c *udmClient) SDMGetUeContextInSmfData() (problemDetails *models.ProblemDetails, err error) {
	/*
	client := udm_sdm_client(ue)

	data, httpResp, localErr :=
		client.UEContextInSMFDataRetrievalApi.GetUeContextInSmfData(org_context.Background(), ue.Supi, nil)
	if localErr == nil {
		ue.GetUdmInfo().UeContextInSmfData = &data
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
	*/
	return
}

func (c *udmClient) SDMSubscribe() (problemDetails *models.ProblemDetails, err error) {
	/*
	client := udm_sdm_client(ue)

	sdmSubscription := models.SdmSubscription{
		NfInstanceId: c.amf.NfId(),
		PlmnId:       &ue.PlmnId,
	}

	resSubscription, httpResp, localErr := client.SubscriptionCreationApi.Subscribe(
		org_context.Background(), ue.Supi, sdmSubscription)
	if localErr == nil {
		ue.GetNssfInfo().SdmSubscriptionId = resSubscription.SubscriptionId
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
	*/
	return
}

func (c *udmClient) SDMGetSliceSelectionSubscriptionData() (problemDetails *models.ProblemDetails, err error) {
	/*
	client := udm_sdm_client(ue)
	udminfo := ue.GetUdmInfo()
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

			udminfo.SubscribedNssai = append(udminfo.SubscribedNssai, subscribedSnssai)
		}
		for _, snssai := range nssai.SingleNssais {
			subscribedSnssai := models.SubscribedSnssai{
				SubscribedSnssai: &models.Snssai{
					Sst: snssai.Sst,
					Sd:  snssai.Sd,
				},
				DefaultIndication: false,
			}
			udminfo.SubscribedNssai = append(udminfo.SubscribedNssai, subscribedSnssai)
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
	*/
	return 
}

func (c *udmClient) SDMUnsubscribe() (problemDetails *models.ProblemDetails, err error) {
	/*
	client := udm_sdm_client(ue)

	httpResp, localErr := client.SubscriptionDeletionApi.Unsubscribe(org_context.Background(), ue.Supi, ue.GetNssfInfo().SdmSubscriptionId)
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
	*/
	return
}

//uecontext management

func (c *udmClient) UeCmRegistration(accessType models.AccessType, initialRegistrationInd bool) (
	*models.ProblemDetails, error) {
	/*
	client := udm_ucm_client(ue)

	switch accessType {
	case models.AccessType__3_GPP_ACCESS:
		registrationData := models.Amf3GppAccessRegistration{
			AmfInstanceId:          c.amf.NfId(),
			InitialRegistrationInd: initialRegistrationInd,
			Guami:                  c.amf.ServedGuami(),
			RatType:                ue.GetLocInfo().RatType,
			// TODO: not support Homogenous Support of IMS Voice over PS Sessions this stage
			ImsVoPs: models.ImsVoPs_HOMOGENEOUS_NON_SUPPORT,
		}

		_, httpResp, localErr := client.AMFRegistrationFor3GPPAccessApi.Registration(org_context.Background(),
			ue.Supi, registrationData)
		if localErr == nil {
			ue.GetNssfInfo().UeCmRegistered = true
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
			RatType:       ue.GetLocInfo().RatType,
		}

		_, httpResp, localErr :=
			client.AMFRegistrationForNon3GPPAccessApi.Register(org_context.Background(), ue.Supi, registrationData)
		if localErr == nil {
			ue.GetNssfInfo().UeCmRegistered = true
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
	*/
	return nil, nil
}

func (c *udmClient) UeCmDeregistration(accessType models.AccessType) (
	*models.ProblemDetails, error) {
		/*
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
	*/
	return nil, nil
}
