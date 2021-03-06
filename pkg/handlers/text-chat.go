package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-text-chat/pkg/database"
	"github.com/gorilla/mux"
)

// KeyMessage is a key used for the Message object inside context
type KeyMessage struct{}

// KeyConversation is a key used for the Conversation object inside context
type KeyConversation struct{}

type TextChatHandler struct {
	db database.TextChatDB
}

func NewTextChatHandler(db database.TextChatDB) *TextChatHandler {
	return &TextChatHandler{db}
}

// getTextChatID extracts the conversation/message ID from the URL
// The verification of this variable is handled by gorilla/mux
func getTextChatID(request *http.Request) string {
	vars := mux.Vars(request)
	id := vars["id"]

	return id
}
