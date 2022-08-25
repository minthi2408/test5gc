package telemetry

import (
	"time"
)

//TODO: telemetry module will be implemented later

type Report interface {
	Time() time.Time
}

type Writer interface {
	Write(Report)
}

func NewWriter() Writer {
	return &dumpTelemetry{}
}

type dumpTelemetry struct {
}

func (tel *dumpTelemetry) Write(r Report) {
}
