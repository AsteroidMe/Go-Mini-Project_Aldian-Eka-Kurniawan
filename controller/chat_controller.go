package controller

import (
	"eco-journal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChatController struct {
	chatService service.ChatServiceInterface
}

func NewChatController(chatService service.ChatServiceInterface) *ChatController {
	return &ChatController{chatService}
}

func (ctc *ChatController) ChatController(c *gin.Context) {
	var input struct {
		UserInput string `json:"user_input"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Input Data Invalid"})
		return
	}
	chat, err := ctc.chatService.ProccessChat(input.UserInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "message": chat})
}

func (ctc *ChatController) GetAllChats(c *gin.Context) {
	chats, err := ctc.chatService.GetAllChats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "message": chats})
}
