package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"movie_CRUD/internal/models"
	"movie_CRUD/internal/repositories"

	"github.com/gorilla/mux"
)

func setJSONResponseHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	setJSONResponseHeader(w)
	json.NewEncoder(w).Encode(repositories.GetMovies())
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	setJSONResponseHeader(w)
	params := mux.Vars(r)
	movie, found := repositories.GetMovieByID(params["id"])
	if found {
		json.NewEncoder(w).Encode(movie)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Film not found")
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	setJSONResponseHeader(w)
	var movie models.Movie

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Printf("Error decoding request body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request data")
		return
	}

	createdMovie := repositories.CreateMovie(movie)
	json.NewEncoder(w).Encode(createdMovie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	setJSONResponseHeader(w)
	params := mux.Vars(r)
	id := params["id"]

	var updatedMovie models.Movie

	err := json.NewDecoder(r.Body).Decode(&updatedMovie)
	if err != nil {
		log.Printf("Error decoding request body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request data")
		return
	}

	updated, found := repositories.UpdateMovie(id, updatedMovie)
	if found {
		json.NewEncoder(w).Encode(updated)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	setJSONResponseHeader(w)
	params := mux.Vars(r)
	id := params["id"]

	deleted := repositories.DeleteMovie(id)
	if deleted {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
