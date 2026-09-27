package routes

import (
	"cloud-assessment-tool/handlers"
	"cloud-assessment-tool/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine) {

	router.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)

	router.Use(middleware.PrometheusMiddleware())
	router.POST("/register", handlers.Register)

	router.POST("/login", handlers.Login)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})

	api := router.Group("/api")

	api.Use(
		middleware.AuthMiddleware(),
	)

	{
		api.GET(
			"/profile",
			handlers.Profile,
		)
	}
	admin := api.Group("/admin")

	admin.Use(
		middleware.AdminMiddleware(),
	)

	{
		admin.GET(
			"/users",
			handlers.GetUsers,
		)

		admin.PUT(
			"/users/:id/disable",
			handlers.DisableUser,
		)

		admin.PUT(
			"/users/:id/enable",
			handlers.EnableUser,
		)

	}
	api.POST(

		"/cloud/sync",

		middleware.AuthMiddleware(),

		middleware.AdminMiddleware(),

		handlers.SyncCloud,
	)

	api.GET(
		"/vms",
		middleware.AdminMiddleware(),
		handlers.GetAllVMs,
	)
	api.PUT(

		"/vms/:id/terminate",

		middleware.AdminMiddleware(),

		handlers.TerminateVM,
	)
	api.PUT(

		"/vms/:id/start",

		middleware.AdminMiddleware(),

		handlers.StartVM,
	)
	api.GET(
		"/vms/recommendations",
		middleware.AuthMiddleware(),
		handlers.GetRecommendations,
	)
	api.GET(
		"/audit-logs",
		middleware.AuthMiddleware(),
		middleware.AdminMiddleware(),
		handlers.GetAuditLogs,
	)
}
