package data

import "testing"

func TestChecksValidation(t *testing.T) {
	message := &Message{
		ConversationID: "a2181017-5c53-422b-b6bc-036b27c04fc8",
		UserID:         "a2181017-5c53-422b-b6bc-036b27c04fc8",
		Text:           "This is a text message",
	}

	err := message.ValidateMessage()
	if err != nil {
		t.Fatal(err)
	}

	conversation := &Conversation{
		UserID: []string{"a2181017-5c53-422b-b6bc-036b27c04fc8", "e2382ea2-b5fa-4506-aa9d-d338aa52af44"},
		GameID: "",
	}

	err = conversation.ValidateConversation()
	if err != nil {
		t.Fatal(err)
	}
}
