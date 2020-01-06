package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"MovieList",
		"GET",
		"/peliculas",
		MovieList,
	},
	Route{
		"MovieShow",
		"GET",
		"/pelicula/{id}",
		MovieShow,
	},
	Route{
		"MovieAdd",
		"POST",
		"/pelicula",
		MovieAdd,
	},
	Route{
		"MovieUpdate",
		"PUT",
		"/pelicula/{id}",
		MovieUpdate,
	},
	Route{
		"MovieRemove",
		"DELETE",
		"/pelicula/{id}",
		MovieRemove,
	},
}
