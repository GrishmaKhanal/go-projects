package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	// "log"
	// "encoding/json"
	// "math/rand"
	// "net/http"
	// "strconv"
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

// slice of type movie
var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Movie not found with id " + params["id"]})
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			temp := movies[index]
			movies = append(movies[:index], movies[index+1:]...)

			json.NewEncoder(w).Encode(map[string]string{"message": temp.Title + " - Movie Deleted Successfully"})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Movie not found"})
}

//	func getUniqueId() string {
//		uniqueId := strconv.Itoa(rand.IntN(100))
//		for _, item := range movies {
//			if item.ID == uniqueId {
//				getUniqueId()
//			}
//		}
//		return uniqueId
//	}
//
// using a loop instead of recursion to avoid stack overflow
func getUniqueId() string {
	for {
		uniqueId := strconv.Itoa(rand.IntN(100))
		isUnique := true

		for _, item := range movies {
			if item.ID == uniqueId {
				isUnique = false
				break
			}
		}

		if isUnique {
			return uniqueId
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// generate a unique id
	movie.ID = getUniqueId()

	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies,
		Movie{ID: "1", Isbn: "420069", Title: "The Great Adventure", Director: &Director{Firstname: "Grishma", Lastname: "Khanal"}},
		Movie{ID: "2", Isbn: "654321", Title: "Mystery of the Code", Director: &Director{Firstname: "Alice", Lastname: "Smith"}},
		Movie{ID: "3", Isbn: "112233", Title: "Rise of the Phoenix", Director: &Director{Firstname: "Michael", Lastname: "Johnson"}},
		Movie{ID: "4", Isbn: "778899", Title: "The Last Stand", Director: &Director{Firstname: "Emma", Lastname: "Brown"}},
	)

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	// r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting Server at port: 8080")
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
