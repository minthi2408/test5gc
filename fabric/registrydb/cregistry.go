package registrydb

import "etri5gc/fabric/common"

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
	//TODO: start sending registration to the central registry
	return nil
}

func (r *centralizedRegistryClient) Terminate() {
}
