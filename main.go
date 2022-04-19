package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func index(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(movies)
}

func destroy(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(writer).Encode(movies)
}

func view(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range movies {

		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
}

func store(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	json.NewEncoder(writer).Encode(movie)
}

func update(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var movie Movie
	params := mux.Vars(request)
	_ = json.NewDecoder(request.Body).Decode(&movie)
	for index, item := range movies {
		if item.ID == params["id"] {
			movie.ID = params["id"]
			movies[index] = movie
			break
		}

	}

	json.NewEncoder(writer).Encode(movie)
}

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		ISBN:  "1244",
		Title: "menyn atym kozha",
		Director: &Director{
			Firstname: "Shaken",
			Lastname:  "Aimanov",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		ISBN:  "432",
		Title: "kyz zhibek",
		Director: &Director{
			Firstname: "Christopher",
			Lastname:  "Nolan",
		},
	})

	router.HandleFunc("/movies", index).Methods("GET")
	router.HandleFunc("/movies/{id}", view).Methods("GET")
	router.HandleFunc("/movies", store).Methods("POST")
	router.HandleFunc("/movies/{id}", update).Methods("PUT")
	router.HandleFunc("/movies/{id}", destroy).Methods("DELETE")

	fmt.Printf("Starting server on port 8001\n")

	log.Fatal(http.ListenAndServe(":8001", router))

}
