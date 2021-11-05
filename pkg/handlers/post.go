package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
	"go.opentelemetry.io/otel"
)

// AddMessage creates a new message from the received JSON
func (textChatHandler *TextChatHandler) AddMessage(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("text-chat").Start(request.Context(), "addMessage")
	defer span.End()
	log.Info("AddMessage request")
	message := request.Context().Value(KeyMessage{}).(*data.Message)

	err := textChatHandler.db.AddMessage(request.Context(), message)
	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorConversationNotFound:
		log.Error(err, "Conversation not found")
		http.Error(responseWriter, "Conversation not found", http.StatusNotFound)
		return
	case data.ErrorUserNotFound:
		log.Error(err, "UserID doesn't exist")
		http.Error(responseWriter, "UserID doesn't exist", http.StatusBadRequest)
		return
	default:
		log.Error(err, "Error adding message")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

// AddConversation creates a new message from the received JSON
func (textChatHandler *TextChatHandler) AddConversation(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("text-chat").Start(request.Context(), "addConversation")
	defer span.End()
	log.Info("AddConversation request")
	conversation := request.Context().Value(KeyConversation{}).(*data.Conversation)

	conversation, err := textChatHandler.db.AddConversation(request.Context(), conversation)

	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(conversation)
		if err != nil {
			log.Error(err, "Error serializing conversation")
		}
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
