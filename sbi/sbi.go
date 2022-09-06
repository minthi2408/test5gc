package sbi

import (
	"fmt"
	"github.com/free5gc/util/httpwrapper"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "sbi"})
}

type SBI interface {
	Serve() error
	Stop()
}

type sbi struct {
	config *Config
	server *http.Server
}

// interface to add routes to the gin router
type AddRoutesFn func(*gin.Engine) error

// create a SBI server
func CreateSbi(c *Config, fn AddRoutesFn) (SBI, error) {
	ret := &sbi{config: c}
	if err := ret.init(fn); err != nil {
		return nil, err
	}
	return ret, nil
}

// create gin router, add routes, then create the http server
func (s *sbi) init(fn AddRoutesFn) (err error) {
	//router := logger_util.NewGinWithLogrus(nil)
	router := gin.Default()

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
	s.server, err = httpwrapper.NewHttp2Server(addr /* amf.KeyLogPath*/, "", router)
	return
}

// run the http server
func (s *sbi) Serve() error {
	go func() {
		if s.config.Scheme == "http" {
			if err := s.server.ListenAndServe(); err != nil {
				log.Errorf("Http server failed to listen", err)
			}
			return
		}

		if err := s.server.ListenAndServeTLS(s.config.Tls.Pem, s.config.Tls.Key); err != nil {
			log.Errorf("Http server failed to listen", err)
		}
	}()
	log.Info("Sbi server is running")
	return nil
}

func (s *sbi) Stop() {
	s.server.Close()
	log.Info("Sbi server stopped")
}
