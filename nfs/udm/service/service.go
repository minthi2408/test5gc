package service

import (
	"etri5gc/fabric"
	fabric_common "etri5gc/fabric/common"
	fabric_config "etri5gc/fabric/config"

	udm_config "etri5gc/nfs/udm/config"
	"etri5gc/nfs/udm/context"

	"etri5gc/nfs/udm/sbi/producer"
	"fmt"

	"github.com/sirupsen/logrus"
	udmcon "etri5gc/openapi/consumers/udm"
	amfcon "etri5gc/openapi/consumers/amf"
	ausfcon "etri5gc/openapi/consumers/ausf"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "service"})
}

type UDM struct {
	producer *producer.Producer    //handling Sbi requests received at the server
	context  *context.UdmContext   // UDM context
	config   *udm_config.UdmConfig // loaded UDM config
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
	check_consumer_compile()
	return
}

func (nf *UDM) AgentConfig() *fabric_config.AgentConfig {
	//TODO: get it from the nf config
	return nil
}

func (nf *UDM) Services() []fabric_common.Service {
	return nf.producer.Services()
}

func (nf *UDM) Context() *context.UdmContext {
	return nf.context
}

func (nf *UDM) Config() *udm_config.UdmConfig {
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


func check_consumer_compile() {
	udmcon.GetIdTranslationResult(nil, "", nil)	
	udmcon.DeleteEeSubscription(nil, "", "")	
	amfcon.DeleteSubscription(nil,"")
	ausfcon.EapAuthMethod(nil, "", nil)
}

