package service

import (
	"etri5gc/fabric"
	fabric_common "etri5gc/fabric/common"
	fabric_config "etri5gc/fabric/config"

	smf_config "etri5gc/nfs/smf/config"
	"etri5gc/nfs/smf/context"

	"etri5gc/nfs/smf/sbi/producer"
	"fmt"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "service"})
}

type SMF struct {
	producer *producer.Producer    //handling Sbi requests received at the server
	context  *context.SmfContext   // SMF context
	config   *smf_config.SmfConfig // loaded SMF config
	agent    fabric.ServiceAgent
	sender   fabric.Forwarder
}

func New(config *smf_config.SmfConfig) (nf *SMF, err error) {
	nf = &SMF{
		config: config,
	}

	//create sbi producer
	nf.producer = producer.New(nf)

	//createfabric agent

	nf.agent, err = fabric.NewAgent(nf)
	nf.sender = nf.agent.Forwarder()

	// initialize SMF context
	//	nf.context = context.NewSmfContext(config, nf.sender)
	return
}

func (nf *SMF) AgentConfig() *fabric_config.AgentConfig {
	//TODO: get it from the nf config
	return nil
}

func (nf *SMF) Services() []fabric_common.Service {
	return nf.producer.Services()
}

func (nf *SMF) Context() *context.SmfContext {
	return nf.context
}

func (nf *SMF) Config() *smf_config.SmfConfig {
	return nf.config
}
func (nf *SMF) Profile() fabric_common.NfProfile {
	return nil
}
func (nf *SMF) Start() (err error) {
	err = nf.agent.Start()
	return
}

func (nf *SMF) Terminate() {
	//stop sbi server
	nf.agent.Terminate()

	fmt.Println("Kill it")
}
