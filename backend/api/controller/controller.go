package controller

/********************************************************************************
* Description: The controller handles important user functions, and uses the    *
* dao package to actually connect to the database                               *
*********************************************************************************/
import (
	"Go-directory/dao"
	//GeoLocater "Go-directory/services"

	//"context"
	"fmt"
	//"io/ioutil"

	//"log"
	"net/http"
	//"strconv"

	"go.mongodb.org/mongo-driver/mongo"
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
func GetListOfUserCounties(w http.ResponseWriter, r *http.Request, database mongo.Database) []byte {

	// Get query parameter for username
	queryParams := r.URL.Query()

	// Get specific query parameter values (for now it's username)
	username := queryParams.Get("username")

	// Get JSON data from dao
	fmt.Println("Worked")
	JSON := dao.GetUserCounites(username, database)

	return JSON

}

func DeleteUserCounties(w http.ResponseWriter, r *http.Request, counties []string, database mongo.Database) {
	// Parse query parameters from the request
	queryParams := r.URL.Query()

	// Access specific query parameters by name
	username := queryParams.Get("username")

	dao.DeleteCountiesforUser(counties, username, database)
}

// Add counties to database for a user
func Addcounties(w http.ResponseWriter, r *http.Request, counties []string, database mongo.Database) {
	// Parse query parameters from the request
	queryParams := r.URL.Query()

	// Access specific query parameters by name
	username := queryParams.Get("username")

	dao.AddCounitesForUser(counties, username, database)

}
