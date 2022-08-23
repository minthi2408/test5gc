package httpdp
import (
    "etri5gc/fabric/common"
)

// httpForwarder
type httpForwarder struct {
}

func (fw *httpForwarder) DiscoveryThenSend(req common.Request, query common.NfQuery) (common.Response, common.AgentAddr, error) {
	// try until success or out of hope:
	// find a list of producers matched the query from the local registry; then
	// select one among them; get a connection from the connection manager for
	// the selected one; send the request using the connection
	return nil, nil, nil
}

func (fw *httpForwarder) DirectSend(req common.Request, addr common.AgentAddr) (common.Response, error) {
	//get a connection for the given addr from the connection manager then send
	//the request using that connection
	return nil, nil
}


func NewForwarder() *httpForwarder {
    return &httpForwarder{}
}
