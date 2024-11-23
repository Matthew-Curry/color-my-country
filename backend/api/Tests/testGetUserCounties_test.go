package Tests

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

// Tests the GetUserCounties endpoint
func TestGetUserCounties(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/getUserCounties?username=test")

	if err != nil {
		t.Errorf("Error retriving user counties")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	county := body[0]

	if string(county) == "York" {
		fmt.Println("Get user counties test successful")
	} else {
		t.Errorf("Endpoint returned incorrect information")
	}
}
