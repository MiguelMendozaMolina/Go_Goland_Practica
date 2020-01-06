package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var movies = Movies{
	Movie{"Sin limites", 2013, "desconocido"},
	Movie{"El joker", 2019, "pat henkins"},
	Movie{"Dragon Ball", 2010, "Rochi"},
}

// todos los metodos de la api

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola mundo desde mi servidor web desde go")
}

func MovieList(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(movies)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	movieId := params["id"]

	fmt.Fprint(w, "Pelicula cargada numero ", movieId)
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movieData Movie
	err := decoder.Decode(&movieData)

	if err != nil {
		panic(err)
	}

	// defer para limpiar la lectura del body
	defer r.Body.Close()

	log.Println(movieData)
	movies = append(movies, movieData)

}
