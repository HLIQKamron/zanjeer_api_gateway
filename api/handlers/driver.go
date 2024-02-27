package handlers

import (
	"fmt"

	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/gin-gonic/gin"
)

// @Router		/user/edit-info [PATCH]
// @Summary		user modifying information
// @Tags        Driver
// @Description	Here users' info can be modified.
// @Accept      json
// @Produce		json
// @Security    BearerAuth
// @Param       post   body       models.Driver true "admin"
// @Success		200 	{object}  models.Driver
// @Failure     default {object}  models.StandardResponse
func (h *handlerV1) UpdateDriverInfo(c *gin.Context) {
	var driver models.Driver

	if err := c.ShouldBindJSON(&driver); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	fmt.Println("driver id: ", driver.Id)
	if (driver.Id) == "" {
		c.JSON(400, gin.H{"error": "id is required"})
		return
	}

	data, err := h.storage.Postgres().UpdateDriverInfo(driver)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": data})
}

// @Router /user/{id} [GET]
// @Summary Get info about driver
// @Tags Driver
// @Description Here driver's info can be fetched.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} models.Driver
// @Failure default {object} models.StandardResponse
func (h *handlerV1) GetDriverInfo(c *gin.Context) {

	driverId := c.Param("id")

	data, err := h.storage.Postgres().GetDriverInfo(driverId)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "error",
			"data":   nil,
		})
	}

	c.JSON(200, gin.H{
		"status": "success",
		"data":   data,
	})
}

// @Router		/user/{id} [DELETE]
// @Summary		user delete
// @Tags        Driver
// @Description	Here drivers can be deleted.
// @Accept      json
// @Produce		json
// @Security    BearerAuth
// @Param 		id 		path 	 string true "id"
// @Success 	200 	{object} models.StandardResponse
// @Failure 	default {object} models.StandardResponse
func (h *handlerV1) DeleteDriverInfo(c *gin.Context) {
	driverId := c.Param("id")

	if driverId == "" {
		c.JSON(400, gin.H{
			"error": "driver id is required",
		})
		return
	}

	if err := h.storage.Postgres().DeleteDriver(driverId); err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, gin.H{"status": "success"})
}