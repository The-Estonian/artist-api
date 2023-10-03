package urlHandlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//fetch single artist relations and return as struct
func FetchRelations(artistId string) map[string]interface{} {
	AllRelations := map[string]interface{}{}
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + artistId)
	if err != nil {
		fmt.Print("API Connection failed!", err.Error())
	}
	jsonData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Body parsing failed!", err.Error())
	}
	jsonError := json.Unmarshal(jsonData, &AllRelations)
	if jsonError != nil {
		fmt.Println(jsonError.Error())
	}
	return AllRelations
}