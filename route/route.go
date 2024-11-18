package route

import (
	"eco-journal/controller"
	"eco-journal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userController *controller.UserController,
	authorController *controller.AuthorController,
	categoryController *controller.CategoryController,
	journalController *controller.JournalController,
	chatController *controller.ChatController) *gin.Engine {
	r := gin.Default()

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	categoryRoutes := r.Group("/categories")
	categoryRoutes.Use(middleware.AuthMiddleware())
	{
		categoryRoutes.GET("/", categoryController.GetAll)
		categoryRoutes.GET("/:id", categoryController.GetDetails)
		categoryRoutes.POST("/", categoryController.Create)
		categoryRoutes.PUT("/:id", categoryController.Update)
		categoryRoutes.DELETE("/:id", categoryController.Delete)
	}

	authorRoutes := r.Group("/authors")
	authorRoutes.Use(middleware.AuthMiddleware())
	{
		authorRoutes.GET("/", authorController.GetAll)
		authorRoutes.GET("/:id", authorController.GetDetails)
		authorRoutes.POST("/", authorController.Create)
		authorRoutes.PUT("/:id", authorController.Update)
		authorRoutes.DELETE("/:id", authorController.Delete)
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

	chatRoutes := r.Group("/chat")
	chatRoutes.Use(middleware.AuthMiddleware())
	{
		chatRoutes.GET("", chatController.GetAllChats)
		chatRoutes.POST("", chatController.ChatController)
	}

	return r
}
