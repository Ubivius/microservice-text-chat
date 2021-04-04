package database

import (
	"context"
	"os"
	"time"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoTextChat struct {
	client                  *mongo.Client
	messagesCollection      *mongo.Collection
	conversationsCollection *mongo.Collection
}

func NewMongoTextChat() TextChatDB {
	mp := &MongoTextChat{}
	err := mp.Connect()
	// If connect fails, kill the program
	if err != nil {
		log.Error(err, "MongoDB setup failed")
		os.Exit(1)
	}
	return mp
}

func (mp *MongoTextChat) Connect() error {
	// Setting client options
	clientOptions := options.Client().ApplyURI("mongodb://admin:pass@localhost:27888/?authSource=admin")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil || client == nil {
		log.Error(err, "Failed to connect to database. Shutting down service")
		os.Exit(1)
	}

	// Ping DB
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Error(err, "Failed to ping database. Shutting down service")
		os.Exit(1)
	}

	log.Info("Connection to MongoDB established")

	messagesCollection := client.Database("ubivius").Collection("messages")
	conversationsCollection := client.Database("ubivius").Collection("conversations")

	// Assign client and collection to the MongoTextChat struct
	mp.messagesCollection = messagesCollection
	mp.conversationsCollection = conversationsCollection
	mp.client = client
	return nil
}

func (mp *MongoTextChat) CloseDB() {
	err := mp.client.Disconnect(context.TODO())
	if err != nil {
		log.Error(err, "Error while disconnecting from database")
	}
}

func (mp *MongoTextChat) GetMessageByID(id string) (*data.Message, error) {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Holds search result
	var result data.Message

	// Find a single matching item from the database
	err := mp.messagesCollection.FindOne(context.TODO(), filter).Decode(&result)

	// Parse result into the returned message
	return &result, err
}

func (mp *MongoTextChat) GetConversationByID(id string) (*data.Conversation, error) {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Holds search result
	var result data.Conversation

	// Find a single matching item from the database
	err := mp.conversationsCollection.FindOne(context.TODO(), filter).Decode(&result)

	// Parse result into the returned conversation
	return &result, err
}

func (mp *MongoTextChat) GetMessagesByConversationID(id string) (data.Messages, error) {
	// MongoDB search filter
	filter := bson.D{{Key: "conversation_id", Value: id}}

	// messages will hold the array of Messages
	var messages data.Messages

	// Find returns a cursor that must be iterated through
	cursor, err := mp.messagesCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Error(err, "Error getting messages by conversationID from database")
	}

	// Iterating through cursor
	for cursor.Next(context.TODO()) {
		var result data.Message
		err := cursor.Decode(&result)
		if err != nil {
			log.Error(err, "Error decoding messages from database")
		}
		messages = append(messages, &result)
	}

	if err := cursor.Err(); err != nil {
		log.Error(err, "Error in cursor after iteration")
	}

	// Close the cursor once finished
	cursor.Close(context.TODO())

	return messages, err
}

func (mp *MongoTextChat) AddMessage(message *data.Message) error {
	_, err := mp.GetConversationByID(message.ConversationID)
	if err != nil {
		return err
	}

	// TODO: Verify if user exist

	message.ID = uuid.NewString()
	// Adding time information to new message
	message.CreatedOn = time.Now().UTC().String()
	message.UpdatedOn = time.Now().UTC().String()

	// Inserting the new message into the database
	insertResult, err := mp.messagesCollection.InsertOne(context.TODO(), message)
	if err != nil {
		return err
	}

	log.Info("Inserting message", "Inserted ID", insertResult.InsertedID)
	return nil
}

func (mp *MongoTextChat) AddConversation(conversation *data.Conversation) error {
	// TODO: Verify if all user exists
	// TODO: Veryfy if game exist

	conversation.ID = uuid.NewString()
	// Adding time information to new conversation
	conversation.CreatedOn = time.Now().UTC().String()
	conversation.UpdatedOn = time.Now().UTC().String()

	// Inserting the new conversation into the database
	insertResult, err := mp.conversationsCollection.InsertOne(context.TODO(), conversation)
	if err != nil {
		return err
	}

	log.Info("Inserting conversation", "Inserted ID", insertResult.InsertedID)
	return nil
}

func (mp *MongoTextChat) DeleteMessage(id string) error {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Delete a single item matching the filter
	result, err := mp.messagesCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Error(err, "Error deleting message")
	}

	log.Info("Deleted documents in messages collection", "delete_count", result.DeletedCount)
	return nil
}

func (mp *MongoTextChat) DeleteConversation(id string) error {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Delete a single item matching the filter
	result, err := mp.conversationsCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Error(err, "Error deleting conversation")
	}

	log.Info("Deleted documents in conversations collection", "delete_count", result.DeletedCount)
	return nil
}
