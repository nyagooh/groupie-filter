package models

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"groupie/services"
)

// ArtistDetailsHandler serves the details of a single artist
func ArtistDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the artist ID from the URL path

	artistIDStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	artistID, err := strconv.Atoi(artistIDStr)
	if err != nil {
		HandleError(w, badRequest, "badrequest")
		return
	}

	artists, err := services.FetchAndUnmarshalArtists()
	if err != nil {
		HandleError(w, internalServerError, "internalservererror")
		return
	}

	// Find the artist with the given ID
	var artist *services.Artist
	for _, a := range artists {
		if a.ID == artistID {
			artist = &a
			break
		}
	}

	if artist == nil {
		HandleError(w, notFound, "pagenotfound")
		return
	}

	tmpl, err := template.ParseFiles(filepath.Join("templates", "artistdetails.html"))
	if err != nil {
		HandleError(w, internalServerError, "internalservererror")
		return
	}

	err = tmpl.Execute(w, artist)
	if err != nil {
		HandleError(w, internalServerError, "internalservererror")
		return
	}
}
