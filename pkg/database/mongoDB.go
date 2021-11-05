package database

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Ubivius/microservice-text-chat/pkg/data"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

// ErrorEnvVar : Environment variable error
var ErrorEnvVar = fmt.Errorf("missing environment variable")

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
	uri := mongodbURI()

	// Setting client options
	opts := options.Client()
	clientOptions := opts.ApplyURI(uri)
	opts.Monitor = otelmongo.NewMonitor()

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil || client == nil {
		log.Error(err, "Failed to connect to database. Shutting down service")
		os.Exit(1)
	}

	// Ping DB
	err = client.Ping(context.Background(), nil)
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

func (mp *MongoTextChat) PingDB() error {
	return mp.client.Ping(context.Background(), nil)
}

func (mp *MongoTextChat) CloseDB() {
	err := mp.client.Disconnect(context.Background())
	if err != nil {
		log.Error(err, "Error while disconnecting from database")
	}
}

func (mp *MongoTextChat) GetMessageByID(ctx context.Context, id string) (*data.Message, error) {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Holds search result
	var result data.Message

	// Find a single matching item from the database
	err := mp.messagesCollection.FindOne(ctx, filter).Decode(&result)

	// Parse result into the returned message
	return &result, err
}

func (mp *MongoTextChat) GetConversationByID(ctx context.Context, id string) (*data.Conversation, error) {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Holds search result
	var result data.Conversation

	// Find a single matching item from the database
	err := mp.conversationsCollection.FindOne(ctx, filter).Decode(&result)

	// Parse result into the returned conversation
	return &result, err
}

func (mp *MongoTextChat) GetMessagesByConversationID(ctx context.Context, id string) (data.Messages, error) {
	// MongoDB search filter
	filter := bson.D{{Key: "conversation_id", Value: id}}

	// messages will hold the array of Messages
	var messages data.Messages

	// Find returns a cursor that must be iterated through
	cursor, err := mp.messagesCollection.Find(ctx, filter)
	if err != nil {
		log.Error(err, "Error getting messages by conversationID from database")
	}

	// Iterating through cursor
	for cursor.Next(ctx) {
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
	cursor.Close(ctx)

	return messages, err
}

func (mp *MongoTextChat) AddMessage(ctx context.Context, message *data.Message) error {
	_, err := mp.GetConversationByID(ctx, message.ConversationID)
	if err != nil {
		return err
	}

	if !mp.validateUserExist(message.UserID) {
		return data.ErrorUserNotFound
	}

	message.ID = uuid.NewString()
	// Adding time information to new message
	message.CreatedOn = time.Now().UTC().String()
	message.UpdatedOn = time.Now().UTC().String()

	// Inserting the new message into the database
	insertResult, err := mp.messagesCollection.InsertOne(ctx, message)
	if err != nil {
		return err
	}

	log.Info("Inserting message", "Inserted ID", insertResult.InsertedID)
	return nil
}

func (mp *MongoTextChat) AddConversation(ctx context.Context, conversation *data.Conversation) (*data.Conversation, error) {
	for _, userID := range conversation.UserID {
		if !mp.validateUserExist(userID) {
			return nil, data.ErrorUserNotFound
		}
	}

	if !mp.validateGameExist(conversation.GameID) {
		return nil, data.ErrorGameNotFound
	}

	conversation.ID = uuid.NewString()
	// Adding time information to new conversation
	conversation.CreatedOn = time.Now().UTC().String()
	conversation.UpdatedOn = time.Now().UTC().String()

	// Inserting the new conversation into the database
	insertResult, err := mp.conversationsCollection.InsertOne(ctx, conversation)
	if err != nil {
		return nil, err
	}

	log.Info("Inserting conversation", "Inserted ID", insertResult.InsertedID)

	return conversation, nil
}

func (mp *MongoTextChat) DeleteMessage(ctx context.Context, id string) error {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Delete a single item matching the filter
	result, err := mp.messagesCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Error(err, "Error deleting message")
	}

	log.Info("Deleted documents in messages collection", "delete_count", result.DeletedCount)
	return nil
}

func (mp *MongoTextChat) DeleteConversation(ctx context.Context, id string) error {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Delete a single item matching the filter
	result, err := mp.conversationsCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Error(err, "Error deleting conversation")
	}

	log.Info("Deleted documents in conversations collection", "delete_count", result.DeletedCount)
	return nil
}

func (mp *MongoTextChat) validateUserExist(userID string) bool {
	getUserByIDPath := data.MicroserviceUserPath + "/users/" + userID
	resp, err := http.Get(getUserByIDPath)
	return err == nil && resp.StatusCode == 200
}

func (mp *MongoTextChat) validateGameExist(gameID string) bool {
	//Verify if game exist
	return true
}

func mongodbURI() string {
	hostname := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	if hostname == "" || port == "" || username == "" || password == "" {
		log.Error(ErrorEnvVar, "Some environment variables are not available for the DB connection. DB_HOSTNAME, DB_PORT, DB_USERNAME, DB_PASSWORD")
		os.Exit(1)
	}

	return "mongodb://" + username + ":" + password + "@" + hostname + ":" + port + "/?authSource=admin"
}
