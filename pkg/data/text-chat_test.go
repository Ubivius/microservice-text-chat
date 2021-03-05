package data

import "testing"

func TestChecksValidation(t *testing.T) {
	message := &Message{
		Text: "This is a text message",
	}

	err := message.ValidateMessage()
	if err != nil {
		t.Fatal(err)
	}

	conversation := &Conversation{}

	err = conversation.ValidateConversation()
	if err != nil {
		t.Fatal(err)
	}
}
