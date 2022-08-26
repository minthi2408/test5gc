package context

import (
	"errors"
	"etri5gc/fabric"
	"etri5gc/fabric/common"
	"etri5gc/fabric/httpdp"
	"etri5gc/openapi"
)

type requestSender struct {
	fw    fabric.Forwarder
	query common.NfQuery
	addr  common.AgentAddr
}

//ask agent's forwarder to send a request, a connection to a selected producer
//is maintained for all requests
func (s *requestSender) Send(request *openapi.Request) (response *openapi.Response, err error) {
	var commonResponse common.Response
	var ok bool
	if s.query == nil {
		panic(errors.New("No destication to send"))
	}
	if s.addr == nil {
		commonResponse, s.addr, err = s.fw.DiscoveryThenSend(request, s.query)
	} else {
		commonResponse, err = s.fw.DirectSend(request, s.addr)
	}
	if err != nil {
		return
	}
	if response, ok = commonResponse.(*openapi.Response); !ok {
		err = errors.New("the response is not in the openapi format")
	}
	return
}

func (s *requestSender) DecodeResponse(resp *openapi.Response) error {
    return httpdp.Encoding().DecodeResponse(resp)
}
