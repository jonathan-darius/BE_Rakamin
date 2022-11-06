package main

import (
	"BE_Rakamin/docs"
	"BE_Rakamin/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

// @title           Gin User Service
// @version         1.0
// @description     A user management service API in Go using Gin framework.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

var (
	router *gin.Engine
	port   string
)

func init() {
	err := godotenv.Load(".env")

	port = os.Getenv("PORT")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if port == "" {
		port = "8000"
	}
	router = gin.Default()
	//router.Use(gin.Logger())
}

func main() {
	version := router.Group("/v1")
	routes.UserRoutes(version)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.BasePath = "/v1"
	errRun := router.Run(":" + port)
	if errRun != nil {
		log.Fatal(errRun)
	}
}
