package urlHandlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type DataArray map[string]interface{}

func HandleLocations(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./templates/locations.html")
	if err != nil {
		http.Error(w, "Template not found!"+err.Error(), http.StatusInternalServerError)
	}
	artistId := r.URL.Query().Get("id")

	AllLocations := FetchLocations(artistId)
	AllArtists := FetchSingleArtist(artistId)

	executeErr := template.Execute(w, DataArray{
		"artists":   AllArtists,
		"locations": AllLocations,
	})
	if executeErr != nil {
		fmt.Println("Template error: ", executeErr)
	}
}
