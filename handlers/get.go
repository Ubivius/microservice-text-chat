package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-text-chat/data"
)

// GetMessages returns the full list of messages
func (textChatJandler *TextChatHandler) GetMessages(responseWriter http.ResponseWriter, request *http.Request) {
	textChatJandler.logger.Println("Handle GET messages")
	messageList := data.GetMessages()
	err := data.ToJSON(messageList, responseWriter)
	if err != nil {
		textChatJandler.logger.Println("[ERROR] serializing message", err)
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// GetMessageByID returns a single message from the database
func (textChatJandler *TextChatHandler) GetMessageByID(responseWriter http.ResponseWriter, request *http.Request) {
	id := getTextChatID(request)

	textChatJandler.logger.Println("[DEBUG] getting id", id)

	message, err := data.GetMessageByID(id)
	switch err {
	case nil:
	case data.ErrorMessageNotFound:
		textChatJandler.logger.Println("[ERROR] fetching message", err)
		http.Error(responseWriter, "Message not found", http.StatusBadRequest)
		return
	default:
		textChatJandler.logger.Println("[ERROR] fetching message", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	err = data.ToJSON(message, responseWriter)
	if err != nil {
		textChatJandler.logger.Println("[ERROR] serializing message", err)
	}
}
