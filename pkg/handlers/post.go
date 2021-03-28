package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
)

// AddMessage creates a new message from the received JSON
func (textChatHandler *TextChatHandler) AddMessage(responseWriter http.ResponseWriter, request *http.Request) {
	textChatHandler.logger.Println("Handle POST Message")
	message := request.Context().Value(KeyMessage{}).(*data.Message)

	err := textChatHandler.db.AddMessage(message)
	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorConversationNotFound :
		textChatHandler.logger.Println("[ERROR] adding, id does not exist")
		http.Error(responseWriter, "Conversation not found", http.StatusNotFound)
		return
	default:
		textChatHandler.logger.Println("[ERROR] adding message", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

// AddConversation creates a new message from the received JSON
func (textChatHandler *TextChatHandler) AddConversation(responseWriter http.ResponseWriter, request *http.Request) {
	textChatHandler.logger.Println("Handle POST Conversation")
	conversation := request.Context().Value(KeyConversation{}).(*data.Conversation)

	err := textChatHandler.db.AddConversation(conversation)

	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	default:
		textChatHandler.logger.Println("[ERROR] adding conversation", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
