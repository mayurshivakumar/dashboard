// Code generated by go-swagger; DO NOT EDIT.

package dashboard_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "dashboard-service/models"
)

// GetHourlyWeatherOKCode is the HTTP code returned for type GetHourlyWeatherOK
const GetHourlyWeatherOKCode int = 200

/*GetHourlyWeatherOK Success

swagger:response getHourlyWeatherOK
*/
type GetHourlyWeatherOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Hourly `json:"body,omitempty"`
}

// NewGetHourlyWeatherOK creates GetHourlyWeatherOK with default headers values
func NewGetHourlyWeatherOK() *GetHourlyWeatherOK {

	return &GetHourlyWeatherOK{}
}

// WithPayload adds the payload to the get hourly weather o k response
func (o *GetHourlyWeatherOK) WithPayload(payload []*models.Hourly) *GetHourlyWeatherOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get hourly weather o k response
func (o *GetHourlyWeatherOK) SetPayload(payload []*models.Hourly) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHourlyWeatherOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Hourly, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetHourlyWeatherBadRequestCode is the HTTP code returned for type GetHourlyWeatherBadRequest
const GetHourlyWeatherBadRequestCode int = 400

/*GetHourlyWeatherBadRequest Bad Request

swagger:response getHourlyWeatherBadRequest
*/
type GetHourlyWeatherBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.SimpleResponse `json:"body,omitempty"`
}

// NewGetHourlyWeatherBadRequest creates GetHourlyWeatherBadRequest with default headers values
func NewGetHourlyWeatherBadRequest() *GetHourlyWeatherBadRequest {

	return &GetHourlyWeatherBadRequest{}
}

// WithPayload adds the payload to the get hourly weather bad request response
func (o *GetHourlyWeatherBadRequest) WithPayload(payload *models.SimpleResponse) *GetHourlyWeatherBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get hourly weather bad request response
func (o *GetHourlyWeatherBadRequest) SetPayload(payload *models.SimpleResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHourlyWeatherBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHourlyWeatherUnauthorizedCode is the HTTP code returned for type GetHourlyWeatherUnauthorized
const GetHourlyWeatherUnauthorizedCode int = 401

/*GetHourlyWeatherUnauthorized Unauthorized

swagger:response getHourlyWeatherUnauthorized
*/
type GetHourlyWeatherUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.SimpleResponse `json:"body,omitempty"`
}

// NewGetHourlyWeatherUnauthorized creates GetHourlyWeatherUnauthorized with default headers values
func NewGetHourlyWeatherUnauthorized() *GetHourlyWeatherUnauthorized {

	return &GetHourlyWeatherUnauthorized{}
}

// WithPayload adds the payload to the get hourly weather unauthorized response
func (o *GetHourlyWeatherUnauthorized) WithPayload(payload *models.SimpleResponse) *GetHourlyWeatherUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get hourly weather unauthorized response
func (o *GetHourlyWeatherUnauthorized) SetPayload(payload *models.SimpleResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHourlyWeatherUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHourlyWeatherForbiddenCode is the HTTP code returned for type GetHourlyWeatherForbidden
const GetHourlyWeatherForbiddenCode int = 403

/*GetHourlyWeatherForbidden Forbidden

swagger:response getHourlyWeatherForbidden
*/
type GetHourlyWeatherForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.SimpleResponse `json:"body,omitempty"`
}

// NewGetHourlyWeatherForbidden creates GetHourlyWeatherForbidden with default headers values
func NewGetHourlyWeatherForbidden() *GetHourlyWeatherForbidden {

	return &GetHourlyWeatherForbidden{}
}

// WithPayload adds the payload to the get hourly weather forbidden response
func (o *GetHourlyWeatherForbidden) WithPayload(payload *models.SimpleResponse) *GetHourlyWeatherForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get hourly weather forbidden response
func (o *GetHourlyWeatherForbidden) SetPayload(payload *models.SimpleResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHourlyWeatherForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHourlyWeatherNotFoundCode is the HTTP code returned for type GetHourlyWeatherNotFound
const GetHourlyWeatherNotFoundCode int = 404

/*GetHourlyWeatherNotFound Not Found

swagger:response getHourlyWeatherNotFound
*/
type GetHourlyWeatherNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.SimpleResponse `json:"body,omitempty"`
}

// NewGetHourlyWeatherNotFound creates GetHourlyWeatherNotFound with default headers values
func NewGetHourlyWeatherNotFound() *GetHourlyWeatherNotFound {

	return &GetHourlyWeatherNotFound{}
}

// WithPayload adds the payload to the get hourly weather not found response
func (o *GetHourlyWeatherNotFound) WithPayload(payload *models.SimpleResponse) *GetHourlyWeatherNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get hourly weather not found response
func (o *GetHourlyWeatherNotFound) SetPayload(payload *models.SimpleResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHourlyWeatherNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHourlyWeatherInternalServerErrorCode is the HTTP code returned for type GetHourlyWeatherInternalServerError
const GetHourlyWeatherInternalServerErrorCode int = 500

/*GetHourlyWeatherInternalServerError Internal Server Error

swagger:response getHourlyWeatherInternalServerError
*/
type GetHourlyWeatherInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.SimpleResponse `json:"body,omitempty"`
}

// NewGetHourlyWeatherInternalServerError creates GetHourlyWeatherInternalServerError with default headers values
func NewGetHourlyWeatherInternalServerError() *GetHourlyWeatherInternalServerError {

	return &GetHourlyWeatherInternalServerError{}
}

// WithPayload adds the payload to the get hourly weather internal server error response
func (o *GetHourlyWeatherInternalServerError) WithPayload(payload *models.SimpleResponse) *GetHourlyWeatherInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get hourly weather internal server error response
func (o *GetHourlyWeatherInternalServerError) SetPayload(payload *models.SimpleResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHourlyWeatherInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
