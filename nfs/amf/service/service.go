package service

import (
	"etri5gc/fabric"
	"etri5gc/nfs/amf/config"
	"etri5gc/nfs/amf/context"
	"etri5gc/nfs/amf/nas"
	"etri5gc/nfs/amf/nfselect"
	"etri5gc/nfs/amf/ngap"
	"etri5gc/nfs/amf/sbi/communication"
	"etri5gc/nfs/amf/sbi/consumer"
	"etri5gc/nfs/amf/sbi/eventexposure"
	"etri5gc/nfs/amf/sbi/httpcallback"
	"etri5gc/nfs/amf/sbi/location"
	"etri5gc/nfs/amf/sbi/mt"
	nfselectinf "etri5gc/nfs/amf/sbi/nfselect"
	"etri5gc/nfs/amf/sbi/oam"
	"etri5gc/nfs/amf/sbi/producer"
	"etri5gc/sbi"
	"fmt"
	"time"

	"github.com/free5gc/openapi/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "service"})
}

type AMF struct {
	sbi      sbi.SBI              //sbi server
	consumer *consumer.Consumer   //sbi consumer interacting with other NFs
	producer *producer.Handler    //handling Sbi requests received at the server
	selector *nfselect.NfSelector //Nf selection implementation
	ngap     *ngap.Server         //ngap server handling Ran connections and ngap messages
	context  *context.AMFContext  // AMF context
	conf     *config.Config       // loaded AMF config
	agent    fabric.ServiceAgent
	sender   fabric.Forwarder
}

func CreateAMF(cfg *config.Config) (nf *AMF, err error) {
	nf = &AMF{
		conf: cfg,
	}

	// initialize AMF context
	nf.context = context.CreateAmfContext(cfg)

	//create network function selector
	nf.selector = nfselect.NewNfSelector(nf)
	//create ngap server
	nf.ngap = ngap.NewServer(nf)
	//create sbi producer
	nf.producer = producer.NewHandler(nf, nf.ngap.Sender(), nf.ngap.Nas())
	//create sbi consumer
	nf.consumer = consumer.NewConsumer(nf)
	// create SBI server
	nf.sbi, err = sbi.CreateSbi(cfg.Configuration.Sbi, nf.makeroutes)

	return
}

func (nf *AMF) Context() *context.AMFContext {
	return nf.context
}

func (nf *AMF) Config() *config.Config {
	return nf.conf
}

func (nf *AMF) Producer() *producer.Handler {
	return nf.producer
}

func (nf *AMF) Consumer() *consumer.Consumer {
	return nf.consumer
}

func (nf *AMF) Nas() *nas.Nas {
	return nf.ngap.Nas()
}

func (nf *AMF) Ngap() *ngap.Server {
	return nf.ngap
}

func (nf *AMF) NfSelector() nfselectinf.NfSelector {
	return nf.selector
}

//add routes to the gin router
func (nf *AMF) makeroutes(router *gin.Engine) error {
	var gn string
	var routes sbi.Routes
	gn, routes = httpcallback.MakeRoutes(nf.producer)
	sbi.AddHttpRoutes(router, gn, routes)
	gn, routes = oam.MakeRoutes(nf.producer)
	sbi.AddHttpRoutes(router, gn, routes)
	for _, serviceName := range nf.conf.Configuration.ServiceNameList {
		switch models.ServiceName(serviceName) {
		case models.ServiceName_NAMF_COMM:
			gn, routes = communication.MakeRoutes(nf.producer)
			sbi.AddHttpRoutes(router, gn, routes)
		case models.ServiceName_NAMF_EVTS:
			gn, routes = eventexposure.MakeRoutes(nf.producer)
			sbi.AddHttpRoutes(router, gn, routes)
		case models.ServiceName_NAMF_MT:
			gn, routes = mt.MakeRoutes(nf.producer)
			sbi.AddHttpRoutes(router, gn, routes)
		case models.ServiceName_NAMF_LOC:
			gn, routes = location.MakeRoutes(nf.producer)
			sbi.AddHttpRoutes(router, gn, routes)
		}
	}
	return nil
}

func (nf *AMF) Start() (err error) {

	// start ngap server
	log.Info("Starting ngap server")
	nf.ngap.Run(nf.conf.Configuration.NgapIpList, 38412)

	// Register to NRF
	//TODO: probably nrf registration should be implemented in a separated go
	//routine. The Amf is functioning only after a registration success.

	log.Info("Registering to NRF")
	cnt := 1
LOOP:
	for cnt < 5 {
		if _, nfid, ierr := nf.consumer.Nrf().SendRegisterNFInstance(); ierr != nil {
			log.Errorf("Fail to register with NRF (attemp #%d) %s", cnt, ierr.Error())
			cnt++
			err = ierr
			timer := time.NewTimer(2 * time.Second)
			<-timer.C
		} else {
			if len(nfid) > 0 {
				log.Info("Amf is registered, need to update AMF info")
				err = nil
				//TODO: Update NF identity
				nf.context.SetNfId(nfid)
			}
			break LOOP
		}
	}

	log.Info("Amf is registered, it is ok now")
	if err != nil {
		nf.ngap.Stop()
		return
	}

	log.Info("Starting sbi server")
	//start sbi server
	nf.sbi.Serve()
	return
}

func (nf *AMF) Terminate() {
	//TODO: notify Ran of connection termination?
	//stop ngap server
	nf.ngap.Stop()
	//stop sbi server
	nf.sbi.Stop()

	if _, err := nf.consumer.Nrf().SendDeregisterNFInstance(); err != nil {
		log.Errorf("Fail to send degistration to the Nrf: %s", err.Error())
	}

	fmt.Println("Kill it")
}
