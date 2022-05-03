package consumer

import (
	"encoding/json"
	org_context "context"

	"etri5gc/nfs/amf/context"
	"github.com/antihax/optional"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/Nnssf_NSSelection"
)

type nssfConsumer struct {
	amf		*context.AMFContext
}


func newNssfConsumer(amf *context.AMFContext) *nssfConsumer {
	return &nssfConsumer{amf: amf}
}

func nssf_nssel_client(uri string) (*Nnssf_NSSelection.APIClient) {
	conf := Nnssf_NSSelection.NewConfiguration()
	conf.SetBasePath(uri)
	return Nnssf_NSSelection.NewAPIClient(conf)
}

func (c *nssfConsumer) NSSelectionGetForRegistration(ue *context.AmfUe, requestedNssai []models.MappingOfSnssai) (
	*models.ProblemDetails, error) {

	client := nssf_nssel_client(ue.NssfUri)

	sliceInfo := models.SliceInfoForRegistration{
		SubscribedNssai: ue.GetUdmInfo().SubscribedNssai,
	}

	for _, snssai := range requestedNssai {
		sliceInfo.RequestedNssai = append(sliceInfo.RequestedNssai, *snssai.ServingSnssai)
		if snssai.HomeSnssai != nil {
			sliceInfo.MappingOfNssai = append(sliceInfo.MappingOfNssai, snssai)
		}
	}

	var paramOpt Nnssf_NSSelection.NSSelectionGetParamOpts
	if e, err := json.Marshal(sliceInfo); err != nil {
	//	logger.ConsumerLog.Warnf("json marshal failed: %+v", err)
	} else {
		paramOpt = Nnssf_NSSelection.NSSelectionGetParamOpts{
			SliceInfoRequestForRegistration: optional.NewInterface(string(e)),
		}
	}
	res, httpResp, localErr := client.NetworkSliceInformationDocumentApi.NSSelectionGet(org_context.Background(),
		models.NfType_AMF, c.amf.NfId(), &paramOpt)
	if localErr == nil {
		ue.NetworkSliceInfo = &res
		for _, allowedNssai := range res.AllowedNssaiList {
			ue.AllowedNssai[allowedNssai.AccessType] = allowedNssai.AllowedSnssaiList
		}
		ue.ConfiguredNssai = res.ConfiguredNssai
	} else if httpResp != nil {
		if httpResp.Status != localErr.Error() {
			err := localErr
			return nil, err
		}
		problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
		return &problem, nil
	} else {
		return nil, openapi.ReportError("NSSF No Response")
	}

	return nil, nil
}

func (c *nssfConsumer) NSSelectionGetForPduSession(ue *context.AmfUe, snssai models.Snssai) (
	*models.AuthorizedNetworkSliceInfo, *models.ProblemDetails, error) {

	client := nssf_nssel_client(ue.NssfUri)

	sliceInfoForPduSession := models.SliceInfoForPduSession{
		SNssai:            &snssai,
		RoamingIndication: models.RoamingIndication_NON_ROAMING, // not support roaming
	}

	e, err := json.Marshal(sliceInfoForPduSession)
	if err != nil {
		//logger.ConsumerLog.Warnf("json marshal failed: %+v", err)
	}
	paramOpt := Nnssf_NSSelection.NSSelectionGetParamOpts{
		SliceInfoRequestForPduSession: optional.NewInterface(string(e)),
	}
	res, httpResp, localErr := client.NetworkSliceInformationDocumentApi.NSSelectionGet(org_context.Background(),
		models.NfType_AMF, c.amf.NfId(), &paramOpt)
	if localErr == nil {
		return &res, nil, nil
	} else if httpResp != nil {
		if httpResp.Status != localErr.Error() {
			return nil, nil, localErr
		}
		problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
		return nil, &problem, nil
	} else {
		return nil, nil, openapi.ReportError("NSSF No Response")
	}
}
