package service

import (
	"etri5gc/fabric"
	fabric_common "etri5gc/fabric/common"
	fabric_config "etri5gc/fabric/config"

	udr_config "etri5gc/nfs/udr/config"
	"etri5gc/nfs/udr/context"

	"etri5gc/nfs/udr/sbi/producer"
	"fmt"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "service"})
}

type UDR struct {
	producer *producer.Producer    //handling Sbi requests received at the server
	context  *context.UdrContext   // UDR context
	config   *udr_config.UdrConfig // loaded UDR config
	agent    fabric.ServiceAgent
	sender   fabric.Forwarder
}

func New(config *udr_config.UdrConfig) (nf *UDR, err error) {
	nf = &UDR{
		config: config,
	}

	//create sbi producer
	nf.producer = producer.New(nf)

	//createfabric agent

	nf.agent, err = fabric.NewAgent(nf)
	nf.sender = nf.agent.Forwarder()

	// initialize UDR context
	//	nf.context = context.NewUdrContext(config, nf.sender)
	return
}

func (nf *UDR) AgentConfig() *fabric_config.AgentConfig {
	//TODO: get it from the nf config
	return nil
}

func (nf *UDR) Services() []fabric_common.Service {
	return nf.producer.Services()
}

func (nf *UDR) Context() *context.UdrContext {
	return nf.context
}

func (nf *UDR) Config() *udr_config.UdrConfig {
	return nf.config
}
func (nf *UDR) Profile() fabric_common.NfProfile {
	return nil
}
func (nf *UDR) Start() (err error) {
	err = nf.agent.Start()
	return
}

func (nf *UDR) Terminate() {
	//stop sbi server
	nf.agent.Terminate()

	fmt.Println("Kill it")
}
