package urlHandlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//Date api struct DO NOT MODIFY!
type Dates struct {
	Id       int      `json:"id"`
	Dates []string `json:"dates"`
}

//fetch single artist dates and return as struct
func FetchDates(artistId string) Dates {
	var AllDates Dates
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + artistId)
	if err != nil {
		fmt.Print("API Connection failed!", err.Error())
	}
	jsonData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Body parsing failed!", err.Error())
	}
	jsonError := json.Unmarshal(jsonData, &AllDates)
	if jsonError != nil {
		fmt.Println(jsonError.Error())
	}
	return AllDates
}
