package main

import (
	"eco-journal/config"
	"eco-journal/controller"
	"eco-journal/repository"
	"eco-journal/route"
	"eco-journal/service"
	"log"
)

func main() {
	config.ConnectDB()
	config.MigrateDB()

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

	chatRepo := repository.NewChatRepository(config.DB)
	chatService := service.NewChatService(chatRepo)
	chatController := controller.NewChatController(chatService)

	r := route.SetupRouter(userController, authorController, categoryController, journalController, chatController)
	if err := r.Run(":8000"); err != nil {
		log.Fatal("Server Run Failed:", err)
	}
}
