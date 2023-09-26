package urlHandlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Locations struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
	Dates    string   `json:"dates"`
}

func FetchLocations(artistId string) Locations {
	var AllLocations Locations
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + artistId)
	if err != nil {
		fmt.Print("API Connection failed!", err.Error())
	}
	jsonData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Body parsing failed!", err.Error())
	}
	jsonError := json.Unmarshal(jsonData, &AllLocations)
	if jsonError != nil {
		fmt.Println("Location Struct error")
		fmt.Println(jsonError.Error())
	}
	return AllLocations
}
