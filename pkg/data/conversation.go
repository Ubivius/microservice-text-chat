package data

import (
	"fmt"
)

// ErrorConversationNotFound : Conversation specific errors
var ErrorConversationNotFound = fmt.Errorf("Conversation not found")

// Conversation defines the structure for an API conversation.
type Conversation struct {
	ID        string    `json:"id" bson:"_id"`
	UserID    []string  `json:"user_id" validate:"required"`
	GameID    string    `json:"game_id"`
	CreatedOn string    `json:"-"`
	UpdatedOn string    `json:"-"`
}

// Conversations is a collection of Conversation
type Conversations []*Conversation
