package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Ubivius/microservice-text-chat/pkg/database"
	"github.com/Ubivius/microservice-text-chat/pkg/handlers"
	"github.com/Ubivius/microservice-text-chat/pkg/router"
	"github.com/Ubivius/pkg-telemetry/metrics"
	"github.com/Ubivius/pkg-telemetry/tracing"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var log = logf.Log.WithName("text-chat-main")

func main() {
	// Starting k8s logger
	opts := zap.Options{}
	opts.BindFlags(flag.CommandLine)
	newLogger := zap.New(zap.UseFlagOptions(&opts), zap.WriteTo(os.Stdout))
	logf.SetLogger(newLogger.WithName("log"))

	// Starting tracer provider
	tp := tracing.CreateTracerProvider(os.Getenv("JAEGER_ENDPOINT"), "microservice-text-chat-traces")

	// Starting metrics exporter
	metrics.StartPrometheusExporterWithName("text-chat")

	// Database init
	db := database.NewMongoTextChat()

	// Creating handlers
	textChatHandler := handlers.NewTextChatHandler(db)

	// Mux route handling with gorilla/mux
	r := router.New(textChatHandler)

	// Server setup
	server := &http.Server{
		Addr:        ":9090",
		Handler:     r,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
	}

	go func() {
		log.Info("Starting server", "port", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			log.Error(err, "Server error")
		}
	}()

	// Mux route that is only exposed internally for the game server
	internalRouter := router.NewInternalRouter(textChatHandler)
	internalServer := &http.Server{
		Addr:        ":9091",
		Handler:     internalRouter,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
	}

	go func() {
		log.Info("Starting internal server", "port", internalServer.Addr)
		err := internalServer.ListenAndServe()
		if err != nil {
			log.Error(err, "Internal server error")
		}
	}()

	// Handle shutdown signals from operating system
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	receivedSignal := <-signalChannel

	log.Info("Received terminate, beginning graceful shutdown", "received_signal", receivedSignal.String())

	// DB connection shutdown
	db.CloseDB()

	// Context cancelling
	timeoutContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Cleanly shutdown and flush telemetry on shutdown
	defer func(ctx context.Context) {
		if err := tp.Shutdown(ctx); err != nil {
			log.Error(err, "Error shutting down tracer provider")
		}
	}(timeoutContext)

	// Server shutdown
	_ = server.Shutdown(timeoutContext)
	_ = internalServer.Shutdown(timeoutContext)
}
