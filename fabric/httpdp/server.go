package httpdp

import (
	"errors"
	"etri5gc/fabric/common"
	"fmt"
	"net/http"
	"sync"

	"github.com/free5gc/util/httpwrapper"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type HttpRoute struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

type HttpRoutes []HttpRoute

type HttpService struct {
	Group  string
	Routes []HttpRoute
}

func (s HttpService) Name() string {
	return s.Group
}

//httpServer
type httpServer struct {
	config *ServerConfig
	server *http.Server
	wg     sync.WaitGroup
}

type HttpAddr struct {
	IPv4 string
	Port int
	//TODO: add certificate
}

func (addr *HttpAddr) Protocol() common.DataPlaneProtocol {
	return common.DATAPLANE_HTTP
}

type ServerConfig struct {
	BindingIPv4 string
	Port        int
	//TODO: add tls
}

func NewHttpServer(config *ServerConfig, services []common.Service) (*httpServer, error) {
	ret := &httpServer{
		config: config,
	}
	if err := ret.register(services); err != nil {
		return nil, err
	}
	return ret, nil
}

// create a http server, register services and their handlers
func (s *httpServer) register(services []common.Service) (err error) {
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

	addr := fmt.Sprintf("%s:%d", s.config.BindingIPv4, s.config.Port)

	for _, sv := range services {
		if httpservice, ok := sv.(*HttpService); !ok {
			panic(errors.New("Not a HttpService"))
		} else {
			addHttpRoutes(router, httpservice.Group, httpservice.Routes)
		}
	}
	s.server, err = httpwrapper.NewHttp2Server(addr /* amf.KeyLogPath*/, "", router)
	return
}

func (s *httpServer) Serve() {
	go func() {
		s.server.ListenAndServe()

		/*
			if s.config.Scheme == "http" {
				if err := s.server.ListenAndServe(); err != nil {
					//log.Errorf("Http server failed to listen", err)
				}
				return
			}

			if err :=s.server.ListenAndServeTLS(s.config.Tls.Pem, s.config.Tls.Key); err != nil {
				//log.Errorf("Http server failed to listen", err)
			}
		*/

		s.wg.Add(1)
	}()
	//log.Info("http server is running")
}

func (s *httpServer) Terminate() {
	if s.server == nil {
		panic(errors.New("http server has not been started"))
	}
	s.server.Close()
	s.wg.Wait()
}

func addHttpRoutes(engine *gin.Engine, groupname string, routes []HttpRoute) *gin.RouterGroup {
	group := engine.Group(groupname)

	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			group.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			group.DELETE(route.Pattern, route.HandlerFunc)
		}
	}
	return group
}

// IndexHandler is the index handler.
func HttpIndexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello from EtriB5GC!")
}
