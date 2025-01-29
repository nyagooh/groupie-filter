package models

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"groupie/services"
)

// FilteredArtistsHandler processes the filtering of artists based on the query parameters
func FilteredArtistsHandler(w http.ResponseWriter, r *http.Request) {
	// Parse filter criteria from query parameters
	criteria := services.FilterCriteria{
		CreationDateFrom:  0,
		CreationDateTo:    0,
		FirstAlbumFrom:    r.URL.Query().Get("firstAlbumFrom"),
		FirstAlbumTo:      r.URL.Query().Get("firstAlbumTo"),
		MemberCount:       0,
		LocationSubstring: r.URL.Query().Get("locationSubstring"),
	}

	// Convert integer fields (CreationDateFrom, CreationDateTo, MemberCount)
	if creationDateFromStr := r.URL.Query().Get("creationDateFrom"); creationDateFromStr != "" {
		creationDateFrom, err := strconv.Atoi(creationDateFromStr)
		if err != nil {
			http.Error(w, "Invalid creationDateFrom value", http.StatusBadRequest)
			return
		}
		criteria.CreationDateFrom = creationDateFrom
	}
	if creationDateToStr := r.URL.Query().Get("creationDateTo"); creationDateToStr != "" {
		creationDateTo, err := strconv.Atoi(creationDateToStr)
		if err != nil {
			http.Error(w, "Invalid creationDateTo value", http.StatusBadRequest)
			return
		}
		criteria.CreationDateTo = creationDateTo
	}
	if memberCountStr := r.URL.Query().Get("memberCount"); memberCountStr != "" {
		memberCount, err := strconv.Atoi(memberCountStr)
		if err != nil {
			http.Error(w, "Invalid memberCount value", http.StatusBadRequest)
			return
		}
		criteria.MemberCount = memberCount
	}

	// Sample list of artists
	artists := []services.Artist{
		// Your artists data
	}

	// Call the FilterArtists function to apply the filtering
	filteredArtists, err := services.FilterArtists(artists, criteria)
	if err != nil {
		http.Error(w, "Error filtering artists", http.StatusInternalServerError)
		return
	}

	// Render the filtered data in a template

	tmpl, err := template.ParseFiles("templates/filter.html") /// this handler renders to filter.html template
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Render template with filtered data
	err = tmpl.Execute(w, filteredArtists)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
