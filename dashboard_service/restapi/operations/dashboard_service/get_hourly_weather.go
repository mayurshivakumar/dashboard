// Code generated by go-swagger; DO NOT EDIT.

package dashboard_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetHourlyWeatherHandlerFunc turns a function with the right signature into a get hourly weather handler
type GetHourlyWeatherHandlerFunc func(GetHourlyWeatherParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetHourlyWeatherHandlerFunc) Handle(params GetHourlyWeatherParams) middleware.Responder {
	return fn(params)
}

// GetHourlyWeatherHandler interface for that can handle valid get hourly weather params
type GetHourlyWeatherHandler interface {
	Handle(GetHourlyWeatherParams) middleware.Responder
}

// NewGetHourlyWeather creates a new http.Handler for the get hourly weather operation
func NewGetHourlyWeather(ctx *middleware.Context, handler GetHourlyWeatherHandler) *GetHourlyWeather {
	return &GetHourlyWeather{Context: ctx, Handler: handler}
}

/*GetHourlyWeather swagger:route GET /weather dashboard-service getHourlyWeather

Get weather

*/
type GetHourlyWeather struct {
	Context *middleware.Context
	Handler GetHourlyWeatherHandler
}

func (o *GetHourlyWeather) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetHourlyWeatherParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
