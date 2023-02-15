package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	errorHandler "todobackend2/utils"
)



func HandleTodoDB(URI string) *mongo.Collection {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	errorHandler.HandleError(err)
	collection := client.Database("learnfiber").Collection("todos")
 
	return collection



}


func HandleUserDB(URI string) *mongo.Collection {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	errorHandler.HandleError(err)
	collection := client.Database("learnfiber").Collection("users")

	return collection

}






