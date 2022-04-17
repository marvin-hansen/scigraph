package gdn_client

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"io"
	"net/url"
	"time"
)

// used for JSON unmarshaling
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// shutdownRequest doesn't get an actual response from the server and thus
// doesn't process a server reply. It only returns an error if there was
// 1) A network transmission error i.e. offline
// 2) An incorrect server URI
// 3) The server was already taken offline before
// https://support.objectivity.com/sites/default/files/docs/ig/latest/index.html#page/topics%2Frest%2FrestVersionShutdownPOST.html%23
func (c *Client) requestWithoutReturnValue(r Requester) error {
	req := c.newRequest(r)
	res := fasthttp.AcquireResponse()
	err := c.HTTPC.DoTimeout(req, res, c.HTTPTimeout)
	return checkError(err)
}

func (c *Client) request(req Requester, results Responder) error {
	res, reqErr := c.do(req)
	if reqErr != nil {
		return reqErr
	}

	results.SetRawMessage(res.Body())

	decErr := decode(res.Body(), results)
	if decErr != nil {
		return decErr
	}

	return nil
}

func (c *Client) requestQuery(req Requester, results Responder) error {
	res, reqErr := c.do(req)
	if reqErr != nil {
		println("Request Error")
		return reqErr
	}

	// Query returns an array of various types, so we have to prefix
	// an extra node "entries" to Marshal the result into a struct that contains the raw message
	queryRes := []byte("{" + "\"entries\":" + string(res.Body()) + "}")
	results.SetRawMessage(queryRes)

	decErr := decode(queryRes, results)
	if decErr != nil {
		println("Decode error")
		return decErr
	}

	return nil
}

// do build & executes the actual request from the rquester
// requester - implementation
// targetStatusCode the expected http status code i.e. 200
func (c *Client) do(r Requester) (*fasthttp.Response, error) {
	req := c.newRequest(r)
	// fmt.Printf("Path: %+v\n", r.Path())

	// fasthttp for http2.0
	res := fasthttp.AcquireResponse()
	err := c.HTTPC.DoTimeout(req, res, c.HTTPTimeout)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("%+v\n", string(res.Body()))
	// no usefull headers

	if res.StatusCode() != r.ResponseCode() {
		var resp = new(Response)
		if jsonErr := json.Unmarshal(res.Body(), resp); jsonErr != nil {
			return nil, &APIError{
				Status:  res.StatusCode(),
				Message: jsonErr.Error(),
			}
		}

		if !resp.Success {
			return nil, &APIError{
				Status:  res.StatusCode(),
				Message: resp.Error,
			}
		}
	}

	return res, nil
}

func (c *Client) newRequest(r Requester) *fasthttp.Request {
	// avoid Pointer's butting
	u, _ := url.ParseRequestURI(c.Endpoint)
	u.Path = u.Path + r.Path()

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(r.Method())
	req.SetRequestURI(u.String())
	body := r.Payload()
	req.SetBody(body)

	nonce := fmt.Sprintf("%d", int64(time.Now().UTC().UnixNano()/int64(time.Millisecond)))
	payload := nonce + r.Method() + u.Path

	u.RawQuery = r.Query()
	if u.RawQuery != "" {
		payload += "?" + u.RawQuery
	}

	payload += string(body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("authorization", c.apiKey)

	return req
}

func decode(inputJson []byte, outputObject interface{}) error {

	var out io.Writer
	enc := json.NewEncoder(out)
	enc.SetIndent("", "    ")
	if err := enc.Encode(inputJson); err != nil {
		return err
	}

	err := json.Unmarshal(inputJson, outputObject)
	if err != nil {
		return fmt.Errorf("decode error")
	} else {
		return nil
	}
}
