package telemetry

import (
    "time"
)
//TODO: to be defined
type Report interface {
	Time() time.Time
}

type Writer interface {
	Write(Report)
}

