package service

import (
	"etri5gc/fabric"
	fabric_common "etri5gc/fabric/common"
	fabric_config "etri5gc/fabric/config"

	ausf_config "etri5gc/nfs/ausf/config"
	"etri5gc/nfs/ausf/context"

	"etri5gc/nfs/ausf/sbi/producer"
	"fmt"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "service"})
}

type AUSF struct {
	producer *producer.Producer      //handling Sbi requests received at the server
	context  *context.AusfContext    // AUSF context
	config   *ausf_config.AusfConfig // loaded AUSF config
	agent    fabric.ServiceAgent
	sender   fabric.Forwarder
}

func New(config *ausf_config.AusfConfig) (nf *AUSF, err error) {
	nf = &AUSF{
		config: config,
	}

	//create sbi producer
	nf.producer = producer.New(nf)

	//createfabric agent

	nf.agent, err = fabric.NewAgent(nf)
	nf.sender = nf.agent.Forwarder()

	// initialize AUSF context
	//	nf.context = context.NewAusfContext(config, nf.sender)
	return
}

func (nf *AUSF) AgentConfig() *fabric_config.AgentConfig {
	//TODO: get it from the nf config
	return nil
}

func (nf *AUSF) Services() []fabric_common.Service {
	return nf.producer.Services()
}

func (nf *AUSF) Context() *context.AusfContext {
	return nf.context
}

func (nf *AUSF) Config() *ausf_config.AusfConfig {
	return nf.config
}
func (nf *AUSF) Profile() fabric_common.NfProfile {
	return nil
}
func (nf *AUSF) Start() (err error) {
	err = nf.agent.Start()
	return
}

func (nf *AUSF) Terminate() {
	//stop sbi server
	nf.agent.Terminate()

	fmt.Println("Kill it")
}
