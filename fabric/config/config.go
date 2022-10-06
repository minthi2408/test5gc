package config

import (
	"etrib5gc/fabric/common"
	"etrib5gc/fabric/httpdp"
	"etrib5gc/fabric/registrydb"
)

// configuration parameters for an agent
type AgentConfig struct {
	NfType           common.NetworkFunctionType
	DProto           common.DataPlaneProtocol //dataplane protocol
	HttpServerConfig *httpdp.ServerConfig     //http-based service server configuration
	RegistryConfig   registrydb.Config        // registry configuration
	Context          common.NfContext
	Queries          []common.NfQuery
}
