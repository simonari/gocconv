package core

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ExchangeRatesResponse struct {
	Result         string             `json:"result"`
	TimeLastUpdate int                `json:"time_last_update_unix"`
	Token          string             `json:"base_code"`
	Rates          map[string]float64 `json:"rates"`
}

func makeRequest(from, to string) *ExchangeRatesResponse {
	baseUrl := "https://open.er-api.com/v6/latest/"
	token := from

	url := baseUrl + token

	response, err := http.Get(url)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	ratesInfo := ExchangeRatesResponse{}
	json.Unmarshal(body, &ratesInfo)

	return &ratesInfo
}

type ExchangeRateResponse struct {
	From           string
	To             string
	TimeLastUpdate int
	Rate           float64
}

func (resp *ExchangeRateResponse) GetRate() *CurrencyRate {
	return &CurrencyRate{
		From: resp.From,
		To:   resp.To,
		Rate: float32(resp.Rate),
	}
}

func GetRateInfo(from, to string) ExchangeRateResponse {
	rs := makeRequest(from, to)

	rate := rs.Rates[to]

	if rate == 0 {
		log.Fatal("[!] Error: rate was not found in web")
	}

	return ExchangeRateResponse{
		From:           rs.Token,
		To:             to,
		TimeLastUpdate: rs.TimeLastUpdate,
		Rate:           rate,
	}
}
