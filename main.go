package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"groupie/models"
	"groupie/services"
)

func main() {
	staticDir := filepath.Join(".", "static")
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/dates", models.DatesHandler)
	http.HandleFunc("/locations", models.LocationsHandler)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/artist/", models.ArtistDetailsHandler)

	http.HandleFunc("/search", models.SearchHandler)
	http.HandleFunc("/searchLocation", models.SearchLocationHandler)

	//
	//http.HandleFunc("/searchMap", models.SearchMapHandler)

	http.HandleFunc("/searchDate", models.SearchDatesHandler)
	// http.HandleFunc("/filter", models.FilteredArtistsHandler)
	http.HandleFunc("/error", models.ErrorHandler)

	fmt.Println("Server is starting...")
	fmt.Println("Go on http://localhost:8080/")
	fmt.Println("To shut down the server press CTRL + C")
	http.ListenAndServe(":8080", nil)
	if err := http.ListenAndServe(":8080", nil); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err) // Fatal stops execution if server crashes
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := services.FetchAndUnmarshalArtists()
	if err != nil {
		models.HandleError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if r.URL.Path != "/" {
		models.HandleError(w, http.StatusNotFound, "Page not found")
		models.Feedback = "This page does not exist"
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, artists)
}
