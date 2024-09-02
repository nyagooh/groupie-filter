package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Artist represents the JSON structure for each artist
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

// FetchAndUnmarshal fetches JSON data from the URL and unmarshals it
func FetchAndUnmarshalArtists(url string) ([]Artist, error) {
	// Create an HTTP client with a timeout
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

	// Decode the JSON data into a slice of Artist structs
	var artists []Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON data: %w", err)
	}

	return artists, nil
}
