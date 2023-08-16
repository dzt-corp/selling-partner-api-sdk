// Package smallAndLight provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package smallAndLight

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	runt "runtime"
	"strings"

	"github.com/amzapi/selling-partner-api-sdk/pkg/runtime"
)

// RequestBeforeFn  is the function signature for the RequestBefore callback function
type RequestBeforeFn func(ctx context.Context, req *http.Request) error

// ResponseAfterFn  is the function signature for the ResponseAfter callback function
type ResponseAfterFn func(ctx context.Context, rsp *http.Response) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Endpoint string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestBefore RequestBeforeFn

	// A callback for modifying response which are generated before sending over
	// the network.
	ResponseAfter ResponseAfterFn

	// The user agent header identifies your application, its version number, and the platform and programming language you are using.
	// You must include a user agent header in each request submitted to the sales partner API.
	UserAgent string
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(endpoint string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Endpoint: endpoint,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the endpoint URL always has a trailing slash
	if !strings.HasSuffix(client.Endpoint, "/") {
		client.Endpoint += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	// setting the default useragent
	if client.UserAgent == "" {
		client.UserAgent = fmt.Sprintf("selling-partner-api-sdk/v1.0 (Language=%s; Platform=%s-%s)", strings.Replace(runt.Version(), "go", "go/", -1), runt.GOOS, runt.GOARCH)
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithUserAgent set up useragent
// add user agent to every request automatically
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.UserAgent = userAgent
		return nil
	}
}

// WithRequestBefore allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestBefore(fn RequestBeforeFn) ClientOption {
	return func(c *Client) error {
		c.RequestBefore = fn
		return nil
	}
}

// WithResponseAfter allows setting up a callback function, which will be
// called right after get response the request. This can be used to log.
func WithResponseAfter(fn ResponseAfterFn) ClientOption {
	return func(c *Client) error {
		c.ResponseAfter = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetSmallAndLightEligibilityBySellerSKU request
	GetSmallAndLightEligibilityBySellerSKU(ctx context.Context, sellerSKU string, params *GetSmallAndLightEligibilityBySellerSKUParams) (*http.Response, error)

	// DeleteSmallAndLightEnrollmentBySellerSKU request
	DeleteSmallAndLightEnrollmentBySellerSKU(ctx context.Context, sellerSKU string, params *DeleteSmallAndLightEnrollmentBySellerSKUParams) (*http.Response, error)

	// GetSmallAndLightEnrollmentBySellerSKU request
	GetSmallAndLightEnrollmentBySellerSKU(ctx context.Context, sellerSKU string, params *GetSmallAndLightEnrollmentBySellerSKUParams) (*http.Response, error)

	// PutSmallAndLightEnrollmentBySellerSKU request
	PutSmallAndLightEnrollmentBySellerSKU(ctx context.Context, sellerSKU string, params *PutSmallAndLightEnrollmentBySellerSKUParams) (*http.Response, error)

	// GetSmallAndLightFeePreview request  with any body
	GetSmallAndLightFeePreviewWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	GetSmallAndLightFeePreview(ctx context.Context, body GetSmallAndLightFeePreviewJSONRequestBody) (*http.Response, error)
}

func (c *Client) GetSmallAndLightEligibilityBySellerSKU(ctx context.Context, sellerSKU string, params *GetSmallAndLightEligibilityBySellerSKUParams) (*http.Response, error) {
	req, err := NewGetSmallAndLightEligibilityBySellerSKURequest(c.Endpoint, sellerSKU, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) DeleteSmallAndLightEnrollmentBySellerSKU(ctx context.Context, sellerSKU string, params *DeleteSmallAndLightEnrollmentBySellerSKUParams) (*http.Response, error) {
	req, err := NewDeleteSmallAndLightEnrollmentBySellerSKURequest(c.Endpoint, sellerSKU, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetSmallAndLightEnrollmentBySellerSKU(ctx context.Context, sellerSKU string, params *GetSmallAndLightEnrollmentBySellerSKUParams) (*http.Response, error) {
	req, err := NewGetSmallAndLightEnrollmentBySellerSKURequest(c.Endpoint, sellerSKU, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) PutSmallAndLightEnrollmentBySellerSKU(ctx context.Context, sellerSKU string, params *PutSmallAndLightEnrollmentBySellerSKUParams) (*http.Response, error) {
	req, err := NewPutSmallAndLightEnrollmentBySellerSKURequest(c.Endpoint, sellerSKU, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetSmallAndLightFeePreviewWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewGetSmallAndLightFeePreviewRequestWithBody(c.Endpoint, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetSmallAndLightFeePreview(ctx context.Context, body GetSmallAndLightFeePreviewJSONRequestBody) (*http.Response, error) {
	req, err := NewGetSmallAndLightFeePreviewRequest(c.Endpoint, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// NewGetSmallAndLightEligibilityBySellerSKURequest generates requests for GetSmallAndLightEligibilityBySellerSKU
func NewGetSmallAndLightEligibilityBySellerSKURequest(endpoint string, sellerSKU string, params *GetSmallAndLightEligibilityBySellerSKUParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "sellerSKU", sellerSKU)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/fba/smallAndLight/v1/eligibilities/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", true, "marketplaceIds", params.MarketplaceIds); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewDeleteSmallAndLightEnrollmentBySellerSKURequest generates requests for DeleteSmallAndLightEnrollmentBySellerSKU
func NewDeleteSmallAndLightEnrollmentBySellerSKURequest(endpoint string, sellerSKU string, params *DeleteSmallAndLightEnrollmentBySellerSKUParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "sellerSKU", sellerSKU)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/fba/smallAndLight/v1/enrollments/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", true, "marketplaceIds", params.MarketplaceIds); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("DELETE", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetSmallAndLightEnrollmentBySellerSKURequest generates requests for GetSmallAndLightEnrollmentBySellerSKU
func NewGetSmallAndLightEnrollmentBySellerSKURequest(endpoint string, sellerSKU string, params *GetSmallAndLightEnrollmentBySellerSKUParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "sellerSKU", sellerSKU)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/fba/smallAndLight/v1/enrollments/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", true, "marketplaceIds", params.MarketplaceIds); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPutSmallAndLightEnrollmentBySellerSKURequest generates requests for PutSmallAndLightEnrollmentBySellerSKU
func NewPutSmallAndLightEnrollmentBySellerSKURequest(endpoint string, sellerSKU string, params *PutSmallAndLightEnrollmentBySellerSKUParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "sellerSKU", sellerSKU)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/fba/smallAndLight/v1/enrollments/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", true, "marketplaceIds", params.MarketplaceIds); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("PUT", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetSmallAndLightFeePreviewRequest calls the generic GetSmallAndLightFeePreview builder with application/json body
func NewGetSmallAndLightFeePreviewRequest(endpoint string, body GetSmallAndLightFeePreviewJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewGetSmallAndLightFeePreviewRequestWithBody(endpoint, "application/json", bodyReader)
}

// NewGetSmallAndLightFeePreviewRequestWithBody generates requests for GetSmallAndLightFeePreview with any type of body
func NewGetSmallAndLightFeePreviewRequestWithBody(endpoint string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/fba/smallAndLight/v1/feePreviews")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(endpoint string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(endpoint, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Endpoint = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetSmallAndLightEligibilityBySellerSKU request
	GetSmallAndLightEligibilityBySellerSKUWithResponse(ctx context.Context, sellerSKU string, params *GetSmallAndLightEligibilityBySellerSKUParams) (*GetSmallAndLightEligibilityBySellerSKUResp, error)

	// DeleteSmallAndLightEnrollmentBySellerSKU request
	DeleteSmallAndLightEnrollmentBySellerSKUWithResponse(ctx context.Context, sellerSKU string, params *DeleteSmallAndLightEnrollmentBySellerSKUParams) (*DeleteSmallAndLightEnrollmentBySellerSKUResp, error)

	// GetSmallAndLightEnrollmentBySellerSKU request
	GetSmallAndLightEnrollmentBySellerSKUWithResponse(ctx context.Context, sellerSKU string, params *GetSmallAndLightEnrollmentBySellerSKUParams) (*GetSmallAndLightEnrollmentBySellerSKUResp, error)

	// PutSmallAndLightEnrollmentBySellerSKU request
	PutSmallAndLightEnrollmentBySellerSKUWithResponse(ctx context.Context, sellerSKU string, params *PutSmallAndLightEnrollmentBySellerSKUParams) (*PutSmallAndLightEnrollmentBySellerSKUResp, error)

	// GetSmallAndLightFeePreview request  with any body
	GetSmallAndLightFeePreviewWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*GetSmallAndLightFeePreviewResp, error)

	GetSmallAndLightFeePreviewWithResponse(ctx context.Context, body GetSmallAndLightFeePreviewJSONRequestBody) (*GetSmallAndLightFeePreviewResp, error)
}

type GetSmallAndLightEligibilityBySellerSKUResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SmallAndLightEligibility
	JSON400      *ErrorList
	JSON403      *ErrorList
	JSON404      *ErrorList
	JSON413      *ErrorList
	JSON415      *ErrorList
	JSON429      *ErrorList
	JSON500      *ErrorList
	JSON503      *ErrorList
}

// Status returns HTTPResponse.Status
func (r GetSmallAndLightEligibilityBySellerSKUResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetSmallAndLightEligibilityBySellerSKUResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteSmallAndLightEnrollmentBySellerSKUResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *ErrorList
	JSON403      *ErrorList
	JSON404      *ErrorList
	JSON413      *ErrorList
	JSON415      *ErrorList
	JSON429      *ErrorList
	JSON500      *ErrorList
	JSON503      *ErrorList
}

// Status returns HTTPResponse.Status
func (r DeleteSmallAndLightEnrollmentBySellerSKUResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteSmallAndLightEnrollmentBySellerSKUResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetSmallAndLightEnrollmentBySellerSKUResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SmallAndLightEnrollment
	JSON400      *ErrorList
	JSON403      *ErrorList
	JSON404      *ErrorList
	JSON413      *ErrorList
	JSON415      *ErrorList
	JSON429      *ErrorList
	JSON500      *ErrorList
	JSON503      *ErrorList
}

// Status returns HTTPResponse.Status
func (r GetSmallAndLightEnrollmentBySellerSKUResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetSmallAndLightEnrollmentBySellerSKUResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PutSmallAndLightEnrollmentBySellerSKUResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SmallAndLightEnrollment
	JSON400      *ErrorList
	JSON403      *ErrorList
	JSON404      *ErrorList
	JSON413      *ErrorList
	JSON415      *ErrorList
	JSON429      *ErrorList
	JSON500      *ErrorList
	JSON503      *ErrorList
}

// Status returns HTTPResponse.Status
func (r PutSmallAndLightEnrollmentBySellerSKUResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PutSmallAndLightEnrollmentBySellerSKUResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetSmallAndLightFeePreviewResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SmallAndLightFeePreviews
	JSON400      *ErrorList
	JSON401      *ErrorList
	JSON403      *ErrorList
	JSON404      *ErrorList
	JSON429      *ErrorList
	JSON500      *ErrorList
	JSON503      *ErrorList
}

// Status returns HTTPResponse.Status
func (r GetSmallAndLightFeePreviewResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetSmallAndLightFeePreviewResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetSmallAndLightEligibilityBySellerSKUWithResponse request returning *GetSmallAndLightEligibilityBySellerSKUResponse
func (c *ClientWithResponses) GetSmallAndLightEligibilityBySellerSKUWithResponse(ctx context.Context, sellerSKU string, params *GetSmallAndLightEligibilityBySellerSKUParams) (*GetSmallAndLightEligibilityBySellerSKUResp, error) {
	rsp, err := c.GetSmallAndLightEligibilityBySellerSKU(ctx, sellerSKU, params)
	if err != nil {
		return nil, err
	}
	return ParseGetSmallAndLightEligibilityBySellerSKUResp(rsp)
}

// DeleteSmallAndLightEnrollmentBySellerSKUWithResponse request returning *DeleteSmallAndLightEnrollmentBySellerSKUResponse
func (c *ClientWithResponses) DeleteSmallAndLightEnrollmentBySellerSKUWithResponse(ctx context.Context, sellerSKU string, params *DeleteSmallAndLightEnrollmentBySellerSKUParams) (*DeleteSmallAndLightEnrollmentBySellerSKUResp, error) {
	rsp, err := c.DeleteSmallAndLightEnrollmentBySellerSKU(ctx, sellerSKU, params)
	if err != nil {
		return nil, err
	}
	return ParseDeleteSmallAndLightEnrollmentBySellerSKUResp(rsp)
}

// GetSmallAndLightEnrollmentBySellerSKUWithResponse request returning *GetSmallAndLightEnrollmentBySellerSKUResponse
func (c *ClientWithResponses) GetSmallAndLightEnrollmentBySellerSKUWithResponse(ctx context.Context, sellerSKU string, params *GetSmallAndLightEnrollmentBySellerSKUParams) (*GetSmallAndLightEnrollmentBySellerSKUResp, error) {
	rsp, err := c.GetSmallAndLightEnrollmentBySellerSKU(ctx, sellerSKU, params)
	if err != nil {
		return nil, err
	}
	return ParseGetSmallAndLightEnrollmentBySellerSKUResp(rsp)
}

// PutSmallAndLightEnrollmentBySellerSKUWithResponse request returning *PutSmallAndLightEnrollmentBySellerSKUResponse
func (c *ClientWithResponses) PutSmallAndLightEnrollmentBySellerSKUWithResponse(ctx context.Context, sellerSKU string, params *PutSmallAndLightEnrollmentBySellerSKUParams) (*PutSmallAndLightEnrollmentBySellerSKUResp, error) {
	rsp, err := c.PutSmallAndLightEnrollmentBySellerSKU(ctx, sellerSKU, params)
	if err != nil {
		return nil, err
	}
	return ParsePutSmallAndLightEnrollmentBySellerSKUResp(rsp)
}

// GetSmallAndLightFeePreviewWithBodyWithResponse request with arbitrary body returning *GetSmallAndLightFeePreviewResponse
func (c *ClientWithResponses) GetSmallAndLightFeePreviewWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*GetSmallAndLightFeePreviewResp, error) {
	rsp, err := c.GetSmallAndLightFeePreviewWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseGetSmallAndLightFeePreviewResp(rsp)
}

func (c *ClientWithResponses) GetSmallAndLightFeePreviewWithResponse(ctx context.Context, body GetSmallAndLightFeePreviewJSONRequestBody) (*GetSmallAndLightFeePreviewResp, error) {
	rsp, err := c.GetSmallAndLightFeePreview(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseGetSmallAndLightFeePreviewResp(rsp)
}

// ParseGetSmallAndLightEligibilityBySellerSKUResp parses an HTTP response from a GetSmallAndLightEligibilityBySellerSKUWithResponse call
func ParseGetSmallAndLightEligibilityBySellerSKUResp(rsp *http.Response) (*GetSmallAndLightEligibilityBySellerSKUResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetSmallAndLightEligibilityBySellerSKUResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SmallAndLightEligibility
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 413:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON413 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseDeleteSmallAndLightEnrollmentBySellerSKUResp parses an HTTP response from a DeleteSmallAndLightEnrollmentBySellerSKUWithResponse call
func ParseDeleteSmallAndLightEnrollmentBySellerSKUResp(rsp *http.Response) (*DeleteSmallAndLightEnrollmentBySellerSKUResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &DeleteSmallAndLightEnrollmentBySellerSKUResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 413:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON413 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseGetSmallAndLightEnrollmentBySellerSKUResp parses an HTTP response from a GetSmallAndLightEnrollmentBySellerSKUWithResponse call
func ParseGetSmallAndLightEnrollmentBySellerSKUResp(rsp *http.Response) (*GetSmallAndLightEnrollmentBySellerSKUResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetSmallAndLightEnrollmentBySellerSKUResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SmallAndLightEnrollment
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 413:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON413 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParsePutSmallAndLightEnrollmentBySellerSKUResp parses an HTTP response from a PutSmallAndLightEnrollmentBySellerSKUWithResponse call
func ParsePutSmallAndLightEnrollmentBySellerSKUResp(rsp *http.Response) (*PutSmallAndLightEnrollmentBySellerSKUResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &PutSmallAndLightEnrollmentBySellerSKUResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SmallAndLightEnrollment
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 413:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON413 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseGetSmallAndLightFeePreviewResp parses an HTTP response from a GetSmallAndLightFeePreviewWithResponse call
func ParseGetSmallAndLightFeePreviewResp(rsp *http.Response) (*GetSmallAndLightFeePreviewResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetSmallAndLightFeePreviewResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SmallAndLightFeePreviews
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}
