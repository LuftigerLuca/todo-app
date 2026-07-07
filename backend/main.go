//go:generate swag init --parseDependency --parseInternal

// @title           Todo App API
// @description     A Todo List Application
// @version         1.0
// @host            localhost:8080
// @BasePath        /api
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name            Authorization
package main

import (
	"log"
	"os"
	"todo-app/controller"
	_ "todo-app/docs"
	"todo-app/middleware"
	"todo-app/model"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	model.ConnectDataBase()

	router := gin.Default()

	router.Use(corsMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authCtrl := controller.AuthController{}
	listCtrl := controller.ListController{}

	public := router.Group("/api/auth")
	{
		public.POST("/register", authCtrl.Register)
		public.POST("/login", authCtrl.Login)
	}

	protected := router.Group("/api")
	protected.Use(middleware.AuthRequired())
	{
		list := protected.Group("/list")
		{
			list.POST("/create", listCtrl.CreateList)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
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
