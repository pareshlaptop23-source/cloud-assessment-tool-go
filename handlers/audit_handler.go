package handlers

import (
	"cloud-assessment-tool/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAuditLogs godoc
// @Summary Get audit logs
// @Description Returns audit logs for administrative actions
// @Tags Audit
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/audit-logs [get]
func GetAuditLogs(c *gin.Context) {

	logs, err := repository.GetAuditLogs()

	if err != nil {

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			})

		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"message": "Audit Logs",
			"data":    logs,
		})
}
