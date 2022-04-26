package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"etri5gc/sbi"
	"etri5gc/nfs/amf/config"
	"etri5gc/nfs/amf/context"
	"etri5gc/nfs/amf/sbi/consumer"
	"etri5gc/nfs/amf/sbi/producer"
	"etri5gc/nfs/amf/sbi/communication"
	"etri5gc/nfs/amf/sbi/httpcallback"
	"etri5gc/nfs/amf/sbi/mt"
	"etri5gc/nfs/amf/sbi/oam"
	"etri5gc/nfs/amf/sbi/location"
	"etri5gc/nfs/amf/sbi/eventexposure"

	"github.com/free5gc/openapi/models"
)

type AMF struct {
	sbi			sbi.SBI //sbi server
	consumer	*consumer.Consumer
	producer	*producer.Handler
	context		*context.AMFContext // AMF contex
	conf			*config.Config // loaded AMF config
}


func CreateAMF(cfg *config.Config) (nf *AMF, err error) {
	nf = &AMF{
		conf:	cfg,
	}

	nf.consumer = consumer.NewConsumer(nf)
	nf.producer = producer.NewHandler(nf)
	// initialize AMF context
	nf.context = context.CreateAmfContext(cfg)

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
//add routes to the gin router
func (nf *AMF)makeroutes(router *gin.Engine) error {
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

func (nf *AMF) Start() error {
	//logger.InitLog.Infoln("Server started")


	// start ngap server
/*
	ngapHandler := ngap_service.NGAPHandler{
		HandleMessage:      ngap.Dispatch,
		HandleNotification: ngap.HandleSCTPNotification,
	}

	ngap_service.Run(self.NgapIpList, 38412, ngapHandler)
*/
	// Register to NRF
	
	if _, _, err := nf.consumer.NRF().SendRegisterNFInstance(); err != nil {
		return err
	} else {
		//TODO: Update NF identity
	}

	//start sbi server
	nf.sbi.Serve()
	return nil
}


func (nf *AMF) Terminate() {
	fmt.Println("Kill it")
}
