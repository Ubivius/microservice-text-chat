package router

import (
	"net/http"

	"github.com/Ubivius/microservice-text-chat/pkg/handlers"
	"github.com/Ubivius/pkg-telemetry/metrics"
	tokenValidation "github.com/Ubivius/shared-authentication/pkg/auth"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

// Mux route handling with gorilla/mux
func New(textChatHandler *handlers.TextChatHandler) *mux.Router {
	log.Info("Starting router")
	router := mux.NewRouter()
	router.Use(otelmux.Middleware("text-chat"))
	router.Use(metrics.RequestCountMiddleware)

	// Get Router
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.Use(tokenValidation.Middleware)
	getRouter.HandleFunc("/messages/{id:[0-9a-z-]+}", textChatHandler.GetMessageByID)
	getRouter.HandleFunc("/conversations/{id:[0-9a-z-]+}", textChatHandler.GetConversationByID)
	getRouter.HandleFunc("/messages/conversation/{id:[0-9a-z-]+}", textChatHandler.GetMessagesByConversationID)

	//Health Check
	getRouter.HandleFunc("/health/live", textChatHandler.LivenessCheck)
	getRouter.HandleFunc("/health/ready", textChatHandler.ReadinessCheck)

	// Message post router
	messagePostRouter := router.Methods(http.MethodPost).Subrouter()
	messagePostRouter.Use(tokenValidation.Middleware)
	messagePostRouter.HandleFunc("/messages", textChatHandler.AddMessage)
	messagePostRouter.Use(textChatHandler.MiddlewareMessageValidation)

	// Conversation post router
	conversationPostRouter := router.Methods(http.MethodPost).Subrouter()
	conversationPostRouter.Use(tokenValidation.Middleware)
	conversationPostRouter.HandleFunc("/conversations", textChatHandler.AddConversation)
	conversationPostRouter.Use(textChatHandler.MiddlewareConversationValidation)

	// Delete router
	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.Use(tokenValidation.Middleware)
	deleteRouter.HandleFunc("/messages/{id:[0-9a-z-]+}", textChatHandler.DeleteMessage)
	deleteRouter.HandleFunc("/conversations/{id:[0-9a-z-]+}", textChatHandler.DeleteConversation)

	// Conversation put router
	conversationPutRouter := router.Methods(http.MethodPut).Subrouter()
	conversationPostRouter.Use(tokenValidation.Middleware)
	conversationPutRouter.HandleFunc("/conversations", textChatHandler.AddUserToConversation)
	conversationPutRouter.Use(textChatHandler.MiddlewareConversationValidation)

	return router
}

func NewInternalRouter(textChatHandler *handlers.TextChatHandler) *mux.Router {
	log.Info("Starting router")
	router := mux.NewRouter()
	router.Use(otelmux.Middleware("text-chat"))
	router.Use(metrics.RequestCountMiddleware)

	// Conversation post router
	conversationPostRouter := router.Methods(http.MethodPost).Subrouter()
	conversationPostRouter.HandleFunc("/conversations", textChatHandler.AddConversation)
	conversationPostRouter.Use(textChatHandler.MiddlewareConversationValidation)

	// Conversation put router
	conversationPutRouter := router.Methods(http.MethodPut).Subrouter()
	conversationPutRouter.HandleFunc("/conversations", textChatHandler.AddUserToConversation)
	conversationPutRouter.Use(textChatHandler.MiddlewareConversationValidation)

	return router
}
