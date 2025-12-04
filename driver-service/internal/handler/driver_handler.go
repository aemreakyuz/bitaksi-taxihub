package handler

import (
	"net/http"
	"strconv"

	"github.com/aemreakyuz/bitaksi-taxihub/driver-service/internal/model"
	"github.com/aemreakyuz/bitaksi-taxihub/driver-service/internal/service"
	"github.com/gin-gonic/gin"
)

type DriverHandler struct {
	service *service.DriverService
}

func NewDriverHandler(service *service.DriverService) *DriverHandler {
	return &DriverHandler{
		service: service,
	}
}

// CreateDriver godoc
// @Summary      Create a new driver
// @Description  Register a new driver in the system
// @Tags         drivers
// @Accept       json
// @Produce      json
// @Param        driver  body      model.Driver  true  "Driver data"
// @Success      201     {object}  model.Driver
// @Failure      400     {object}  map[string]string
// @Router       /drivers [post]
func (h *DriverHandler) CreateDriver(c *gin.Context) {
	var driver model.Driver

	if err := c.ShouldBindJSON(&driver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateDriver(&driver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, driver)
}

// UpdateDriver godoc
// @Summary      Update a driver
// @Description  Update driver information by ID
// @Tags         drivers
// @Accept       json
// @Produce      json
// @Param        id      path      string        true  "Driver ID"
// @Param        driver  body      model.Driver  true  "Driver data"
// @Success      200     {object}  model.Driver
// @Failure      400     {object}  map[string]string
// @Router       /drivers/{id} [put]
func (h *DriverHandler) UpdateDriver(c *gin.Context) {
	id := c.Param("id")

	var driver model.Driver
	if err := c.ShouldBindJSON(&driver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateDriver(id, &driver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, driver)
}

// GetAllDrivers godoc
// @Summary      List all drivers
// @Description  Get a paginated list of all drivers
// @Tags         drivers
// @Accept       json
// @Produce      json
// @Param        page      query     int  false  "Page number"  default(1)
// @Param        pageSize  query     int  false  "Page size"    default(20)
// @Success      200       {array}   model.Driver
// @Failure      500       {object}  map[string]string
// @Router       /drivers [get]
func (h *DriverHandler) GetAllDrivers(c *gin.Context) {
	page := 1
	pageSize := 20

	if p := c.Query("page"); p != "" {
		if val, err := strconv.Atoi(p); err == nil {
			page = val
		}
	}

	if ps := c.Query("pageSize"); ps != "" {
		if val, err := strconv.Atoi(ps); err == nil {
			pageSize = val
		}
	}

	drivers, err := h.service.GetAllDrivers(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, drivers)
}

// GetNearbyDrivers godoc
// @Summary      Find nearby drivers
// @Description  Find drivers within 6km radius of given coordinates
// @Tags         drivers
// @Accept       json
// @Produce      json
// @Param        lat       query     number  true   "Latitude"
// @Param        lon       query     number  true   "Longitude"
// @Param        taxiType  query     string  false  "Taxi type filter"
// @Success      200       {array}   model.Driver
// @Failure      400       {object}  map[string]string
// @Router       /drivers/nearby [get]
func (h *DriverHandler) GetNearbyDrivers(c *gin.Context) {
	lat, err := strconv.ParseFloat(c.Query("lat"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lat parameter"})
		return
	}

	lon, err := strconv.ParseFloat(c.Query("lon"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lon parameter"})
		return
	}

	taxiType := c.Query("taxiType")

	drivers, err := h.service.GetNearbyDrivers(lat, lon, taxiType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, drivers)
}
