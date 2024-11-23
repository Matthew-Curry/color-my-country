package Tests

import (
	api "Go-directory/API_Init"
	serverinit "Go-directory/Server_Init"
	"context"
	"net/http"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

var API *api.API

// Testing file. Docker image must be composed locally to use database connection for testing
func TestMain(m *testing.M) {
	// Setup code (e.g., initialize database connections, environment variables, etc.)
	API, server, client := setup()

	// Run all tests
	code := m.Run()

	// Teardown code (e.g., close database connections, clean up files, etc.)
	teardown(API, server, client)

	// Exit with the test code result
	os.Exit(code)
}

func setup() (*api.API, *http.Server, *mongo.Client) {
	// start up API server
	API, server, client := serverinit.InitializeServer()
	return API, server, client
}

func teardown(API *api.API, server *http.Server, client *mongo.Client) {
	// Code to clean up after tests
	println("Cleaning up after tests...")
	// Gracefully shutdown server
	server.Shutdown(context.Background())
	// Close database connection\
	client.Disconnect(context.Background())
}
