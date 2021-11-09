package controller

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://dample:madhav@cluster0.xbc9d.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connect with mongodb

func inti() {
	// client options
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	mongo.Connect(context.TODO(), clientOption)
}