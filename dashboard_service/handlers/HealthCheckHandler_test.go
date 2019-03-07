package handlers

import (
	"dashboard-service/restapi/operations/health_checks"
	"reflect"
	"testing"
)

func TestHealthCheckHandlerImpl_Handle(t *testing.T) {
	healthCheck := HealthCheckHandlerImpl{}
	params := health_checks.NewHealthCheckParams()
	response := healthCheck.Handle(params)

	if reflect.TypeOf(response) == nil || reflect.TypeOf(response) != reflect.TypeOf(health_checks.NewHealthCheckOK()) {
		t.Error("Expected response of type NewHealthCheckOK, got:", reflect.TypeOf(response))
	}
}
