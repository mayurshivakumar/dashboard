package dataLayer

import (
	"bufio"
	"dashboard-service/source"
	"encoding/csv"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func GetInvestments() []*source.Investments {
	f, err := os.Open("symbols.csv")
	if err != nil {
		log.Fatalf("%s", err)
	}
	r := csv.NewReader(bufio.NewReader(f))
	investments := []*source.Investments{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		investment := source.Investments{}
		investment.Symbol = record[0]
		investment.NumberBought, _ = strconv.Atoi(record[1])
		investment.PricePerShareBought, _ = strconv.ParseFloat(record[2], 32)
		investment.PricePerShareBought = math.Round(investment.PricePerShareBought*100) / 100
		investments = append(investments, &investment)
	}
	return investments
}
