package urlHandlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"io"
	"01.kood.tech/git/jsaar/groupie-tracker/helpers"
	"strconv"
)

//execute location template with fetched location, 
//artist and date data by applying artist ID from URL
func HandleLocations(w http.ResponseWriter, r *http.Request) {
	temp, err := template.New("locations.html").Funcs(template.FuncMap{
		"Split" : strings.Split,
		"CapitalizeAll": helpers.CapitalizeAll,
		"Replace": helpers.Replace,
	}).ParseFiles("./templates/locations.html")

	if err != nil {
		http.Error(w, "Template not found!"+err.Error(), http.StatusInternalServerError)
	}
	artistId := r.URL.Query().Get("id")
	AllArtists := FetchSingleArtist(artistId)
	AllRelations := FetchRelations(artistId)
	AllRelations["artists"] = AllArtists
	executeErr := temp.Execute(w, AllRelations)
	if executeErr != nil {
		fmt.Println("Template error: ", executeErr)
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, "Internal Server Template error on *Location Template* \nError Code: "+strconv.Itoa(http.StatusInternalServerError))
		_ = temp.Execute(w, nil)
		return
	}
}
