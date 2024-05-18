package main

import (
	
	"fmt"
	"greenlight/bytemoves/internal/data"
	"net/http"
	"time"
)


func(app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	// anonymous struct to hold the information that we expect to be in the hhttp req body
	var input struct {
		Title   string   `json:"title"`
        Year    int32    `json:"year"`
        Runtime int32    `json:"runtime"`
        Genres  []string `json:"genres"`
		

	}
	err := app.readJSON(w, r, &input)
    if err != nil {
        app.errorResponse(w, r, http.StatusBadRequest, err.Error())
        return
    }

    fmt.Fprintf(w, "%+v\n", input)
	

}


func (app *application) showMovieHandler (w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
    if err != nil {
        app.notFoundResponse(w,r)
        return
    }

	movie := data.Movie {
		ID: id,
		CreatedAt: time.Now(),
		Title: "casablanca",
		Runtime: 102,
		Genres: []string{"drama","romance","war"},
		Version: 1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"movie":movie}, nil)

	if err != nil {
		app.notFoundResponse(w,r,err)
	}

	
	
}


