package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/budhirajamadhav/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://dample:<password>@cluster0.xbc9d.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connect with mongodb

func init() {
	// client options
	clientOption := options.Client().ApplyURI(connectionString)
				
	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = (*mongo.Collection)(client.Database(dbName).Collection(colName))

	// collection instance
	fmt.Println("Collection instance is ready")

}

// MONGODB helpers - file

// inset 1 record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with id:", inserted.InsertedID)

}


// update 1 record
func updateOneMovie(movieID string) {
			// converts string to _id
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id":id}
	update := bson.M{"$set":bson.M{"watched":true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count:", result.ModifiedCount)

}