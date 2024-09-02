package models

import (
	"html/template"
	"net/http"

	"groupie/services"
)

func RelationsHandler(w http.ResponseWriter, r *http.Request) {
	relations, err := services.FetchAndUnmarshalRelations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("templates/relations.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with the dates data
	err = tmpl.Execute(w, relations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
