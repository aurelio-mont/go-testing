package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router  {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Pattern).
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
		"Pelicula",
		"GET",
		"/peliculas",
		MovieLits,
	},
	Route{
		"Index",
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
	Route{
		"PersonAdd",
		"POST",
		"/person",
		PersonAdd,
	},
}