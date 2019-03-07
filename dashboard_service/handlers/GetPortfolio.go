package handlers

import (
	"dashboard-service/dataLayer"
	"dashboard-service/models"
	"dashboard-service/restapi/operations/dashboard_service"
	"github.com/go-openapi/runtime/middleware"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
)

func GetPortfilioHandlerFuncImpl(params dashboard_service.GetPortfolioParams) middleware.Responder {
	investments := dataLayer.GetInvestments()
	res := models.SimpleResponse{}
	portfolio := []*models.Portfolio{}
	for _, investment := range investments {
		p := models.Portfolio{}
		p.Symbol = investment.Symbol
		p.PricePerShareBought = investment.PricePerShareBought
		response, err := http.Get("https://api.iextrading.com/1.0/stock/" + investment.Symbol + "/price")
		if err != nil {
			res.Message = err.Error()
			return dashboard_service.NewGetHourlyWeatherInternalServerError().WithPayload(&res)
		}
		defer response.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		p.PricePerShareNow, _ = strconv.ParseFloat(string(bodyBytes), 32)
		p.PricePerShareNow = math.Round(p.PricePerShareNow*100) / 100
		p.TotalValueBought = investment.PricePerShareBought * float64(investment.NumberBought)
		p.TotalValueNow = p.PricePerShareNow * float64(investment.NumberBought)
		p.SharesHeld = float64(investment.NumberBought)
		if (p.TotalValueNow - p.TotalValueBought) > 0 {
			p.Profit = p.TotalValueNow - p.TotalValueBought
		} else {
			p.Loss = p.TotalValueNow - p.TotalValueBought
		}

		portfolio = append(portfolio, &p)
	}
	return dashboard_service.NewGetPortfolioOK().WithPayload(portfolio)
}
