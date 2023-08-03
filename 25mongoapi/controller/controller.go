package controller

import (
	"MongoAPI/model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: change password to <password>
const connectionString = "mongodb+srv://doopstech:doops@cluster0.6vl9klf.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Netflix"
const collectionName = "Watchlist"

// MOST IMPORTANT: reference of mongodb collection
var collection *mongo.Collection

// Connect to MongoDB
func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Collection instance created")
}

// MONGODG Helper functions

// insert one document into mongodb
func InsertOneMovie(movie model.Netflix) {
	insertResult, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

// update one document in mongodb
func UpdateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Matched ", updateResult.MatchedCount, " documents and updated ", updateResult.ModifiedCount, " documents.")
}
