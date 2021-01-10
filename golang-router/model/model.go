package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Note lol
type Note struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	MyPassword string             `json:"mypassword" bson:"mypassword,omitempty"`
	User       string             `json:"user" bson:"user,omitempty"`
}

//Pasal Controller is for auth logic Pasal
type Pasal struct{}

//Pass Login is to Pass login request Pasal
func (lol *Pasal) Pass(j string) Note {

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
	passwordResult, err := passwordCollection.Find(ctx, bson.M{"user": j})

	if err != nil {
		log.Fatal(err)
	}
	var book Note
	for passwordResult.Next(ctx) {

		// create a value into which the single document can be decoded

		// & character returns the memory address of the following variable.
		err := passwordResult.Decode(&book) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

	}
	return book
}
