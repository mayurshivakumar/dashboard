// Code generated by go-swagger; DO NOT EDIT.

package dashboard_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetPortfolioHandlerFunc turns a function with the right signature into a get portfolio handler
type GetPortfolioHandlerFunc func(GetPortfolioParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPortfolioHandlerFunc) Handle(params GetPortfolioParams) middleware.Responder {
	return fn(params)
}

// GetPortfolioHandler interface for that can handle valid get portfolio params
type GetPortfolioHandler interface {
	Handle(GetPortfolioParams) middleware.Responder
}

// NewGetPortfolio creates a new http.Handler for the get portfolio operation
func NewGetPortfolio(ctx *middleware.Context, handler GetPortfolioHandler) *GetPortfolio {
	return &GetPortfolio{Context: ctx, Handler: handler}
}

/*GetPortfolio swagger:route GET /portfolio dashboard-service getPortfolio

Get portfolio

*/
type GetPortfolio struct {
	Context *middleware.Context
	Handler GetPortfolioHandler
}

func (o *GetPortfolio) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPortfolioParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
