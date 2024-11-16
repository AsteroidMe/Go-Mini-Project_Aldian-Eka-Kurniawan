package route

import (
	"eco-journal/config"
	"eco-journal/controller"
	"eco-journal/middleware"
	"eco-journal/repository"
	"eco-journal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	authorRepo := repository.NewAuthorRepository(config.DB)
	authorService := service.NewAuthorService(authorRepo)
	authorController := controller.NewAuthorController(authorService)

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	authorRoutes := r.Group("/authors")
	authorRoutes.Use(middleware.AuthMiddleware())
	{
		authorRoutes.GET("/", authorController.GetAll)
		authorRoutes.GET("/:id", authorController.GetDetails)
		authorRoutes.POST("/", authorController.Create)
		authorRoutes.PUT("/:id", authorController.Update)
		authorRoutes.DELETE("/:id", authorController.Delete)
	}

	return r
}
