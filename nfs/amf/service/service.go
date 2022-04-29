package service

import (
	"fmt"
	"etri5gc/sbi"
	"etri5gc/nfs/amf/config"
	"etri5gc/nfs/amf/ngap"
	"etri5gc/nfs/amf/nas"
	"etri5gc/nfs/amf/context"
	"etri5gc/nfs/amf/sbi/nfselect"
	"etri5gc/nfs/amf/sbi/consumer"
	"etri5gc/nfs/amf/sbi/producer"
	"etri5gc/nfs/amf/sbi/communication"
	"etri5gc/nfs/amf/sbi/httpcallback"
	"etri5gc/nfs/amf/sbi/mt"
	"etri5gc/nfs/amf/sbi/oam"
	"etri5gc/nfs/amf/sbi/location"
	"etri5gc/nfs/amf/sbi/eventexposure"

	"github.com/free5gc/openapi/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log	*logrus.Entry
func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "service"})
}

type AMF struct {
	sbi			sbi.SBI //sbi server
	consumer	*consumer.Consumer
	producer	*producer.Handler
	selector	nfselect.NfSelector
	ngap		*ngap.Server
	//nas			*nas.Nas
	context		*context.AMFContext // AMF contex
	conf			*config.Config // loaded AMF config
}


func CreateAMF(cfg *config.Config) (nf *AMF, err error) {
	nf = &AMF{
		conf:	cfg,
	}

	// initialize AMF context
	nf.context = context.CreateAmfContext(cfg)

	nf.selector = NewNfSelector(nf)
	nf.ngap = ngap.NewServer(nf)
	nf.producer = producer.NewHandler(nf,nf.ngap.Sender(), nf.ngap.Nas())
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

func (nf *AMF) NfSelector() nfselect.NfSelector {
	return nf.selector
}

func (nf *AMF) makeroutes1(router *gin.Engine) error {
	return nil
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
	log.Info("Registering to NRF")
	cnt := 1
	for cnt < 5 {
		if sid, _, ierr := nf.consumer.Nrf().SendRegisterNFInstance(); ierr != nil {
			log.Errorf("Fail to register with NRF (attemp #%d) %s", cnt, ierr.Error())
			cnt++
			err = ierr
		} else {
			if len(sid) > 0 {
				log.Info("Amf is registered, need to update AMF info")
				err = nil 
				//TODO: Update NF identity
			}	
			break
		}
	}
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
	nf.ngap.Stop()
	nf.sbi.Stop()
	fmt.Println("Kill it")
}
