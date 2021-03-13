package data

import (
	"github.com/go-playground/validator"
)

func (message *Message) ValidateMessage() error {
	validate := validator.New()

	// Validate that the conversation ID exist
	err1 := validate.RegisterValidation("isconversation", validateIsConversation)
	// Validate that the user ID exist
	err2 := validate.RegisterValidation("isuser", validateIsUser)

	if err1 != nil {
		panic(err1)
	} else if err2 != nil {
		panic(err2)
	}

	return validate.Struct(message)
}

func (conversation *Conversation) ValidateConversation() error {
	validate := validator.New()

	// Validate that the users ID exist
	err1 := validate.RegisterValidation("isuser", validateIsUser)
	// Validate that the game ID exist
	err2 := validate.RegisterValidation("isgame", validateIsGame)

	if err1 != nil {
		panic(err1)
	} else if err2 != nil {
		panic(err2)
	}
	return validate.Struct(conversation)
}

func validateIsConversation(fieldLevel validator.FieldLevel) bool {
	_, err := GetConversationByID(int(fieldLevel.Field().Int()))
	return err == nil
}

func validateIsUser(fieldLevel validator.FieldLevel) bool {
	// validation of the UserID with a call to microservice-user
	return true
}

func validateIsGame(fieldLevel validator.FieldLevel) bool {
	// validation of the GameID with a call to microservice-user
	return true
}
