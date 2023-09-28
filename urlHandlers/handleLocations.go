package urlHandlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"01.kood.tech/git/jsaar/groupie-tracker/helpers"
)

//execute location template with fetched location, 
//artist and date data by applying artist ID from URL
func HandleLocations(w http.ResponseWriter, r *http.Request) {
	temp, err := template.New("locations.html").Funcs(template.FuncMap{
		"Split" : strings.Split,
		"Upper" : strings.ToUpper,
		"Capitalize": helpers.Capitalize,
		"CapitalizeAll": helpers.CapitalizeAll,
		"Replace": helpers.Replace,
	}).ParseFiles("./templates/locations.html")

	if err != nil {
		http.Error(w, "Template not found!"+err.Error(), http.StatusInternalServerError)
	}
	artistId := r.URL.Query().Get("id")

	dataMap := make(map[string][]string)

	AllLocations := FetchLocations(artistId)
	AllArtists := FetchSingleArtist(artistId)
	AllDates := FetchDates(artistId)


	for j := 0; j < len(AllArtists.Members); j++ {
		dataMap["artistMembers"] = append(dataMap["artistMembers"], AllArtists.Members[j])
	}

	for i := 0; i < len(AllLocations.Location); i++ {
		dataMap["artistLocations"] = append(dataMap["artistLocations"], AllLocations.Location[i])
		dataMap["artistDates"] = append(dataMap["artistDates"], AllDates.Dates[i])
	}
	
	fmt.Println(dataMap)
	fmt.Println("---------------------------------------------------------")
	executeErr := temp.Execute(w, dataMap)
	if executeErr != nil {
		fmt.Println("Template error: ", executeErr)
	}
}
