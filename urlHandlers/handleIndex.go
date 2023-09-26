package urlHandlers

import (
	"fmt"
	"net/http"
	"html/template"
	"strings"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "Template not found!"+err.Error(), http.StatusInternalServerError)
	}
	if r.URL.Path != "/" && r.URL.Path != "/artists" && r.URL.Path != "/locations" && !strings.HasPrefix(r.URL.Path, "/locations") {
		http.Error(w, "Bad Request: 404", http.StatusNotFound)
		return
	}

	dataStream := `Welcome to the site!
	Press Artists Menu button hard to fetch some sweet api data and present it to you
	in a bijutiful way!`

	executeErr := template.Execute(w, dataStream)
	if executeErr != nil {
		fmt.Println("Template error: ", executeErr)
	}
}
