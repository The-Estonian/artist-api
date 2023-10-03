package main

import (
	"fmt"
	"net/http"

	"01.kood.tech/git/jsaar/groupie-tracker/urlHandlers"
)

const PORT = "80"

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", urlHandlers.HandleIndex)
	http.HandleFunc("/artists", urlHandlers.HandleArtists)
	http.HandleFunc("/locations", urlHandlers.HandleLocations)

	fmt.Println("Server hosted at: http://localhost")
	fmt.Println("To Kill Server press Ctrl+C")
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		panic(err)
	}
}
