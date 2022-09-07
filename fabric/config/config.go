package config

import (
	"etri5gc/fabric/common"
	"etri5gc/fabric/httpdp"
	"etri5gc/fabric/registrydb"
)

// configuration parameters for an agent
type AgentConfig struct {
    NfType          common.NetworkFunctionType
	DProto           common.DataPlaneProtocol //dataplane protocol
	HttpServerConfig *httpdp.ServerConfig     //http-based service server configuration
	RegistryConfig   registrydb.Config // registry configuration
    Context         common.NfContext
    Queries         []common.NfQuery
}
