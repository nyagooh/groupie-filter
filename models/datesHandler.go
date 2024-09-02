package models

import (
	"html/template"
	"net/http"

	"groupie/services"
)

func DatesHandler(w http.ResponseWriter, r *http.Request) {
	dates, err := services.FetchAndUnmarshalDates()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Parse the HTML template file
	tmpl, err := template.ParseFiles("templates/dates.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with the dates data
	err = tmpl.Execute(w, dates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
