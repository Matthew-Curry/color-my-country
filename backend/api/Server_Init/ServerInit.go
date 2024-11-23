package serverinit

import (
	api "Go-directory/API_Init"
	//"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeServer() (*api.API, *http.Server, *mongo.Client) {

	//create a new API
	API, err, client := api.NewAPI()
	//if there is an error, log the error
	if err != nil {
		log.Fatal(err)
	}

	//initialize the router so it can be used to redirect
	router := API.SetupRoutes()

	// Start the HTTP server
	addr := ":8080"
	server := &http.Server{
		Addr:         addr,
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server listening on %s", addr)
	server.ListenAndServe()
	return API, server, client
}
