package main

import (
	"todo-app/controller"
	"todo-app/model"

	"github.com/gin-gonic/gin"
)

func main() {
	model.ConnectDataBase()

	router := gin.Default()

	public := router.Group("/api/auth")
	public.POST("/register", controller.Register)
	public.POST("/login", controller.Login)

	router.Run(":8080")
}
