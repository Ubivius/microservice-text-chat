package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Ubivius/microservice-text-chat/handlers"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/exporters/stdout"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	// Logger
	logger := log.New(os.Stdout, "Text-Chat", log.LstdFlags)

	// Initialising open telemetry

	// Creating console exporter
	exporter, err := stdout.NewExporter(
		stdout.WithPrettyPrint(),
	)
	if err != nil {
		logger.Fatal("Failed to initialize stdout export pipeline : ", err)
	}

	// Creating tracer provider
	ctx := context.Background()
	batchSpanProcessor := sdktrace.NewBatchSpanProcessor(exporter)
	tracerProvider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(batchSpanProcessor))
	defer func() { _ = tracerProvider.Shutdown(ctx) }()

	// Creating handlers
	textChatHandler := handlers.NewTextChatHandler(logger)

	// Mux route handling with gorilla/mux
	router := mux.NewRouter()

	// Get Router
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/messages", textChatHandler.GetMessages)
	getRouter.HandleFunc("/messages/{id:[0-9]+}", textChatHandler.GetMessageByID)
	getRouter.HandleFunc("/messages/conversation/{id:[0-9]+}", textChatHandler.GetMessagesByConversationID)

	// -> Post conversation
	// -> Post message: Envoie un message écrie dans la bd

	// Put router
	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/messages", textChatHandler.UpdateMessages)
	putRouter.HandleFunc("/conversations", textChatHandler.UpdateConversation)
	putRouter.Use(textChatHandler.MiddlewareMessageValidation)

	// Post router
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/messages", textChatHandler.AddMessage)
	postRouter.HandleFunc("/conversations", textChatHandler.AddConversation)
	postRouter.Use(textChatHandler.MiddlewareMessageValidation)

	// Delete router
	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/messages/{id:[0-9]+}", textChatHandler.Delete)
	deleteRouter.HandleFunc("/conversations/{id:[0-9]+}", textChatHandler.Delete)

	// Server setup
	server := &http.Server{
		Addr:        ":9090",
		Handler:     router,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
	}

	go func() {
		logger.Println("Starting server on port ", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			logger.Println("Error starting server : ", err)
			logger.Fatal(err)
		}
	}()

	// Handle shutdown signals from operating system
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	receivedSignal := <-signalChannel

	logger.Println("Received terminate, beginning graceful shutdown", receivedSignal)

	// Server shutdown
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_ = server.Shutdown(timeoutContext)
}
