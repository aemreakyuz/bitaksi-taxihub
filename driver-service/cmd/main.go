package main

import (
	"github.com/aemreakyuz/bitaksi-taxihub/driver-service/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to MongoDB
	client := config.ConnectDB()
	defer client.Disconnect(nil)

	// Create router
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "driver-service",
		})
	})

	router.Run(":8081")
}
