package controller

import (
	"Go-directory/dao"
	GeoLocater "Go-directory/services"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//"fmt"

//"strconv"
//"time"

// create database connection, and test to see if correct data is in database
func ConnectToDatabase() {
	//struct that will be used to hold user data
	type User struct {
		_id      primitive.ObjectID
		userID   string
		username string
		counties []string
	}

	//connect to the database container
	dataBaseContainerURI := "mongodb://backend-db-1:27017"

	//setup mongo connection through database container
	client, err := mongo.NewClient(options.Client(), options.Client().ApplyURI(dataBaseContainerURI))
	//if the connection cannot be established, log the error
	if err != nil {
		log.Fatal(err)
	}
	//give the mongo driver a defined timeout
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//connect to the database
	err = client.Connect(ctx)
	//if the connection cannot be established, log the error
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

	fmt.Println("Tests were successful!")

	//test counties
	//get county collection
	counties := database.Collection("counties")

	//get county count
	count, err := counties.CountDocuments(context.Background(), bson.D{})
	//print error(if there is one)
	if err != nil {
		log.Fatal("Error get county data")
	}
	//print county count
	fmt.Printf("%d Counties exist in the database", count)

}

// will use GeoLocater package to get counties for a user given their google maps JSON, and then call the database to add these counties to the users list, if they do not already exist
func GetListOfUserCounties(w http.ResponseWriter, r *http.Request) {

	// Parse query parameters from the request
	queryParams := r.URL.Query()

	// Access specific query parameters by name
	userId := queryParams.Get("userid")
	//convert userId to int
	ID, err := strconv.Atoi(userId)

	//gets user json from http request body
	UserJson, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Print("error")
	}
	//gives user json to geolocater, which returns a list of the counties within json
	counties := GeoLocater.GetCounties(UserJson)
	//calls database method to add counties to user object (if they dont already exist)
	dao.AddCounitesForUser(counties, ID)

}

func DeleteUserCounties(w http.ResponseWriter, r *http.Request, counties []string) {
	// Parse query parameters from the request
	queryParams := r.URL.Query()

	// Access specific query parameters by name
	userId := queryParams.Get("userid")
	//convert userId to int
	ID, err := strconv.Atoi(userId)
	//handle error
	if err != nil {
		fmt.Print("error")
	}
	dao.DeleteCountiesforUser(counties, ID)
}

// will county be a string? Int? Generic type? (for now string)
func Addcounties(w http.ResponseWriter, r *http.Request, counties []string) {
	// Parse query parameters from the request
	queryParams := r.URL.Query()

	// Access specific query parameters by name
	userId := queryParams.Get("userid")
	//convert userId to int
	ID, err := strconv.Atoi(userId)
	//handle error
	if err != nil {
		fmt.Print("error")
	}
	dao.AddCounitesForUser(counties, ID)

}
