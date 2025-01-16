package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"groupie/services"
)

// Mock data for locations
var locationsMock = services.LocationData{
	Index: []services.Index{
		{ID: 1, Locations: []string{"Location1", "Location2"}},
		{ID: 2, Locations: []string{"Location3", "Location4"}},
	},
}

// Helper function to create a mock server for locations API
func createMockLocationsServer(mockData services.LocationData) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockData)
	})
	return httptest.NewServer(handler)
}

// Test function for FetchAndUnmarshalLocations
func TestFetchAndUnmarshalLocations(t *testing.T) {
	// Create mock server for the locations API
	mockServer := createMockLocationsServer(locationsMock)
	defer mockServer.Close()

	// Override the URL in the function to use the mock server

	// Call the function to test
	data, err := services.FetchAndUnmarshalLocations()

	// Assert no error occurred
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Assert correct data is returned
	if len(data.Index) != 52 {
		t.Fatalf("expected 52 location entries, got %d", len(data.Index))
	}

	// Check the content of the mock data
	if len(data.Index[0].Locations) != 8 {
		t.Errorf("expected 8 locations, got %d", len(data.Index[0].Locations))
	}

}
