package data

import (
	"fmt"
	"time"
)

// ErrorConversationNotFound : Conversation specific errors
var ErrorConversationNotFound = fmt.Errorf("Conversation not found")

// Conversation defines the structure for an API conversation.
// Formatting done with json tags to the right. "-" : don't include when encoding to json
type Conversation struct {
	ID        int    `json:"id"`
	UserID    []int  `json:"userid"`
	GameID    int    `json:"gameid"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
}

// Conversations is a collection of Conversation
type Conversations []*Conversation

// All of these functions will become database calls in the future
// GETTING PRODUCTS

// GetConversations returns the list of conversations
func GetConversations() Conversations {
	return conversationList
}

// GetConversationByID returns a single conversation with the given id
func GetConversationByID(id int) (*Conversation, error) {
	index := findIndexByConversationID(id)
	if id == -1 {
		return nil, ErrorConversationNotFound
	}
	return conversationList[index], nil
}

// UPDATING PRODUCTS

// UpdateConversation updates the conversation specified in received JSON
func UpdateConversation(conversation *Conversation) error {
	index := findIndexByConversationID(conversation.ID)
	if index == -1 {
		return ErrorConversationNotFound
	}
	conversationList[index] = conversation
	return nil
}

// AddConversation creates a new conversation
func AddConversation(conversation *Conversation) {
	conversation.ID = getNextID()
	conversationList = append(conversationList, conversation)
}

// DeleteConversation deletes the conversation with the given id
func DeleteConversation(id int) error {
	index := findIndexByConversationID(id)
	if index == -1 {
		return ErrorConversationNotFound
	}

	// This should not work, probably needs ':' after index+1. To test
	conversationList = append(conversationList[:index], conversationList[index+1])

	return nil
}

// Returns the index of a conversation in the database
// Returns -1 when no conversation is found
func findIndexByConversationID(id int) int {
	for index, conversation := range conversationList {
		if conversation.ID == id {
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
func getNextID() int {
	lastConversation := conversationList[len(conversationList)-1]
	return lastConversation.ID + 1
}

// conversationList is a hard coded list of conversations for this
// example data source. Should be replaced by database connection
var conversationList = []*Conversation{
	{
		ID:        1,
		UserID:    []int{1, 3},
		GameID:    -1,
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	{
		ID:        2,
		UserID:    []int{1, 3, 4, 5},
		GameID:    1,
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}
