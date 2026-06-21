package model

import "time"

type Candle struct {
	Symbol   string    `json:"symbol"`
	DateTime time.Time `json:"datetime"`
	Open     float64   `json:"open"`
	High     float64   `json:"high"`
	Low      float64   `json:"low"`
	Close    float64   `json:"close"`
	Volume   int64     `json:"volume"`
}

type CandleResponse struct {
	Symbol    string   `json:"symbol"`
	Timeframe string   `json:"timeframe"`
	Candles   []Candle `json:"candles"`
	Count     int      `json:"count"`
}
