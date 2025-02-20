package models

import (
	"html/template"
	"net/http"

	"groupie/services"
)

// ArtistsHandler handles the request to fetch and display artist data
func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := services.FetchAndUnmarshalArtists()
	if err != nil {
		HandleError(w, notFound, "Page not found")
		Feedback = "This page does not exist"
		return
	}

	// Parse the HTML template file
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		HandleError(w, notFound, "Page not found")
		Feedback = "This page does not exist"
		return
	}

	// Execute the template with the artist data
	err = tmpl.Execute(w, artists)
	if err != nil {
		HandleError(w, internalServerError, "internalServerError")
		Feedback = "This page does not exist"
		return
	}
}
