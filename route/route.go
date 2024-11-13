package route

import (
	"eco-journal/config"
	"eco-journal/controller"
	"eco-journal/repository"
	"eco-journal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	return r
}
