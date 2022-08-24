package fabric

import (
	"etri5gc/fabric/common"
)

// a service abstraction to expose service requesting to upper layers
// if the agent implements a http-based data plane protocol then the forwarder
// works like an http client. However, load balancing, fault recover, security,
// telemetry will be implemented by the forwarder with transparancy.
type Forwarder interface {
	//select an producer; send a request to get a response; retry if selected producer is dead
	DiscoveryThenSend(common.Request, common.NfQuery) (common.Response, common.AgentAddr, error)
	DirectSend(common.Request, common.AgentAddr) (common.Response, error)
}

// a server abstraction that listen to requests for registered services
type ServiceServer interface {
	//register services to handling incomming requests
	//must be called before the server starts
	Register([]common.Service) error
	//start server
	Serve()
	//terminate server
	Terminate()
}

// an abstraction that exposes the agent to the upper layers
//
type ServiceAgent interface {
	// start agent; just once
	Start() error
	// terminate the agent; only after it has started
	Terminate()

	// expose the forwarder
	Forwarder() Forwarder

	//  expose service registering method
	Register([]common.Service) error
}

//////////////////////////////
//Registry

//a registry to query for producers
type AgentRegistry interface {
	//search agents which match a query
	Search(common.NfQuery) []common.AgentProfile

	//drop a dead agent
	Drop(common.AgentProfile)
}

///////////////////////
//Selector
// selector do the load balancing on the list of agent profiles
type LoadBalancer interface {
	Select([]common.AgentProfile) common.AgentProfile
}

/////////////////////////
//connection

type ConnectionManager interface {
	// create a new connection or resuse an existing one
	// add security layer if needs
	// NOTE: it is just a connection preparation, no interaction is done now
	Connect(common.AgentAddr) (RemoteConnection, error)

	// drop an unresponsive connection
	Drop(common.AgentAddr)
}
