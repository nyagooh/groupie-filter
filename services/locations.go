package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Index struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type LocationData struct {
	Index []Index `json:"index"`
}

func FetchAndUnmarshalLocations() (*LocationData, error) {
	// Set locations URL
	url := "https://groupietrackers.herokuapp.com/api/locations"
	client := &http.Client{
		Timeout: 10 * time.Second, // Set the timeout duration to 15 seconds
	}

	// Fetch JSON data from the API
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching locations data: %w", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	// Decode the JSON data
	var data LocationData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("error decoding locations data: %w", err)
	}

	return &data, nil
}
