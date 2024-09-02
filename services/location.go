package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type LocationsData struct {
	Index []Locations `json:"index"`
}

func FetchAndUnmarshalLocations() (*LocationsData, error) {
	url := "https://groupietrackers.herokuapp.com/api/locations"

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching JSON data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	var data LocationsData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON data: %w", err)
	}

	return &data, nil
}
