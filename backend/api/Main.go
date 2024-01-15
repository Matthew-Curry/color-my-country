//Should work with test user for now

//query database, make sure data is built and correct(run loop until condition is true)
//initialize services
//need to write your own docker file, and point to it in docker-compose for the api

package main

import (
	"Go-directory/controller"
)

//"GeoLocator"

//first need to understand mongo db, how to connect and query database. Then need to properly run curl commands to test main to make sure it querys database

//Then need to figure out service initiation

//Then docker

func main() {
	//initialize database connection
	controller.ConnectToDatabase()

	//initialize services

}
