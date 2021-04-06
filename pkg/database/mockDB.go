package database

import (
	"time"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
	"github.com/google/uuid"
)

type MockTextChat struct {
}

func NewMockTextChat() TextChatDB {
	log.Info("Connecting to mock database")
	return &MockTextChat{}
}

func (mp *MockTextChat) Connect() error {
	return nil
}

func (mp *MockTextChat) CloseDB() {
	log.Info("Mocked DB connection closed")
}

func (mp *MockTextChat) GetMessageByID(id string) (*data.Message, error) {
	index := findIndexByMessageID(id)
	if index == -1 {
		return nil, data.ErrorMessageNotFound
	}
	return messageList[index], nil
}

func (mp *MockTextChat) GetMessagesByConversationID(id string) (data.Messages, error) {
	var messages data.Messages
	for _, v := range messageList {
		if v.ConversationID == id {
			messages = append(messages, v)
		}
	}
	if len(messages) <= 0 {
		return nil, data.ErrorMessageNotFound
	}
	return messages, nil
}

func (mp *MockTextChat) GetConversationID(userID []string) string {
	return uuid.NewString()
}

func (mp *MockTextChat) GetConversationByID(id string) (*data.Conversation, error) {
	index := findIndexByConversationID(id)
	if index == -1 {
		return nil, data.ErrorConversationNotFound
	}
	return conversationList[index], nil
}

func (mp *MockTextChat) AddMessage(message *data.Message) error {
	_, err := mp.GetConversationByID(message.ConversationID)
	if err != nil {
		return err
	}

	// TODO: Verify if user exist

	message.ID = uuid.NewString()
	messageList = append(messageList, message)
	return nil
}

func (mp *MockTextChat) AddConversation(conversation *data.Conversation) error {
	// TODO: Verify if all user exists
	// TODO: Veryfy if game exist
	conversation.ID = uuid.NewString()
	conversationList = append(conversationList, conversation)
	return nil
}

func (mp *MockTextChat) DeleteMessage(id string) error {
	index := findIndexByMessageID(id)
	if index == -1 {
		return data.ErrorMessageNotFound
	}

	messageList = append(messageList[:index], messageList[index+1:]...)

	return nil
}

func (mp *MockTextChat) DeleteConversation(id string) error {
	index := findIndexByConversationID(id)
	if index == -1 {
		return data.ErrorConversationNotFound
	}

	// This should not work, probably needs ':' after index+1. To test
	conversationList = append(conversationList[:index], conversationList[index+1])

	return nil
}

// Returns the index of a message in the database
// Returns -1 when no message is found
func findIndexByMessageID(id string) int {
	for index, message := range messageList {
		if message.ID == id {
			return index
		}
	}
	return -1
}

// Returns the index of a conversation in the database
// Returns -1 when no conversation is found
func findIndexByConversationID(id string) int {
	for index, conversation := range conversationList {
		if conversation.ID == id {
			return index
		}
	}
	return -1
}

////////////////////////////////////////////////////////////////////////////////
/////////////////////////// Mocked database ///////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

var messageList = []*data.Message{
	{
		ID:             "a2181017-5c53-422b-b6bc-036b27c04fc8",
		UserID:         "a2181017-5c53-422b-b6bc-036b27c04fc8",
		ConversationID: "a2181017-5c53-422b-b6bc-036b27c04fc8",
		Text:           "This is a message",
		CreatedOn:      time.Now().UTC().String(),
		UpdatedOn:      time.Now().UTC().String(),
	},
	{
		ID:             "e2382ea2-b5fa-4506-aa9d-d338aa52af44",
		UserID:         "e2382ea2-b5fa-4506-aa9d-d338aa52af44",
		ConversationID: "a2181017-5c53-422b-b6bc-036b27c04fc8",
		Text:           "This is an other message",
		CreatedOn:      time.Now().UTC().String(),
		UpdatedOn:      time.Now().UTC().String(),
	},
	{
		ID:             "2aee2975-6b76-4340-b679-e81661b1cdb5",
		UserID:         "e2382ea2-b5fa-4506-aa9d-d338aa52af44",
		ConversationID: "e2382ea2-b5fa-4506-aa9d-d338aa52af44",
		Text:           "This is a third message",
		CreatedOn:      time.Now().UTC().String(),
		UpdatedOn:      time.Now().UTC().String(),
	},
}

var conversationList = []*data.Conversation{
	{
		ID: "a2181017-5c53-422b-b6bc-036b27c04fc8",
		UserID: []string{
			"a2181017-5c53-422b-b6bc-036b27c04fc8",
			"2aee2975-6b76-4340-b679-e81661b1cdb5",
		},
		GameID:    "",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	{
		ID: "e2382ea2-b5fa-4506-aa9d-d338aa52af44",
		UserID: []string{
			"a2181017-5c53-422b-b6bc-036b27c04fc8",
			"2aee2975-6b76-4340-b679-e81661b1cdb5",
			"3a1c152e-f172-41de-a5ab-ca21f6573bf3",
			"c6e6a2b2-bd25-4151-ace1-611accc15a50",
		},
		GameID:    "a2181017-5c53-422b-b6bc-036b27c04fc8",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}
