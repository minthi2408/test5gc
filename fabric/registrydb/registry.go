package registrydb

import (
	"etri5gc/fabric/common"
)

type distributedRegistry struct {
}

func NewDistributedAgentRegistry() *distributedRegistry {
	ret := &distributedRegistry{}
	return ret
}

func (r *distributedRegistry) Search(query common.NfQuery) []common.AgentProfile {
	return nil
}

func (r *distributedRegistry) Drop(agent common.AgentProfile) {
}
