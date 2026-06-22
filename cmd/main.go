package main

import (
	"log"
	"stock-service/internal/handler"
	"stock-service/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	err := repository.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET(
		"/api/v1/candles",
		handler.GetCandles,
	)

	router.Run(":8081")
}
