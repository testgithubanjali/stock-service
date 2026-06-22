# Stock Market Data Aggregation Service

## Overview

A RESTful service built in Go that reads stock market time-series data from Apache Cassandra and returns aggregated OHLCV candlestick data for different timeframes.

## Tech Stack

* Golang
* Apache Cassandra
* Gin Framework
* gocql Driver

## Project Structure

submission/
├── server/
├── client/
├── schema.cql
└── README.md

## Prerequisites

* Go 1.24+
* Docker Desktop
* Apache Cassandra

## Cassandra Setup

Pull Cassandra Image:

docker pull cassandra:4.1

Run Cassandra:

docker run --name cassandra -p 9042:9042 -d cassandra:4.1

Open Cassandra Shell:

docker exec -it cassandra cqlsh

## Schema Setup

Execute schema.cql:

SOURCE 'schema.cql';

## Data Ingestion

Run:

go run scripts/ingest.go

Expected Output:

Data inserted successfully

## Start API Server

Run:

go run cmd/main.go

Server starts on:

http://localhost:8081

## API Endpoint

GET /api/v1/candles

Query Parameters:

* symbol
* timeframe
* start_date
* end_date

Example:

http://localhost:8081/api/v1/candles?symbol=TCS&timeframe=15m&start_date=2026-01-01%2009:15:00&end_date=2026-01-01%2009:45:00

## Sample Response

{
"symbol": "TCS",
"timeframe": "15m",
"count": 3
}

## Client Application

Run:

go run client/main.go

Expected Output:

=== Fetched Candle Data ===

Symbol: TCS | Timeframe: 15m | Total Candles: 3

...

===========================

## Design Decisions

* Cassandra partition key: symbol
* Cassandra clustering key: datetime
* Layered architecture:

  * Handler
  * Service
  * Repository
* Timeframe aggregation performed in service layer

## Assumptions

* Input CSV contains 1-minute candles.
* Cassandra is running on localhost:9042.
* API runs on localhost:8081.
