package config

import (
	"etri5gc/fabric/common"
	"etri5gc/fabric/httpdp"
	"etri5gc/fabric/registrydb"
)

// configuration parameters for an agent
type AgentConfig struct {
	DProto           common.DataPlaneProtocol //dataplane protocol
	HttpServerConfig *httpdp.ServerConfig     //http-based service server configuration
	Profile          common.AgentProfile
	RegistryConfig   registrydb.Config // registry configuration
}
