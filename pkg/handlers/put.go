package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
)

// UpdateMessages updates the message with the ID specified in the received JSON message
func (textChatHandler *TextChatHandler) UpdateMessage(responseWriter http.ResponseWriter, request *http.Request) {
	message := request.Context().Value(KeyMessage{}).(data.Message)
	textChatHandler.logger.Println("Handle PUT message", message.ID)

	// Update message
	err := data.UpdateMessage(&message)
	if err == data.ErrorMessageNotFound {
		textChatHandler.logger.Println("[ERROR} message not found", err)
		http.Error(responseWriter, "Message not found", http.StatusNotFound)
		return
	}

	// Returns status, no content required
	responseWriter.WriteHeader(http.StatusNoContent)
}
