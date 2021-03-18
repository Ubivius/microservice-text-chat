package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
	"github.com/gorilla/mux"
)

func TestValidationMiddlewareWithValidBody(t *testing.T) {
	// Creating request body
	body := &data.Message{
		ID:             1,
		UserID:         1,
		ConversationID: 1,
		Text:           "This is a test message",
	}

	bodyBytes, _ := json.Marshal(body)
	request := httptest.NewRequest(http.MethodPost, "/messages", strings.NewReader(string(bodyBytes)))
	response := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(NewTestLogger())

	// Create a router for middleware because function attachment is handled by gorilla/mux
	router := mux.NewRouter()

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/messages", textChatHandler.AddMessage)
	postRouter.Use(textChatHandler.MiddlewareMessageValidation)

	// Server http on our router
	postRouter.ServeHTTP(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestValidationMiddlewareWithNoMessage(t *testing.T) {
	// Creating request body
	body := &data.Message{
		ID:             1,
		UserID:         1,
		ConversationID: 1,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Error("Body passing to test is not a valid json struct : ", err)
	}

	request := httptest.NewRequest(http.MethodPost, "/messages", strings.NewReader(string(bodyBytes)))
	response := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(NewTestLogger())

	// Create a router for middleware because linking is handled by gorilla/mux
	router := mux.NewRouter()
	router.HandleFunc("/messages", textChatHandler.AddMessage)
	router.Use(textChatHandler.MiddlewareMessageValidation)

	// Server http on our router
	router.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Field validation for 'Text' failed on the 'required' tag") {
		t.Error("Expected error on field validation for Text but got : ", response.Body.String())
	}
}
