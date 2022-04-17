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
		//println("Request error")
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
		//println("http req error")
		apiErr := getApiError(res)
		return nil, apiErr
	}

	fmt.Printf("%+v\n", string(res.Body()))

	if res.StatusCode() != r.ResponseCode() {
		var resp = new(Response)
		if jsonErr := json.Unmarshal(res.Body(), resp); jsonErr != nil {
			apiErr := getApiError(res)
			return nil, apiErr
		}

		if !resp.Success {
			//println("Response code not success")
			apiErr := getApiError(res)
			return nil, apiErr
		}
	}

	return res, nil
}

func getApiError(res *fasthttp.Response) *APIError {
	apiErr := &APIError{}
	_ = decode(res.Body(), apiErr)
	return apiErr
}

func getUri(endpoint string, r Requester) *url.URL {
	var u = new(url.URL)
	if r.HasQueryParameter() {
		u, _ = url.ParseRequestURI(endpoint + r.GetQueryParameter())
	} else {
		u, _ = url.ParseRequestURI(endpoint)
	}
	u.Path = r.Path()
	return u
}

func (c *Client) newRequest(r Requester) *fasthttp.Request {

	u := getUri(c.Endpoint, r)
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(r.Method())
	req.SetRequestURI(u.String())
	body := r.Payload()
	req.SetBody(body)

	nonce := fmt.Sprintf("%d", int64(time.Now().UTC().UnixNano()/int64(time.Millisecond)))
	payload := nonce + r.Method() + u.Path

	if u.RawQuery != "" {
		u.RawQuery = r.Query()
		payload += "?" + u.RawQuery
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "apikey "+c.getApiKey())

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
