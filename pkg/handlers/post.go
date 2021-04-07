package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
)

// AddMessage creates a new message from the received JSON
func (textChatHandler *TextChatHandler) AddMessage(responseWriter http.ResponseWriter, request *http.Request) {
	log.Info("AddMessage request")
	message := request.Context().Value(KeyMessage{}).(*data.Message)

	err := textChatHandler.db.AddMessage(message)
	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorConversationNotFound :
		log.Error(err, "Conversation not found")
		http.Error(responseWriter, "Conversation not found", http.StatusNotFound)
		return
	default:
		log.Error(err, "Error adding message")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

// AddConversation creates a new message from the received JSON
func (textChatHandler *TextChatHandler) AddConversation(responseWriter http.ResponseWriter, request *http.Request) {
	log.Info("AddConversation request")
	conversation := request.Context().Value(KeyConversation{}).(*data.Conversation)

	conversation, err := textChatHandler.db.AddConversation(conversation)

	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(conversation)
		if err != nil {
			log.Error(err, "Error serializing conversation")
		}
		return
	default:
		log.Error(err, "Error adding conversation")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
