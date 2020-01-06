package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Conexion de go a MongoDB

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}

func responseMovie(w http.ResponseWriter, status int, results Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

func responseMovies(w http.ResponseWriter, status int, results []Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

// se utiliza para poder reeinvocar la variable cuantas veces quiera
var collection = getSession().DB("curso_go").C("movies")

// Todos los metodos de la api

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola mundo desde mi servidor web desde go")
}

func MovieList(w http.ResponseWriter, r *http.Request) {

	var results []Movie
	err := collection.Find(nil).Sort("-_id").All(&results)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Resultados: ", results)
	}

	responseMovies(w, 200, results)

}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	results := Movie{}
	err := collection.FindId(oid).One(&results)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseMovie(w, 200, results)
}
func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movieData Movie
	err := decoder.Decode(&movieData)

	if err != nil {
		panic(err)
	}

	// Cerrar la lectura del body
	defer r.Body.Close()

	//log.Println(movieData) // lectura por consola de los datos que llegan por POST

	err = collection.Insert(movieData)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	responseMovie(w, 200, movieData)

}

func MovieUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	document := bson.M{"_id": oid}
	change := bson.M{"$set": movie_data}
	err = collection.Update(document, change)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseMovie(w, 200, movie_data)
}

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// punteros es una copia y no apunta a la variable principal (*) siempre tiene ese signo
func (this *Message) setStatus(data string) {
	this.Status = data
}

func (this *Message) setMessage(data string) {
	this.Message = data
}

// nuevo metodo eliminar
func MovieRemove(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	err := collection.RemoveId(oid)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	//results := Message{"success", "La pelicula con el ID" + movie_id + "ha sido borrada correctamente"}

	message := new(Message)
	message.setStatus("success")
	message.setMessage("La pelicula con el ID" + movie_id + "ha sido borrada correctamente")

	results := message

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}
