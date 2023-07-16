package repositories

import (
	"math/rand"
	"strconv"

	"movie_CRUD/internal/models"
)

var movies []models.Movie

func InitMovies() {
	movies = append(movies, models.Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &models.Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, models.Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &models.Director{Firstname: "Steve", Lastname: "Smith"}})
}

func GetMovies() []models.Movie {
	return movies
}

func GetMovieByID(id string) (models.Movie, bool) {
	for _, movie := range movies {
		if movie.ID == id {
			return movie, true
		}
	}

	return models.Movie{}, false
}

func CreateMovie(movie models.Movie) models.Movie {
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)

	return movie
}

func UpdateMovie(id string, updatedMovie models.Movie) (models.Movie, bool) {
	for index, movie := range movies {
		if movie.ID == id {
			if updatedMovie.Title != "" {
				movies[index].Title = updatedMovie.Title
			}

			if updatedMovie.Isbn != "" {
				movies[index].Isbn = updatedMovie.Isbn
			}

			if updatedMovie.Director != nil {
				movies[index].Director = updatedMovie.Director
			}

			return movies[index], true
		}
	}

	return models.Movie{}, false
}

func DeleteMovie(id string) bool {
	for index, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			return true
		}
	}

	return false
}
