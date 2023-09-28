package urlHandlers

import (
	"fmt"
	"html/template"

	"net/http"
)

//execute artists template with fetched artists data
func HandleArtists(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./templates/artists.html")
	if err != nil {
		http.Error(w, "Template not found!"+err.Error(), http.StatusInternalServerError)
	}

	AllArtists := FetchAllArtists()

	executeErr := template.Execute(w, AllArtists)

	if executeErr != nil {
		fmt.Println("Template error: ", executeErr)
	}
}
