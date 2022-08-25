package service

import (
	"etri5gc/fabric"
	"etri5gc/nfs/amf/config"
	"etri5gc/nfs/amf/context"
	"etri5gc/nfs/amf/nas"
	"etri5gc/nfs/amf/ngap"

	"etri5gc/nfs/amf/sbi/producer"
	"fmt"

	//"time"

	//	"github.com/free5gc/openapi/models"
	//	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "service"})
}

type AMF struct {
	producer *producer.Producer  //handling Sbi requests received at the server
	ngap     *ngap.Server        //ngap server handling Ran connections and ngap messages
	context  *context.AMFContext // AMF context
	conf     *config.Config      // loaded AMF config
	agent    fabric.ServiceAgent
	sender   fabric.Forwarder
}

func CreateAMF(cfg *config.Config) (nf *AMF, err error) {
	nf = &AMF{
		conf: cfg,
	}

	//create ngap server
	nf.ngap = ngap.NewServer(nf)

	//create sbi producer
	nf.producer = producer.NewProducer(nf, nf.ngap.Sender(), nf.ngap.Nas())

	//createfabric agent
	agentconfig := &fabric.AgentConfig{} //TODO: should be loaded from a configuration file
	agentconfig.SetServices(nf.producer.Services())
	nf.agent, err = fabric.CreateServiceAgent(agentconfig)
	nf.sender = nf.agent.Forwarder()

	// initialize AMF context
	nf.context = context.CreateAmfContext(cfg, nf.sender)

	return
}

func (nf *AMF) Context() *context.AMFContext {
	return nf.context
}

func (nf *AMF) Config() *config.Config {
	return nf.conf
}

func (nf *AMF) Producer() *producer.Producer {
	return nf.producer
}

func (nf *AMF) Nas() *nas.Nas {
	return nf.ngap.Nas()
}

func (nf *AMF) Ngap() *ngap.Server {
	return nf.ngap
}

func (nf *AMF) Start() (err error) {

	// start ngap server
	log.Info("Starting ngap server")
	nf.ngap.Run(nf.conf.Configuration.NgapIpList, 38412)

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

	/*
		if _, err := nf.consumer.Nrf().SendDeregisterNFInstance(); err != nil {
			log.Errorf("Fail to send degistration to the Nrf: %s", err.Error())
		}
	*/

	fmt.Println("Kill it")
}
