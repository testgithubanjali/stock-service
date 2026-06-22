package handler

import (
	"net/http"
	"stock-service/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

func GetCandles(c *gin.Context) {

	symbol := c.Query("symbol")
	timeframe := c.Query("timeframe")
	startStr := c.Query("start_date")
	endStr := c.Query("end_date")

	// Validation
	if symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "symbol is required",
		})
		return
	}

	if timeframe == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "timeframe is required",
		})
		return
	}

	if startStr == "" || endStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "start_date and end_date are required",
		})
		return
	}

	start, err := time.Parse(
		"2006-01-02 15:04:05",
		startStr,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid start_date format. Use YYYY-MM-DD HH:MM:SS",
		})
		return
	}

	end, err := time.Parse(
		"2006-01-02 15:04:05",
		endStr,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid end_date format. Use YYYY-MM-DD HH:MM:SS",
		})
		return
	}

	resp, err := service.AggregateCandles(
		symbol,
		timeframe,
		start,
		end,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
