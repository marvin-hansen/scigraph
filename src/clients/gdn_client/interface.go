package gdn_client

type Requester interface {
	Path() string      // path relative to URI i.e. /data/function/
	Method() string    // http method i.e. get, post, put, delete etc
	Query() string     // query to append after ? i.e. query?"my query"
	Payload() []byte   // payload i.e. file to upload
	ResponseCode() int // expected http status code the server should normally return.
}

type Responder interface {
	GetRawMessage() []byte
	SetRawMessage(raw []byte)
}
