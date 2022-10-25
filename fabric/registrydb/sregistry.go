package registrydb

import "etrib5gc/fabric/common"

//static registry
//list of profiles are loaded from a config file

type staticRegistry struct {
	profile common.NfProfile
	others  []*common.NfProfileStruct
}

func newStaticRegistry(profile common.NfProfile, others []*common.NfProfileStruct) *staticRegistry {
	ret := &staticRegistry{
		profile: profile,
		others:  others,
	}
	return ret
}

func (r *staticRegistry) Search(query common.NfQuery) []common.NfProfile {
	var results []common.NfProfile
	var profile common.NfProfile

	results = append(results, profile)
	return results
}

func (r *staticRegistry) Drop(agent common.NfProfile) {
}

func (r *staticRegistry) Start() error {
	return nil
}

func (r *staticRegistry) Terminate() {
}
