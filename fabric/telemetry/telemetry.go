package telemetry

import (
	"etrib5gc/fabric/common"
	"time"

	log "github.com/sirupsen/logrus"
)

//TODO: telemetry module will be implemented later

type Report interface {
	Time() time.Time
}

type Writer interface {
	common.InternalService
	Write(Report)
}

func NewWriter() Writer {
	log.Info("NewWriter")
	return &dumpTelemetry{}
}

type dumpTelemetry struct {
}

func (tel *dumpTelemetry) Write(r Report) {
}

func (tel *dumpTelemetry) Start() error {
	return nil
}

func (tel *dumpTelemetry) Terminate() {
}
