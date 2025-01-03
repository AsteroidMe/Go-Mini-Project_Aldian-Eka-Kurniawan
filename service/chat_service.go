package service

import (
	"context"
	"eco-journal/entities"
	"eco-journal/repository"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type ChatServiceInterface interface {
	ProccessChat(userInput string) (entities.Chat, error)
	GetAllChats() ([]entities.Chat, error)
}

type chatService struct {
	chatRepo repository.ChatRepoInterface
}

func NewChatService(chatRepo repository.ChatRepoInterface) *chatService {
	return &chatService{chatRepo}
}

func (cuc *chatService) ProccessChat(userInput string) (entities.Chat, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return entities.Chat{}, err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(userInput))
	if err != nil {
		return entities.Chat{}, err
	}
	if len(resp.Candidates) == 0 {
		return entities.Chat{}, err
	}
	aiResponse := ""
	for _, candidate := range resp.Candidates {
		if candidate.Content == nil {
			continue
		}
		for _, part := range candidate.Content.Parts {
			aiResponse += fmt.Sprintf("%v", part)
		}
	}
	chat := entities.Chat{
		UserInput: userInput,
		AiRespon:  aiResponse,
	}
	if err := cuc.chatRepo.SaveChat(chat); err != nil {
		return entities.Chat{}, err
	}

	return chat, nil

}

func (cts *chatService) GetAllChats() ([]entities.Chat, error) {
	return cts.chatRepo.GetAllChat()
}
