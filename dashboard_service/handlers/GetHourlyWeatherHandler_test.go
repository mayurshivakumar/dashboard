package handlers

import (
	"dashboard-service/restapi/operations/dashboard_service"
	"reflect"
	"testing"
)

func TestGetHourlyWeatherHandlerFuncImpl(t *testing.T) {
	// happy path
	params := dashboard_service.NewGetHourlyWeatherParams()
	params.State = "TX"
	params.City = "Houston"

	response := GetHourlyWeatherHandlerFuncImpl(params)
	if reflect.TypeOf(response) == nil || reflect.TypeOf(response) != reflect.TypeOf(dashboard_service.NewGetHourlyWeatherOK()) {
		t.Error("Expected response of type NewGetHourlyWeatherOK, got:", reflect.TypeOf(response))
	}
}

//TODO write tests for errors
