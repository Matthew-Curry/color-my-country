package GeoLocater

/**************************************************************************
*Credits:                                                                 *
*-Nerdcademy                                                              *
*	Source code can be found at https://github.com/NerdCademyDev/golang   *
*-Chat GPT                                                                *
*-https://golangdocs.com/golang-read-json-file                            *
**************************************************************************/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"os"
	"time"
)

var client *http.Client

// for getting county/area location after parsing google maps JSON for visited locations
type GeocodeResponse struct {
	PlaceID     int    `json:"place_id"`
	License     string `json:"licence"`
	PoweredBy   string `json:"powered_by"`
	OSMType     string `json:"osm_type"`
	OSMID       int64  `json:"osm_id"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`

	Emergency     string   `json:"emergency"`
	Address       address  `json:"address"`
	Road          string   `json:"road"`
	Neighbourhood string   `json:"neighbourhood"`
	City          string   `json:"city"`
	State         string   `json:"state"`
	Postcode      string   `json:"postcode"`
	Country       string   `json:"country"`
	CountryCode   string   `json:"country_code"`
	BoundingBox   []string `json:"boundingbox"`
}

type address struct {
	City    string `json:"city"`
	County  string `json:"county"`
	State   string `json:"state"`
	Zipcode string `json:"postcode"`
}

//structs for unmarshaling google maps json

type PlaceVisit struct {
	Location struct {
		LatitudeE7            int     `json:"latitudeE7"`
		LongitudeE7           int     `json:"longitudeE7"`
		PlaceID               string  `json:"placeId"`
		Address               string  `json:"address"`
		Name                  string  `json:"name"`
		LocationConfidence    float64 `json:"locationConfidence"`
		CalibratedProbability float64 `json:"calibratedProbability"`
	} `json:"location"`
	Duration struct {
		StartTimestamp string `json:"startTimestamp"`
		EndTimestamp   string `json:"endTimestamp"`
	} `json:"duration"`
	PlaceConfidence        string `json:"placeConfidence"`
	VisitConfidence        int    `json:"visitConfidence"`
	EditConfirmationStatus string `json:"editConfirmationStatus"`
	LocationConfidence     int    `json:"locationConfidence"`
	PlaceVisitType         string `json:"placeVisitType"`
	LocationAssertionType  string `json:"locationAssertionType"`
	LastEditedTimestamp    string `json:"lastEditedTimestamp"`
	PlaceVisitImportance   string `json:"placeVisitImportance"`
	EditActionMetadata     struct {
		EditHistory struct {
			EditEvent []struct {
				EditOperation []string `json:"editOperation"`
			} `json:"editEvent"`
		} `json:"editHistory"`
	} `json:"editActionMetadata"`
}

type TimelineObject struct {
	PlaceVisit PlaceVisit `json:"placeVisit"`
}

type TimelineData struct {
	TimelineObjects []TimelineObject `json:"timelineObjects"`
}

func getJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&target)

}

func getCoordInfo(lat float64, lon float64) (string, error) {
	//test latitude and longitude (for York pa)
	//lat := "39.96"
	//lng := "-76.72"

	// Construct the URL for the geocoding API request
	Url := fmt.Sprintf("https://geocode.maps.co/reverse?lat=%s&lon=%s", lat, lon)
	fmt.Printf(Url)
	var inputStruct GeocodeResponse

	err := getJson(Url, &inputStruct)

	if err != nil {
		fmt.Print("HTTP error")
		return "", err
	} else {
		county := inputStruct.Address.County
		return county, nil
	}

}

func GetCounties(GoogleJson []byte) []string {
	//content, err := os.ReadFile("./Example_GoogleData/2023_SEPTEMBER.json")
	/*
		if err != nil {
			fmt.Print("Problem reading file")
		}

		//fmt.Print(content)
	*/

	var userLocation TimelineData
	err := json.Unmarshal(GoogleJson, &userLocation)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	countyArr := make([]string, len(userLocation.TimelineObjects))

	//converts to proper latitude
	//go through Timeline objects, and extract all counties
	for i := 0; i < len(userLocation.TimelineObjects); i++ {
		Latitude := float64((userLocation.TimelineObjects[0].PlaceVisit.Location.LatitudeE7)) / 10000000.0
		fmt.Printf("%.8f\n", Latitude)
		//converts to proper longitude
		Longitude := float64((userLocation.TimelineObjects[0].PlaceVisit.Location.LongitudeE7)) / 10000000.0
		fmt.Printf("%.8f\n", Longitude)

		County, err := getCoordInfo(Latitude, Longitude)
		if err != nil {
			fmt.Print("Error getting coordinate information")
		} else {
			countyArr[i] = County
		}

	}

	client = &http.Client{Timeout: 10 * time.Second}

	return countyArr

}
