package main

import (
	"REST_API/Controller"
	utility "REST_API/Utility"
	"context"
	"time"

	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate swagger generate spec -m -o ./swagger.json

func main() {

	fmt.Println("Go  Tutorial")
	ctx, canc := context.WithTimeout(context.Background(), 10*time.Second)
	defer canc()
	clientOptions := options.Client().ApplyURI("mongodb://db-service:27017")
	utility.Client, _ = mongo.Connect(ctx, clientOptions)

	fmt.Println("connected")
	// Handle Subsequent requests

	// defer utility.Client.Disconnect(ctx)
	Controller.HandleRequests()

}
