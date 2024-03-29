// Package service provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

const (
	ApiKeyAuthScopes = "ApiKeyAuth.Scopes"
)

// Network defines model for Network.
type Network struct {
	// Embedded fields due to inline allOf schema

	// Let ZeroTier modify the system's DNS settings
	AllowDNS *bool `json:"allowDNS,omitempty"`

	// Let ZeroTier to modify the system's default route.
	AllowDefault *bool `json:"allowDefault,omitempty"`

	// Let ZeroTier to manage IP addresses and Route assignments that aren't in private ranges (rfc1918).
	AllowGlobal *bool `json:"allowGlobal,omitempty"`

	// Let ZeroTier to manage IP addresses and Route assignments.
	AllowManaged *bool `json:"allowManaged,omitempty"`
	// Embedded fields due to inline allOf schema

	AssignedAddresses *[]string `json:"assignedAddresses,omitempty"`
	Bridge            *bool     `json:"bridge,omitempty"`
	BroadcastEnabled  *bool     `json:"broadcastEnabled,omitempty"`
	Dns               *struct {
		Domain  *string   `json:"domain,omitempty"`
		Servers *[]string `json:"servers,omitempty"`
	} `json:"dns,omitempty"`
	Id *string `json:"id,omitempty"`

	// MAC address for this network's interface
	Mac                    *string `json:"mac,omitempty"`
	Mtu                    *int    `json:"mtu,omitempty"`
	MulticastSubscriptions *[]struct {
		Adi *int64  `json:"adi,omitempty"`
		Mac *string `json:"mac,omitempty"`
	} `json:"multicastSubscriptions,omitempty"`
	Name            *string  `json:"name,omitempty"`
	NetconfRevision *int     `json:"netconfRevision,omitempty"`
	PortDeviceName  *string  `json:"portDeviceName,omitempty"`
	PortError       *float32 `json:"portError,omitempty"`
	Routes          *[]struct {
		Flags  *float32 `json:"flags,omitempty"`
		Metric *float32 `json:"metric,omitempty"`
		Target *string  `json:"target,omitempty"`
		Via    *string  `json:"via,omitempty"`
	} `json:"routes,omitempty"`
	Status *string `json:"status,omitempty"`
	Type   *string `json:"type,omitempty"`
}

// Peer defines model for Peer.
type Peer struct {
	Address  *string `json:"address,omitempty"`
	IsBonded *bool   `json:"isBonded,omitempty"`
	Latency  *int    `json:"latency,omitempty"`
	Paths    *[]struct {
		Active        *bool   `json:"active,omitempty"`
		Address       *string `json:"address,omitempty"`
		Expired       *bool   `json:"expired,omitempty"`
		LastReceive   *int    `json:"lastReceive,omitempty"`
		LastSend      *int    `json:"lastSend,omitempty"`
		Preferred     *bool   `json:"preferred,omitempty"`
		TrustedPathId *int    `json:"trustedPathId,omitempty"`
	} `json:"paths,omitempty"`
	Role         *string `json:"role,omitempty"`
	Version      *string `json:"version,omitempty"`
	VersionMajor *int    `json:"versionMajor,omitempty"`
	VersionMinor *int    `json:"versionMinor,omitempty"`
	VersionRev   *int    `json:"versionRev,omitempty"`
}

// Status defines model for Status.
type Status struct {
	Address *string `json:"address,omitempty"`
	Clock   *int    `json:"clock,omitempty"`
	Config  *struct {
		Settings *struct {
			AllowTcpFallbackRelay *bool `json:"allowTcpFallbackRelay,omitempty"`
			PortMappingEnabled    *bool `json:"portMappingEnabled,omitempty"`
			PrimaryPort           *int  `json:"primaryPort,omitempty"`
		} `json:"settings,omitempty"`
	} `json:"config,omitempty"`
	Online               *bool    `json:"online,omitempty"`
	PlanetWorldId        *float32 `json:"planetWorldId,omitempty"`
	PlanetWorldTimestamp *float32 `json:"planetWorldTimestamp,omitempty"`
	PublicIdentity       *string  `json:"publicIdentity,omitempty"`
	TcpFallbackActive    *bool    `json:"tcpFallbackActive,omitempty"`
	Version              *string  `json:"version,omitempty"`
	VersionBuild         *int     `json:"versionBuild,omitempty"`
	VersionMajor         *int     `json:"versionMajor,omitempty"`
	VersionMinor         *int     `json:"versionMinor,omitempty"`
	VersionRev           *int     `json:"versionRev,omitempty"`
}

// UpdateNetworkJSONBody defines parameters for UpdateNetwork.
type UpdateNetworkJSONBody struct {

	// Let ZeroTier modify the system's DNS settings
	AllowDNS *bool `json:"allowDNS,omitempty"`

	// Let ZeroTier to modify the system's default route.
	AllowDefault *bool `json:"allowDefault,omitempty"`

	// Let ZeroTier to manage IP addresses and Route assignments that aren't in private ranges (rfc1918).
	AllowGlobal *bool `json:"allowGlobal,omitempty"`

	// Let ZeroTier to manage IP addresses and Route assignments.
	AllowManaged *bool `json:"allowManaged,omitempty"`
}

// UpdateNetworkJSONRequestBody defines body for UpdateNetwork for application/json ContentType.
type UpdateNetworkJSONRequestBody UpdateNetworkJSONBody

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
	// GetNetworks request
	GetNetworks(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteNetwork request
	DeleteNetwork(ctx context.Context, networkID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetNetwork request
	GetNetwork(ctx context.Context, networkID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateNetwork request  with any body
	UpdateNetworkWithBody(ctx context.Context, networkID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateNetwork(ctx context.Context, networkID string, body UpdateNetworkJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetPeers request
	GetPeers(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetPeer request
	GetPeer(ctx context.Context, address string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetStatus request
	GetStatus(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetNetworks(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetNetworksRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteNetwork(ctx context.Context, networkID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteNetworkRequest(c.Server, networkID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetNetwork(ctx context.Context, networkID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetNetworkRequest(c.Server, networkID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateNetworkWithBody(ctx context.Context, networkID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateNetworkRequestWithBody(c.Server, networkID, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateNetwork(ctx context.Context, networkID string, body UpdateNetworkJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateNetworkRequest(c.Server, networkID, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetPeers(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPeersRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetPeer(ctx context.Context, address string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPeerRequest(c.Server, address)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetStatus(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetStatusRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetNetworksRequest generates requests for GetNetworks
func NewGetNetworksRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/network")
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewDeleteNetworkRequest generates requests for DeleteNetwork
func NewDeleteNetworkRequest(server string, networkID string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "networkID", runtime.ParamLocationPath, networkID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/network/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetNetworkRequest generates requests for GetNetwork
func NewGetNetworkRequest(server string, networkID string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "networkID", runtime.ParamLocationPath, networkID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/network/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateNetworkRequest calls the generic UpdateNetwork builder with application/json body
func NewUpdateNetworkRequest(server string, networkID string, body UpdateNetworkJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateNetworkRequestWithBody(server, networkID, "application/json", bodyReader)
}

// NewUpdateNetworkRequestWithBody generates requests for UpdateNetwork with any type of body
func NewUpdateNetworkRequestWithBody(server string, networkID string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "networkID", runtime.ParamLocationPath, networkID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/network/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetPeersRequest generates requests for GetPeers
func NewGetPeersRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/peer")
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetPeerRequest generates requests for GetPeer
func NewGetPeerRequest(server string, address string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "address", runtime.ParamLocationPath, address)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/peer/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetStatusRequest generates requests for GetStatus
func NewGetStatusRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/status")
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
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
	// GetNetworks request
	GetNetworksWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetNetworksResponse, error)

	// DeleteNetwork request
	DeleteNetworkWithResponse(ctx context.Context, networkID string, reqEditors ...RequestEditorFn) (*DeleteNetworkResponse, error)

	// GetNetwork request
	GetNetworkWithResponse(ctx context.Context, networkID string, reqEditors ...RequestEditorFn) (*GetNetworkResponse, error)

	// UpdateNetwork request  with any body
	UpdateNetworkWithBodyWithResponse(ctx context.Context, networkID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateNetworkResponse, error)

	UpdateNetworkWithResponse(ctx context.Context, networkID string, body UpdateNetworkJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateNetworkResponse, error)

	// GetPeers request
	GetPeersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetPeersResponse, error)

	// GetPeer request
	GetPeerWithResponse(ctx context.Context, address string, reqEditors ...RequestEditorFn) (*GetPeerResponse, error)

	// GetStatus request
	GetStatusWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetStatusResponse, error)
}

type GetNetworksResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Network
}

// Status returns HTTPResponse.Status
func (r GetNetworksResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetNetworksResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteNetworkResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r DeleteNetworkResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteNetworkResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetNetworkResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Network
}

// Status returns HTTPResponse.Status
func (r GetNetworkResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetNetworkResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateNetworkResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Network
}

// Status returns HTTPResponse.Status
func (r UpdateNetworkResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateNetworkResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetPeersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Peer
}

// Status returns HTTPResponse.Status
func (r GetPeersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPeersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetPeerResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Peer
}

// Status returns HTTPResponse.Status
func (r GetPeerResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPeerResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetStatusResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Status
}

// Status returns HTTPResponse.Status
func (r GetStatusResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetStatusResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetNetworksWithResponse request returning *GetNetworksResponse
func (c *ClientWithResponses) GetNetworksWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetNetworksResponse, error) {
	rsp, err := c.GetNetworks(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetNetworksResponse(rsp)
}

// DeleteNetworkWithResponse request returning *DeleteNetworkResponse
func (c *ClientWithResponses) DeleteNetworkWithResponse(ctx context.Context, networkID string, reqEditors ...RequestEditorFn) (*DeleteNetworkResponse, error) {
	rsp, err := c.DeleteNetwork(ctx, networkID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteNetworkResponse(rsp)
}

// GetNetworkWithResponse request returning *GetNetworkResponse
func (c *ClientWithResponses) GetNetworkWithResponse(ctx context.Context, networkID string, reqEditors ...RequestEditorFn) (*GetNetworkResponse, error) {
	rsp, err := c.GetNetwork(ctx, networkID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetNetworkResponse(rsp)
}

// UpdateNetworkWithBodyWithResponse request with arbitrary body returning *UpdateNetworkResponse
func (c *ClientWithResponses) UpdateNetworkWithBodyWithResponse(ctx context.Context, networkID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateNetworkResponse, error) {
	rsp, err := c.UpdateNetworkWithBody(ctx, networkID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateNetworkResponse(rsp)
}

func (c *ClientWithResponses) UpdateNetworkWithResponse(ctx context.Context, networkID string, body UpdateNetworkJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateNetworkResponse, error) {
	rsp, err := c.UpdateNetwork(ctx, networkID, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateNetworkResponse(rsp)
}

// GetPeersWithResponse request returning *GetPeersResponse
func (c *ClientWithResponses) GetPeersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetPeersResponse, error) {
	rsp, err := c.GetPeers(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPeersResponse(rsp)
}

// GetPeerWithResponse request returning *GetPeerResponse
func (c *ClientWithResponses) GetPeerWithResponse(ctx context.Context, address string, reqEditors ...RequestEditorFn) (*GetPeerResponse, error) {
	rsp, err := c.GetPeer(ctx, address, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPeerResponse(rsp)
}

// GetStatusWithResponse request returning *GetStatusResponse
func (c *ClientWithResponses) GetStatusWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetStatusResponse, error) {
	rsp, err := c.GetStatus(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetStatusResponse(rsp)
}

// ParseGetNetworksResponse parses an HTTP response from a GetNetworksWithResponse call
func ParseGetNetworksResponse(rsp *http.Response) (*GetNetworksResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetNetworksResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Network
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseDeleteNetworkResponse parses an HTTP response from a DeleteNetworkWithResponse call
func ParseDeleteNetworkResponse(rsp *http.Response) (*DeleteNetworkResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &DeleteNetworkResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParseGetNetworkResponse parses an HTTP response from a GetNetworkWithResponse call
func ParseGetNetworkResponse(rsp *http.Response) (*GetNetworkResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetNetworkResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Network
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseUpdateNetworkResponse parses an HTTP response from a UpdateNetworkWithResponse call
func ParseUpdateNetworkResponse(rsp *http.Response) (*UpdateNetworkResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &UpdateNetworkResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Network
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetPeersResponse parses an HTTP response from a GetPeersWithResponse call
func ParseGetPeersResponse(rsp *http.Response) (*GetPeersResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetPeersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Peer
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetPeerResponse parses an HTTP response from a GetPeerWithResponse call
func ParseGetPeerResponse(rsp *http.Response) (*GetPeerResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetPeerResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Peer
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetStatusResponse parses an HTTP response from a GetStatusWithResponse call
func ParseGetStatusResponse(rsp *http.Response) (*GetStatusResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetStatusResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Status
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
