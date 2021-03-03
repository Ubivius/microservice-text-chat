package data

import (
	"github.com/go-playground/validator"
)

func (message *Message) ValidateMessage() error {
	validate := validator.New()

	return validate.Struct(message)
}

func (conversation *Conversation) ValidateConversation() error {
	validate := validator.New()

	return validate.Struct(conversation)
}
