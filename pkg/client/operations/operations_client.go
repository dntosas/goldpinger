// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CheckAllPods(params *CheckAllPodsParams) (*CheckAllPodsOK, error)

	CheckServicePods(params *CheckServicePodsParams) (*CheckServicePodsOK, error)

	Healthz(params *HealthzParams) (*HealthzOK, error)

	Ping(params *PingParams) (*PingOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CheckAllPods Queries the API server for all other pods in this service, and makes all of them query all of their neighbours, using their pods IPs. Calls their /check endpoint.
*/
func (a *Client) CheckAllPods(params *CheckAllPodsParams) (*CheckAllPodsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckAllPodsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "checkAllPods",
		Method:             "GET",
		PathPattern:        "/check_all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CheckAllPodsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CheckAllPodsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for checkAllPods: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CheckServicePods Queries the API server for all other pods in this service, and pings them via their pods IPs. Calls their /ping endpoint
*/
func (a *Client) CheckServicePods(params *CheckServicePodsParams) (*CheckServicePodsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckServicePodsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "checkServicePods",
		Method:             "GET",
		PathPattern:        "/check",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CheckServicePodsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CheckServicePodsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for checkServicePods: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Healthz The healthcheck endpoint provides detailed information about the health of a web service. If each of the components required by the service are healthy, then the service is considered healthy and will return a 200 OK response. If any of the components needed by the service are unhealthy, then a 503 Service Unavailable response will be provided.
*/
func (a *Client) Healthz(params *HealthzParams) (*HealthzOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHealthzParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "healthz",
		Method:             "GET",
		PathPattern:        "/healthz",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HealthzReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HealthzOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for healthz: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Ping return query stats
*/
func (a *Client) Ping(params *PingParams) (*PingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ping",
		Method:             "GET",
		PathPattern:        "/ping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PingReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ping: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
