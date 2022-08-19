package fabric

type HttpServer interface {
}

type RegistryManager interface {
}

type httpAgent struct {
	forwarder Forwarder
	tm        TelemetryManager
	cm        ConnectionManager
	rm        RegistryManager
	server    HttpServer
}

func (agent *httpAgent) Terminate() {
}

func (agent *httpAgent) Send(Request, NfContext) (Response, error) {
	return nil, nil
}

func (agent *httpAgent) Register(service Service) error {
	return nil
}
