package database

import (
	"testing"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
)

func TestMongoDBConnectionAndShutdownIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoTextChat()
	if mp == nil {
		t.Fail()
	}
	mp.CloseDB()
}

func TestMongoDBAddMessageIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	message := &data.Message{
		UserID:         "a2181017-5c53-422b-b6bc-036b27c04fc8",
		ConversationID: "a2181017-5c53-422b-b6bc-036b27c04fc8",
		Text:           "testText",
	}

	mp := NewMongoTextChat()
	err := mp.AddMessage(message)
	if err != nil {
		t.Errorf("Failed to add message to database")
	}
	mp.CloseDB()
}

func TestMongoDBAddConversationIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	conversation := &data.Conversation{
		UserID: []string{
			"a2181017-5c53-422b-b6bc-036b27c04fc8",
			"2aee2975-6b76-4340-b679-e81661b1cdb5",
		},
		GameID: "",
	}

	mp := NewMongoTextChat()
	err := mp.AddConversation(conversation)
	if err != nil {
		t.Errorf("Failed to add conversation to database")
	}
	mp.CloseDB()
}

func TestMongoDBGetMessageByIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoTextChat()
	_, err := mp.GetMessageByID("a2181017-5c53-422b-b6bc-036b27c04fc8")
	if err != nil {
		t.Fail()
	}

	mp.CloseDB()
}

func TestMongoDBGetConversationByIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoTextChat()
	_, err := mp.GetConversationByID("a2181017-5c53-422b-b6bc-036b27c04fc8")
	if err != nil {
		t.Fail()
	}

	mp.CloseDB()
}

func TestMongoDBGetMessagesByConversationIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoTextChat()
	_, err := mp.GetMessagesByConversationID("a2181017-5c53-422b-b6bc-036b27c04fc8")
	if err != nil {
		t.Fail()
	}

	mp.CloseDB()
}
