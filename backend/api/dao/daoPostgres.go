package dao

/********************************************************************************
* Description: daoPostgres implements the dao interface and is used to interact *
* with the database                                                             *
*********************************************************************************/
import (
	//"Go-directory/dao"
	//GeoLocater "Go-directory/services"
	"context"
	"encoding/json"

	"fmt"

	//"io/ioutil"
	"log"
	//"net/http"
	//"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`      // Maps to the _id field, use primitive.ObjectID for ObjectId types
	UserID   int                `bson:"userID"`   // Maps to the userID field
	Username string             `bson:"username"` // Maps to the username field
	Counties []string           `bson:"counties"` // Maps to the counties field, assuming it's an array of strings
}

type GeoFeature struct {
	ID         string `json:"_id"`
	Type       string `json:"type"`
	Properties struct {
		GEOID      string  `json:"GEO_ID"`
		State      string  `json:"STATE"`
		County     string  `json:"COUNTY"`
		Name       string  `json:"NAME"`
		LSAD       string  `json:"LSAD"`
		CensusArea float64 `json:"CENSUSAREA"`
	} `json:"properties"`
}

// add counties to a users database
func AddCounitesForUser(counties []string, username string, database mongo.Database) {
	//Get the users collection
	collection := database.Collection("users")
	//find the user
	filter := bson.D{{"username", username}}
	//define update operations(append to county array multiple counties)
	update := bson.M{
		"$push": bson.M{
			"counties": bson.M{"$each": counties},
		},
	}

	ctx := context.Background()

	//execute the update
	updateResult, err := collection.UpdateOne(ctx, filter, update)
	//log the error if there is one
	if err != nil {
		log.Fatal(err)
	}
	//print whether or not any of the array elements were modified
	if updateResult.ModifiedCount == 0 {
		fmt.Println("No documents matched the filter.")
		return
	}

	fmt.Printf("Updated %v document(s)\n", updateResult.ModifiedCount)

}

// retrive counties for a user
// return json from database

func GetUserCounites(userID string, database mongo.Database) []byte {

	collection := database.Collection("users")
	// Define a filter to find the user by username
	filter := bson.D{{"username", userID}}

	// Create an instance of the User struct to store the result
	var user User

	ctx := context.Background()
	// Perform the query
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		fmt.Println("User not found.")
		return nil
	} else if err != nil {
		log.Fatal(err)
	}

	// Access the array of counties for the user
	counties := user.Counties
	//convert counties to a json
	JSON, err := json.Marshal(counties)

	//return the user counties
	return JSON

}

// delete counties for a user
func DeleteCountiesforUser(counties []string, username string, database mongo.Database) {
	// Get the users collection
	collection := database.Collection("users")

	// Find the user
	filter := bson.D{{"username", username}}

	// Define update operations (remove multiple counties from array)
	update := bson.M{
		"$pullAll": bson.M{
			"counties": counties,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	// Execute the update
	updateResult, err := collection.UpdateOne(ctx, filter, update)

	// Log the error if there is one
	if err != nil {
		log.Fatal(err)
	}

	// Print whether or not any of the array elements were modified
	if updateResult.ModifiedCount == 0 {
		fmt.Println("No documents matched the filter.")
		return
	}

	fmt.Printf("Updated %v document(s)\n", updateResult.ModifiedCount)

}
