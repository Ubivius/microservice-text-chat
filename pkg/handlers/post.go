package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
)

// AddMessage creates a new message from the received JSON
func (textChatHandler *TextChatHandler) AddMessage(responseWriter http.ResponseWriter, request *http.Request) {
	textChatHandler.logger.Println("Handle POST Message")
	message := request.Context().Value(KeyMessage{}).(*data.Message)

	data.AddMessage(message)

	responseWriter.WriteHeader(http.StatusNoContent)
}

func (textChatHandler *TextChatHandler) AddConversation(responseWriter http.ResponseWriter, request *http.Request) {
	textChatHandler.logger.Println("Handle POST Conversation")
	conversation := request.Context().Value(KeyConversation{}).(*data.Conversation)

	data.AddConversation(conversation)

	responseWriter.WriteHeader(http.StatusNoContent)
}
