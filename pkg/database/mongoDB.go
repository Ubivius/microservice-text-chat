package database

import (
	"context"
	"log"
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
	logger                  *log.Logger
}

func NewMongoTextChat(l *log.Logger) TextChatDB {
	mp := &MongoTextChat{logger: l}
	err := mp.Connect()
	// If connect fails, kill the program
	if err != nil {
		mp.logger.Fatal(err)
	}
	return mp
}

func (mp *MongoTextChat) Connect() error {
	// Setting client options
	clientOptions := options.Client().ApplyURI("mongodb://admin:pass@localhost:27888/?authSource=admin")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil || client == nil {
		mp.logger.Fatalln("Failed to connect to database. Shutting down service")
	}

	// Ping DB
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		mp.logger.Fatal(err)
	}

	log.Println("Connection to MongoDB established")

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
		mp.logger.Println(err)
	} else {
		log.Println("Connection to MongoDB closed.")
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
		log.Fatal(err)
	}

	// Iterating through cursor
	for cursor.Next(context.TODO()) {
		var result data.Message
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		messages = append(messages, &result)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cursor.Close(context.TODO())

	return messages, err
}

func (mp *MongoTextChat) GetConversationID(userID []string) string {
	return uuid.NewString()
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

	log.Println("Inserting a document: ", insertResult.InsertedID)
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

	log.Println("Inserting a document: ", insertResult.InsertedID)
	return nil
}

func (mp *MongoTextChat) DeleteMessage(id string) error {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Delete a single item matching the filter
	result, err := mp.messagesCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Deleted %v documents in the messages collection\n", result.DeletedCount)
	return nil
}

func (mp *MongoTextChat) DeleteConversation(id string) error {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Delete a single item matching the filter
	result, err := mp.conversationsCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Deleted %v documents in the conversations collection\n", result.DeletedCount)
	return nil
}
