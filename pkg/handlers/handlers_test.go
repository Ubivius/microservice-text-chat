package handlers

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
	"github.com/gorilla/mux"
)

// Move to util package in Sprint 9, should be a testing specific logger
func NewTestLogger() *log.Logger {
	return log.New(os.Stdout, "Tests", log.LstdFlags)
}

func TestGetMessages(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/messages", nil)
	response := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(NewTestLogger())
	textChatHandler.GetMessages(response, request)

	if response.Code != 200 {
		t.Errorf("Expected status code 200 but got : %d", response.Code)
	}
	if !strings.Contains(response.Body.String(), "\"id\":2") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetConversations(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/conversations", nil)
	response := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(NewTestLogger())
	textChatHandler.GetConversations(response, request)

	if response.Code != 200 {
		t.Errorf("Expected status code 200 but got : %d", response.Code)
	}
	if !strings.Contains(response.Body.String(), "\"id\":2") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetExistingMessageByID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/messages/1", nil)
	response := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "1",
	}
	request = mux.SetURLVars(request, vars)

	textChatHandler.GetMessageByID(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got : %d", http.StatusOK, response.Code)
	}
	if !strings.Contains(response.Body.String(), "\"id\":1") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetExistingConversationByID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/conversations/1", nil)
	response := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "1",
	}
	request = mux.SetURLVars(request, vars)

	textChatHandler.GetMessageByID(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got : %d", http.StatusOK, response.Code)
	}
	if !strings.Contains(response.Body.String(), "\"id\":1") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetNonExistingMessageByID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/messages/4", nil)
	response := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "4",
	}
	request = mux.SetURLVars(request, vars)

	textChatHandler.GetMessageByID(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got : %d", http.StatusBadRequest, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Product not found") {
		t.Error("Expected response : Product not found")
	}
}

func TestGetNonExistingConversationByID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/conversations/4", nil)
	response := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "4",
	}
	request = mux.SetURLVars(request, vars)

	textChatHandler.GetConversationByID(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got : %d", http.StatusBadRequest, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Product not found") {
		t.Error("Expected response : Product not found")
	}
}

func TestDeleteNonExistantMessage(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/products/4", nil)
	response := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "4",
	}
	request = mux.SetURLVars(request, vars)

	textChatHandler.DeleteMessage(response, request)
	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got : %d", http.StatusNotFound, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Product not found") {
		t.Error("Expected response : Product not found")
	}
}

func TestDeleteNonExistantConversation(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/conversations/4", nil)
	response := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "4",
	}
	request = mux.SetURLVars(request, vars)

	textChatHandler.DeleteMessage(response, request)
	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got : %d", http.StatusNotFound, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Product not found") {
		t.Error("Expected response : Product not found")
	}
}

func TestAddMessage(t *testing.T) {
	// Creating request body
	body := &data.Message{
		UserID:         1,
		ConversationID: 1,
		Text:           "This is a test message",
	}

	request := httptest.NewRequest(http.MethodPost, "/messages", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyMessage{}, body)
	request = request.WithContext(ctx)

	textChatHandler := NewTextChatHandler(NewTestLogger())
	textChatHandler.AddMessage(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestAddConversation(t *testing.T) {
	// Creating request body
	body := &data.Conversation{
		UserID: []int{1, 3},
		GameID: -1,
	}

	request := httptest.NewRequest(http.MethodPost, "/conversations", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyMessage{}, body)
	request = request.WithContext(ctx)

	textChatHandler := NewTextChatHandler(NewTestLogger())
	textChatHandler.AddMessage(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestDeleteExistingMessage(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/messages/1", nil)
	response := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "1",
	}
	request = mux.SetURLVars(request, vars)

	textChatHandler.DeleteMessage(response, request)
	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d but got : %d", http.StatusNoContent, response.Code)
	}
}

func TestDeleteExistingConversation(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/conversations/1", nil)
	response := httptest.NewRecorder()

	textChatHandler := NewTextChatHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "1",
	}
	request = mux.SetURLVars(request, vars)

	textChatHandler.DeleteMessage(response, request)
	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d but got : %d", http.StatusNoContent, response.Code)
	}
}
