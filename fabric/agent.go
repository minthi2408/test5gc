package fabric

import (
	"errors"
	"etri5gc/fabric/common"
	"etri5gc/fabric/httpdp"
	"etri5gc/fabric/registrydb"
	"etri5gc/fabric/telemetry"
)

// configuration parameters for an agent
// there should be core parameters
// there should be plug-in parameters that are NF-dependent
type AgentConfig struct {
	NfType   common.NetworkFunctionType
	DProto   common.DataPlaneProtocol
	HttpConf *httpdp.HttpServerConfig

	services []common.Service
}

func (conf *AgentConfig) SetServices(services []common.Service) {
	conf.services = services
}

type serviceAgent struct {
	reporter    telemetry.Writer
	connections ConnectionManager
	registry    AgentRegistry
	server      ServiceServer
	forwarder   Forwarder
}

func (agent *serviceAgent) Start() (err error) {
	if agent.server == nil {
		panic(errors.New("server has not been created"))
	}
	agent.server.Serve()
	return
}

func (agent *serviceAgent) Terminate() {
	agent.server.Terminate()
}

func (agent *serviceAgent) Forwarder() Forwarder {
	return agent.forwarder
}

func (agent *serviceAgent) Register(services []common.Service) error {
	return agent.server.Register(services)
}

// a factory method to create an agent
// it returns an nil value and an error in case of failure
// otherwise, internal go routines are running. The caller should tell the
// agent to terminate when exiting the application

func CreateServiceAgent(config *AgentConfig) (ServiceAgent, error) {
	agent := &serviceAgent{
		registry: registrydb.NewDistributedAgentRegistry(),
	}
	if config.DProto == common.DATAPLANE_HTTP {
		agent.forwarder = httpdp.NewForwarder()
		agent.server = httpdp.NewHttpServer(config.HttpConf)
		return agent, nil
	} else {
		return nil, errors.New("the input data plane protocol is not supported")
	}
}
