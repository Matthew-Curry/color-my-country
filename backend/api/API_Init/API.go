package api

/********************************************************************************
* Description: the api package contains all neccessary functions and utilities  *
* that the api for color-your-country will need. It contains a logger, a        *
* database connectionand a router. It will be initialized as an object and used *
* as such.                                                                      *
*********************************************************************************/
import (
	"Go-directory/controller"
	"fmt"
	"net/http"

	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"Go-directory/dao"
	//GeoLocater "Go-directory/services"
	"context"
	"log"
	"time"

	//"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// API structs
type userCounties struct {
}

type API struct {
	DB     *mongo.Database
	Logger *log.Logger
	//services
	//Geolocator *GeoLocater

}

// Initialize and return a database connection
func NewDatabase() (*mongo.Database, error, context.Context) {
	//connect to the database container
	dataBaseContainerURI := "mongodb://backend-db-1:27017"

	//setup mongo connection through database container
	client, err := mongo.NewClient(options.Client(), options.Client().ApplyURI(dataBaseContainerURI))
	//if the connection cannot be established, log the error
	if err != nil {
		return nil, fmt.Errorf("Can't establish connection to database"), nil
	}
	//give the mongo driver a defined timeout
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//connect to the database
	err = client.Connect(ctx)
	//if the connection cannot be established, log the error
	if err != nil {
		return nil, fmt.Errorf("Can't establish connection to database"), nil
	}
	//disconnect if the connection cannot be established
	defer client.Disconnect(ctx)
	//ping the database to see if connection is established
	err = client.Ping(ctx, readpref.Primary())
	//if there is an error, print error
	if err != nil {
		return nil, fmt.Errorf("Can't establish connection to database"), nil
	}
	//setup database connection
	database := client.Database("color-my-country-db")
	//return the database connection
	return database, nil, ctx
}

// Initialize and return a logger
func NewLogger() (*log.Logger, error) {
	return &log.Logger{}, nil

}

// initialize a new api
func NewAPI() (*API, error) {
	//create a database connection
	db, err, ctx := NewDatabase()
	if err != nil {
		return nil, err
	}

	//so ctx error goes away
	fmt.Print(ctx)

	//create a logger
	logger, err := NewLogger()
	if err != nil {
		return nil, err
	}
	//create an api from the API struct
	api := &API{
		DB:     db,
		Logger: logger,
		// Initialize other services

	}
	//return the api
	return api, nil
}

// create routes and set handler functions for each route
func (api *API) SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Define your routes using the router and API methods
	router.HandleFunc("/uploadJSON", api.handleGoogleJson).Methods("POST")
	router.HandleFunc("/getUserCounties", api.getCountiesforUser).Methods("GET")

	// Add other routes
	//return the router
	return router
}

// calls Geoservice to upload user JSON
func (api *API) handleGoogleJson(w http.ResponseWriter, r *http.Request) {
	controller.HandleGoogleJson(w, r)
}

func (api *API) getCountiesforUser(w http.ResponseWriter, r *http.Request) {
	JSON := controller.GetListOfUserCounties(w, r, *api.DB)

	// Set Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the response body
	_, err := w.Write(JSON)
	if err != nil {
		log.Fatal("Error writing JSON to response: ", err)
	}
}
