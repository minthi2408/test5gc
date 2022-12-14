package fabric

import (
	"errors"

	"etrib5gc/fabric/common"
	"etrib5gc/fabric/conman"
	"etrib5gc/fabric/registrydb"
)

const (
	FORWARDER_MAX_NUM_TRIES   = 3
	FORWARDER_REQUEST_TIMEOUT = 2000 //miliseconds
)

type LoadBalancer interface {
	Select([]common.NfProfile) common.NfProfile
}

type forwarderImpl struct {
	//registry to search for remote agents
	reg registrydb.AgentRegistry
	//a load balancer implementation
	lb LoadBalancer
	//for creating a connection to a selected remote agent
	conman conman.ConnectionManager
	//protocol-sepecific implementation of service client
	//	sender conman.RemoteConnection
}

func New(reg registrydb.AgentRegistry, lb LoadBalancer, conman conman.ConnectionManager) *forwarderImpl {
	ret := &forwarderImpl{
		reg:    reg,
		lb:     lb,
		conman: conman,
	}

	return ret
}

// discover and select a remote agent to send a request
func (fw *forwarderImpl) DiscoveryThenSend(request common.Request, query common.NfQuery) (
	response common.Response, addr common.AgentAddrStruct, err error) {

	var sender conman.RemoteConnection
	//TODO: add request timeout (requesting should be non-blocking -> need a
	//worker -> do it later
	for {
		//1. discover agents
		var profiles []common.NfProfile
		if profiles = fw.reg.Search(query); len(profiles) == 0 {
			err = errors.New("Out of hope, no agent reply; can't send the request")
			return
		}
		//2. select one
		agent := fw.lb.Select(profiles)

		addr = agent.Addr()

		//3. get a connection to the selected one
		if sender, err = fw.conman.Connect(addr); err != nil {
			// something serious happens, just stop
			fw.reg.Drop(agent)
			return
		}

		//3. send a request
		for i := 0; i < FORWARDER_MAX_NUM_TRIES; i++ {
			if response, err = sender.Send(request); err == nil {
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
func (fw *forwarderImpl) DirectSend(req common.Request, addr common.AgentAddrStruct) (
	common.Response, error) {

	// NOTE: conman should return an existing sender (if available)
	if sender, err := fw.conman.Connect(addr); err != nil {
		return nil, err
	} else {
		return sender.Send(req)
	}
}

func (fw *forwarderImpl) Start() error {
	return nil
}

func (fw *forwarderImpl) Terminate() {
}
