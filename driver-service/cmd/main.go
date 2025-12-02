package main

import (
	"github.com/aemreakyuz/bitaksi-taxihub/driver-service/internal/config"
	"github.com/aemreakyuz/bitaksi-taxihub/driver-service/internal/handler"
	"github.com/aemreakyuz/bitaksi-taxihub/driver-service/internal/repository"
	"github.com/aemreakyuz/bitaksi-taxihub/driver-service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {

	client := config.ConnectDB()
	defer client.Disconnect(nil)

	driverRepo := repository.NewDriverRepository(client)
	driverService := service.NewDriverService(driverRepo)
	driverHandler := handler.NewDriverHandler(driverService)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "driver-service",
		})
	})

	router.POST("/drivers", driverHandler.CreateDriver)
	router.PUT("/drivers/:id", driverHandler.UpdateDriver)
	router.GET("/drivers", driverHandler.GetAllDrivers)
	router.GET("/drivers/nearby", driverHandler.GetNearbyDrivers)

	router.Run(":8081")
}
