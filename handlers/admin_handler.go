package handlers

import (
	"cloud-assessment-tool/config"
	"cloud-assessment-tool/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary Get all users
// @Description Returns all registered users
// @Tags Administration
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /api/admin/users [get]
func GetUsers(c *gin.Context) {

	var users []models.User

	config.DB.Find(&users)

	c.JSON(
		http.StatusOK,
		users,
	)
}
