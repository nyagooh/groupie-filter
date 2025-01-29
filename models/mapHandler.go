package models

import (
	"encoding/json"
	"net/http"
	"strings"
	"fmt"

	"groupie/services"
)

// SearchMapHandler handles the search for locations and artists
func SearchMapHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the search query from the request URL (e.g., /search?query=artist_name)
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Query parameter is required", http.StatusBadRequest)
		return
	}

	// Fetch artist and locations data
	artists, err := services.FetchAndUnmarshalArtists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	locationsData, err := services.FetchAndUnmarshalLocations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a list of matching artist locations
	var matchedLocations []map[string]interface{} 
	for _, loc := range locationsData.Index {
		// Search for matching artist name or location in the query
		if strings.Contains(strings.ToLower(loc.Locations[0]), strings.ToLower(query)) {
			// Find the artist's name for the location
			artistName := "Unknown"
			for _, artist := range artists {
				if artist.ID == loc.ID {
					artistName = artist.Name
					break
				}
			}

			// Add matching locations to the result list
			matchedLocations = append(matchedLocations, map[string]interface{}{
				"name":      artistName,
				"locations": loc.Locations,
			})
		}

		fmt.Println(matchedLocations)// test for locations that match
	}

	// Return the results as JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(matchedLocations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
