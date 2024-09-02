package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RelationsData struct {
	Index []Relations `json:"index"`
}

func FetchAndUnmarshalRelations() ([]Relations, error) {
	url := "https://groupietrackers.herokuapp.com/api/relation"
	// Create an HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second, // Set the timeout duration
	}

	// Fetch JSON data from the provided URL using the custom client
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching JSON data: %w", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	// Decode the JSON data
	var data RelationsData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON data: %w", err)
	}

	return data.Index, nil
}
