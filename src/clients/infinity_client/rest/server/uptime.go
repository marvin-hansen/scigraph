package server

import (
	"net/http"
	"time"
)

func NewRequestForUptime() *RequestForUptime {
	return &RequestForUptime{}
}

type RequestForUptime struct{}

func (req *RequestForUptime) Path() string {
	return ""
}

func (req *RequestForUptime) Method() string {
	return http.MethodGet
}

func (req *RequestForUptime) Query() string {
	return ""
}

func (req *RequestForUptime) Payload() []byte {
	return nil
}

func NewResponseForUptime() *ResponseForUptime {
	return &ResponseForUptime{}
}

type ResponseForUptime Uptime

type Uptime struct {
	Uptime   time.Time `json:"uptime"`
	Versions []string  `json:"versions"`
}
