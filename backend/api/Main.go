package main
/********************************************************************************
* Description: The controller handles important user functions, and uses the    *
* dao package to actually connect to the database                               *
*********************************************************************************/
import (
	api "Go-directory/API_Init"
	"log"
	"net/http"
	"time"
)

//first need to understand mongo db, how to connect and query database. Then need to properly run curl commands to test main to make sure it querys database

//Then need to figure out service initiation

//Then docker

func main() {
	//create a new API
	API, err := api.NewAPI()
	//if there is an error, log the error
	if err != nil {
		log.Fatal(err)
	}
	//initialize the router so it can be used to redirect
	router := API.SetupRoutes()

	// Start the HTTP server
	addr := ":8080"
	server := &http.Server{
		Addr:    addr,
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server listening on %s", addr)
	log.Fatal(server.ListenAndServe())

}
