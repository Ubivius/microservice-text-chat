package router

import (
	"log"
	"net/http"

	"github.com/Ubivius/microservice-text-chat/pkg/handlers"
	"github.com/gorilla/mux"
)

// Mux route handling with gorilla/mux
func New(textChatHandler *handlers.TextChatHandler, logger *log.Logger) *mux.Router {
	router := mux.NewRouter()

	// Get Router
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/messages/{id:[0-9a-z-]+}", textChatHandler.GetMessageByID)
	getRouter.HandleFunc("/conversations/{id:[0-9a-z-]+}", textChatHandler.GetConversationByID)
	getRouter.HandleFunc("/messages/conversation/{id:[0-9a-z-]+}", textChatHandler.GetMessagesByConversationID)

	// Message post router
	messagePostRouter := router.Methods(http.MethodPost).Subrouter()
	messagePostRouter.HandleFunc("/messages", textChatHandler.AddMessage)
	messagePostRouter.Use(textChatHandler.MiddlewareMessageValidation)

	// Conversation post router
	conversationPostRouter := router.Methods(http.MethodPost).Subrouter()
	conversationPostRouter.HandleFunc("/conversations", textChatHandler.AddConversation)
	conversationPostRouter.Use(textChatHandler.MiddlewareConversationValidation)

	// Delete router
	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/messages/{id:[0-9a-z-]+}", textChatHandler.DeleteMessage)
	deleteRouter.HandleFunc("/conversations/{id:[0-9a-z-]+}", textChatHandler.DeleteConversation)

	return router
}
