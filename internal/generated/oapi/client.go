// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package oapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

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
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
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

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// Provision request
	Provision(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetIoTInfoWithBody request with any body
	GetIoTInfoWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	GetIoTInfo(ctx context.Context, body GetIoTInfoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetDevices request
	GetDevices(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BeginLoginWithBody request with any body
	BeginLoginWithBody(ctx context.Context, params *BeginLoginParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	BeginLogin(ctx context.Context, params *BeginLoginParams, body BeginLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetUserStatusWithBody request with any body
	GetUserStatusWithBody(ctx context.Context, params *GetUserStatusParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	GetUserStatus(ctx context.Context, params *GetUserStatusParams, body GetUserStatusJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CompleteLoginWithBody request with any body
	CompleteLoginWithBody(ctx context.Context, params *CompleteLoginParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CompleteLogin(ctx context.Context, params *CompleteLoginParams, body CompleteLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) Provision(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewProvisionRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetIoTInfoWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetIoTInfoRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetIoTInfo(ctx context.Context, body GetIoTInfoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetIoTInfoRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetDevices(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetDevicesRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) BeginLoginWithBody(ctx context.Context, params *BeginLoginParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBeginLoginRequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) BeginLogin(ctx context.Context, params *BeginLoginParams, body BeginLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBeginLoginRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetUserStatusWithBody(ctx context.Context, params *GetUserStatusParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserStatusRequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetUserStatus(ctx context.Context, params *GetUserStatusParams, body GetUserStatusJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserStatusRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CompleteLoginWithBody(ctx context.Context, params *CompleteLoginParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCompleteLoginRequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CompleteLogin(ctx context.Context, params *CompleteLoginParams, body CompleteLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCompleteLoginRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewProvisionRequest generates requests for Provision
func NewProvisionRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/provisioningservice/application/Android/version")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetIoTInfoRequest calls the generic GetIoTInfo builder with application/json body
func NewGetIoTInfoRequest(server string, body GetIoTInfoJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewGetIoTInfoRequestWithBody(server, "application/json", bodyReader)
}

// NewGetIoTInfoRequestWithBody generates requests for GetIoTInfo with any type of body
func NewGetIoTInfoRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v2/authorize/iot-credentials")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetDevicesRequest generates requests for GetDevices
func NewGetDevicesRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v3/manifest")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewBeginLoginRequest calls the generic BeginLogin builder with application/json body
func NewBeginLoginRequest(server string, params *BeginLoginParams, body BeginLoginJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewBeginLoginRequestWithBody(server, params, "application/json", bodyReader)
}

// NewBeginLoginRequestWithBody generates requests for BeginLogin with any type of body
func NewBeginLoginRequestWithBody(server string, params *BeginLoginParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v3/userregistration/email/auth")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Country != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "country", runtime.ParamLocationQuery, *params.Country); err != nil {
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

		}

		if params.Culture != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "culture", runtime.ParamLocationQuery, *params.Culture); err != nil {
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

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		var headerParam0 string

		headerParam0, err = runtime.StyleParamWithLocation("simple", false, "User-Agent", runtime.ParamLocationHeader, params.UserAgent)
		if err != nil {
			return nil, err
		}

		req.Header.Set("User-Agent", headerParam0)

	}

	return req, nil
}

// NewGetUserStatusRequest calls the generic GetUserStatus builder with application/json body
func NewGetUserStatusRequest(server string, params *GetUserStatusParams, body GetUserStatusJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewGetUserStatusRequestWithBody(server, params, "application/json", bodyReader)
}

// NewGetUserStatusRequestWithBody generates requests for GetUserStatus with any type of body
func NewGetUserStatusRequestWithBody(server string, params *GetUserStatusParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v3/userregistration/email/userstatus")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Country != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "country", runtime.ParamLocationQuery, *params.Country); err != nil {
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

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		var headerParam0 string

		headerParam0, err = runtime.StyleParamWithLocation("simple", false, "User-Agent", runtime.ParamLocationHeader, params.UserAgent)
		if err != nil {
			return nil, err
		}

		req.Header.Set("User-Agent", headerParam0)

	}

	return req, nil
}

// NewCompleteLoginRequest calls the generic CompleteLogin builder with application/json body
func NewCompleteLoginRequest(server string, params *CompleteLoginParams, body CompleteLoginJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCompleteLoginRequestWithBody(server, params, "application/json", bodyReader)
}

// NewCompleteLoginRequestWithBody generates requests for CompleteLogin with any type of body
func NewCompleteLoginRequestWithBody(server string, params *CompleteLoginParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v3/userregistration/email/verify")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Country != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "country", runtime.ParamLocationQuery, *params.Country); err != nil {
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

		}

		if params.Culture != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "culture", runtime.ParamLocationQuery, *params.Culture); err != nil {
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

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		var headerParam0 string

		headerParam0, err = runtime.StyleParamWithLocation("simple", false, "User-Agent", runtime.ParamLocationHeader, params.UserAgent)
		if err != nil {
			return nil, err
		}

		req.Header.Set("User-Agent", headerParam0)

	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
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
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// ProvisionWithResponse request
	ProvisionWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ProvisionResponse, error)

	// GetIoTInfoWithBodyWithResponse request with any body
	GetIoTInfoWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*GetIoTInfoResponse, error)

	GetIoTInfoWithResponse(ctx context.Context, body GetIoTInfoJSONRequestBody, reqEditors ...RequestEditorFn) (*GetIoTInfoResponse, error)

	// GetDevicesWithResponse request
	GetDevicesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetDevicesResponse, error)

	// BeginLoginWithBodyWithResponse request with any body
	BeginLoginWithBodyWithResponse(ctx context.Context, params *BeginLoginParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BeginLoginResponse, error)

	BeginLoginWithResponse(ctx context.Context, params *BeginLoginParams, body BeginLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*BeginLoginResponse, error)

	// GetUserStatusWithBodyWithResponse request with any body
	GetUserStatusWithBodyWithResponse(ctx context.Context, params *GetUserStatusParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*GetUserStatusResponse, error)

	GetUserStatusWithResponse(ctx context.Context, params *GetUserStatusParams, body GetUserStatusJSONRequestBody, reqEditors ...RequestEditorFn) (*GetUserStatusResponse, error)

	// CompleteLoginWithBodyWithResponse request with any body
	CompleteLoginWithBodyWithResponse(ctx context.Context, params *CompleteLoginParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CompleteLoginResponse, error)

	CompleteLoginWithResponse(ctx context.Context, params *CompleteLoginParams, body CompleteLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*CompleteLoginResponse, error)
}

type ProvisionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Provision
}

// Status returns HTTPResponse.Status
func (r ProvisionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ProvisionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetIoTInfoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *IoTData
}

// Status returns HTTPResponse.Status
func (r GetIoTInfoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetIoTInfoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetDevicesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Devices
}

// Status returns HTTPResponse.Status
func (r GetDevicesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetDevicesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type BeginLoginResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *LoginChallenge
}

// Status returns HTTPResponse.Status
func (r BeginLoginResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r BeginLoginResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserStatusResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *UserStatus
}

// Status returns HTTPResponse.Status
func (r GetUserStatusResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetUserStatusResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CompleteLoginResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *LoginComplete
}

// Status returns HTTPResponse.Status
func (r CompleteLoginResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CompleteLoginResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ProvisionWithResponse request returning *ProvisionResponse
func (c *ClientWithResponses) ProvisionWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ProvisionResponse, error) {
	rsp, err := c.Provision(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseProvisionResponse(rsp)
}

// GetIoTInfoWithBodyWithResponse request with arbitrary body returning *GetIoTInfoResponse
func (c *ClientWithResponses) GetIoTInfoWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*GetIoTInfoResponse, error) {
	rsp, err := c.GetIoTInfoWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetIoTInfoResponse(rsp)
}

func (c *ClientWithResponses) GetIoTInfoWithResponse(ctx context.Context, body GetIoTInfoJSONRequestBody, reqEditors ...RequestEditorFn) (*GetIoTInfoResponse, error) {
	rsp, err := c.GetIoTInfo(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetIoTInfoResponse(rsp)
}

// GetDevicesWithResponse request returning *GetDevicesResponse
func (c *ClientWithResponses) GetDevicesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetDevicesResponse, error) {
	rsp, err := c.GetDevices(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetDevicesResponse(rsp)
}

// BeginLoginWithBodyWithResponse request with arbitrary body returning *BeginLoginResponse
func (c *ClientWithResponses) BeginLoginWithBodyWithResponse(ctx context.Context, params *BeginLoginParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BeginLoginResponse, error) {
	rsp, err := c.BeginLoginWithBody(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBeginLoginResponse(rsp)
}

func (c *ClientWithResponses) BeginLoginWithResponse(ctx context.Context, params *BeginLoginParams, body BeginLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*BeginLoginResponse, error) {
	rsp, err := c.BeginLogin(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBeginLoginResponse(rsp)
}

// GetUserStatusWithBodyWithResponse request with arbitrary body returning *GetUserStatusResponse
func (c *ClientWithResponses) GetUserStatusWithBodyWithResponse(ctx context.Context, params *GetUserStatusParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*GetUserStatusResponse, error) {
	rsp, err := c.GetUserStatusWithBody(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserStatusResponse(rsp)
}

func (c *ClientWithResponses) GetUserStatusWithResponse(ctx context.Context, params *GetUserStatusParams, body GetUserStatusJSONRequestBody, reqEditors ...RequestEditorFn) (*GetUserStatusResponse, error) {
	rsp, err := c.GetUserStatus(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserStatusResponse(rsp)
}

// CompleteLoginWithBodyWithResponse request with arbitrary body returning *CompleteLoginResponse
func (c *ClientWithResponses) CompleteLoginWithBodyWithResponse(ctx context.Context, params *CompleteLoginParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CompleteLoginResponse, error) {
	rsp, err := c.CompleteLoginWithBody(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCompleteLoginResponse(rsp)
}

func (c *ClientWithResponses) CompleteLoginWithResponse(ctx context.Context, params *CompleteLoginParams, body CompleteLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*CompleteLoginResponse, error) {
	rsp, err := c.CompleteLogin(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCompleteLoginResponse(rsp)
}

// ParseProvisionResponse parses an HTTP response from a ProvisionWithResponse call
func ParseProvisionResponse(rsp *http.Response) (*ProvisionResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ProvisionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Provision
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetIoTInfoResponse parses an HTTP response from a GetIoTInfoWithResponse call
func ParseGetIoTInfoResponse(rsp *http.Response) (*GetIoTInfoResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetIoTInfoResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest IoTData
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetDevicesResponse parses an HTTP response from a GetDevicesWithResponse call
func ParseGetDevicesResponse(rsp *http.Response) (*GetDevicesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetDevicesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Devices
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseBeginLoginResponse parses an HTTP response from a BeginLoginWithResponse call
func ParseBeginLoginResponse(rsp *http.Response) (*BeginLoginResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &BeginLoginResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest LoginChallenge
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetUserStatusResponse parses an HTTP response from a GetUserStatusWithResponse call
func ParseGetUserStatusResponse(rsp *http.Response) (*GetUserStatusResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetUserStatusResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest UserStatus
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseCompleteLoginResponse parses an HTTP response from a CompleteLoginWithResponse call
func ParseCompleteLoginResponse(rsp *http.Response) (*CompleteLoginResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CompleteLoginResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest LoginComplete
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}