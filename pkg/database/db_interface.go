package database

import (
	"context"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
)

// The interface that any kind of database must implement
type TextChatDB interface {
	GetMessageByID(ctx context.Context, id string) (*data.Message, error)
	GetConversationByID(ctx context.Context, id string) (*data.Conversation, error)
	GetMessagesByConversationID(ctx context.Context, id string) (data.Messages, error)
	AddMessage(ctx context.Context, message *data.Message) error
	AddConversation(ctx context.Context, conversation *data.Conversation) (*data.Conversation, error)
	AddUserToConversation(ctx context.Context, conversation *data.Conversation) error
	DeleteMessage(ctx context.Context, id string) error
	DeleteConversation(ctx context.Context, id string) error
	Connect() error
	PingDB() error
	CloseDB()
}
