package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gocql/gocql"
)

func main() {

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "demo"

	session, err := cluster.CreateSession()

	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	file, err := os.Open("D:\\Assignment (1)\\stock_data.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	rows, _ := reader.ReadAll()

	for i, row := range rows {

		if i == 0 {
			continue
		}

		dt, _ := time.Parse(
			"2006-01-02 15:04:05",
			row[1],
		)

		open, _ := strconv.ParseFloat(row[2], 64)
		high, _ := strconv.ParseFloat(row[3], 64)
		low, _ := strconv.ParseFloat(row[4], 64)
		closeVal, _ := strconv.ParseFloat(row[5], 64)
		volume, _ := strconv.ParseInt(row[6], 10, 64)

		session.Query(
			`INSERT INTO stock_data
			(symbol, datetime, open, high, low, close, volume)
			VALUES (?, ?, ?, ?, ?, ?, ?)`,
			row[0],
			dt,
			open,
			high,
			low,
			closeVal,
			volume,
		).Exec()
	}

	log.Println("Data inserted successfully")
}
