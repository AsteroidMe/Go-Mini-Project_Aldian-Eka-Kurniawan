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

	categoryRepo := repository.NewCategoryRepository(config.DB)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryController := controller.NewCategoryController(categoryService)

	journalRepo := repository.NewJournalRepository(config.DB)
	journalService := service.NewJournalService(journalRepo)
	journalController := controller.NewJournalController(journalService)

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

	categoryRoutes := r.Group("/categories")
	categoryRoutes.Use(middleware.AuthMiddleware())
	{
		categoryRoutes.GET("/", categoryController.GetAll)
		categoryRoutes.GET("/:id", categoryController.GetDetails)
		categoryRoutes.POST("/", categoryController.Create)
		categoryRoutes.PUT("/:id", categoryController.Update)
		categoryRoutes.DELETE("/:id", categoryController.Delete)
	}

	journalRoutes := r.Group("/journals")
	journalRoutes.Use(middleware.AuthMiddleware())
	{
		journalRoutes.GET("/", journalController.GetAll)
		journalRoutes.GET("/:id", journalController.GetDetails)
		journalRoutes.POST("/", journalController.Create)
		journalRoutes.PUT("/:id", journalController.Update)
		journalRoutes.DELETE("/:id", journalController.Delete)
	}

	return r
}
