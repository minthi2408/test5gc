package registrydb

import "etri5gc/fabric/common"

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
