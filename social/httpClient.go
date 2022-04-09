/*
httpClient.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package social

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// HttpClient is an HTTP Request builder and sender.
type HttpClient struct {
	// http HttpClient for doing requests
	httpClient *http.Client
	// HTTP method (GET, POST, etc.)
	method string
	// raw url string for the requests
	rawURL string
	// stores key-values pairs to add to request's Headers
	header http.Header
	// url tagged query structs
	query []interface{}
	// HTTP Body
	body io.Reader
	// response decoder: default json decoder
	responseDecoder ResponseDecoder
}

// NewClient returns a new http client with a http DefaultClient.
func NewClient() *HttpClient {
	return &HttpClient{
		httpClient:      http.DefaultClient,
		method:          http.MethodGet,
		header:          make(http.Header),
		query:           make([]interface{}, 0),
		responseDecoder: jsonDecoder{},
	}
}

// parentClient := client.New().Base("https://api.de/")
// childClient1 := parentClient.New().Get("foo/")
// childClient2 := parentClient.New().POST("bar/")
//
// childClient1 and childClient2 will both use the same client with the same host
// but will send request to https://api.io/foo/ and https://api.io/bar/.
//
// If pointer values are used in client, mutating the parent original parent client
// will mutate the properties of child as well.
func (c *HttpClient) New() *HttpClient {
	// copy Headers pairs into new Header map
	headerCopy := make(http.Header)
	for k, v := range c.header {
		headerCopy[k] = v
	}
	return &HttpClient{
		httpClient:      c.httpClient,
		method:          c.method,
		rawURL:          c.rawURL,
		header:          headerCopy,
		query:           append([]interface{}{}, c.query...),
		responseDecoder: c.responseDecoder,
	}
}

// Base sets the rawURL.
func (c *HttpClient) Base(rawURL string) *HttpClient {
	c.rawURL = rawURL
	return c
}

// Body sets the HttpClient's body.
// If the provided body is also an io.Closer, the request Body will be closed
// by http.Client methods.
func (c *HttpClient) Body(body io.Reader) *HttpClient {
	if body == nil {
		return c
	}
	c.body = body

	return c
}

// Path extends the rawURL with the given path.
// If parsing errors occur, the rawURL is left unmodified.
func (c *HttpClient) Path(path string) *HttpClient {
	baseURL, baseErr := url.Parse(c.rawURL)
	pathURL, pathErr := url.Parse(path)
	if baseErr == nil && pathErr == nil {
		c.rawURL = baseURL.ResolveReference(pathURL).String()
	}
	return c
}

func (c *HttpClient) AddQuery(query interface{}) *HttpClient {
	if query != nil {
		c.query = append(c.query, query)
	}
	return c
}

// Get sets the Clients method to GET and sets the given pathURL.
func (c *HttpClient) Get(pathURL string) *HttpClient {
	c.method = "GET"
	return c.Path(pathURL)
}

// Post sets the Clients method to POST and sets the given pathURL.
func (c *HttpClient) Post(pathURL string) *HttpClient {
	c.method = "POST"
	return c.Path(pathURL)
}

// Put sets the Clients method to PUT and sets the given pathURL.
func (c *HttpClient) Put(pathURL string) *HttpClient {
	c.method = "PUT"
	return c.Path(pathURL)
}

// Patch sets the Clients method to PATCH and sets the given pathURL.
func (c *HttpClient) Patch(pathURL string) *HttpClient {
	c.method = "PATCH"
	return c.Path(pathURL)
}

// Delete sets the Clients method to DELETE and sets the given pathURL.
func (c *HttpClient) Delete(pathURL string) *HttpClient {
	c.method = "DELETE"
	return c.Path(pathURL)
}

// Request returns a new http.Request with the HttpClient properties.
// Returns errors if parsing the rawURL, encoding the query, encoding
// the body, or creating the http.Request.
func (c *HttpClient) Request() (*http.Request, error) {
	reqURL, err := url.Parse(c.rawURL)
	if err != nil {
		return nil, err
	}

	err = addQueryStructs(reqURL, c.query)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(c.method, reqURL.String(), c.body)
	if err != nil {
		return nil, err
	}

	addHeaders(req, c.header)
	return req, err
}

// Add adds the key, value pair in Headers, appending values for existing keys
// to the key's values. Header keys are canonicalized.
func (c *HttpClient) Add(key, value string) *HttpClient {
	c.header.Add(key, value)
	return c
}

// Set sets the key, value pair in Headers, replacing existing values
// associated with key. Header keys are canonicalized.
func (c *HttpClient) Set(key, value string) *HttpClient {
	c.header.Set(key, value)
	return c
}

// addQueryStructs parses url tagged query structs using go-querystring to
// encode them to url.Values and format them onto the url.RawQuery. Any
// query parsing or encoding errors are returned.
func addQueryStructs(reqURL *url.URL, queryStructs []interface{}) error {
	urlValues, err := url.ParseQuery(reqURL.RawQuery)
	if err != nil {
		return err
	}
	// encodes query structs into a url.Values map and merges maps
	for _, queryStruct := range queryStructs {
		queryValues, err := query.Values(queryStruct)
		if err != nil {
			return err
		}
		for key, values := range queryValues {
			for _, value := range values {
				urlValues.Add(key, value)
			}
		}
	}
	// url.Values format to a sorted "url encoded" string, e.g. "key=val&foo=bar"
	reqURL.RawQuery = urlValues.Encode()
	return nil
}

// addHeaders adds the key, value pairs from the given http.Header to the
// request. Values for existing keys are appended to the keys values.
func addHeaders(req *http.Request, header http.Header) {
	for key, values := range header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
}

// Do sends an HTTP request and returns the response. Success responses (2XX)
// are JSON decoded into the value pointed to by successV and other responses
// are JSON decoded into the value pointed to by failureV.
// If the status code of response is 204(no content), decoding is skipped.
// Any error sending the request or decoding the response is returned.
func (c *HttpClient) Do(req *http.Request, success interface{}, failure interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return resp, err
	}

	// when err is nil, resp contains a non-nil resp.Body which must be closed
	defer resp.Body.Close()

	/*bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)*/

	// The default HTTP client'c Transport may not
	// reuse HTTP/1.x "keep-alive" TCP connections if the Body is
	// not read to completion and closed. So simply dump it afterwards.
	// See: https://golang.org/pkg/net/http/#Response
	defer io.Copy(ioutil.Discard, resp.Body)

	// Don't try to decode on 204s
	if resp.StatusCode == http.StatusNoContent {
		return resp, nil
	}

	// Decode from json
	if success != nil {
		err = decodeResponse(resp, c.responseDecoder, success, failure)
	}
	return resp, err
}

// decodeResponse decodes response Body into the value pointed to by successV
// if the response is a success (2XX) or into the value pointed to by failureV
// otherwise. If the successV or failureV argument to decode into is nil,
// decoding is skipped.
// Caller is responsible for closing the resp.Body.
func decodeResponse(resp *http.Response, decoder ResponseDecoder, success interface{}, failure interface{}) error {
	if code := resp.StatusCode; 200 <= code && code <= 299 {
		if success != nil {
			return decoder.Decode(resp, success)
		}
	} else {
		return decoder.Decode(resp, failure)
	}
	return nil
}

// ResponseDecoder decodes http responses into struct values.
type ResponseDecoder interface {
	// Decode decodes the response into the value pointed to by v.
	Decode(resp *http.Response, v interface{}) error
}

// jsonDecoder decodes http response JSON into a JSON-tagged struct value.
type jsonDecoder struct{}

// Decode decodes the Response Body into the value pointed to by v.
// Caller must provide a non-nil v and close the resp.Body.
func (d jsonDecoder) Decode(resp *http.Response, v interface{}) error {
	return json.NewDecoder(resp.Body).Decode(v)
}
