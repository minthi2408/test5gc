package common
import (
    "time"
)
const (
	DATAPLANE_HTTP = iota
	DATAPLANE_HTTPS
	DATALANE_MAX_PROTOCOLS
)

type DataPlaneProtocol int

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

type AgentAddr interface {
    Protocol() DataPlaneProtocol
}

//the type of a network function (amf, smf, etc.)
type NetworkFunctionType string

// the abstraction of querries that are sent by an upper layer caller to select
// a remote producer
// Note: we have not sure how the query should be implemented yet. Should it be a
// complex sql-like query, or should it be any other simplified form?

type NfQuery interface {
	GetNfType() NetworkFunctionType
}

//define a context where a network function operate
//a NfQuery should be map to a single NfContext so that a list of currently
//deployed (running) NFs can be retrieved
type NfContext struct {
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
}


type TelemetryReport interface {
	Time() time.Time
}

