package common

const (
	DATAPLANE_HTTP = iota
	DATAPLANE_HTTPS
	DATALANE_MAX_PROTOCOLS

	SERVICE_MSG_TYPE_OPENAPI = iota
	SERVICE_MSG_TYPE_MAX
)

type DataPlaneProtocol int

// anything that can be marshalled/unmarshalled
// it is up to its concrete implementation to determined an encoding method
type Marshallable interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type ServiceMessage interface {
	MsgType() int
}

// the abstraction of service requests composed by the upper layer
type Request interface {
	ServiceMessage
	//	Marshallable
}

// the abstraction of service responses composed by the upper layer
type Response interface {
	ServiceMessage
	//	Marshallable
}

type AgentAddr interface {
	Protocol() DataPlaneProtocol
	String() string //identity in string format
}

// the abstraction of a service supported by an NF
// it should be created and registered to the service agent and lives there
type Service interface {
	// name of the service. Prefixing it with a protocol transport address of
	// its NF instance should produce the service url
	Name() string
}

type AgentProfile interface {
	NfType() NetworkFunctionType
	Addr() AgentAddr
	Load() int //workload - for load balancer to make comparisons
}
