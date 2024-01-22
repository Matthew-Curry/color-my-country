package controller
/********************************************************************************
* Description: The controller handles important user functions, and uses the    *
* dao package to actually connect to the database                               *
*********************************************************************************/
import (
	"Go-directory/dao"
	GeoLocater "Go-directory/services"
	//"context"
	"fmt"
	"io/ioutil"
	//"log"
	"net/http"
	"strconv"
	//"time"

	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
)

//"fmt"

//"strconv"
//"time"



// will use GeoLocater package to get counties for a user given their google maps JSON, and then call the database to add these counties to the users list, if they do not already exist
func GetListOfUserCounties(w http.ResponseWriter, r *http.Request) {

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

func DeleteUserCounties(w http.ResponseWriter, r *http.Request, counties []string) {
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

// Add counties to database for a user
func Addcounties(w http.ResponseWriter, r *http.Request, counties []string) {
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

func handleGoogleJson(w http.ResponseWriter, r *http.Request) {
	GeoLocater.GeoService(w, r)
}
