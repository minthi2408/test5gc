package context

import (
	"etri5gc/fabric"
	"etri5gc/fabric/common"
	"etri5gc/openapi"
)

type requestSender struct {
	fw    fabric.Forwarder
	query common.NfQuery
	addr  common.AgentAddr
}

//ask agent's forwarder to send a request, a connection to a selected producer
//is maintained for all requests
func (s *requestSender) Send(request *openapi.Request) (*openapi.Response, error) {
	return nil, nil
}
