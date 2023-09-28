package main

import (
	"fmt"
	"net/http"
	"01.kood.tech/git/jsaar/groupie-tracker/urlHandlers"
)

var PORT = "8080"

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", urlHandlers.HandleIndex)
	http.HandleFunc("/artists", urlHandlers.HandleArtists)
	http.HandleFunc("/locations", urlHandlers.HandleLocations)

	fmt.Println("Server hosted at: http://localhost:8080")
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		panic(err)
	}
}
