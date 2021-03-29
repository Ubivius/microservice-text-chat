package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
)

// AddMessage creates a new message from the received JSON
func (textChatHandler *TextChatHandler) AddMessage(responseWriter http.ResponseWriter, request *http.Request) {
	textChatHandler.logger.Println("Handle POST Message")
	message := request.Context().Value(KeyMessage{}).(*data.Message)

	err := data.AddMessage(message)
	if err == data.ErrorConversationNotFound {
		textChatHandler.logger.Println("[ERROR] adding, id does not exist")
		http.Error(responseWriter, "Conversation not found", http.StatusNotFound)
		return
	}
	responseWriter.WriteHeader(http.StatusNoContent)
}

// AddConversation creates a new message from the received JSON
func (textChatHandler *TextChatHandler) AddConversation(responseWriter http.ResponseWriter, request *http.Request) {
	textChatHandler.logger.Println("Handle POST Conversation")
	conversation := request.Context().Value(KeyConversation{}).(*data.Conversation)

	data.AddConversation(conversation)

	responseWriter.WriteHeader(http.StatusNoContent)
}
