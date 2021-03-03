package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-text-chat/data"
)

// Delete a message with specified id from the database
func (textChatJandler *TextChatHandler) DeleteMessage(responseWriter http.ResponseWriter, request *http.Request) {
	id := getTextChatID(request)
	textChatJandler.logger.Println("Handle DELETE message", id)

	err := data.DeleteMessage(id)
	if err == data.ErrorMessageNotFound {
		textChatJandler.logger.Println("[ERROR] deleting, id does not exist")
		http.Error(responseWriter, "Message not found", http.StatusNotFound)
		return
	}

	if err != nil {
		textChatJandler.logger.Println("[ERROR] deleting message", err)
		http.Error(responseWriter, "Erro deleting message", http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)
}

// Delete a message with specified id from the database
func (textChatJandler *TextChatHandler) DeleteConversation(responseWriter http.ResponseWriter, request *http.Request) {
	id := getTextChatID(request)
	textChatJandler.logger.Println("Handle DELETE conversation", id)

	// Delete de tous les messages de la conversation

	err := data.DeleteConversation(id)
	if err == data.ErrorConversationNotFound {
		textChatJandler.logger.Println("[ERROR] deleting, id does not exist")
		http.Error(responseWriter, "Conversation not found", http.StatusNotFound)
		return
	}

	if err != nil {
		textChatJandler.logger.Println("[ERROR] deleting conversation", err)
		http.Error(responseWriter, "Error deleting conversation", http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)
}
