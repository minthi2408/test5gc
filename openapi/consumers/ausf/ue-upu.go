package ausf

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	"fmt"
	"net/http"
)

func (consumer *ausfConsumerImpl) SupiUeUpuPost(supi string, upuInfo models.UpuInfo) (
	result models.UpuSecurityInfo, httpresp *http.Response, err error) {

	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/%s/%s", supi, openapi.AUSF_UPUPROTECTION, openapi.AUSF_UE_UPUPROTECTION)
	req.Body = &upuInfo

	//send the request
	var resp *openapi.Response
	if resp, err = consumer.client.Send(req); err != nil {
		return
	}

	//handle the response
	httpresp = resp.Response.(*http.Response)
	encoding := httpresp.Header.Get("Content-Type")

	var prob models.ProblemDetails

	switch httpresp.StatusCode {
	case 200:
		err = consumer.client.DecodeBody(&result, resp.BodyBytes, encoding)
	case 503:
		if err = consumer.client.DecodeBody(&prob, resp.BodyBytes, encoding); err == nil {
			err = openapi.NewError(err.Error(), &prob)
		}
	default:
		err = fmt.Errorf("invalid status code: %d", httpresp.StatusCode)
	}
	return
}
