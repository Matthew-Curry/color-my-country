package api

/********************************************************************************
* Description: the api package contains all neccessary functions and utilities  *
* that the api for color-your-country will need. It contains a logger, a        *
* database connectionand a router. It will be initialized as an object and used *
* as such.                                                                      *
*********************************************************************************/
import (
	"Go-directory/controller"
	GeoLocater "Go-directory/services"
	"encoding/json"
	"fmt"
	"net/http"

	"context"
	"io/ioutil"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// API structs
type userCounties struct {
	Counties []string `json:"Counties"`
}

type API struct {
	DB     *mongo.Database
	Logger *log.Logger
	//services
	//Geolocator *GeoLocater

}

// Initialize and return a database connection
func NewDatabase() (*mongo.Database, error, context.Context, *mongo.Client) {
	// check to see if environment is dockerized. If so, set connection string appropriately
	dockerEnv := os.Getenv("DOCKERIZED")
	dataBaseContainerURI := ""
	if dockerEnv != "" {
		dataBaseContainerURI = "mongodb://db:27017" // docker connection
	} else {
		dataBaseContainerURI = "mongodb://localhost:27017" // local connection
	}
	//setup mongo connection through database container
	client, err := mongo.NewClient(options.Client(), options.Client().ApplyURI(dataBaseContainerURI))
	//if the connection cannot be established, log the error
	if err != nil {
		return nil, fmt.Errorf("Can't establish connection to database"), nil, nil
	}
	//give the mongo driver a defined timeout
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//connect to the database
	err = client.Connect(ctx)
	//if the connection cannot be established, log the error
	if err != nil {
		return nil, fmt.Errorf("Can't establish connection to database"), nil, nil
	}
	//disconnect if the connection cannot be established
	//defer client.Disconnect(ctx)
	//ping the database to see if connection is established
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	//setup database connection
	database := client.Database("color-my-country-db")
	//return the database connection
	return database, nil, ctx, client
}

// Initialize and return a logger
func NewLogger() (*log.Logger, error) {
	return &log.Logger{}, nil

}

// initialize a new api
func NewAPI() (*API, error, *mongo.Client) {
	//create a database connection
	db, err, ctx, client := NewDatabase()
	if err != nil {
		return nil, err, nil
	}

	//so ctx error goes away
	fmt.Print(ctx)

	//create a logger
	logger, err := NewLogger()
	if err != nil {
		return nil, err, nil
	}
	//create an api from the API struct
	api := &API{
		DB:     db,
		Logger: logger,
		// Initialize other services

	}
	//return the api
	return api, nil, client
}

// create routes and set handler functions for each route
func (api *API) SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Define the routes using the router and API methods
	router.HandleFunc("/uploadJSON", api.handleGoogleJson).Methods("POST")
	router.HandleFunc("/getUserCounties", api.getCountiesforUser).Methods("GET")
	router.HandleFunc("/uploadCounties", api.uploadUserCounties).Methods("POST")
	router.HandleFunc("/deleteCounties", api.deleteUserCounties).Methods("POST")

	//return the router
	return router
}

// calls Geoservice to upload user JSON
func (api *API) handleGoogleJson(w http.ResponseWriter, r *http.Request) {
	//check to see if method is post
	//if so return method not allowed status
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if the Content-Type is "application/json", if not return
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
		return
	}

	// Read the request body
	userJson, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	GeoLocater.GeoService(w, r, userJson, *api.DB)
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

func (api *API) uploadUserCounties(w http.ResponseWriter, r *http.Request) {
	//check to see if method is post
	//if so return method not allowed status
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if the Content-Type is "application/json", if not return
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
		return
	}

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	//unmarshall the json (retrive the county array)
	var counties userCounties

	err = json.Unmarshal(body, &counties)

	controller.Addcounties(w, r, counties.Counties, *api.DB)
}

func (api *API) deleteUserCounties(w http.ResponseWriter, r *http.Request) {
	//check to see if method is post
	//if so return method not allowed status
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if the Content-Type is "application/json", if not return
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
		return
	}

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	//unmarshall the json (retrive the county array)
	var counties userCounties

	err = json.Unmarshal(body, &counties)

	controller.DeleteUserCounties(w, r, counties.Counties, *api.DB)
}
