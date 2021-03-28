package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
)

func (textChatHandler *TextChatHandler) GetMessageByID(responseWriter http.ResponseWriter, request *http.Request) {
	id := getTextChatID(request)

	textChatHandler.logger.Println("[DEBUG] getting id", id)

	message, err := textChatHandler.db.GetMessageByID(id)
	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(message)
		if err != nil {
			textChatHandler.logger.Println("[ERROR] serializing message", err)
		}
		return
	case data.ErrorMessageNotFound:
		textChatHandler.logger.Println("[ERROR] fetching message", err)
		http.Error(responseWriter, "Message not found", http.StatusBadRequest)
		return
	default:
		textChatHandler.logger.Println("[ERROR] fetching message", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (textChatHandler *TextChatHandler) GetConversationByID(responseWriter http.ResponseWriter, request *http.Request) {
	id := getTextChatID(request)

	textChatHandler.logger.Println("[DEBUG] getting id", id)

	conversation, err := textChatHandler.db.GetConversationByID(id)
	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(conversation)
		if err != nil {
			textChatHandler.logger.Println("[ERROR] serializing conversation", err)
		}
		return
	case data.ErrorConversationNotFound:
		textChatHandler.logger.Println("[ERROR] fetching conversation", err)
		http.Error(responseWriter, "Conversation not found", http.StatusBadRequest)
		return
	default:
		textChatHandler.logger.Println("[ERROR] fetching conversation", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (textChatHandler *TextChatHandler) GetMessagesByConversationID(responseWriter http.ResponseWriter, request *http.Request) {
	id := getTextChatID(request)

	textChatHandler.logger.Println("[DEBUG] getting id", id)

	messages, err := textChatHandler.db.GetMessagesByConversationID(id)
	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(messages)
		if err != nil {
			textChatHandler.logger.Println("[ERROR] serializing messages", err)
		}
		return
	case data.ErrorMessageNotFound:
		textChatHandler.logger.Println("[ERROR] fetching messages", err)
		http.Error(responseWriter, "Message not found", http.StatusBadRequest)
		return
	default:
		textChatHandler.logger.Println("[ERROR] fetching messages", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
