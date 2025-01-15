package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	// Update with your actual package path
)

func TestFetchAndUnmarshalArtists(t *testing.T) {
	// Mock data for artists, locations, and relations
	artistsMock := []Artist{
		{
			ID:           1,
			Image:        "image_url",
			Name:         "Band 1",
			Members:      []string{"Member 1", "Member 2"},
			CreationDate: 2000,
			FirstAlbum:   "Album 1",
		},
	}
	locationsMock := LocationsData{
		Index: []Locations{
			{ID: 1, Locations: []string{"City 1", "City 2"}},
		},
	}
	relationsMock := RelationsData{
		Index: []Relations{
			{ID: 1, DatesLocations: map[string][]string{"2024": {"City 1", "City 2"}}},
		},
	}

	// Create mock servers
	artistsServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(artistsMock)
	}))
	defer artistsServer.Close()

	locationsServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(locationsMock)
	}))
	defer locationsServer.Close()

	relationsServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(relationsMock)
	}))
	defer relationsServer.Close()

	// Override URLs temporarily
	artistsURL := "https://groupietrackers.herokuapp.com/api/artists"
	locationsURL := "https://groupietrackers.herokuapp.com/api/locations"
	relationsURL := "https://groupietrackers.herokuapp.com/api/relation"

	originalArtistURL := artistsURL
	originalLocationsURL := locationsURL
	originalRelationsURL := relationsURL

	artistsURL = artistsServer.URL
	locationsURL = locationsServer.URL
	relationsURL = relationsServer.URL

	defer func() {
		artistsURL = originalArtistURL
		locationsURL = originalLocationsURL
		relationsURL = originalRelationsURL
	}()

	// Call the function to test
	artists, err := FetchAndUnmarshalArtists()

	// Check for errors
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Simplified checks
	if len(artists) != 52 {
		t.Fatalf("Expected 52 artists, got %d", len(artists))
	}

	if artists[0].Name != "Queen" {
		t.Errorf("Expected artist name 'Queen', got %s", artists[0].Name)
	}

	expectedLocations := []string{"north_carolina-usa", "georgia-usa"}
	for i, loc := range expectedLocations {
		if artists[0].Locations[i] != loc {
			t.Errorf("Expected location %s, got %s", loc, artists[0].Locations[i])
		}
	}

	// expectedRelations := map[string][]string{"2024": {"City 1", "City 2"}}
	// for year, cities := range expectedRelations {
	// 	if cities[0] != artists[0].Relations[year][0] {
	// 		t.Errorf("Expected relation %v, got %v", cities, artists[0].Relations[year])
	// 	}
	// }
}
