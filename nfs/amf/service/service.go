package service

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"etrib5gc/fabric"
	fabric_common "etrib5gc/fabric/common"
	fabric_config "etrib5gc/fabric/config"
	amf_config "etrib5gc/nfs/amf/config"
	"etrib5gc/nfs/amf/context"
	"etrib5gc/nfs/amf/ngap"
	"etrib5gc/nfs/amf/sbi/producer"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "service"})
}

type AMF struct {
	producer *producer.Producer    //handling Sbi requests received at the server
	ngap     *ngap.Server          //ngap server handling Ran connections and ngap messages
	context  *context.AmfContext   // AMF context
	config   *amf_config.AmfConfig // loaded AMF config
	agent    fabric.ServiceAgent
	sender   fabric.Forwarder
}

func New(config *amf_config.AmfConfig) (nf *AMF, err error) {
	nf = &AMF{
		config: config,
	}
	log.Info(nf.config.Agent.DProto)
	//create NGAP server (including a NAS handler)
	nf.ngap = ngap.NewServer(nf)

	//create sbi producer
	nf.producer = producer.New(nf, nf.ngap.Sender(), nf.ngap.Nas())
	//createfabric agent

	nf.agent, err = fabric.NewAgent(nf)
	nf.sender = nf.agent.Forwarder()

	// initialize AMF context
	nf.context = context.NewAmfContext(config, nf.sender)

	return
}

func (nf *AMF) AgentConfig() *fabric_config.AgentConfig {
	return nf.config.Agent
}

func (nf *AMF) Services() []fabric_common.Service {
	return nf.producer.Services()
}

func (nf *AMF) Context() *context.AmfContext {
	return nf.context
}

func (nf *AMF) Config() *amf_config.AmfConfig {
	return nf.config
}
func (nf *AMF) Profile() fabric_common.NfProfile {
	return &fabric_common.NfProfileStruct{
		NfType_: nf.config.Agent.NfType,

		Addr_: fabric_common.AgentAddrStruct{Dproto: int(nf.config.Agent.DProto),

			IpAddress: nf.config.Agent.HttpServerConfig.IPv4,
			Port:      nf.config.Agent.HttpServerConfig.Port},

		Context_: nf.config.Agent.Context}
}
func (nf *AMF) Start() (err error) {
	// start ngap server
	log.Info("Starting NGAP server")
	nf.ngap.Run(nf.config.NgapIpList, 38412)

	if err != nil {
		nf.ngap.Stop()
		return
	}

	if err = nf.agent.Start(); err != nil {
		nf.ngap.Stop()
	}
	return

}

func (nf *AMF) Terminate() {
	//TODO: notify Ran of connection termination?
	//stop ngap server
	nf.ngap.Stop()
	//stop sbi server
	nf.agent.Terminate()

	fmt.Println("Kill it")
}
