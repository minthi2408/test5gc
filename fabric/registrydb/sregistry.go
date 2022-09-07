package registrydb

import "etri5gc/fabric/common"

//static registry
//list of profiles are loaded from a config file

type staticRegistry struct {
	profile common.NfProfile
	others  []common.NfProfile
}

func newStaticRegistry(profile common.NfProfile, others []common.NfProfile) *staticRegistry {
	ret := &staticRegistry{
		profile: profile,
		others:  others,
	}
	return ret
}

func (r *staticRegistry) Search(query common.NfQuery) []common.NfProfile {
	return nil
}

func (r *staticRegistry) Drop(agent common.NfProfile) {
}

func (r *staticRegistry) Start() error {
	return nil
}

func (r *staticRegistry) Terminate() {
}
