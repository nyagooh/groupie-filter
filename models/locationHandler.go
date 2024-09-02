package models

import (
	"html/template"
	"net/http"

	"groupie/services"
)

func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	locations, err := services.FetchAndUnmarshalLocations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // http 500
		return
	}
	tmpl, err := template.ParseFiles("templates/locations.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with the dates data
	err = tmpl.Execute(w, locations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
