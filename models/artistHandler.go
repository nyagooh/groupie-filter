package models

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strconv"

	"groupie/services"
)

type Data struct {
	Success      bool
	Result       string
	UserInput    string
	StatusCode   int
	ErrorMessage string
	Feedback     string
}

var feedback string

const (
	notFound            = http.StatusNotFound
	internalServerError = http.StatusInternalServerError
	methodNotAllowed    = http.StatusMethodNotAllowed
	badRequest          = http.StatusBadRequest
)

func handleError(writer http.ResponseWriter, statusCode int, message string) {
	// Construct the URL for the error page with query parameters
	target := fmt.Sprintf("/error?code=%d&message=%s", statusCode, url.QueryEscape(message))
	http.Redirect(writer, &http.Request{URL: &url.URL{Path: target}}, target, http.StatusSeeOther)
}

// ArtistsHandler handles the request to fetch and display artist data
func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := services.FetchAndUnmarshalArtists()
	if err != nil {
		handleError(w, notFound, "Page not found")
		feedback = "This page does not exist"
		return
	}

	// Parse the HTML template file
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		handleError(w, notFound, "Page not found")
		feedback = "This page does not exist"
		return
	}

	// Execute the template with the artist data
	err = tmpl.Execute(w, artists)
	if err != nil {
		handleError(w, internalServerError, "internalServerError")
		feedback = "This page does not exist"
		return
	}
}

func ErrorHandler(writer http.ResponseWriter, reader *http.Request) {
	statusCodeStr := reader.URL.Query().Get("code")
	statusCode, err := strconv.Atoi(statusCodeStr)
	if err != nil {
		statusCode = http.StatusInternalServerError
	}
	message := reader.URL.Query().Get("message")
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		handleError(writer, notFound, "Page not found")
		feedback = "This page does not exist"
		return
	}

	err = tmpl.Execute(writer, Data{ErrorMessage: message, StatusCode: statusCode, Feedback: feedback})
	if err != nil {
		http.Error(writer, "Error rendering error page", http.StatusInternalServerError)
		fmt.Printf("Error executing error template: %s\n", err)
	}
}
