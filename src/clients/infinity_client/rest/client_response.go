package rest

type Response struct {
	Result  interface{} `json:"result,omitempty"`
	Error   string      `json:"error,omitempty"`
	Success bool        `json:"success"`
}

func (c *Client) newResponse() *Response {
	return &Response{}
}
