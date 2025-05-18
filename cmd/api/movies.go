package main

import (
	"fmt"
	"net/http"

	"greenlight.mancuoj.me/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(w, r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:      id,
		Title:   "Casablanca",
		Year:    0,
		Runtime: 102,
		Genres:  []string{"drama", "romance", "war"},
		Version: 1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
