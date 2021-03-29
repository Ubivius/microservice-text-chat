package data

import "testing"

func TestChecksValidation(t *testing.T) {
	message := &Message{
		ConversationID: 1,
		UserID:         1,
		Text:           "This is a text message",
	}

	err := message.ValidateMessage()
	if err != nil {
		t.Fatal(err)
	}

	conversation := &Conversation{
		UserID: []int{1, 2},
		GameID: -1,
	}

	err = conversation.ValidateConversation()
	if err != nil {
		t.Fatal(err)
	}
}
