package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
)

// Delete a message with specified id from the database
func (textChatHandler *TextChatHandler) DeleteMessage(responseWriter http.ResponseWriter, request *http.Request) {
	id := getTextChatID(request)
	textChatHandler.logger.Println("Handle DELETE message", id)

	err := data.DeleteMessage(id)
	if err == data.ErrorMessageNotFound {
		textChatHandler.logger.Println("[ERROR] deleting, id does not exist")
		http.Error(responseWriter, "Message not found", http.StatusNotFound)
		return
	}

	if err != nil {
		textChatHandler.logger.Println("[ERROR] deleting message", err)
		http.Error(responseWriter, "Erro deleting message", http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)
}

// Delete a conversation with specified id from the database
func (textChatHandler *TextChatHandler) DeleteConversation(responseWriter http.ResponseWriter, request *http.Request) {
	id := getTextChatID(request)
	textChatHandler.logger.Println("Handle DELETE conversation", id)

	// TODO: Delete all messages from the conversation

	err := data.DeleteConversation(id)
	if err == data.ErrorConversationNotFound {
		textChatHandler.logger.Println("[ERROR] deleting, id does not exist")
		http.Error(responseWriter, "Conversation not found", http.StatusNotFound)
		return
	}

	if err != nil {
		textChatHandler.logger.Println("[ERROR] deleting conversation", err)
		http.Error(responseWriter, "Error deleting conversation", http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)
}
