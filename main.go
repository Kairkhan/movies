package main

import (
	"fmt"

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

func main() {
	router := mux.NewRouter()

	router.HandlerFunc("/movies", index).Methods("GET")
	router.HandleFunc("/movies/{id}", view).Methods("GET")
	router.HandleFunc("/movies", store).Methods("POST")
	router.HandleFunc("/movies/{id}", update).Methods("PUT")
	router.HandleFunc("/movies/{id}", destroy).Methods("DELETE")

	fmt.Printf("Starting server on port 8000\n")

}
