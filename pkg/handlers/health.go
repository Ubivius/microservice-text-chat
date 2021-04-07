package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
)

// LivenessCheck determine when the application needs to be restarted
func (textChatHandler *TextChatHandler) LivenessCheck(responseWriter http.ResponseWriter, request *http.Request) {
	log.Info("LivenessCheck")
	responseWriter.WriteHeader(http.StatusOK)
}

//ReadinessCheck verifies that the application is ready to accept requests
func (textChatHandler *TextChatHandler) ReadinessCheck(responseWriter http.ResponseWriter, request *http.Request) {
	log.Info("ReadinessCheck")

	readinessProbeMicroserviceUser := data.MicroserviceUserPath + "/health/ready"

	_, errMicroserviceUser := http.Get(readinessProbeMicroserviceUser)

	err := textChatHandler.db.PingDB()

	if err != nil {
		log.Error(err, "DB unavailable")
		http.Error(responseWriter, "DB unavailable", http.StatusServiceUnavailable)
		return
	}

	if errMicroserviceUser != nil {
		log.Error(errMicroserviceUser, "Microservice-user unavailable")
		http.Error(responseWriter, "Microservice-user unavailable", http.StatusServiceUnavailable)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
}
