package fabric

import (
	"errors"
	"etri5gc/fabric/common"
	"etri5gc/fabric/conman"
	"etri5gc/fabric/registrydb"
)

const (
	FORWARDER_MAX_NUM_TRIES   = 3
	FORWARDER_REQUEST_TIMEOUT = 2000 //miliseconds
)

type forwarderImpl struct {
	//registry to search for remote agents
	reg registrydb.AgentRegistry
	//a load balancer implementation
	lb LoadBalancer
	//for creating a connection to a selected remote agent
	conman conman.ConnectionManager
	//protocol-sepecific implementation of service client
	sender conman.RemoteConnection
}

func newForwarder(reg registrydb.AgentRegistry, lb LoadBalancer, conman conman.ConnectionManager) Forwarder {
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
			err = errors.New("Out of hope, no agent reply; can't send the request")
			return
		}
		//2. select one
		agent := fw.lb.Select(profiles)

		addr = agent.Addr()

		//3. get a connection to the selected one
		if fw.sender, err = fw.conman.Connect(addr); err != nil {
			// something serious happens, just stop
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
}

// send a request directly to a remote agent with given address
func (fw *forwarderImpl) DirectSend(req common.Request, addr common.AgentAddr) (
	common.Response, error) {

	// NOTE: addr might have been returned to the caller in a previous call, so
	// make sure we can reuse the same client
	if fw.sender != nil && addr == fw.sender.Addr() {
		return fw.sender.Send(req)
	} else {
		if sender, err := fw.conman.Connect(addr); err != nil {
			return nil, err
		} else {
			return sender.Send(req)
		}
	}
}
