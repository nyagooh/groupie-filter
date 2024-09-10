package models

import (
	"html/template"
	"net/http"
	"path/filepath"

	"groupie/services"
)

// ArtistLocation represents the combined data of artist names and locations
type ArtistLocation struct {
	Name      string   `json:"name"`
	Locations []string `json:"locations"`
}

// LocationsHandler serves the locations with artist names
func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch artist data
	artists, err := services.FetchAndUnmarshalArtists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a map of artist IDs to names
	artistNameMap := make(map[int]string)
	for _, artist := range artists {
		artistNameMap[artist.ID] = artist.Name
	}

	// Fetch locations data
	locationsData, err := services.FetchAndUnmarshalLocations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Combine artist names with locations
	var artistLocations []ArtistLocation
	for _, loc := range locationsData.Index {
		name, found := artistNameMap[loc.ID]
		if !found {
			name = "Unknown"
		}
		artistLocations = append(artistLocations, ArtistLocation{
			Name:      name,
			Locations: loc.Locations,
		})
	}

	// Parse the HTML template file
	tmpl, err := template.ParseFiles(filepath.Join("templates", "locations.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with the combined data
	err = tmpl.Execute(w, struct {
		Index []ArtistLocation
	}{
		Index: artistLocations,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
