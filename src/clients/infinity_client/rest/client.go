package rest

import (
	"github.com/valyala/fasthttp"
	"time"
)

const DefaultEndpoint = "http://localhost:8185/"

type Client struct {
	Endpoint    string
	HTTPC       *fasthttp.Client
	HTTPTimeout time.Duration
}

func NewClient(config *ClientConfig) *Client {
	return &Client{
		Endpoint:    getEndpoint(config),
		HTTPC:       new(fasthttp.Client),
		HTTPTimeout: 5 * time.Second,
	}
}
