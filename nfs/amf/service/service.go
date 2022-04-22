package service

import (
	"fmt"
	"github.com/urfave/cli"
	"etri5gc/sbi"
	"etri5gc/common"
	"etri5gc/nfs/amf/config"
	"etri5gc/nfs/amf/context"
)

type AMF struct {
	sbi			sbi.SBI //sbi server
	nrf			common.NrfInterface //interface to NRF
	context		*context.Amf // AMF contex
	cgf			*config.Config // loaded AMF config
}


func CreateAMF(cfg *config.Config) (nf *AMF, err error) {
	nf = &AMF{
		cfg:	cfg,
	}

	// initialize AMF context
	nf.context = context.CreateAmfContext(cfg)

	// create SBI server
	nf.sbi, err = sbi.CreateSbi(cfg.Sbi, func(router *gin.Engine) error {
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
	})

	// create an interface to NRF
	nf.nrf = common.CreateNrfInterface(amf.context.BuildProfile)
	return
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
	
	if sid, nfid, err := nf.sbi.Register(); err != nil {
		return err
	} else {
		//Update NF identity
	}
	//start sbi server
	nf.sbi.Serve()
}


func (nf *AMF) Terminate() {
	fmt.Println("Kill it")
}
