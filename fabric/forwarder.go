package fabric

import (
	"errors"
	"etri5gc/fabric/common"
)

const (
	FORWARDER_MAX_NUM_TRIES   = 3
	FORWARDER_REQUEST_TIMEOUT = 2000 //miliseconds
)

//abstraction of protocol-specific service clients
type RemoteConnection interface {
	Send(common.Request) (common.Response, error)
	Addr() common.AgentAddr
}

type forwarderImpl struct {
	//registry to search for remote agents
	reg AgentRegistry
	//a load balancer implementation
	lb LoadBalancer
	//for creating a connection to a selected remote agent
	conman ConnectionManager
	//protocol-sepecific implementation of service client
	sender RemoteConnection
}

func newForwarder(reg AgentRegistry, lb LoadBalancer, conman ConnectionManager) Forwarder {
	ret := &forwarderImpl{
		reg:    reg,
		lb:     lb,
		conman: conman,
	}

	return ret
}

// discover and select a remote agent to send a request
func (fw *forwarderImpl) DiscoveryThenSend(request common.Request, query common.NfQuery) (
	response common.Response, addr common.AgentAddr, err error) {

	//TODO: add request timeout (requesting should be non-blocking -> need a
	//worker -> do it later
	for {
		//1. discover agents
		var profiles []common.AgentProfile
		if profiles = fw.reg.Search(query); len(profiles) == 0 {
			err = errors.New("No agent is found")
			return
		}
		//2. select one
		agent := fw.lb.Select(profiles)

		addr = agent.Addr()

		//3. get a connection to the selected one
		if fw.sender, err = fw.conman.Connect(addr); err != nil {
			// something serious happends, just stop
			fw.reg.Drop(agent)
			return
		}

		//3. send a request
		for i := 0; i < FORWARDER_MAX_NUM_TRIES; i++ {
			if response, err = fw.sender.Send(request); err == nil {
				//done requesting the service, just exit
				return
			}
		}
		//4. the remote agent is dead, tell the registry and the connection
		//mananger to drop it
		fw.reg.Drop(agent)
		fw.conman.Drop(addr)
	}
	err = errors.New("Out of hope, can't send the request")
	return
}

// send a request directly to a remote agent with given address
func (fw *forwarderImpl) DirectSend(req common.Request, addr common.AgentAddr) (
	common.Response, error) {

	if fw.sender != nil && addr == fw.sender.Addr() {
		//reuse connection
		return fw.sender.Send(req)
	} else {
		if sender, err := fw.conman.Connect(addr); err != nil {
			return nil, err
		} else {
			return sender.Send(req)
		}
	}
}
