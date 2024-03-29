package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
	"github.com/Ubivius/microservice-text-chat/pkg/database"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func newTextChatDB() database.TextChatDB {
	return database.NewMockTextChat()
}

func TestGetExistingMessageByID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/messages/a2181017-5c53-422b-b6bc-036b27c04fc8", nil)
	response := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(newTextChatDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}
	request = mux.SetURLVars(request, vars)

	textChatHandler.GetMessageByID(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got : %d", http.StatusOK, response.Code)
	}
	if !strings.Contains(response.Body.String(), "a2181017-5c53-422b-b6bc-036b27c04fc8") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetExistingConversationByID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/conversations/a2181017-5c53-422b-b6bc-036b27c04fc8", nil)
	response := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(newTextChatDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}
	request = mux.SetURLVars(request, vars)

	textChatHandler.GetMessageByID(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got : %d", http.StatusOK, response.Code)
	}
	if !strings.Contains(response.Body.String(), "a2181017-5c53-422b-b6bc-036b27c04fc8") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetNonExistingMessageByID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/messages/4", nil)
	response := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(newTextChatDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": uuid.NewString(),
	}
	request = mux.SetURLVars(request, vars)

	textChatHandler.GetMessageByID(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got : %d", http.StatusBadRequest, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Message not found") {
		t.Error("Expected response : Message not found")
	}
}

func TestGetNonExistingConversationByID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/conversations/4", nil)
	response := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(newTextChatDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": uuid.NewString(),
	}
	request = mux.SetURLVars(request, vars)

	textChatHandler.GetConversationByID(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got : %d", http.StatusBadRequest, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Conversation not found") {
		t.Error("Expected response : Conversation not found")
	}
}

func TestDeleteNonExistantMessage(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/messages/4", nil)
	response := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(newTextChatDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": uuid.NewString(),
	}
	request = mux.SetURLVars(request, vars)

	textChatHandler.DeleteMessage(response, request)
	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got : %d", http.StatusNotFound, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Message not found") {
		t.Error("Expected response : Message not found")
	}
}

func TestDeleteNonExistingConversation(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/conversations/4", nil)
	response := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(newTextChatDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": uuid.NewString(),
	}
	request = mux.SetURLVars(request, vars)

	textChatHandler.DeleteConversation(response, request)
	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got : %d", http.StatusNotFound, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Conversation not found") {
		t.Error("Expected response : Conversation not found")
	}
}

func TestAddMessage(t *testing.T) {
	// Creating request body
	body := &data.Message{
		UserID:         "a2181017-5c53-422b-b6bc-036b27c04fc8",
		ConversationID: "a2181017-5c53-422b-b6bc-036b27c04fc8",
		Text:           "This is a test message",
	}

	request := httptest.NewRequest(http.MethodPost, "/messages", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyMessage{}, body)
	request = request.WithContext(ctx)

	textChatHandler := NewTextChatHandler(newTextChatDB())
	textChatHandler.AddMessage(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestAddMessageNonExistingConversation(t *testing.T) {
	// Creating request body
	body := &data.Message{
		UserID:         "a2181017-5c53-422b-b6bc-036b27c04fc8",
		ConversationID: uuid.NewString(),
		Text:           "This is a test message",
	}

	request := httptest.NewRequest(http.MethodPost, "/messages", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyMessage{}, body)
	request = request.WithContext(ctx)

	textChatHandler := NewTextChatHandler(newTextChatDB())
	textChatHandler.AddMessage(response, request)

	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, response.Code)
	}
}

func TestAddConversation(t *testing.T) {
	// Creating request body
	body := &data.Conversation{
		UserID: []string{"a2181017-5c53-422b-b6bc-036b27c04fc8", "2aee2975-6b76-4340-b679-e81661b1cdb5"},
		GameID: "",
	}

	request := httptest.NewRequest(http.MethodPost, "/conversations", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyConversation{}, body)
	request = request.WithContext(ctx)

	textChatHandler := NewTextChatHandler(newTextChatDB())
	textChatHandler.AddConversation(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, response.Code)
	}
}

func TestDeleteExistingMessage(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/messages/a2181017-5c53-422b-b6bc-036b27c04fc8", nil)
	response := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(newTextChatDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}
	request = mux.SetURLVars(request, vars)

	textChatHandler.DeleteMessage(response, request)
	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d but got : %d", http.StatusNoContent, response.Code)
	}
}

func TestDeleteExistingConversation(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/conversations/a2181017-5c53-422b-b6bc-036b27c04fc8", nil)
	response := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(newTextChatDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}
	request = mux.SetURLVars(request, vars)

	textChatHandler.DeleteConversation(response, request)
	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d but got : %d", http.StatusNoContent, response.Code)
	}
}

func TestAddUserToExistingConversation(t *testing.T) {
	// Creating request body
	body := &data.Conversation{
		ID: "e2382ea2-b5fa-4506-aa9d-d338aa52af44",
		UserID: []string{"a2181017-5c53-422b-b6bc-036b27c04fc8",
			"2aee2975-6b76-4340-b679-e81661b1cdb5",
			"3a1c152e-f172-41de-a5ab-ca21f6573bf3",
			"c6e6a2b2-bd25-4151-ace1-611accc15a50",
			"newUser"},
		GameID: "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}

	request := httptest.NewRequest(http.MethodPut, "/conversations", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyConversation{}, body)
	request = request.WithContext(ctx)

	textChatHandler := NewTextChatHandler(newTextChatDB())
	textChatHandler.AddUserToConversation(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}
