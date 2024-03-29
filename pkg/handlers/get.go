package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
	"go.opentelemetry.io/otel"
)

func (textChatHandler *TextChatHandler) GetMessageByID(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("text-chat").Start(request.Context(), "getMessageById")
	defer span.End()
	id := getTextChatID(request)

	log.Info("GetMessageByID request for ID", "id", id)

	message, err := textChatHandler.db.GetMessageByID(request.Context(), id)
	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(message)
		if err != nil {
			log.Error(err, "Error serializing message")
		}
		return
	case data.ErrorMessageNotFound:
		log.Error(err, "Message not found")
		http.Error(responseWriter, "Message not found", http.StatusBadRequest)
		return
	default:
		log.Error(err, "Error fetching message")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (textChatHandler *TextChatHandler) GetConversationByID(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("text-chat").Start(request.Context(), "getConversationById")
	defer span.End()
	id := getTextChatID(request)

	log.Info("GetConversationByID request for ID", "id", id)

	conversation, err := textChatHandler.db.GetConversationByID(request.Context(), id)
	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(conversation)
		if err != nil {
			log.Error(err, "Error serializing conversation")
		}
		return
	case data.ErrorConversationNotFound:
		log.Error(err, "Conversation not found")
		http.Error(responseWriter, "Conversation not found", http.StatusBadRequest)
		return
	default:
		log.Error(err, "Error fetching conversation")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (textChatHandler *TextChatHandler) GetMessagesByConversationID(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("text-chat").Start(request.Context(), "getMessagesByConversationId")
	defer span.End()
	id := getTextChatID(request)

	log.Info("GetMessagesByConversationID request for conversationID", "id", id)

	messages, err := textChatHandler.db.GetMessagesByConversationID(request.Context(), id)
	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(messages)
		if err != nil {
			log.Error(err, "Error serializing messages")
		}
		return
	case data.ErrorMessageNotFound:
		log.Error(err, "Messages not found")
		http.Error(responseWriter, "Messages not found", http.StatusBadRequest)
		return
	default:
		log.Error(err, "Error fetching messages")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
