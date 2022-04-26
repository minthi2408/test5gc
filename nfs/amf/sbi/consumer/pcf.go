package consumer

import (
	"regexp"
	org_context "context"

	"etri5gc/nfs/amf/context"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/Npcf_AMPolicy"
)

type PcfConsumer interface {
}

type pcfConsumer struct {
	amf		*context.AMFContext
}


func newPcfConsumer(amf *context.AMFContext) PcfConsumer {
	return &pcfConsumer{amf: amf}
}

func pcf_pol_client(uri string) (*Npcf_AMPolicy.APIClient) {
	conf := Npcf_AMPolicy.NewConfiguration()
	conf.SetBasePath(uri)
	return Npcf_AMPolicy.NewAPIClient(conf)
}

func (c *pcfConsumer) AMPolicyControlCreate(ue *context.AmfUe, anType models.AccessType) (*models.ProblemDetails, error) {
	client := pcf_pol_client(ue.PcfUri)

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

	if ue.AccessAndMobilitySubscriptionData != nil {
		policyAssociationRequest.Rfsp = ue.AccessAndMobilitySubscriptionData.RfspIndex
	}

	res, httpResp, localErr := client.DefaultApi.PoliciesPost(org_context.Background(), policyAssociationRequest)
	if localErr == nil {
		locationHeader := httpResp.Header.Get("Location")
		//logger.ConsumerLog.Debugf("location header: %+v", locationHeader)
		ue.AmPolicyUri = locationHeader

		re := regexp.MustCompile("/policies/.*")
		match := re.FindStringSubmatch(locationHeader)

		ue.PolicyAssociationId = match[0][10:]
		ue.AmPolicyAssociation = &res

		if res.Triggers != nil {
			for _, trigger := range res.Triggers {
				if trigger == models.RequestTrigger_LOC_CH {
					ue.RequestTriggerLocationChange = true
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
	return nil, nil
}

func (c *pcfConsumer) AMPolicyControlUpdate(ue *context.AmfUe, updateRequest models.PolicyAssociationUpdateRequest) (
	problemDetails *models.ProblemDetails, err error) {
	
	client := pcf_pol_client(ue.PcfUri)

	res, httpResp, localErr := client.DefaultApi.PoliciesPolAssoIdUpdatePost(
		org_context.Background(), ue.PolicyAssociationId, updateRequest)
	if localErr == nil {
		if res.ServAreaRes != nil {
			ue.AmPolicyAssociation.ServAreaRes = res.ServAreaRes
		}
		if res.Rfsp != 0 {
			ue.AmPolicyAssociation.Rfsp = res.Rfsp
		}
		ue.AmPolicyAssociation.Triggers = res.Triggers
		ue.RequestTriggerLocationChange = false
		for _, trigger := range res.Triggers {
			if trigger == models.RequestTrigger_LOC_CH {
				ue.RequestTriggerLocationChange = true
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
}

func (c *pcfConsumer) AMPolicyControlDelete(ue *context.AmfUe) (problemDetails *models.ProblemDetails, err error) {
	client := pcf_pol_client(ue.PcfUri)

	httpResp, localErr := client.DefaultApi.PoliciesPolAssoIdDelete(org_context.Background(), ue.PolicyAssociationId)
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

	return
}
