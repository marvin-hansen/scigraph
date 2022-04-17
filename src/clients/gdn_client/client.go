package gdn_client

import (
	"github.com/valyala/fasthttp"
	"time"
)

type Client struct {
	apiKey      string
	Endpoint    string
	HTTPC       *fasthttp.Client
	HTTPTimeout time.Duration
}

func NewClient(config *ClientConfig) *Client {
	return &Client{
		apiKey:      apiKEY,
		Endpoint:    getEndpoint(config),
		HTTPTimeout: getTimeout(config),
		HTTPC:       new(fasthttp.Client),
	}
}
