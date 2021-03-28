package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
)

// Delete a message with specified id from the database
func (textChatHandler *TextChatHandler) DeleteMessage(responseWriter http.ResponseWriter, request *http.Request) {
	id := getTextChatID(request)
	textChatHandler.logger.Println("Handle DELETE message", id)

	err := textChatHandler.db.DeleteMessage(id)

	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorMessageNotFound :
		textChatHandler.logger.Println("[ERROR] deleting, id does not exist")
		http.Error(responseWriter, "Message not found", http.StatusNotFound)
		return
	default:
		textChatHandler.logger.Println("[ERROR] deleting message", err)
		http.Error(responseWriter, "Erro deleting message", http.StatusInternalServerError)
		return
	}
}

// Delete a conversation with specified id from the database
func (textChatHandler *TextChatHandler) DeleteConversation(responseWriter http.ResponseWriter, request *http.Request) {
	id := getTextChatID(request)
	textChatHandler.logger.Println("Handle DELETE conversation", id)

	// TODO: Delete all messages from the conversation

	err := textChatHandler.db.DeleteConversation(id)

	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorConversationNotFound :
		textChatHandler.logger.Println("[ERROR] deleting, id does not exist")
		http.Error(responseWriter, "Conversation not found", http.StatusNotFound)
		return
	default:
		textChatHandler.logger.Println("[ERROR] deleting conversation", err)
		http.Error(responseWriter, "Error deleting conversation", http.StatusInternalServerError)
		return
	}
}
