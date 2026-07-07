package main

import (
	"os"
	"todo-app/controller"
	"todo-app/model"

	"github.com/gin-gonic/gin"
)

func main() {
	model.ConnectDataBase()

	router := gin.Default()

	router.Use(corsMiddleware())

	authCtrl := controller.AuthController{}

	public := router.Group("/api/auth")
	{
		public.POST("/register", authCtrl.Register)
		public.POST("/login", authCtrl.Login)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
