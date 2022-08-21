package context

import (
	"github.com/free5gc/openapi/models"
)

type pcfClient struct {
	ue	*AmfUe
}

func (c *pcfClient) AMPolicyControlCreate(anType models.AccessType) (*models.ProblemDetails, error) {
	/*
	pcfinfo := ue.GetPcfInfo()
	client := pcf_pol_client(pcfinfo.PcfUri)

	policyAssociationRequest := models.PolicyAssociationRequest{
		NotificationUri: c.amf.GetIPv4Uri() + "/namf-callback/v1/am-policy/",
		Supi:            ue.Supi,
		Pei:             ue.Pei,
		Gpsi:            ue.Gpsi,
		AccessType:      anType,
		ServingPlmn: &models.NetworkId{
			Mcc: ue.PlmnId.Mcc,
			Mnc: ue.PlmnId.Mnc,
		},
		Guami: c.amf.ServedGuami(),
	}

	udminfo := ue.GetUdmInfo()
	if udminfo.AccessAndMobilitySubscriptionData != nil {
		policyAssociationRequest.Rfsp = udminfo.AccessAndMobilitySubscriptionData.RfspIndex
	}

	res, httpResp, localErr := client.DefaultApi.PoliciesPost(org_context.Background(), policyAssociationRequest)
	if localErr == nil {
		locationHeader := httpResp.Header.Get("Location")
		//logger.ConsumerLog.Debugf("location header: %+v", locationHeader)
		pcfinfo.AmPolicyUri = locationHeader

		re := regexp.MustCompile("/policies/.*")
		match := re.FindStringSubmatch(locationHeader)

		pcfinfo.PolicyAssociationId = match[0][10:]
		pcfinfo.AmPolicyAssociation = &res

		if res.Triggers != nil {
			for _, trigger := range res.Triggers {
				if trigger == models.RequestTrigger_LOC_CH {
					pcfinfo.RequestTriggerLocationChange = true
				}
				//if trigger == models.RequestTrigger_PRA_CH {
				// TODO: Presence Reporting Area handling (TS 23.503 6.1.2.5, TS 23.501 5.6.11)
				//}
			}
		}

		//logger.ConsumerLog.Debugf("UE AM Policy Association ID: %s", ue.PolicyAssociationId)
		//logger.ConsumerLog.Debugf("AmPolicyAssociation: %+v", ue.AmPolicyAssociation)
	} else if httpResp != nil {
		if httpResp.Status != localErr.Error() {
			return nil, localErr
		}
		problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
		return &problem, nil
	} else {
		return nil, openapi.ReportError("server no response")
	}
	*/
	return nil, nil
}

func (c *pcfClient) AMPolicyControlUpdate(updateRequest models.PolicyAssociationUpdateRequest) (
	problemDetails *models.ProblemDetails, err error) {
		/*
	pcfinfo := ue.GetPcfInfo()	
	client := pcf_pol_client(pcfinfo.PcfUri)

	res, httpResp, localErr := client.DefaultApi.PoliciesPolAssoIdUpdatePost(
		org_context.Background(), pcfinfo.PolicyAssociationId, updateRequest)
	if localErr == nil {
		if res.ServAreaRes != nil {
			pcfinfo.AmPolicyAssociation.ServAreaRes = res.ServAreaRes
		}
		if res.Rfsp != 0 {
			pcfinfo.AmPolicyAssociation.Rfsp = res.Rfsp
		}
		pcfinfo.AmPolicyAssociation.Triggers = res.Triggers
		pcfinfo.RequestTriggerLocationChange = false
		for _, trigger := range res.Triggers {
			if trigger == models.RequestTrigger_LOC_CH {
				pcfinfo.RequestTriggerLocationChange = true
			}
			// if trigger == models.RequestTrigger_PRA_CH {
			// TODO: Presence Reporting Area handling (TS 23.503 6.1.2.5, TS 23.501 5.6.11)
			// }
		}
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
	return problemDetails, err
	*/
	return
}

func (c *pcfClient) AMPolicyControlDelete() (problemDetails *models.ProblemDetails, err error) {
	/*
	pcfinfo := ue.GetPcfInfo()
	client := pcf_pol_client(pcfinfo.PcfUri)

	httpResp, localErr := client.DefaultApi.PoliciesPolAssoIdDelete(org_context.Background(), pcfinfo.PolicyAssociationId)
	if localErr == nil {
		ue.RemoveAmPolicyAssociation()
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
