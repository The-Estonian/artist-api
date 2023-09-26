package urlHandlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
}

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
