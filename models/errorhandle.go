// handlers/error_handler.go

package models

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"text/template"
)

// Data holds the data to be passed to the error template.
type Data struct {
	Success      bool
	Result       string
	UserInput    string
	StatusCode   int
	ErrorMessage string
	Feedback     string
}

var Feedback string

const (
	notFound            = http.StatusNotFound
	internalServerError = http.StatusInternalServerError
	methodNotAllowed    = http.StatusMethodNotAllowed
	badRequest          = http.StatusBadRequest
)

// ErrorHandler handles errors by rendering an error page.
func ErrorHandler(writer http.ResponseWriter, reader *http.Request) {
	statusCodeStr := reader.URL.Query().Get("code")
	statusCode, err := strconv.Atoi(statusCodeStr)
	if err != nil {
		statusCode = internalServerError
	}
	message := reader.URL.Query().Get("message")
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		HandleError(writer, notFound, "Page not found")
		Feedback = "This page does not exist"
		return
	}
	err = tmpl.Execute(writer, Data{ErrorMessage: message, StatusCode: statusCode, Feedback: Feedback})
	if err != nil {
		http.Error(writer, "Error rendering error page", internalServerError)
		fmt.Printf("Error executing error template: %s\n", err)
	}
}

// handleError redirects to the error page with the specified status code and message.
func HandleError(writer http.ResponseWriter, statusCode int, message string) {
	// Construct the URL for the error page with query parameters
	target := fmt.Sprintf("/error?code=%d&message=%s", statusCode, url.QueryEscape(message))
	http.Redirect(writer, &http.Request{URL: &url.URL{Path: target}}, target, http.StatusSeeOther)
}
