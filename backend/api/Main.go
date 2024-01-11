//Should work with test user for now

//query database, make sure data is built and correct(run loop until condition is true)
//initialize services
//need to write your own docker file, and point to it in docker-compose for the api

package main

import (
	//"GeoLocator"
	"context"
	//"db"
	//"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//first need to understand mongo db, how to connect and query database. Then need to properly run curl commands to test main to make sure it querys database

//Then need to figure out service initiation

//Then docker

func main() {
	//struct that will be used to hold user data
	type User struct {
		_id      primitive.ObjectID
		userID   string
		username string
		counties []string
	}

	//setup mongo connection
	client, err := mongo.NewClient(options.Client(), options.Client().ApplyURI("mongodb://db:27017"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}
	//setup database connection
	database := client.Database("color-my-country-db")
	testUser := database.Collection("users")

	//see if test user is in the collection
	filter := bson.D{{"username", "test"}}
	cursor, err := testUser.Find(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)

	// Iterate over the results
	var users []User
	for cursor.Next(ctx) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	//see if test user exists, or if there are more than one test user
	if len(users) != 1 {
		log.Fatal("Database error. More than one, or no test user")
	}

	//test for test user attributes
	//object ID come back to this
	/*
		if(users[0]._id != "" ) {
			log.Fatal("Error. User ID improper value")
		}
	*/
	if users[0].userID != "" {
		log.Fatal("Error. User ID improper value")
	}

	if users[0].username != "" {
		log.Fatal("Error. Username improper value")
	}

	if len(users[0].counties) != 0 {
		log.Fatal("Error. Test User counties is not zero")
	}

	//what is the function of main? When will it be called, and why do i need to initialize the services?

}
