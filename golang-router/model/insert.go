package model

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://practice_user2:practiceA438@<domain>.com/practice_db"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database("practice_db")
	passwordCollection := quickstartDatabase.Collection("passwords")
	//episodeCollection:=quickstartDatabase.Collection("episodes")
	passwordResult, err := passwordCollection.InsertOne(ctx, bson.D{

		{Key: "user", Value: "myuser"},
		{Key: "mypassword", Value: "mypassword"},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(passwordResult.InsertedID)
}
