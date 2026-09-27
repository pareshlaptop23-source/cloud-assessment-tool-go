package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	//
	return func(c *gin.Context) {

		role, exists :=
			c.Get("role")

		if !exists {

			c.JSON(
				http.StatusForbidden,
				gin.H{
					"error": "Role Missing",
				},
			)

			c.Abort()
			return
		}

		if role != "ADMIN" {

			c.JSON(
				http.StatusForbidden,
				gin.H{
					"error": "Admin Access Required",
				},
			)

			c.Abort()
			return
		}

		c.Next()
	}
}
