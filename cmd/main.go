package main

import (
	"log"
	"net/http"

	"movie_CRUD/internal/repositories"
	"movie_CRUD/internal/routers"
)

func main() {
	repositories.InitMovies()

	r := routers.SetupRouter()

	log.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
