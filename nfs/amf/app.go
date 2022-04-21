package amf

import (
	"fmt"
	"github.com/urfave/cli"
	"etri5gc/common"
)


type GenericNF interface {
	Register() error //register to NRF
	Serve() error //serve http
}


type AMF struct {
	common		GenericNF
	context		AmfContext
	Config		Config
}

func NewAMF() *AMF {
	return &AMF{}
}

func (nf *AMF) Initialize(c *cli.Context) error {
}


func (nf *AMF) Start() error {
	logger.InitLog.Infoln("Server started")

	router := logger_util.NewGinWithLogrus(logger.GinLog)
	router.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{
			"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host",
			"Token", "X-Requested-With",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           86400,
	}))


	// all routes
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

	pemPath := util.AmfDefaultPemPath
	keyPath := util.AmfDefaultKeyPath
	sbi := factory.AmfConfig.Configuration.Sbi
	if sbi.Tls != nil {
		pemPath = sbi.Tls.Pem
		keyPath = sbi.Tls.Key
	}

	self := context.AMF_Self()
	util.InitAmfContext(self)

	addr := fmt.Sprintf("%s:%d", self.BindingIPv4, self.SBIPort)

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

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				// Print stack for panic to log. Fatalf() will let program exit.
				logger.InitLog.Fatalf("panic: %v\n%s", p, string(debug.Stack()))
			}
		}()

		<-signalChannel
		amf.Terminate()
		os.Exit(0)
	}()

	server, err := httpwrapper.NewHttp2Server(addr, amf.KeyLogPath, router)

	if server == nil {
		logger.InitLog.Errorf("Initialize HTTP server failed: %+v", err)
		return
	}

	if err != nil {
		logger.InitLog.Warnf("Initialize HTTP server: %+v", err)
	}

	serverScheme := factory.AmfConfig.Configuration.Sbi.Scheme
	if serverScheme == "http" {
		err = server.ListenAndServe()
	} else if serverScheme == "https" {
		err = server.ListenAndServeTLS(pemPath, keyPath)
	}

	if err != nil {
		logger.InitLog.Fatalf("HTTP server setup failed: %+v", err)
	}

}


func (nf *AMF) Terminate() {
	fmt.Println("Kill it")
}
