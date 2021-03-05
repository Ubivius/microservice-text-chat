package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
)

func (textChatHandler *TextChatHandler) GetMessages(responseWriter http.ResponseWriter, request *http.Request) {
	textChatHandler.logger.Println("Handle GET messages")
	messageList := data.GetMessages()
	err := json.NewEncoder(responseWriter).Encode(messageList)
	if err != nil {
		textChatHandler.logger.Println("[ERROR] serializing message", err)
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (textChatHandler *TextChatHandler) GetConversations(responseWriter http.ResponseWriter, request *http.Request) {
	textChatHandler.logger.Println("Handle GET conversation")
	conversationList := data.GetMessages()
	err := json.NewEncoder(responseWriter).Encode(conversationList)
	if err != nil {
		textChatHandler.logger.Println("[ERROR] serializing conversation", err)
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (textChatHandler *TextChatHandler) GetMessageByID(responseWriter http.ResponseWriter, request *http.Request) {
	id := getTextChatID(request)

	textChatHandler.logger.Println("[DEBUG] getting id", id)

	message, err := data.GetMessageByID(id)
	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(message)
		if err != nil {
			textChatHandler.logger.Println("[ERROR] serializing message", err)
		}
	case data.ErrorMessageNotFound:
		textChatHandler.logger.Println("[ERROR] fetching message", err)
		http.Error(responseWriter, "Message not found", http.StatusBadRequest)
		return
	default:
		textChatHandler.logger.Println("[ERROR] fetching message", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(responseWriter).Encode(message)
	if err != nil {
		textChatHandler.logger.Println("[ERROR] serializing message", err)
	}
}

func (textChatHandler *TextChatHandler) GetConversationByID(responseWriter http.ResponseWriter, request *http.Request) {
	id := getTextChatID(request)

	textChatHandler.logger.Println("[DEBUG] getting id", id)

	conversation, err := data.GetConversationByID(id)
	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(conversation)
		if err != nil {
			textChatHandler.logger.Println("[ERROR] serializing conversation", err)
		}
	case data.ErrorConversationNotFound:
		textChatHandler.logger.Println("[ERROR] fetching conversation", err)
		http.Error(responseWriter, "Conversation not found", http.StatusBadRequest)
		return
	default:
		textChatHandler.logger.Println("[ERROR] fetching conversation", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(responseWriter).Encode(conversation)
	if err != nil {
		textChatHandler.logger.Println("[ERROR] serializing conversation", err)
	}
}

func (textChatHandler *TextChatHandler) GetMessagesByConversationID(responseWriter http.ResponseWriter, request *http.Request) {
	id := getTextChatID(request)

	textChatHandler.logger.Println("[DEBUG] getting id", id)

	messages, err := data.GetMessagesByConversationID(id)
	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(messages)
		if err != nil {
			textChatHandler.logger.Println("[ERROR] serializing messages", err)
		}
	case data.ErrorMessageNotFound:
		textChatHandler.logger.Println("[ERROR] fetching messages", err)
		http.Error(responseWriter, "Message not found", http.StatusBadRequest)
		return
	default:
		textChatHandler.logger.Println("[ERROR] fetching messages", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(responseWriter).Encode(messages)
	if err != nil {
		textChatHandler.logger.Println("[ERROR] serializing messages", err)
	}
}
