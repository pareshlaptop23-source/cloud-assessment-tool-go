package handlers

import (
	"cloud-assessment-tool/config"
	"cloud-assessment-tool/models"
	"cloud-assessment-tool/utils"
	"net/http"

	"cloud-assessment-tool/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register godoc
// @Summary Register a new user
// @Description Creates a new user account
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body models.User true "User registration"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
func Register(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})

		return
	}

	// Check whether the email is already registered
	_, err := repository.GetUserByEmail(user.Email)

	if err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "User already exists",
		})
		return
	}

	hashedPassword, err :=
		bcrypt.GenerateFromPassword(
			[]byte(user.Password),
			bcrypt.DefaultCost,
		)

	if err != nil {

		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Hashing failed"})

		return
	}

	user.Password = string(hashedPassword)

	user.Role = "USER"
	user.Status = "ACTIVE"

	if err := repository.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Registered",
	})
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
func Login(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})

		return
	}

	user, err := repository.GetUserByEmail(req.Email)

	if err != nil {

		c.JSON(http.StatusUnauthorized,
			gin.H{"error": "Invalid credentials"})

		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {

		c.JSON(http.StatusUnauthorized,
			gin.H{"error": "Invalid credentials"})

		return
	}

	token, _ := utils.GenerateToken(user.ID, user.Role)

	c.JSON(http.StatusOK,
		gin.H{
			"token": token,
		})
}

// Profile godoc
// @Summary Get user profile
// @Description Returns the authenticated user's ID
// @Tags Authentication
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Router /api/profile [get]
func Profile(c *gin.Context) {

	userID, _ := c.Get("userID")

	c.JSON(http.StatusOK,
		gin.H{
			"message": "JWT Valid",
			"userID":  userID,
		})
}

// DisableUser godoc
// @Summary Disable a user
// @Description Disables an existing user account
// @Tags Administration
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/admin/users/{id}/disable [put]
func DisableUser(c *gin.Context) {

	id := c.Param("id")

	var user models.User

	if err :=
		config.DB.
			First(&user, id).Error; err != nil {

		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "User Not Found",
			},
		)

		return
	}

	user.Status = "DISABLED"

	config.DB.Save(&user)

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "User Disabled",
		},
	)
}

// EnableUser godoc
// @Summary Enable a user
// @Description Enables an existing user account
// @Tags Administration
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/admin/users/{id}/enable [put]
func EnableUser(c *gin.Context) {

	id := c.Param("id")

	var user models.User

	if err :=
		config.DB.
			First(&user, id).Error; err != nil {

		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "User Not Found",
			},
		)

		return
	}

	user.Status = "ACTIVE"

	config.DB.Save(&user)

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "User Enabled",
		},
	)
}
