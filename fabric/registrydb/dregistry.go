package registrydb

import "etrib5gc/fabric/common"

///distributed registry

type distributedRegistry struct {
	profile common.NfProfile
	broker  *HttpAddr
}

func newDistributedRegistry(profile common.NfProfile, broker *HttpAddr) *distributedRegistry {
	ret := &distributedRegistry{
		profile: profile,
		broker:  broker,
	}
	return ret
}

func (r *distributedRegistry) Search(query common.NfQuery) []common.NfProfile {
	return nil
}

func (r *distributedRegistry) Drop(agent common.NfProfile) {
}

func (r *distributedRegistry) Start() error {
	//TODO: start announe to the registry broker then fetch profiles of other
	//agents
	return nil
}

func (r *distributedRegistry) Terminate() {
}
