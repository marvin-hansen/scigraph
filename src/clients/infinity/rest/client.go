package rest

import (
	"github.com/valyala/fasthttp"
	"time"
)

const ENDPOINT = "https://ftx.com/api"

type Client struct {
	Endpoint    string
	HTTPC       *fasthttp.Client
	HTTPTimeout time.Duration
}

func New() *Client {
	hc := new(fasthttp.Client)

	return &Client{
		Endpoint:    ENDPOINT,
		HTTPC:       hc,
		HTTPTimeout: 5 * time.Second,
	}
}
