package main

import (
	"encoding/json"
	
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string] string {
		"status": "available",
		"enviroment": app.config.env,
		"version": version,
	}
	js , err := json.Marshal(data)

	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "the server encoutered s proble could not process your req",http.StatusInternalServerError)

		return
	}
    
	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")

    // Use w.Write() to send the []byte slice containing the JSON as the response body.
    w.Write(js)
}