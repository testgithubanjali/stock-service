package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Candle struct {
	Symbol   string  `json:"symbol"`
	DateTime string  `json:"datetime"`
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	Volume   int64   `json:"volume"`
}

type Response struct {
	Symbol    string   `json:"symbol"`
	Timeframe string   `json:"timeframe"`
	Candles   []Candle `json:"candles"`
	Count     int      `json:"count"`
}

func main() {

	url := "http://localhost:8081/api/v1/candles?symbol=TCS&timeframe=15m&start_date=2026-01-01%2009:15:00&end_date=2026-01-01%2009:45:00"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result Response

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== Fetched Candle Data ===")
	fmt.Printf(
		"Symbol: %s | Timeframe: %s | Total Candles: %d\n\n",
		result.Symbol,
		result.Timeframe,
		result.Count,
	)

	for i, candle := range result.Candles {

		fmt.Printf(
			"%d %s | O: %.2f | H: %.2f | L: %.2f | C: %.2f | V: %d\n",
			i+1,
			candle.DateTime,
			candle.Open,
			candle.High,
			candle.Low,
			candle.Close,
			candle.Volume,
		)
	}

	fmt.Println("===========================")
}
