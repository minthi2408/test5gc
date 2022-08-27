package httpdp

import (
	"errors"
	"io/ioutil"
	"net/http"

	"etri5gc/fabric/common"
	"etri5gc/fabric/httpdp/encoding"
	"etri5gc/openapi"
)

type remoteAgent struct {
	addr   common.AgentAddr
	client *http.Client
}

func (c *remoteAgent) Send(req common.Request) (resp common.Response, err error) {
	if req.MsgType() != common.SERVICE_MSG_TYPE_OPENAPI {
		return nil, errors.New("only openapi request format is supported")
	}
	openapiReq := req.(*openapi.Request)
    var httpreq *http.Request
    var httpresp *http.Response
	//now turn an openapi request into a http request
	if httpreq, err = c.prepareHttpRequest(openapiReq); err != nil {
        return
    }
	//send the request and get a response
	if httpresp, err = c.client.Do(httpreq); err != nil {
		return
	} else {
		//read body of the response and prepare the openapi response
        openapiResp := &openapi.Response{
			Response: httpresp,
            StatusCode: httpresp.StatusCode,
            Status: httpresp.Status,
		}
		openapiResp.BodyBytes, err = ioutil.ReadAll(httpresp.Body)
		httpresp.Body.Close()
        if err == nil {
            resp = openapiResp
        }
		return 
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
func (c *remoteAgent) prepareHttpRequest(req *openapi.Request) (httpreq *http.Request, err error) {
	//netvision should implement this method (refer to the
	//free5gc/openapi/client.go)
    //1. inject the ipaddr:port into the request path
    //2. encode the request
    if err = encoding.New().EncodeRequest(req); err == nil {
        httpreq = req.Request.(*http.Request)
    }
	return 
}

func (c *remoteAgent) Addr() common.AgentAddr {
	return c.addr
}
