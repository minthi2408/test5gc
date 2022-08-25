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

//distributed registry
type distributedRegistry struct {
	profile common.AgentProfile
	broker  *HttpAddr
}

func newDistributedRegistry(profile common.AgentProfile, broker *HttpAddr) *distributedRegistry {
	ret := &distributedRegistry{
		profile: profile,
		broker:  broker,
	}
	return ret
}

func (r *distributedRegistry) Search(query common.NfQuery) []common.AgentProfile {
	return nil
}

func (r *distributedRegistry) Drop(agent common.AgentProfile) {
}

func (r *distributedRegistry) Start() error {
	return nil
}

func (r *distributedRegistry) Terminate() {
}

//centralized registry
type centralizedRegistryClient struct {
	profile common.AgentProfile
	server  *HttpAddr
}

func newCentralizedRegistry(profile common.AgentProfile, server *HttpAddr) *centralizedRegistryClient {
	ret := &centralizedRegistryClient{
		profile: profile,
		server:  server,
	}
	return ret
}

func (r *centralizedRegistryClient) Search(query common.NfQuery) []common.AgentProfile {
	return nil
}

func (r *centralizedRegistryClient) Drop(agent common.AgentProfile) {
}

func (r *centralizedRegistryClient) Start() error {
	return nil
}

func (r *centralizedRegistryClient) Terminate() {
}

//static registry
//list of profiles are loaded from a config file

type staticRegistry struct {
	profile common.AgentProfile
	others  []common.AgentProfile
}

func newStaticRegistry(profile common.AgentProfile, others []common.AgentProfile) *staticRegistry {
	ret := &staticRegistry{
		profile: profile,
		others:  others,
	}
	return ret
}

func (r *staticRegistry) Search(query common.NfQuery) []common.AgentProfile {
	return nil
}

func (r *staticRegistry) Drop(agent common.AgentProfile) {
}

func (r *staticRegistry) Start() error {
	return nil
}

func (r *staticRegistry) Terminate() {
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
