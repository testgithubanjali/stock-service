package service

import (
	"errors"
	"stock-service/internal/model"
	"stock-service/internal/repository"
	"time"
)

func AggregateCandles(
	symbol string,
	timeframe string,
	start time.Time,
	end time.Time,
) (*model.CandleResponse, error) {

	candles, err := repository.GetCandles(
		symbol,
		start,
		end,
	)

	if err != nil {
		return nil, err
	}

	if len(candles) == 0 {
		return nil, errors.New("no data found")
	}

	return &model.CandleResponse{
		Symbol:    symbol,
		Timeframe: timeframe,
		Candles:   Aggregate(candles, timeframe),
		Count:     len(Aggregate(candles, timeframe)),
	}, nil
}
func getDuration(tf string) time.Duration {

	switch tf {

	case "1m":
		return time.Minute

	case "5m":
		return 5 * time.Minute

	case "15m":
		return 15 * time.Minute

	case "30m":
		return 30 * time.Minute

	case "1h":
		return time.Hour

	case "1d":
		return 24 * time.Hour
	}

	return time.Minute
}
func Aggregate(
	candles []model.Candle,
	tf string,
) []model.Candle {

	duration := getDuration(tf)

	buckets := make(map[time.Time][]model.Candle)

	for _, c := range candles {

		bucket := c.DateTime.Truncate(duration)

		buckets[bucket] = append(
			buckets[bucket],
			c,
		)
	}

	var result []model.Candle

	for bucket, items := range buckets {

		first := items[0]
		last := items[len(items)-1]

		high := first.High
		low := first.Low

		var volume int64

		for _, item := range items {

			if item.High > high {
				high = item.High
			}

			if item.Low < low {
				low = item.Low
			}

			volume += item.Volume
		}

		result = append(
			result,
			model.Candle{
				Symbol:   first.Symbol,
				DateTime: bucket,
				Open:     first.Open,
				High:     high,
				Low:      low,
				Close:    last.Close,
				Volume:   volume,
			},
		)
	}

	return result
}
