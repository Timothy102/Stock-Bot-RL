package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const apiKey = "pIBGUlr3Gd8Laqjyz9masQr8m_LgZM_N"

func makeReq(url string) ([]byte, error) {
	req, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not create request at %s: %v", url, err)
	}

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read at %s: %v", url, err)
	}
	if err := req.Body.Close(); err != nil {
		return nil, fmt.Errorf("could not close response body: %v", err)
	}
	return res, nil
}

func getTickers(url string) (*Master, error) {
	//url := "https://api.polygon.io/v3/reference/tickers?active=true&sort=ticker&order=asc&limit=1000&apiKey=pIBGUlr3Gd8Laqjyz9masQr8m_LgZM_N"
	res, err := makeReq(url)
	if err != nil {
		return nil, fmt.Errorf("could not make req at %s : %v", url, err)
	}
	var master Master
	if err := json.Unmarshal(res, &master); err != nil {
		return nil, fmt.Errorf("could not unmarshal: %v", err)
	}
	return &master, nil
}

func getAggregate(ticker, timestamp string) (*Aggregate, error) {
	aggUrl := "https://api.polygon.io/v2/aggs/ticker/" + ticker + "/range/" + timestamp + "/2020-10-14/2020-10-14?adjusted=true&sort=asc&limit=497&apiKey=pIBGUlr3Gd8Laqjyz9masQr8m_LgZM_N"
	res, err := makeReq(aggUrl)
	if err != nil {
		return nil, fmt.Errorf("could not make req at %s : %v", aggUrl, err)
	}
	var agg Aggregate
	if err := json.Unmarshal(res, &agg); err != nil {
		return nil, fmt.Errorf("could not unmarshal: %v", err)
	}
	return &agg, nil
}

func main() {

}

// Aggregates struct
type Aggregate struct {
	Ticker       string  `json:"ticker"`
	QueryCount   int     `json:"queryCount"`
	ResultsCount int     `json:"resultsCount"`
	Adjusted     bool    `json:"adjusted"`
	Results      []OHTCL `json:"results"`
}

type OHTCL struct {
	V  float64 `json:"v"`
	Vw float64 `json:"vw"`
	O  float64 `json:"o"`
	C  float64 `json:"c"`
	H  float64 `json:"h"`
	L  float64 `json:"l"`
	T  float64 `json:"t"`
	N  float64 `json:"n"`
}

// poglej tole za OHTCL vrednosti : https://polygon.io/docs/get_v2_aggs_ticker__cryptoTicker__range__multiplier___timespan___from___to__anchor

// Ticker structs
type Master struct {
	Results []Result `json:"results"`
}
type Result struct {
	Status    string `json:"status"`
	RequestID string `json:"request_id"`
	Count     int    `json:"count"`
	NextURL   string `json:"next_url"`
	Tickers   []Ticker
}
type Ticker struct {
	Ticker          string    `json:"ticker"`
	Name            string    `json:"name"`
	Market          string    `json:"market"`
	Locale          string    `json:"locale"`
	PrimaryExchange string    `json:"primary_exchange"`
	Type            string    `json:"type"`
	Active          bool      `json:"active"`
	CurrencyName    string    `json:"currency_name"`
	Cik             string    `json:"cik"`
	CompositeFigi   string    `json:"composite_figi"`
	ShareClassFigi  string    `json:"share_class_figi"`
	LastUpdatedUtc  time.Time `json:"last_updated_utc"`
}
