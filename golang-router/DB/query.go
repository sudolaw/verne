package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Note struct {
	ID       primitive.ObjectID
	Password string
	User     string
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://practice_user2:practiceA438@otomator.com/practice_db"))
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
	passwordResult, err := passwordCollection.Find(ctx, bson.M{"user": "myuser"})

	if err != nil {
		log.Fatal(err)
	}
	var episodesFiltered []bson.M
	if err = passwordResult.All(ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T", episodesFiltered)

	n := Note{}
	for passwordResult.Next(ctx) {
		passwordResult.Decode(&n)

	}

	fmt.Println(n)

}
