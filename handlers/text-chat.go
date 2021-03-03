package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// KeyMessage is a key used for the Message object inside context
type KeyMessage struct{}
type KeyConversation struct{}

type TextChatHandler struct {
	logger *log.Logger
}

func NewTextChatHandler(logger *log.Logger) *TextChatHandler {
	return &TextChatHandler{logger}
}

// getMessageID extracts the message ID from the URL
// The verification of this variable is handled by gorilla/mux
// We panic if it is not valid because that means gorilla is failing
func getTextChatID(request *http.Request) int {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	return id
}
