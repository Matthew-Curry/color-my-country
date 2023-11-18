package controller

import (
	"Go-directory/dao"
	GeoLocater "Go-directory/services"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//"fmt"

//"strconv"
//"time"

// will use GeoLocater package to get counties for a user given their google maps JSON, and then call the database to add these counties to the users list, if they do not already exist
func getListOfUserCounties(w http.ResponseWriter, r *http.Request) {

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

func deleteUserCounties(w http.ResponseWriter, r *http.Request, counties []string) {
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
//will county be a string? Int? Generic type? (for now string)
func addcounties(w http.ResponseWriter, r *http.Request, counties []string) {
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
