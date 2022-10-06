package service

import (
	"etri5gc/fabric"
	fabric_common "etri5gc/fabric/common"
	fabric_config "etri5gc/fabric/config"

	pcf_config "etri5gc/nfs/pcf/config"
	"etri5gc/nfs/pcf/context"

	"etri5gc/nfs/pcf/sbi/producer"
	"fmt"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "service"})
}

type PCF struct {
	producer *producer.Producer    //handling Sbi requests received at the server
	context  *context.PcfContext   // PCF context
	config   *pcf_config.PcfConfig // loaded PCF config
	agent    fabric.ServiceAgent
	sender   fabric.Forwarder
}

func New(config *pcf_config.PcfConfig) (nf *PCF, err error) {
	nf = &PCF{
		config: config,
	}

	//create sbi producer
	nf.producer = producer.New(nf)

	//createfabric agent

	nf.agent, err = fabric.NewAgent(nf)
	nf.sender = nf.agent.Forwarder()

	// initialize PCF context
	//	nf.context = context.NewPcfContext(config, nf.sender)
	return
}

func (nf *PCF) AgentConfig() *fabric_config.AgentConfig {
	//TODO: get it from the nf config
	return nil
}

func (nf *PCF) Services() []fabric_common.Service {
	return nf.producer.Services()
}

func (nf *PCF) Context() *context.PcfContext {
	return nf.context
}

func (nf *PCF) Config() *pcf_config.PcfConfig {
	return nf.config
}
func (nf *PCF) Profile() fabric_common.NfProfile {
	return nil
}
func (nf *PCF) Start() (err error) {
	err = nf.agent.Start()
	return
}

func (nf *PCF) Terminate() {
	//stop sbi server
	nf.agent.Terminate()

	fmt.Println("Kill it")
}
