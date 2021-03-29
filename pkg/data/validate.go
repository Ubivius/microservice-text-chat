package data

import (
	"github.com/go-playground/validator"
)

func (message *Message) ValidateMessage() error {
	validate := validator.New()
	// To be discussed (depend on how we manage our id)
	return validate.Struct(message)
}

func (conversation *Conversation) ValidateConversation() error {
	validate := validator.New()
	// To be discussed (depend on how we manage our id)
	return validate.Struct(conversation)
}
