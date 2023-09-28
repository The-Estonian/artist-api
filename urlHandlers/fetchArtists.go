package urlHandlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//Artist api struct DO NOT MODIFY!
type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
}

//fetch all artists and return array of structs
func FetchAllArtists() []Artists {
	var AllArtists []Artists
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print("API Connection failed!", err.Error())
	}
	jsonData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Body parsing failed!", err.Error())
	}

	jsonError := json.Unmarshal(jsonData, &AllArtists)
	if jsonError != nil {
		fmt.Println(jsonError.Error())
	}
	return AllArtists
}

//fetch single artist and return as struct
func FetchSingleArtist(artistId string) Artists {
	var AllArtists Artists
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + artistId)
	if err != nil {
		fmt.Print("API Connection failed!", err.Error())
	}
	jsonData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Body parsing failed!", err.Error())
	}

	jsonError := json.Unmarshal(jsonData, &AllArtists)
	if jsonError != nil {
		fmt.Println(jsonError.Error())
	}
	return AllArtists
}
