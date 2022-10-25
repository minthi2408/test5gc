package common

const (
	DATAPLANE_HTTP = iota
	DATAPLANE_HTTPS
	DATALANE_MAX_PROTOCOLS

	SERVICE_MSG_TYPE_OPENAPI = iota
	SERVICE_MSG_TYPE_MAX
)

type DataPlaneProtocol int
type ServiceMsgType int

// anything that can be marshalled/unmarshalled
// it is up to its concrete implementation to determined an encoding method
type Marshallable interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

// currently only openapi message format is support, but the fabric should
// be designed to incorporate other formats in the future.
type ServiceMessage interface {
	MsgType() ServiceMsgType
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
	Protocol() int
	String() string //identity in string format
}

type AgentAddrStruct struct {
	Dproto    int
	IpAddress string
	Port      int
	Identity  string //identity in string format
}

func (a AgentAddrStruct) Protocol() int {
	return a.Dproto
}

func (a AgentAddrStruct) String() string {
	return a.Identity
}

// the abstraction of a group of services supported by an NF
// it should be created and registered to the service agent and lives there
type Service interface {
	// name of the service. Prefixing it with a protocol transport address of
	// its NF instance should produce the service url
	Name() string //just a signature method, no use
}

type NfProfile interface {
	NfType() NetworkFunctionType
	Addr() AgentAddrStruct //dataplane transport address
	Load() int             //workload - for load balancer to make comparisons
	Context() NfContext
}

type NfProfileStruct struct {
	NfType_  NetworkFunctionType `json:"nftype"`
	Addr_    AgentAddrStruct     `json:"addr"` //dataplane transport address
	Load_    int                 `json:"load"` //workload - for load balancer to make comparisons
	Context_ NfContext           `json:"context"`
}

func (nf *NfProfileStruct) NfType() NetworkFunctionType {
	return nf.NfType_
}

func (nf *NfProfileStruct) Addr() AgentAddrStruct {
	return nf.Addr_
}

func (nf *NfProfileStruct) Load() int {
	return nf.Load_
}

func (nf *NfProfileStruct) Context() NfContext {
	return nf.Context_
}

// any service that runs within an agents; for easy initiation and clean
// termination of the agent
type InternalService interface {
	Start() error
	Terminate()
}
