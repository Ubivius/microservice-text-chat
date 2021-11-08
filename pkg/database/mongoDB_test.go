package database

import (
	"context"
	"os"
	"testing"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
)

func integrationTestSetup(t *testing.T) {
	t.Log("Test setup")

	if os.Getenv("DB_USERNAME") == "" {
		os.Setenv("DB_USERNAME", "admin")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		os.Setenv("DB_PASSWORD", "pass")
	}
	if os.Getenv("DB_PORT") == "" {
		os.Setenv("DB_PORT", "27888")
	}
	if os.Getenv("DB_HOSTNAME") == "" {
		os.Setenv("DB_HOSTNAME", "localhost")
	}

	err1, err2 := deleteAllConversationsAndMessagesFromMongoDB()
	if err1 != nil || err2 != nil {
		t.Errorf("Failed to delete existing items from database during setup")
	}
}

func TestMongoDBConnectionAndShutdownIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

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
	integrationTestSetup(t)

	message := &data.Message{
		UserID:         "a2181017-5c53-422b-b6bc-036b27c04fc8",
		ConversationID: "a2181017-5c53-422b-b6bc-036b27c04fc8",
		Text:           "testText",
	}

	mp := NewMongoTextChat()
	err := mp.AddMessage(context.Background(), message)
	if err != nil {
		t.Errorf("Failed to add message to database")
	}

	messages, err := mp.GetMessagesByConversationID(context.Background(), message.ConversationID)
	if err != nil {
		t.Error("Failed to retrieve messages with error : " + err.Error())
	}
	if messages == nil {
		t.Error("messages slice is nil")
	}
	if len(messages) != 1 {
		t.Errorf("Incorrect number of messages returned. Message count : %d", len(messages))
	}
	if messages != nil && messages[0].Text != message.Text {
		t.Errorf("Incorrect message returned, expected %s but received %s", message.Text, messages[0].Text)
	}
	mp.CloseDB()
}

func TestMongoDBAddConversationIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	conversation := &data.Conversation{
		UserID: []string{
			"a2181017-5c53-422b-b6bc-036b27c04fc8",
			"2aee2975-6b76-4340-b679-e81661b1cdb5",
		},
		GameID: "",
	}

	mp := NewMongoTextChat()
	_, err := mp.AddConversation(context.Background(), conversation)
	if err != nil {
		t.Errorf("Failed to add conversation to database")
	}
	mp.CloseDB()
}

func TestMongoDBGetMessageByIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	mp := NewMongoTextChat()
	_, err := mp.GetMessageByID(context.Background(), "a2181017-5c53-422b-b6bc-036b27c04fc8")
	if err != nil {
		t.Error("Error getting message from database")
	}

	mp.CloseDB()
}

func TestMongoDBGetConversationByIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	mp := NewMongoTextChat()
	_, err := mp.GetConversationByID(context.Background(), "a2181017-5c53-422b-b6bc-036b27c04fc8")
	if err != nil {
		t.Fail()
	}

	mp.CloseDB()
}

func TestMongoDBGetMessagesByConversationIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	mp := NewMongoTextChat()
	_, err := mp.GetMessagesByConversationID(context.Background(), "a2181017-5c53-422b-b6bc-036b27c04fc8")
	if err != nil {
		t.Fail()
	}

	mp.CloseDB()
}
