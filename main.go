package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"groupie/models"
	"groupie/services"
)

func main() {
	// Serve static files
	staticDir := filepath.Join(".", "static")
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/dates", models.DatesHandler)
	http.HandleFunc("/locations", models.LocationsHandler)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/artist/", models.ArtistDetailsHandler)

	http.HandleFunc("/search", models.SearchHandler)
	http.HandleFunc("/searchLocation", models.SearchLocationHandler)
	http.HandleFunc("/searchDate", models.SearchDatesHandler)
	fmt.Println("Server is starting...")
	fmt.Println("Go on http://localhost:8080/")
	fmt.Println("To shut down the server press CTRL + C")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := services.FetchAndUnmarshalArtists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, artists)
}
