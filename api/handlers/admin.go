package handlers

import (
	"strconv"

	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Router		/superadmin/add/admin [POST]
// @Summary		Add admin by superadmin
// @Tags        Admin
// @Description	Here admins can be created.
// @Accept      json
// @Produce		json
// @Security    BearerAuth
// @Param       post   body       models.Admin true "admin"
// @Success		200 	{object}  models.Admin
// @Failure     default {object}  models.StandardResponse
func (h *handlerV1) CreateAdmin(c *gin.Context) {
	var resp models.Admin

	if err := c.ShouldBindJSON(&resp); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(resp.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "server error"})
		return
	}
	resp.Password = string(hashedPassword)

	data, err := h.storage.Postgres().CreateAdmin(resp)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		},
		)
		return
	}
	c.JSON(200, gin.H{
		"status":  "OK",
		"message": "Admin created successfully",
		"data":    data,
	})
}

// @Router /superadmin/get/admins [GET]
// @Summary Get all admins
// @Tags Admin
// @Description Here all admins can be fetched.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query string true "limit"
// @Param offset query string true "offset"
// @Success 200 {object} []models.Admin
// @Failure default {object} models.StandardResponse
func (h *handlerV1) GetAdmins(c *gin.Context) {

	limit, ok := c.GetQuery("limit")
	if !ok {
		limit = "10"
	}

	page, ok := c.GetQuery("page")
	if !ok {
		page = "0"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid argument"})
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid argument"})
		return
	}

	data, err := h.storage.Postgres().GetAdmins(models.GetAdmins{
		Limit: limitInt,
		Page:  pageInt,
	})
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		},
		)
		return
	}
	c.JSON(200, gin.H{
		"status":  "OK",
		"message": "Admins fetched successfully",
		"data":    data,
	})
}
