package repository

import (
	"eco-journal/entities"

	"gorm.io/gorm"
)

type ChatRepoInterface interface {
	SaveChat(chat entities.Chat) error
	GetAllChat() ([]entities.Chat, error)
}

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) *chatRepository {
	return &chatRepository{db}
}

func (r *chatRepository) SaveChat(chat entities.Chat) error {
	return r.db.Debug().Create(&chat).Error
}

func (r *chatRepository) GetAllChat() ([]entities.Chat, error) {
	var chats []entities.Chat
	err := r.db.Find(&chats).Error
	return chats, err
}
