package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
	"go.opentelemetry.io/otel"
)

// AddConversation creates a new message from the received JSON
func (textChatHandler *TextChatHandler) AddUserToConversation(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("text-chat").Start(request.Context(), "addUserToConversation")
	defer span.End()
	log.Info("Add User to Conversation request")
	conversation := request.Context().Value(KeyConversation{}).(*data.Conversation)

	err := textChatHandler.db.AddUserToConversation(request.Context(), conversation)

	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorUserNotFound:
		log.Error(err, "A UserID doesn't exist")
		http.Error(responseWriter, "A UserID doesn't exist", http.StatusBadRequest)
		return
	case data.ErrorGameNotFound:
		log.Error(err, "GameID doesn't exist")
		http.Error(responseWriter, "GameID doesn't exist", http.StatusBadRequest)
		return
	default:
		log.Error(err, "Error adding conversation")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
