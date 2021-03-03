package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Ubivius/microservice-text-chat/data"
)

// Errors should be templated in the future.
// A good starting reference can be found here : https://github.com/nicholasjackson/building-microservices-youtube/blob/episode_7/message-api/handlers/middleware.go
// We want our validation errors to have a standard format

// MiddlewareMessageValidation is used to validate incoming message JSONS
func (textChatJandler *TextChatHandler) MiddlewareMessageValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		message := &data.Message{}

		err := data.FromJSON(message, request.Body)
		if err != nil {
			textChatJandler.logger.Println("[ERROR] deserializing message", err)
			http.Error(responseWriter, "Error reading message", http.StatusBadRequest)
			return
		}

		// validate the message
		err = message.ValidateMessage()
		if err != nil {
			textChatJandler.logger.Println("[ERROR] validating message", err)
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
