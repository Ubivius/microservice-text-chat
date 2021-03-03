package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-text-chat/data"
)

// AddMessage creates a new message from the received JSON
func (textChatJandler *TextChatHandler) AddMessage(responseWriter http.ResponseWriter, request *http.Request) {
	textChatJandler.logger.Println("Handle POST Message")
	message := request.Context().Value(KeyMessage{}).(*data.Message)

	data.AddMessage(message)
}
