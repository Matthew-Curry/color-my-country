package Tests

import (
	"fmt"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/bson"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo"

	"Go-directory/dao"
	"context"
)

// Test to see if initial test data is contained within database
func TestDatabase(*testing.T) bool {
	// Initialize database. Should be accessible from the API variable declare in TestInit
	database := API.DB

	ctx := context.Background()

	testUser := database.Collection("users")

	//see if test user is in the collection
	filter := bson.D{{"username", "test"}}
	cursor, err := testUser.Find(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)

	// Iterate over the results
	var users []dao.User
	for cursor.Next(ctx) {
		var user dao.User
		if err := cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	//see if test user exists, or if there are more than one test user
	if len(users) != 1 {
		log.Fatal("Database error. More than one, or no test user")
	}

	if users[0].UserID != 123 {
		log.Fatal("Error. User ID improper value")
	}

	if users[0].Username != "test" {
		log.Fatal("Error. Username improper value")
	}

	if len(users[0].Counties) == 0 {
		log.Fatal("Error. Test User counties is not zero")

	}
	fmt.Printf("Tests were successful! User ID: %s. Username: %s. User counties = %v", users[0].UserID, users[0].Username, users[0].Counties)

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
