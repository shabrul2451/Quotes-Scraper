package storage

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var quotesDb = "quotesdb"
var quotesColl = "quotes"

type Quote struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Text      string             `bson:"text"`
	Author    string             `bson:"author"`
	CreatedAt time.Time          `bson:"createdAt"`
}

func GetQuotes(client *mongo.Client) ([]Quote, error) {
	collection := client.Database(quotesDb).Collection(quotesColl)

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var quotes []Quote
	for cursor.Next(context.Background()) {
		var quote Quote
		if err := cursor.Decode(&quote); err != nil {
			return nil, err
		}
		quotes = append(quotes, quote)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return quotes, nil
}

func QuoteExists(client *mongo.Client, quote Quote) bool {
	collection := client.Database(quotesDb).Collection(quotesColl)

	filter := bson.M{"text": quote.Text, "author": quote.Author}
	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return false
	}

	return count > 0
}

func StoreQuote(client *mongo.Client, quote Quote) error {
	collection := client.Database(quotesDb).Collection(quotesColl)

	_, err := collection.InsertOne(context.Background(), quote)
	return err
}
