package internal

import (
	"net/http"
)

func SearchPageHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		SearchPageGetHandler(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func SearchPageGetHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/search" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	err := templates.ExecuteTemplate(w, "search.html", nil)
	if err != nil {
		http.Error(w, "Go back to the main page", 500)
		return
	}
}

func SearchResultPageHandler(w http.ResponseWriter, req *http.Request) {

}
func ArtistInfoPageHandler(w http.ResponseWriter, req *http.Request) {

}
