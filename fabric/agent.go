package fabric

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/free5gc/util/httpwrapper"
)


type RegistryManager interface {
}

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
	Group			string
	Routes			[]HttpRoute
}

func (s HttpService) Name() string {
	return s.Group
}

type httpAgent struct {
	tm        TelemetryManager
	cm        ConnectionManager
	rm        RegistryManager
	server    *httpServer
	forwarder *httpForwarder
}

func (agent *httpAgent) Start() (err error) {
	if agent.server == nil {
		panic(errors.New("http server has not been created"))
	}
	agent.server.serve()
	return
}

func (agent *httpAgent) Terminate() {
	agent.server.close()
	//log.Info("Sbi server stopped")
}

func (agent *httpAgent) Forwarder() Forwarder {
	return agent.forwarder
}

func (agent *httpAgent) Server() ServiceServer {
	return agent.server
}

// a factory method to create an agent
// it returns an nil value and an error in case of failure
// otherwise, internal go routines are running. The caller should tell the
// agent to terminate when exiting the application

func CreateServiceAgent(config	*AgentConfig) (ServiceAgent, error) {
	agent := &httpAgent{
		forwarder:		&httpForwarder{},
		server:			newHttpServer(config.http),
	}

	return agent, nil
}

//httpServer
type httpServer struct {
	config		*HttpServerConfig
	server		*http.Server	
	wg			sync.WaitGroup
}

type HttpServerConfig struct {
	BindingIPv4			string
	Port				int
}

func newHttpServer(config *HttpServerConfig) *httpServer {
	return &httpServer {
		config:	config,
	}
}

// create a http server, register services and their handlers
func (s *httpServer) Register(services []Service) (err error) {
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
	s.server, err = httpwrapper.NewHttp2Server(addr,/* amf.KeyLogPath*/"", router)
	return 
}

func (s *httpServer) serve() {
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

func (s *httpServer) close() {
	if s.server == nil {
		panic(errors.New("http server has not been started"))
	}
	s.server.Close()
	s.wg.Wait()
}


func addHttpRoutes(engine *gin.Engine, groupname string, routes []HttpRoute) *gin.RouterGroup{
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

// httpForwarder
type httpForwarder struct {
}

func (fw *httpForwarder) Send(Request, NfQuery) (Response, error) {
	return nil, nil
}
