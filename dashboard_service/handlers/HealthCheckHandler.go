package handlers

import (
	"dashboard-service/models"
	"dashboard-service/restapi/operations/health_checks"
	"github.com/go-openapi/runtime/middleware"
)

var _healthy = models.SimpleResponse{Message: "OK"}

type HealthCheckHandlerImpl struct {
}

// HealthCheckHandlerFuncImpl func
func (*HealthCheckHandlerImpl) Handle(params health_checks.HealthCheckParams) middleware.Responder {
	return health_checks.NewHealthCheckOK().WithPayload(&_healthy)
}
