package service

import (
	"fmt"
	"github.com/urfave/cli"
	"etri5gc/common"
	"etri5gc/nfs/amf/config"
)


type GenericNF interface {
	Register() error //register to NRF
	Serve() error //serve http
}


type AMF struct {
	common		GenericNF
	sbi			common.SBI
	context		AmfContext
	cgf			*config.Config
}


func CreateAMF(cfg *config.Config) (amf *AMF, err error) {
	amf = &AMF{
		cfg:	cfg,
	}
	amf.sbi, err = common.CreateSbi(cfg, func(router *gin.Engine) error {
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
	return
}


func (nf *AMF) Start() error {
	logger.InitLog.Infoln("Server started")

//	self := context.AMF_Self()
//	util.InitAmfContext(self)


	ngapHandler := ngap_service.NGAPHandler{
		HandleMessage:      ngap.Dispatch,
		HandleNotification: ngap.HandleSCTPNotification,
	}

	ngap_service.Run(self.NgapIpList, 38412, ngapHandler)

	// Register to NRF
	var profile models.NfProfile
	if profileTmp, err := consumer.BuildNFInstance(self); err != nil {
		logger.InitLog.Error("Build AMF Profile Error")
	} else {
		profile = profileTmp
	}

	if _, nfId, err := consumer.SendRegisterNFInstance(self.NrfUri, self.NfId, profile); err != nil {
		logger.InitLog.Warnf("Send Register NF Instance failed: %+v", err)
	} else {
		self.NfId = nfId
	}
	nf.sbi.Serve()
}


func (nf *AMF) Terminate() {
	fmt.Println("Kill it")
}
