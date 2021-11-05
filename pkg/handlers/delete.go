package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
)

// Delete a message with specified id from the database
func (textChatHandler *TextChatHandler) DeleteMessage(responseWriter http.ResponseWriter, request *http.Request) {
	id := getTextChatID(request)
	log.Info("Delete message by ID request", "id", id)

	err := textChatHandler.db.DeleteMessage(request.Context(), id)

	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorMessageNotFound:
		log.Error(err, "Error deleting message, id does not exist")
		http.Error(responseWriter, "Message not found", http.StatusNotFound)
		return
	default:
		log.Error(err, "Error deleting message")
		http.Error(responseWriter, "Erro deleting message", http.StatusInternalServerError)
		return
	}
}

// Delete a conversation with specified id from the database
func (textChatHandler *TextChatHandler) DeleteConversation(responseWriter http.ResponseWriter, request *http.Request) {
	id := getTextChatID(request)
	log.Info("Delete conversation by ID request", "id", id)

	// TODO: Delete all messages from the conversation

	err := textChatHandler.db.DeleteConversation(request.Context(), id)

	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorConversationNotFound:
		log.Error(err, "Error deleting conversation, id does not exist")
		http.Error(responseWriter, "Conversation not found", http.StatusNotFound)
		return
	default:
		log.Error(err, "Error deleting conversation")
		http.Error(responseWriter, "Error deleting conversation", http.StatusInternalServerError)
		return
	}
}
