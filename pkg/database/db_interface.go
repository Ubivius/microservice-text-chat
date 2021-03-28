package database

import (
	"github.com/Ubivius/microservice-text-chat/pkg/data"
)

// The interface that any kind of database must implement
type TextChatDB interface {
	GetMessageByID(id string) (*data.Message, error)
	GetConversationByID(id string) (*data.Conversation, error)
	GetMessagesByConversationID(id string) (data.Messages, error)
	AddMessage(message *data.Message) error
	AddConversation(conversation *data.Conversation) error
	DeleteMessage(id string) error
	DeleteConversation(id string) error
	GetConversationID(userID []string) string
	Connect() error
	CloseDB()
}
