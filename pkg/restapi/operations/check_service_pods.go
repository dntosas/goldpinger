// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CheckServicePodsHandlerFunc turns a function with the right signature into a check service pods handler
type CheckServicePodsHandlerFunc func(CheckServicePodsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CheckServicePodsHandlerFunc) Handle(params CheckServicePodsParams) middleware.Responder {
	return fn(params)
}

// CheckServicePodsHandler interface for that can handle valid check service pods params
type CheckServicePodsHandler interface {
	Handle(CheckServicePodsParams) middleware.Responder
}

// NewCheckServicePods creates a new http.Handler for the check service pods operation
func NewCheckServicePods(ctx *middleware.Context, handler CheckServicePodsHandler) *CheckServicePods {
	return &CheckServicePods{Context: ctx, Handler: handler}
}

/*CheckServicePods swagger:route GET /check checkServicePods

Queries the API server for all other pods in this service, and pings them via their pods IPs. Calls their /ping endpoint

*/
type CheckServicePods struct {
	Context *middleware.Context
	Handler CheckServicePodsHandler
}

func (o *CheckServicePods) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCheckServicePodsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
