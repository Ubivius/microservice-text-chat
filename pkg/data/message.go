package data

import (
	"fmt"
)

// ErrorMessageNotFound : Message specific errors
var ErrorMessageNotFound = fmt.Errorf("Message not found")

// Message defines the structure for an API message.
type Message struct {
	ID             string `json:"id" bson:"_id"`
	UserID         string `json:"user_id" validate:"required"`
	ConversationID string `json:"conversation_id" bson:"conversation_id" validate:"required"`
	Text           string `json:"text" validate:"required"`
	CreatedOn      string `json:"-"`
	UpdatedOn      string `json:"-"`
}

// Messages is a collection of Message
type Messages []*Message
