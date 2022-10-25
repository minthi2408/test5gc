package httpdp

import (
	"errors"
	"io/ioutil"
	"net/http"

	"etrib5gc/fabric/common"
	"etrib5gc/sbi"
)

type remoteAgent struct {
	addr   common.AgentAddrStruct
	client *http.Client
}

func (c *remoteAgent) Send(req common.Request) (resp common.Response, err error) {
	if req.MsgType() != common.SERVICE_MSG_TYPE_OPENAPI {
		return nil, errors.New("only sbi request format is supported")
	}
	sbiReq := req.(*sbi.Request)
	var httpreq *http.Request
	var httpresp *http.Response
	//now turn an sbi request into a http request
	if httpreq, err = c.prepareHttpRequest(sbiReq); err != nil {
		return
	}
	//send the request and get a response
	if httpresp, err = c.client.Do(httpreq); err != nil {
		return
	} else {
		//read body of the response and prepare the sbi response
		sbiResp := &sbi.Response{
			Response:   httpresp,
			StatusCode: httpresp.StatusCode,
			Status:     httpresp.Status,
		}
		sbiResp.BodyBytes, err = ioutil.ReadAll(httpresp.Body)
		httpresp.Body.Close()
		if err == nil {
			resp = sbiResp
		}
		return
	}
}

// create a http client to send requests to a remote agent
func NewHttpRemoteAgent(addr common.AgentAddrStruct) *remoteAgent {
	ret := &remoteAgent{
		addr: addr,
	}

	//TODO: netvision should create this client. Added features such as security and connection
	//reuseability should be considered
	ret.client = &http.Client{}
	return ret
}

// build an http request from sbi request
func (c *remoteAgent) prepareHttpRequest(req *sbi.Request) (httpreq *http.Request, err error) {
	//netvision should implement this method (refer to the
	//free5gc/sbi/client.go)
	//1. inject the ipaddr:port into the request path
	//2. encode the request
	if err = Encoding().EncodeRequest(req); err == nil {
		httpreq = req.Request.(*http.Request)
	}
	return
}

func (c *remoteAgent) Addr() common.AgentAddrStruct {
	return c.addr
}
