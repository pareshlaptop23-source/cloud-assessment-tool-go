package middleware

import (
	"cloud-assessment-tool/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {

			c.JSON(http.StatusUnauthorized,
				gin.H{"error": "Token Missing"})

			c.Abort()
			return
		}

		tokenString :=
			strings.TrimPrefix(
				authHeader,
				"Bearer ",
			)

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {
				return utils.SecretKey, nil
			},
		)

		if err != nil || !token.Valid {

			c.JSON(http.StatusUnauthorized,
				gin.H{"error": "Invalid Token"})

			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		userID := uint(claims["user_id"].(float64))

		c.Set("userID", userID)

		c.Set("role", claims["role"])
		c.Next()
	}
}
