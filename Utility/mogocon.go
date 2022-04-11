package utility

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://db-service:27017")

	client, _ := mongo.Connect(context.TODO(), clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return client
	// }

	//fmt.Println("connection succesfully!", client)
	return client
}

var Client *mongo.Client

// Collection student

func NewFunction() *mongo.Collection {
	// Connection()
	collection := Client.Database("Universe").Collection("studentdata")
	return collection
}
