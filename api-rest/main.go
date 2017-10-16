package main

import (
	"net/http"
	"log"
)

func main() {
	/*router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Index)
	router.HandleFunc("/peliculas", MovieLits)
	router.HandleFunc("/pelicula/{id}", MovieShow)
	*/
	
	router := NewRouter()

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)
}