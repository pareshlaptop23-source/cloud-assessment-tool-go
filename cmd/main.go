// @title Cloud Assessment Tool API
// @version 1.0
// @description Cloud Assessment Tool REST API
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter "Bearer <your-jwt-token>".

package main

import (
	"cloud-assessment-tool/config"
	"cloud-assessment-tool/metrics"
	"cloud-assessment-tool/models"
	"cloud-assessment-tool/routes"
	"cloud-assessment-tool/seed"
	"cloud-assessment-tool/services"

	_ "cloud-assessment-tool/docs"

	"github.com/gin-gonic/gin"
)

func main() {
	metrics.Init()
	config.ConnectDB()
	if err := config.ConnectRedis(); err != nil {
		panic("Redis connection failed: " + err.Error())
	}

	config.DB.AutoMigrate(

		&models.User{},

		&models.CloudProvider{},

		&models.VirtualMachine{},
		&models.AuditLog{},
	)
	seed.SeedCloudProviders()
	go services.StartConsumer()

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":8080")
}
