package fabric

import (
	"time"
)

// anything that can be marshalled/unmarshalled
// it is up to its concrete implementation to determined an encoding method
type Marshallable interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

// the abstraction of service requests composed by the upper layer
type Request interface {
	Marshallable
}

// the abstraction of service responses composed by the upper layer
type Response interface {
	Marshallable
}

// the abstraction of querries that are sent by an upper layer caller to select
// a remote producer
// Note: we have not sure how the query should be implemented yet. Should it be a
// complex sql-like query, or should it be any other simplified form?

type NfQuery interface {
	GetNfType() NfType
}

//the type of a network function (amf, smf, etc.)
type NfType string

//define a context where a network function operate
//a NfQuery should be map to a single NfContext so that a list of currently
//deployed (running) NFs can be retrieved
type NfContext struct {
}

// configuration parameters for an agent
// there should be core parameters
// there should be plug-in parameters that are NF-dependent
type AgentConfig struct {
	nfType NfType
}

// the abstraction of a service supported by an NF
// it should be created and registered to the service agent and lives there
type Service interface {
	// name of the service. Prefixing it with a protocol transport address of
	// its NF instance should produce the service url
	Name() string
	// handle a consumer request. this should be implemented in the upper
	// layers
	Handle(Request) (Response, error)
}

// a service abstraction to expose service requesting to upper layers
// if the agent implements a http-based data plane protocol then the forwarder
// works like an http client. However, load balancing, fault recover, security,
// telemetry will be implemented by the forwarder with transparancy.
type Forwarder interface {
	//select an producer; send a request to get a response; retry if selected producer is dead
	Send(Request, NfQuery) (Response, error)
}

// an abstraction that exposes the agent to the upper layers
//
type ServiceAgent interface {
	//terminate the agent
	Terminate()

	// expose the forwarder
	Forwarder() Forwarder

	//register a service to handling incomming requests
	Register(Service) error
}

// a factory method to create an agent
// it returns an nil value and an error in case of failure
// otherwise, internal go routines are running. The caller should tell the
// agent to terminate when exiting the application

func MakeAgent(config AgentConfig) (ServiceAgent, error) {
	return nil, nil
}

//////////////////////////////
//Registry

type AgentProfile interface {
	NfType() NfType
}

//a registry to query for producers
type ServiceRegistry interface {
	Query(NfContext) []AgentProfile
}

///////////////////////
//Selector
// selector do the load balancing on the list of agent profiles
type Selector interface {
	Select([]AgentProfile) AgentProfile
}

////////////////////////////////////
//telemetry

type TelemetryReport interface {
	Time() time.Time
}

type TelemetryManager interface {
	Report(TelemetryReport)
}

/////////////////////////
//connection

type ConnectionManager interface {
}
