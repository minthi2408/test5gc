package telemetry

import (
	"etri5gc/fabric/common"
	"time"
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
