package sbi

import (
	"fmt"
	"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	logger_util "github.com/free5gc/util/logger"
	"github.com/free5gc/util/httpwrapper"
)

type SBI interface {
	Serve() error
}

type sbi struct {
	config 		*Config
	server		*http.Server	
}

// interface to add routes to the gin router
type AddRoutesFn func(*gin.Engine) error

//create a SBI server
func CreateSbi(c *Config, fn AddRoutesFn) (SBI, error) {
	ret := sbi{config: c}
	if err := ret.init(fn); err != nil {
		return nil, err
	}
	return ret, nil
}

//create gin router, add routes, then create the http server
func (s sbi) init(fn AddRoutesFn) (err error) {

	router := logger_util.NewGinWithLogrus(nil)

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

	addr := fmt.Sprintf("%s:%d", s.config.BindingIPv4, s.config.Port)

	s.server, err = httpwrapper.NewHttp2Server(addr,/* amf.KeyLogPath*/"", router)

	return 
}

//run the http server
func (s sbi) Serve() error {
	if s.config.Scheme == "http" {
		return s.server.ListenAndServe()
	}

	return  s.server.ListenAndServeTLS(s.config.Tls.Pem, s.config.Tls.Key)
}
