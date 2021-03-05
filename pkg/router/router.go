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
	getRouter.HandleFunc("/messages", textChatHandler.GetMessages)
	getRouter.HandleFunc("/messages/{id:[0-9]+}", textChatHandler.GetMessageByID)
	getRouter.HandleFunc("/conversations", textChatHandler.GetConversations)
	getRouter.HandleFunc("/conversations/{id:[0-9]+}", textChatHandler.GetConversationByID)
	getRouter.HandleFunc("/messages/conversation/{id:[0-9]+}", textChatHandler.GetMessagesByConversationID)

	// Post router
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/messages", textChatHandler.AddMessage)
	postRouter.HandleFunc("/conversations", textChatHandler.AddConversation)
	postRouter.Use(textChatHandler.MiddlewareMessageValidation)

	// Delete router
	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/messages/{id:[0-9]+}", textChatHandler.DeleteMessage)
	deleteRouter.HandleFunc("/conversations/{id:[0-9]+}", textChatHandler.DeleteConversation)

	return router
}
