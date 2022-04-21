package common


type SBI interface {
	Serve() error
}

type SbiConfig interface {
	GetBindingIP4() string
	GetSbiPort() int
	GetTls() (string, string)
	GetScheme() string
}

type sbi struct {
	config 		SbiConfig
	server		*http2wraper.Http2Server	
}
type AddRoutesFn func(*gin.Engine) error

func CreateSbi(c SbiConfig, fn AddRoutesFn) SBI, error {
	ret := sbi{config: c}
	if err := ret.init(c, fn); err != nil {
		return nil, err
	}
	return ret, nil
}

func (s sbi) init(s SbiConfig, fn AddRoutesFn) err error {

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

	
	// add all routes
	fn(router)	

	addr := fmt.Sprintf("%s:%d", s.config.GetBindingIPv4(), s.config.GetSbiPort())

	s.server, err = httpwrapper.NewHttp2Server(addr,/* amf.KeyLogPath*/"", router)

	return 
}

func (s *sbi) Serve() error {
	pemPath, keyPath := s.config.GetTls()
	serverScheme := c.GetScheme()
	if serverScheme == "http" {
		return s.server.ListenAndServe()
	}

	return  s.server.ListenAndServeTLS(pemPath, keyPath)
}
