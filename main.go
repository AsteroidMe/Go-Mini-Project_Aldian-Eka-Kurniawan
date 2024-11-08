package main

import (
	"log"
	"mini-project/config"
	"mini-project/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	config.MigrateDB()

	r := gin.Default()

	r.POST("/api/v1/register", controller.Register)
	//r.POST("/api/v1/login", controller.Login)

	// api := r.Group("/api/v1")
	// api.Use(middleware.AuthMiddleware)

	// api.GET("/packages", controllers.GetPaketsHandler)
	// api.GET("/packages/:id", controllers.GetDetailPaketsHandler)
	// api.POST("/packages", controllers.AddPaketsHandler)
	// api.PUT("/packages/:id", controllers.UpdatePaketsHandler)
	// api.DELETE("/packages/:id", controllers.DeletePaketsHandler)

	// Memulai server
	log.Println("Server started at :8000")
	r.Run(":8000")
}
