package service

import (
	"fmt"
	"github.com/urfave/cli"
	"etri5gc/sbi"
	"etri5gc/common"
	"etri5gc/nfs/amf/config"
	"etri5gc/nfs/amf/context"
	"etri5gc/nfs/amf/sbi/consumer"
	"etri5gc/nfs/amf/sbi/producer"
)

type AMF struct {
	sbi			sbi.SBI //sbi server
	consumer	consumer.Consumer
	producer	producer.Handler
	context		*context.Amf // AMF contex
	cgf			*config.Config // loaded AMF config
}


func CreateAMF(cfg *config.Config) (nf *AMF, err error) {
	nf = &AMF{
		cfg:	cfg,
	}

	nf.consumer = consumer.NewConsumer(nf)
	nf.producer = producer.NewHandler(nf)
	// initialize AMF context
	nf.context = context.CreateAmfContext(cfg)

	// create SBI server
	nf.sbi, err = sbi.CreateSbi(cfg.Sbi, nf.makeroutes)

	return
}

func (nf *AMF) Context() *context.AmfContext {
	return nf.context
}

func (nf *AMF) Producer() *producer.Handler {
	return nf.producer
}

func (nf *AMF) Consumer() *consumer.Consumer {
	return nf.consumer
}
//add routes to the gin router
func (nf *AMF)makeroutes(router *gin.Engine) error {
	common.AddHttpRoutes(router, httpcallback.MakeRoutes(nf))
	common.AddHttpRoutes(router, oam.MakeRoutes(nf))
	for _, serviceName := range factory.AmfConfig.Configuration.ServiceNameList {
		switch models.ServiceName(serviceName) {
		case models.ServiceName_NAMF_COMM:
			common.AddHttpRoutes(router, communication.MakeRoutes(nf))
		case models.ServiceName_NAMF_EVTS:
			eventexposure.AddService(router)
			common.AddHttpRoutes(router, eventexposure.MakeRoutes(nf))
		case models.ServiceName_NAMF_MT:
			common.AddHttpRoutes(router, mt.MakeRoutes(nf))
		case models.ServiceName_NAMF_LOC:
			common.AddHttpRoutes(router, location.MakeRoutes(nf))
		}
	}
	return nil
}

func (nf *AMF) Start() error {
	logger.InitLog.Infoln("Server started")


	// start ngap server
/*
	ngapHandler := ngap_service.NGAPHandler{
		HandleMessage:      ngap.Dispatch,
		HandleNotification: ngap.HandleSCTPNotification,
	}

	ngap_service.Run(self.NgapIpList, 38412, ngapHandler)
*/
	// Register to NRF
	
	if sid, nfid, err := nf.consumer.NRF().Register(); err != nil {
		return err
	} else {
		//TODO: Update NF identity
	}

	//start sbi server
	nf.sbi.Serve()
}


func (nf *AMF) Terminate() {
	fmt.Println("Kill it")
}
