package handlers

import (
	"cloud-assessment-tool/config"
	"cloud-assessment-tool/models"
	"cloud-assessment-tool/repository"
	"cloud-assessment-tool/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllVMs godoc
// @Summary Get all virtual machines
// @Description Returns all virtual machines
// @Tags Virtual Machines
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /api/vms [get]
func GetAllVMs(c *gin.Context) {

	// 1. Check Redis
	vms, err := services.GetCachedVMs()

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "VM list from cache",
			"data":    vms,
		})
		return
	}

	// If GetCachedVMs failed, check whether the key exists and return raw cached value
	fmt.Println("Redis cache miss or unmarshal error:", err)
	exists, _ := config.RedisClient.Exists(config.RedisCtx, services.VMCacheKey).Result()
	if exists > 0 {
		raw, getErr := config.RedisClient.Get(config.RedisCtx, services.VMCacheKey).Result()
		if getErr == nil {
			// Try unmarshalling here again; if it still fails, return raw value to caller
			var cached []models.VirtualMachine
			if umErr := json.Unmarshal([]byte(raw), &cached); umErr == nil {
				c.JSON(http.StatusOK, gin.H{"message": "VM list from cache", "data": cached})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "VM list from cache", "data_raw": raw})
			return
		}
		fmt.Println("Redis GET after EXISTS returned error:", getErr)
	}
	// 2. Cache miss → Get from MySQL
	vms, err = repository.GetAllVMs()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	// 3. Store result in Redis
	err = services.SetCachedVMs(vms)
	if err != nil {
		fmt.Println("Redis SET error:", err)
	} else {
		fmt.Println("Redis cache SET successful: vms:all")
	}

	// 4. Return response
	c.JSON(http.StatusOK, gin.H{
		"message": "VM list from database",
		"data":    vms,
	})
}

// TerminateVM godoc
// @Summary Terminate a virtual machine
// @Description Stops/terminates a running virtual machine
// @Tags Virtual Machines
// @Produce json
// @Security BearerAuth
// @Param id path int true "VM ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/vms/{id}/terminate [put]
func TerminateVM(c *gin.Context) {

	idStr := c.Param("id")

	id64, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {

		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "Invalid VM ID",
			})

		return
	}
	config.RedisClient.Del(
		config.RedisCtx,
		services.VMCacheKey,
	)

	id := uint(id64)

	vm, err := repository.GetVMByID(id)

	if err != nil {

		c.JSON(http.StatusNotFound,
			gin.H{
				"error": "VM Not Found",
			})

		return
	}

	if vm.Status == "STOPPED" {

		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "VM is already stopped",
			})

		return
	}

	err = repository.StopVM(id)

	if err != nil {

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			})

		return
	}

	userID, _ := c.Get("userID")

	log := models.AuditLog{
		UserID:     userID.(uint),
		Action:     "STOP_VM",
		Resource:   "VirtualMachine",
		ResourceID: id,
	}

	if err := repository.CreateAuditLog(&log); err != nil {

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			})

		return
	}
	message := fmt.Sprintf(
		"User %d stopped VM %d",
		userID.(uint),
		id,
	)
	services.PublishAudit(message)

	c.JSON(http.StatusOK,
		gin.H{
			"message": "VM Terminated Successfully",
		})

}

// StartVM godoc
// @Summary Start a virtual machine
// @Description Starts a stopped virtual machine
// @Tags Virtual Machines
// @Produce json
// @Security BearerAuth
// @Param id path int true "VM ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/vms/{id}/start [put]
func StartVM(c *gin.Context) {

	idStr := c.Param("id")

	id64, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {

		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "Invalid VM ID",
			})

		return
	}
	config.RedisClient.Del(
		config.RedisCtx,
		services.VMCacheKey,
	)

	id := uint(id64)

	vm, err := repository.GetVMByID(id)

	if err != nil {

		c.JSON(http.StatusNotFound,
			gin.H{
				"error": "VM Not Found",
			})

		return
	}

	if vm.Status == "RUNNING" {

		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "VM is already running",
			})

		return
	}

	err = repository.StartVM(id)

	if err != nil {

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			})

		return
	}

	userID, _ := c.Get("userID")

	log := models.AuditLog{

		UserID: userID.(uint),

		Action: "START_VM",

		Resource: "VirtualMachine",

		ResourceID: id,
	}

	if err := repository.CreateAuditLog(&log); err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error": "Failed to create audit log",
			})
		return
	}
	message := fmt.Sprintf(
		"User %d started VM %d",
		userID.(uint),
		id,
	)

	services.PublishAudit(message)

	c.JSON(http.StatusOK,
		gin.H{
			"message": "VM Started Successfully",
		})

}

// GetRecommendations godoc
// @Summary Get VM cost optimization recommendations
// @Description Returns cost optimization recommendations for virtual machines
// @Tags Virtual Machines
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/vms/recommendations [get]
func GetRecommendations(c *gin.Context) {
	vms, err := repository.GetAllVMs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	recommendations := services.GetRecommendations(vms)
	c.JSON(http.StatusOK, gin.H{
		"message": "Cost Optimization Recommendations",
		"data":    recommendations,
	})
}
