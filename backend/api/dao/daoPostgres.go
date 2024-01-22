package dao
/********************************************************************************
* Description: daoPostgres implements the dao interface and is used to interact *
* with the database                                                             *
*********************************************************************************/
import (
	//"Go-directory/dao"
	//GeoLocater "Go-directory/services"
	"context"
	"fmt"

	//"io/ioutil"
	"log"
	//"net/http"
	//"strconv"
	//"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
)

//add counties to a users database
func AddCounitesForUser(counties []string, userId int) {

}
//retrive counties for a user
func GetUserCounites(userID int) {

}
//delete counties for a user
func DeleteCountiesforUser(counties []string, userId int) {

}

// create database connection, and test to see if correct data is in database
func testDatabase(database mongo.Database, ctx context.Context) bool {
	//struct that will be used to hold user data
	type User struct {
		_id      primitive.ObjectID
		userID   string
		username string
		counties []string
	}

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
	//test values for test user

	if users[0].userID != "" {
		log.Fatal("Error. User ID improper value")
	}

	if users[0].username != "" {
		log.Fatal("Error. Username improper value")
	}

	if len(users[0].counties) != 0 {
		log.Fatal("Error. Test User counties is not zero")
	}

	fmt.Printf("Tests were successful! User ID: %s. Username: %s. User counties = %v", users[0].userID, users[0].username, users[0].counties)

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

	return true

}
