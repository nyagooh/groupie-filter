package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Define a struct to represent each entry in the index
type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// Define a struct to represent the whole JSON structure
type DatesData struct {
	Index []Date `json:"index"`
}

func FetchAndUnmarshalDates() (*DatesData, error) {
	url := "https://groupietrackers.herokuapp.com/api/dates"
	client := &http.Client{
		Timeout: 10 * time.Second, // Set the timeout duration
	}

	// Fetch JSON data from the provided URL
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
	var data DatesData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON data: %w", err)
	}

	return &data, nil
}
