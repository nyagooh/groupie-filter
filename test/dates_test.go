package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock data for dates
var datesMock = DatesData{
	Index: []Date{
		{ID: 1, Dates: []string{"2022-01-01", "2022-01-02"}},
		{ID: 2, Dates: []string{"2022-02-01", "2022-02-02"}},
	},
}

// Helper function to create a mock server for dates API
func createMockDatesServer(mockData DatesData) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockData)
	})
	return httptest.NewServer(handler)
}

// Test function for FetchAndUnmarshalDates
func TestFetchAndUnmarshalDates(t *testing.T) {
	// Create mock server for the dates API
	mockServer := createMockDatesServer(datesMock)
	defer mockServer.Close()

	// Override the URL in the function to use the mock server

	// Call the function to test
	data, err := FetchAndUnmarshalDates()

	// Assert no error occurred
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Assert correct data is returned
	if len(data.Index) != 52 {
		t.Fatalf("expected 52 date entries, got %d", len(data.Index))
	}

	// Check the content of the mock data
	if len(data.Index[0].Dates) != 8 {
		t.Errorf("expected %d dates, got 2", len(data.Index[0].Dates))
	}

}
