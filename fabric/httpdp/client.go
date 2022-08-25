package httpdp

import (
	"errors"
	"io/ioutil"
	"net/http"

	"etri5gc/fabric/common"
	"etri5gc/openapi"
)

type remoteAgent struct {
	addr   common.AgentAddr
	client *http.Client
}

func (c *remoteAgent) Send(req common.Request) (common.Response, error) {
	if req.MsgType() != common.SERVICE_MSG_TYPE_OPENAPI {
		return nil, errors.New("unknown request format")
	}
	openapiReq := req.(*openapi.Request)
	//now turn an openapi request into a http request
	httpreq := c.prepareHttpRequest(openapiReq)
	//send the request and get a response
	if httpresp, err := c.client.Do(httpreq); err != nil {
		return nil, err
	} else {
		//read body of the response and prepare the openapi response
		openapiResp := &openapi.Response{
			Raw:        httpresp,
			StatusCode: httpresp.StatusCode,
		}
		var err error
		openapiResp.Body, err = ioutil.ReadAll(httpresp.Body)
		httpresp.Body.Close()
		return openapiResp, err
	}
}

//create a http client to send requests to a remote agent
func NewHttpRemoteAgent(addr common.AgentAddr) *remoteAgent {
	ret := &remoteAgent{
		addr: addr,
	}

	//TODO: netvision should create this client. Added features such as security and connection
	//reuseability should be considered
	ret.client = &http.Client{}
	return ret
}

//build an http request from openapi request
func (c *remoteAgent) prepareHttpRequest(req *openapi.Request) *http.Request {
	//netvision should implement this method (refer to the
	//free5gc/openapi/client.go)
	return nil
}

func (c *remoteAgent) Addr() common.AgentAddr {
	return c.addr
}
