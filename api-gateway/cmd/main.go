package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const driverServiceURL = "http://localhost:8081"

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "api-gateway",
		})
	})

	router.POST("/drivers", proxyToDriverService)
	router.PUT("/drivers/:id", proxyToDriverService)
	router.GET("/drivers", proxyToDriverService)
	router.GET("/drivers/nearby", proxyToDriverService)

	router.Run(":3000")
}

func proxyToDriverService(c *gin.Context) {
	targetURL := driverServiceURL + c.Request.URL.Path
	if c.Request.URL.RawQuery != "" {
		targetURL += "?" + c.Request.URL.RawQuery
	}

	req, err := http.NewRequest(c.Request.Method, targetURL, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create request"})
		return
	}

	req.Header = c.Request.Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "driver-service unavailable"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read response"})
		return
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
