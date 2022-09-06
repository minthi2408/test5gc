package fabric

import (
	"errors"
	"etri5gc/fabric/common"
	fabric_config "etri5gc/fabric/config"
	"etri5gc/fabric/conman"
	fwimpl "etri5gc/fabric/forwarder"
	"etri5gc/fabric/httpdp"
	"etri5gc/fabric/registrydb"
	"etri5gc/fabric/telemetry"
)

type serviceAgent struct {
	config    *fabric_config.AgentConfig
	reporter  telemetry.Writer
	conmngr   conman.ConnectionManager
	registry  registrydb.AgentRegistry
	lb        LoadBalancer
	server    ServiceServer
	forwarder Forwarder
	runnings  []common.InternalService
}

func (agent *serviceAgent) Start() (err error) {
	if agent.server == nil {
		panic(errors.New("server has not been created"))
	}

	//register services to execute
	services := make([]common.InternalService, 5)
	services[0] = agent.server
	services[1] = agent.reporter
	services[2] = agent.forwarder
	services[3] = agent.conmngr
	services[4] = agent.registry
	//execute service sequentially
	for _, service := range services {
		if err = service.Start(); err != nil {
			agent.Terminate()
			return
		}
		agent.runnings = append(agent.runnings, service)
	}
	return
}

func (agent *serviceAgent) Terminate() {
	//terminate all running services
	for _, service := range agent.runnings {
		service.Terminate()
	}
}

func (agent *serviceAgent) Forwarder() Forwarder {
	return agent.forwarder
}
func (agent *serviceAgent) Telemetry() telemetry.Writer {
	return agent.reporter
}

// a factory method to create an agent
// it returns an nil value and an error in case of failure
// otherwise, internal go routines are running. The caller should tell the
// agent to terminate when exiting the application

func NewAgent(nf Application) (ServiceAgent, error) {
	config := nf.AgentConfig()
	agent := &serviceAgent{
		config:   config,
		registry: registrydb.NewRegistry(config.Profile, &config.RegistryConfig),
		reporter: telemetry.NewWriter(),
		conmngr:  conman.NewConnectionManager(),
	}

	if config.DProto == common.DATAPLANE_HTTP {
		var err error
		agent.forwarder = fwimpl.New(agent.registry, agent.lb, agent.conmngr)
		if agent.server, err = httpdp.NewHttpServer(config.HttpServerConfig, nf.Services()); err != nil {
			return nil, err
		}
		return agent, nil
	} else {
		return nil, errors.New("the input data plane protocol is not supported")
	}
}
