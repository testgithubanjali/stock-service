package repository

import (
	"stock-service/internal/model"
	"time"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func InitDB() error {

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "demo"

	session, err := cluster.CreateSession()
	if err != nil {
		return err
	}

	Session = session

	return nil
}

func GetCandles(
	symbol string,
	start time.Time,
	end time.Time,
) ([]model.Candle, error) {

	var candles []model.Candle

	iter := Session.Query(
		`SELECT symbol, datetime, open, high, low, close, volume
		FROM stock_data
		WHERE symbol=?
		AND datetime>=?
		AND datetime<=?`,
		symbol,
		start,
		end,
	).Iter()

	var c model.Candle

	for iter.Scan(
		&c.Symbol,
		&c.DateTime,
		&c.Open,
		&c.High,
		&c.Low,
		&c.Close,
		&c.Volume,
	) {
		candles = append(candles, c)
	}

	return candles, iter.Close()
}
