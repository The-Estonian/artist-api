package urlHandlers

import (
	"fmt"
	"net/http"
	"html/template"
	"strings"
)

//execute index template with hardcoded dummy data
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "Template not found!"+err.Error(), http.StatusInternalServerError)
	}
	if r.URL.Path != "/" && r.URL.Path != "/artists" && r.URL.Path != "/locations" && !strings.HasPrefix(r.URL.Path, "/locations") {
		http.Error(w, "Bad Request: 404", http.StatusNotFound)
		return
	}

	executeErr := template.Execute(w, nil)
	if executeErr != nil {
		fmt.Println("Template error: ", executeErr)
	}
}
