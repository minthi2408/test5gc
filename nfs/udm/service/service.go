package service

import (
	"etri5gc/fabric"
	fabric_common "etri5gc/fabric/common"
	fabric_config "etri5gc/fabric/config"

	udm_config "etri5gc/nfs/udm/config"
	"etri5gc/nfs/amf/context"

	"etri5gc/nfs/udm/sbi/producer"
	"fmt"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "service"})
}

type UDM struct {
	producer *producer.Producer    //handling Sbi requests received at the server
	context  *context.UDMContext   // UDM context
	config   *amf_config.UdmConfig // loaded UDM config
	agent    fabric.ServiceAgent
	sender   fabric.Forwarder
}

func New(config *udm_config.UdmConfig) (nf *UDM, err error) {
	nf = &UDM{
		config: config,
	}

	//create sbi producer
	nf.producer = producer.New(nf)

	//createfabric agent

	nf.agent, err = fabric.NewAgent(nf)
	nf.sender = nf.agent.Forwarder()

	// initialize UDM context
//	nf.context = context.NewUdmContext(config, nf.sender)

	return
}

func (nf *UDM) AgentConfig() *fabric_config.AgentConfig {
	//TODO: get it from the nf config
	return nil
}

func (nf *UDM) Services() []fabric_common.Service {
	return nf.producer.Services()
}

func (nf *UDM) Context() *context.AmfContext {
	return nf.context
}

func (nf *UDM) Config() *amf_config.AmfConfig {
	return nf.config
}
func (nf *UDM) Profile() fabric_common.NfProfile {
	return nil
}
func (nf *UDM) Start() (err error) {
	err = nf.agent.Start()
	return
}

func (nf *UDM) Terminate() {
	//stop sbi server
	nf.agent.Terminate()

	fmt.Println("Kill it")
}
