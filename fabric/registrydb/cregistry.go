package registrydb

import "etrib5gc/fabric/common"

// centralized registry
type centralizedRegistryClient struct {
	profile common.NfProfile
	server  *HttpAddr
}

func newCentralizedRegistry(profile common.NfProfile, server *HttpAddr) *centralizedRegistryClient {
	ret := &centralizedRegistryClient{
		profile: profile,
		server:  server,
	}
	return ret
}

func (r *centralizedRegistryClient) Search(query common.NfQuery) []common.NfProfile {
	return nil
}

func (r *centralizedRegistryClient) Drop(agent common.NfProfile) {
}

func (r *centralizedRegistryClient) Start() error {
	//TODO: start sending registration to the central registry
	return nil
}

func (r *centralizedRegistryClient) Terminate() {
}
