package conman

import (
    "net/http"
    "errors"
    "etri5gc/fabric/common"
    "etri5gc/openapi"
)
//abstraction of protocol-specific service clients
type RemoteConnection interface {
	Send(common.Request) (common.Response, error)
	Addr() common.AgentAddr
}

type ConnectionManager interface {
	// create a new connection or resuse an existing one
	// add security layer if needs
	// NOTE: it is just a connection preparation, no interaction is done now
	Connect(common.AgentAddr) (RemoteConnection, error)

	// drop an unresponsive connection
	Drop(common.AgentAddr)
}

type remoteAgent struct {
    addr    common.AgentAddr
    client  *http.Client
}


func (c *remoteAgent) Send(req common.Request) (common.Response, error) {
    if req.MsgType() != common.SERVICE_MSG_TYPE_OPENAPI {
        return nil, errors.New("unknown request format")
    }
    openapiReq := req.(*openapi.Request)
    //now turn an openapi request into a http request
    httpreq := c.prepareHttpRequest(openapiReq)
    //send the request and get a response
    httpresp, err := c.client.Do(httpreq)
    return openapiReq.Handler()(httpresp, err)

}

func (c *remoteAgent) prepareHttpRequest(req *openapi.Request) *http.Request {
    return nil
}
func (c *remoteAgent) Addr() common.AgentAddr {
    return c.addr
}

type connectionManager struct {
}


func NewConnectionManager() *connectionManager {
    ret := & connectionManager{
    }
    return ret
}


func (cm *connectionManager) Connect(addr common.AgentAddr) (RemoteConnection, error) {
    return nil, nil
}

func (cm *connectionManager) Drop(addr common.AgentAddr) {
}
