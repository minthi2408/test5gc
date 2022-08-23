package fabric

import (
    "errors"
    "etri5gc/fabric/common"
)

//abstraction of protocol-specific service clients
type Sender interface {
    Send(common.Request, common.AgentAddr) (common.Response, error)
}

type forwarderImpl struct {
    reg             AgentRegistry
    lb              LoadBalancer
    connections     ConnectionManager
    sender          Sender
}

func newForwarder(reg AgentRegistry, lb LoadBalancer, conns ConnectionManager) Forwarder {
    ret := &forwarderImpl{
        reg:            reg,
        lb:             lb,
        connections:    conns,
    }

    return ret
}

func (fw *forwarderImpl)  DiscoveryThenSend(request common.Request, query common.NfQuery) (response common.Response, addr common.AgentAddr, err error) {
    var max_tries = 3
    for i:=0; i< max_tries;  i++ {
        var profiles []common.AgentProfile
        if profiles = fw.reg.Search(query); len(profiles) == 0 {
            err = errors.New("No agent is found")
            return
        }
        agent := fw.lb.Select(profiles)
        addr = agent.Addr()
        if response, err = fw.sender.Send(request, addr); err == nil {
            //done requesting the service
            return
        } else {
        }
    }
    err = errors.New("Out of hope, can't send the request")
    return
}


func (fw *forwarderImpl) DirectSend(req common.Request, addr common.AgentAddr) (common.Response, error) {
    return fw.sender.Send(req, addr)
}



