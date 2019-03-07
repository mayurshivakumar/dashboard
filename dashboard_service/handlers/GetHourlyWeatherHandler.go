package handlers

import (
	"dashboard-service/models"
	"dashboard-service/restapi/operations/dashboard_service"
	"dashboard-service/source"
	"encoding/json"
	"github.com/go-openapi/runtime/middleware"
	"io/ioutil"
	"net/http"
)

func GetHourlyWeatherHandlerFuncImpl(params dashboard_service.GetHourlyWeatherParams) middleware.Responder {
	//TODO: move api key to config and inject config into the handler
	response, err := http.Get("https://api.wunderground.com/api/Your_Api/hourly/q/" + params.State + "/" + params.City + ".json")
	res := models.SimpleResponse{}
	if err != nil {
		res.Message = err.Error()
		return dashboard_service.NewGetHourlyWeatherInternalServerError().WithPayload(&res)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		res.Message = err.Error()
		return dashboard_service.NewGetHourlyWeatherInternalServerError().WithPayload(&res)
	}
	weather := source.Response{}
	json.Unmarshal(data, &weather)

	// TODO: Error checking
	hourlyWeather := []*models.Hourly{}
	for _, val := range weather.HourlyForecast {
		hourly := models.Hourly{}
		hourly.Temperature = val.Temperature.Fahrenheit
		hourly.Time = val.FCTtime.Civil
		hourlyWeather = append(hourlyWeather, &hourly)
	}
	return dashboard_service.NewGetHourlyWeatherOK().WithPayload(hourlyWeather)
}
