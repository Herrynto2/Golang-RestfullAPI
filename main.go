package main

import (
	"log"
	"portservices/config"
	"portservices/handler"
	"portservices/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()

	// Initialize database connection
	config.InitDb()

	// Routes for authentication
	authRoute := r.Group("/users")
	{
		authRoute.POST("/register", handler.Register)
		authRoute.POST("/login", handler.Login)
		authRoute.PATCH("/logout", middleware.AuthMiddleware(), handler.Logout)
		authRoute.PATCH("/change-password", handler.ChangePassword)
		authRoute.PATCH("/verify-account", handler.VerifyAccount)
	}
	// Routes for authentication
	clientRoute := r.Group("/master-data")
	{
		clientRoute.POST("/client", middleware.AuthMiddleware(), handler.CreateDataClient)
		clientRoute.GET("/client", middleware.AuthMiddleware(), handler.GetDataClient)
	}
	r.Run(":9000")
}
