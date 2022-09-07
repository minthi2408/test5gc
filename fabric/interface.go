package fabric

import (
	"etri5gc/fabric/common"
	"etri5gc/fabric/config"
)

type Application interface {
	AgentConfig() *config.AgentConfig
	Services() []common.Service
	Profile() common.NfProfile
}

// a service abstraction to expose service requesting to upper layers
// if the agent implements a http-based data plane protocol then the forwarder
// works like an http client. However, load balancing, fault recover, security,
// telemetry will be implemented by the forwarder with transparancy.
type Forwarder interface {
	common.InternalService
	//select an producer; send a request to get a response; retry if selected producer is dead
	DiscoveryThenSend(common.Request, common.NfQuery) (common.Response, common.AgentAddr, error)
	DirectSend(common.Request, common.AgentAddr) (common.Response, error)
}

// a server abstraction that listen to requests for registered services
type ServiceServer common.InternalService

// an abstraction that exposes the agent to the upper layers
type ServiceAgent interface {
	// start agent; just once
	Start() error
	// terminate the agent; only after it has started
	Terminate()

	// expose the forwarder
	Forwarder() Forwarder
}

// /////////////////////
// Selector
// selector do the load balancing on the list of agent profiles
type LoadBalancer interface {
	Select([]common.NfProfile) common.NfProfile
}
