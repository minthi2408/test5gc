package registrydb

import (
	"errors"
	"etri5gc/fabric/common"
)

const (
	REGISTRY_STATIC = iota
	REGISTRY_CENTRALIZED
	REGISTRY_DISTRIBUTED
)

// abstraction of a service registry
type AgentRegistry interface {
	common.InternalService
	//search agents which match a query
	Search(common.NfQuery) []common.AgentProfile

	//drop a dead agent
	Drop(common.AgentProfile)
}

type HttpAddr struct {
	IPv4 string
	Port int
}

//Registry configuration
type Config struct {
	RegType int                   //static, centralized, or distributed
	Addr    *HttpAddr             //address of the server (centralized registry)  or address of the controller (distributed registry)
	Others  []common.AgentProfile // for static registry
}

//factory method to create a registry
func NewRegistry(profile common.AgentProfile, config *Config) AgentRegistry {
	switch config.RegType {
	case REGISTRY_STATIC:
		return newStaticRegistry(profile, config.Others)
	case REGISTRY_CENTRALIZED:
		return newCentralizedRegistry(profile, config.Addr)
	case REGISTRY_DISTRIBUTED:
		return newDistributedRegistry(profile, config.Addr)
	default:
		panic(errors.New("unknown registry type"))
	}
}
