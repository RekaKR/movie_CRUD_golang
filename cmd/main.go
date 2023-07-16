package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func setJSONResponseHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	setJSONResponseHeader(w)
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	setJSONResponseHeader(w)
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	setJSONResponseHeader(w)
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Film not found")
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	setJSONResponseHeader(w)
	var movie Movie

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Printf("Error decoding request body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request data")

		return
	}

	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	setJSONResponseHeader(w)
	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.ID == params["id"] {
			var updatedMovie Movie

			err := json.NewDecoder(r.Body).Decode(&updatedMovie)
			if err != nil {
				log.Printf("Error decoding request body: %v\n", err)
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode("Invalid request data")

				return
			}

			if updatedMovie.Title != "" {
				movies[index].Title = updatedMovie.Title
			}

			if updatedMovie.Isbn != "" {
				movies[index].Isbn = updatedMovie.Isbn
			}

			if updatedMovie.Director != nil {
				movies[index].Director = updatedMovie.Director
			}

			json.NewEncoder(w).Encode(movies[index])
			return
		}
	}

	/* // other, but less "good" approach
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie

			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	*/

	w.WriteHeader(http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	log.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
