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
	conversationBody := &data.Conversation{
		ID:     1,
		UserID: []int{1, 3},
		GameID: -1,
	}

	messageBody := &data.Message{
		ID:             1,
		UserID:         1,
		ConversationID: 1,
		Text:           "This is a test message",
	}

	conversationBodyBytes, _ := json.Marshal(conversationBody)
	conversationRequest := httptest.NewRequest(http.MethodPost, "/conversations", strings.NewReader(string(conversationBodyBytes)))
	conversationResponse := httptest.NewRecorder()

	messageBodyBytes, _ := json.Marshal(messageBody)
	messageRequest := httptest.NewRequest(http.MethodPost, "/messages", strings.NewReader(string(messageBodyBytes)))
	messageResponse := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(NewTestLogger())

	// Create a router for middleware because function attachment is handled by gorilla/mux
	router := mux.NewRouter()

	conversationPostRouter := router.Methods(http.MethodPost).Subrouter()
	conversationPostRouter.HandleFunc("/conversations", textChatHandler.AddConversation)
	conversationPostRouter.Use(textChatHandler.MiddlewareConversationValidation)

	messagePostRouter := router.Methods(http.MethodPost).Subrouter()
	messagePostRouter.HandleFunc("/messages", textChatHandler.AddMessage)
	messagePostRouter.Use(textChatHandler.MiddlewareMessageValidation)

	// Server http on our router
	conversationPostRouter.ServeHTTP(conversationResponse, conversationRequest)
	messagePostRouter.ServeHTTP(messageResponse, messageRequest)

	if messageResponse.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, messageResponse.Code)
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
