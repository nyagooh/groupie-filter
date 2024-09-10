package models

import (
	"html/template"
	"net/http"
	"path/filepath"

	"groupie/services"
)

// ArtistDate represents the combined data of artist names and dates
type ArtistDate struct {
	Name  string   `json:"name"`
	Dates []string `json:"dates"`
}

// DatesHandler serves the tour dates with artist names
func DatesHandler(w http.ResponseWriter, r *http.Request) {
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

	// Fetch dates data
	datesData, err := services.FetchAndUnmarshalDates()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Combine artist names with dates
	var artistDates []ArtistDate
	for _, date := range datesData.Index {
		name, found := artistNameMap[date.ID]
		if !found {
			name = "Unknown"
		}
		artistDates = append(artistDates, ArtistDate{
			Name:  name,
			Dates: date.Dates,
		})
	}

	// Parse the HTML template file
	tmpl, err := template.ParseFiles(filepath.Join("templates", "dates.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with the combined data
	err = tmpl.Execute(w, struct {
		Index []ArtistDate
	}{
		Index: artistDates,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
