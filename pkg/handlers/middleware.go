package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
)

// Errors should be templated in the future.
// A good starting reference can be found here : https://github.com/nicholasjackson/building-microservices-youtube/blob/episode_7/message-api/handlers/middleware.go
// We want our validation errors to have a standard format

// MiddlewareMessageValidation is used to validate incoming message JSONS
func (textChatHandler *TextChatHandler) MiddlewareMessageValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		message := &data.Message{}
		err := json.NewDecoder(request.Body).Decode(message)
		if err != nil {
			textChatHandler.logger.Println("[ERROR] deserializing message", err)
			http.Error(responseWriter, "Error reading message", http.StatusBadRequest)
			return
		}

		// validate the message
		err = message.ValidateMessage()
		if err != nil {
			textChatHandler.logger.Println("[ERROR] validating message", err)
			http.Error(responseWriter, fmt.Sprintf("Error validating message: %s", err), http.StatusBadRequest)
			return
		}

		// Add the message to the context
		ctx := context.WithValue(request.Context(), KeyMessage{}, message)
		newRequest := request.WithContext(ctx)

		// Call the next handler, which can be another middleware or the final handler
		next.ServeHTTP(responseWriter, newRequest)
	})
}

// MiddlewareConversatonValidation is used to validate incoming conversation JSONS
func (textChatHandler *TextChatHandler) MiddlewareConversationValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		conversation := &data.Conversation{}

		err := json.NewDecoder(request.Body).Decode(conversation)
		if err != nil {
			textChatHandler.logger.Println("[ERROR] deserializing conversation", err)
			http.Error(responseWriter, "Error reading conversation", http.StatusBadRequest)
			return
		}

		// validate the conversation
		err = conversation.ValidateConversation()
		if err != nil {
			textChatHandler.logger.Println("[ERROR] validating conversation", err)
			http.Error(responseWriter, fmt.Sprintf("Error validating conversation: %s", err), http.StatusBadRequest)
			return
		}

		// Add the conversation to the context
		ctx := context.WithValue(request.Context(), KeyConversation{}, conversation)
		newRequest := request.WithContext(ctx)

		// Call the next handler, which can be another middleware or the final handler
		next.ServeHTTP(responseWriter, newRequest)
	})
}
