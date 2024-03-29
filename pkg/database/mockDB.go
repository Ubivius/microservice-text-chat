package database

import (
	"context"
	"time"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
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

func (mp *MockTextChat) PingDB() error {
	return nil
}

func (mp *MockTextChat) CloseDB() {
	log.Info("Mocked DB connection closed")
}

func (mp *MockTextChat) GetMessageByID(ctx context.Context, id string) (*data.Message, error) {
	_, span := otel.Tracer("text-chat").Start(ctx, "getMessageByIdDatabase")
	defer span.End()
	index := findIndexByMessageID(id)
	if index == -1 {
		return nil, data.ErrorMessageNotFound
	}
	return messageList[index], nil
}

func (mp *MockTextChat) GetMessagesByConversationID(ctx context.Context, id string) (data.Messages, error) {
	_, span := otel.Tracer("text-chat").Start(ctx, "getMessagesByConversationIdDatabase")
	defer span.End()
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

func (mp *MockTextChat) GetConversationByID(ctx context.Context, id string) (*data.Conversation, error) {
	_, span := otel.Tracer("text-chat").Start(ctx, "getConversationByIdDatabase")
	defer span.End()
	index := findIndexByConversationID(id)
	if index == -1 {
		return nil, data.ErrorConversationNotFound
	}
	return conversationList[index], nil
}

func (mp *MockTextChat) AddMessage(ctx context.Context, message *data.Message) error {
	_, span := otel.Tracer("text-chat").Start(ctx, "addMessageDatabase")
	defer span.End()
	_, err := mp.GetConversationByID(ctx, message.ConversationID)
	if err != nil {
		return err
	}

	if !mp.validateUserExist(message.UserID) {
		return data.ErrorUserNotFound
	}

	message.ID = uuid.NewString()
	messageList = append(messageList, message)
	return nil
}

func (mp *MockTextChat) AddConversation(ctx context.Context, conversation *data.Conversation) (*data.Conversation, error) {
	_, span := otel.Tracer("text-chat").Start(ctx, "addConversationDatabase")
	defer span.End()
	for _, userID := range conversation.UserID {
		if !mp.validateUserExist(userID) {
			return nil, data.ErrorUserNotFound
		}
	}

	if !mp.validateGameExist(conversation.GameID) {
		return nil, data.ErrorGameNotFound
	}

	conversation.ID = uuid.NewString()
	conversationList = append(conversationList, conversation)
	return conversation, nil
}

func (mp *MockTextChat) DeleteMessage(ctx context.Context, id string) error {
	_, span := otel.Tracer("text-chat").Start(ctx, "deleteMessageDatabase")
	defer span.End()
	index := findIndexByMessageID(id)
	if index == -1 {
		return data.ErrorMessageNotFound
	}

	messageList = append(messageList[:index], messageList[index+1:]...)

	return nil
}

func (mp *MockTextChat) DeleteConversation(ctx context.Context, id string) error {
	_, span := otel.Tracer("text-chat").Start(ctx, "deleteConversationDatabase")
	defer span.End()
	index := findIndexByConversationID(id)
	if index == -1 {
		return data.ErrorConversationNotFound
	}

	// This should not work, probably needs ':' after index+1. To test
	conversationList = append(conversationList[:index], conversationList[index+1])

	return nil
}

func (mp *MockTextChat) AddUserToConversation(ctx context.Context, conversation *data.Conversation) error {
	_, span := otel.Tracer("text-chat").Start(ctx, "addUserToConversationTextChat")
	defer span.End()
	for _, userID := range conversation.UserID {
		if !mp.validateUserExist(userID) {
			return data.ErrorUserNotFound
		}
	}

	if !mp.validateGameExist(conversation.GameID) {
		return data.ErrorGameNotFound
	}

	conversationIndex := findIndexByConversationID(conversation.ID)
	conversationToUpdate := conversationList[conversationIndex]
	conversationToUpdate.UpdatedOn = time.Now().UTC().String()
	conversationToUpdate.UserID = make([]string, len(conversation.UserID))
	_ = copy(conversationToUpdate.UserID, conversation.UserID)
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

func (mp *MockTextChat) validateUserExist(userID string) bool {
	return true
}

func (mp *MockTextChat) validateGameExist(gameID string) bool {
	return true
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
