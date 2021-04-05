package data

import (
	"fmt"
	"time"
)

// ErrorMessageNotFound : Message specific errors
var ErrorMessageNotFound = fmt.Errorf("Message not found")

// Message defines the structure for an API message.
// Formatting done with json tags to the right. "-" : don't include when encoding to json
type Message struct {
	ID             int    `json:"id"`
	UserID         int    `json:"userid" validate:"required"`
	ConversationID int    `json:"conversationid" validate:"required"`
	Text           string `json:"text" validate:"required"`
	CreatedOn      string `json:"-"`
	UpdatedOn      string `json:"-"`
	DeletedOn      string `json:"-"`
}

// Messages is a collection of Message
type Messages []*Message

// All of these functions will become database calls in the future

// GetMessageByID returns a single message with the given id
func GetMessageByID(id int) (*Message, error) {
	index := findIndexByMessageID(id)
	if index == -1 {
		return nil, ErrorMessageNotFound
	}
	return messageList[index], nil
}

// GetMessageByConversationID returns an array of messages corresponding to the conversation
func GetMessagesByConversationID(id int) ([]*Message, error) {
	messages := []*Message{}
	for _, v := range messageList {
		if v.ConversationID == id {
			messages = append(messages, v)
		}
	}
	if len(messages) <= 0 {
		return nil, ErrorMessageNotFound
	}
	return messages, nil
}

// AddMessage creates a new message
func AddMessage(message *Message) error {
	_, err := GetConversationByID(message.ConversationID)
	if err != nil {
		return err
	}

	// TODO: Verify if user exist

	message.ID = getNextMessageID()
	messageList = append(messageList, message)
	return nil
}

// DeleteMessage deletes the message with the given id
func DeleteMessage(id int) error {
	index := findIndexByMessageID(id)
	if index == -1 {
		return ErrorMessageNotFound
	}

	// This should not work, probably needs ':' after index+1. To test
	messageList = append(messageList[:index], messageList[index+1])

	return nil
}

// Returns the index of a message in the database
// Returns -1 when no message is found
func findIndexByMessageID(id int) int {
	for index, message := range messageList {
		if message.ID == id {
			return index
		}
	}
	return -1
}

//////////////////////////////////////////////////////////////////////////////
/////////////////////////// Fake database ///////////////////////////////////
///// DB connection setup and docker file will be done in sprint 8 /////////
///////////////////////////////////////////////////////////////////////////

// Finds the maximum index of our fake database and adds 1
func getNextMessageID() int {
	lastMessage := messageList[len(messageList)-1]
	return lastMessage.ID + 1
}

// messageList is a hard coded list of messages for this
// example data source. Should be replaced by database connection
var messageList = []*Message{
	{
		ID:             1,
		UserID:         1,
		ConversationID: 1,
		Text:           "This is a message",
		CreatedOn:      time.Now().UTC().String(),
		UpdatedOn:      time.Now().UTC().String(),
	},
	{
		ID:             2,
		UserID:         2,
		ConversationID: 1,
		Text:           "This is an other message",
		CreatedOn:      time.Now().UTC().String(),
		UpdatedOn:      time.Now().UTC().String(),
	},
	{
		ID:             3,
		UserID:         2,
		ConversationID: 2,
		Text:           "This is a third message",
		CreatedOn:      time.Now().UTC().String(),
		UpdatedOn:      time.Now().UTC().String(),
	},
}