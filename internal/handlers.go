package internal

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

var (
	templates *template.Template
)

func init() {
	templates = template.Must(template.New("t").ParseGlob(filepath.Join(".", "templates", "*.html")))
}

func ArtistsPageHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		ArtistsPageGetHandler(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func ArtistsPageGetHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/artists" {
		http.Error(w, "Artist page error. Go back to the main page", 404)
		return
	}
	artists, err := GetAllArtists()
	if err != nil {
		http.Error(w, "Return to main page", 400)
		return
	}
	js, err2 := json.Marshal(artists)
	if err2 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func ArtistPageHandler(w http.ResponseWriter, req *http.Request) {
	if len(req.URL.Path) <= len("/artist/") {
		http.Error(w, "Artist ID is missing, return to main page", 400)
		return
	}
	_id := req.URL.Path[len("/artist/"):]
	id, errID := strconv.Atoi(_id)
	if errID != nil {
		http.Error(w, "There is no such artist, return to main page", 400)
		return
	}
	if id < 1 || id > 52 {
		http.Error(w, "There is no such artist, Return to main page", 404)
		return
	}

	artist, err := GetArtist(_id)
	if err != nil {
		http.Error(w, "Go back to the main page", 500)
		return
	}
	relation, err2 := GetRelation(_id)
	if err2 != nil {
		http.Error(w, "Go back to the main page", 500)
		return
	}
	data := struct {
		Artist   Artist
		Relation SubRelation
	}{artist, relation}
	js, err2 := json.Marshal(data)
	if err2 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func ShowArtistsHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		ShowArtistsPageGetHandler(w, req)
	case "POST":
		SearchResultPageHandler(w, req)
	default:
		http.Error(w, "Only GET/POST method allowed, return to main page", 405)
	}
}

func ShowArtistsPageGetHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.Error(w, "Main page error.Go back to the main page", 404)
		return
	}
	w.Header().Set("Content-Type", "text/html")

	err := templates.ExecuteTemplate(w, "artists.html", nil)
	if err != nil {
		http.Error(w, "Go back to the main page", 500)
		return
	}

}

func RelationPageHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/relations" {
		http.Error(w, "Relations page error, Go back to the main page", 404)
		return
	}
	relations, err := GetAllRelations()
	if err != nil {
		http.Error(w, "Return to main page", 400)
		return
	}
	allRelations := parseRelations(relations)
	js, err2 := json.Marshal(allRelations)
	if err2 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func ShowRelationHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		ShowRelationPageGetHandler(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func ShowRelationPageGetHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/relation" {
		http.Error(w, "Relation page error.Go back to the main page", 404)
		return
	}
	w.Header().Set("Content-Type", "text/html")

	err := templates.ExecuteTemplate(w, "relation.html", nil)
	if err != nil {
		http.Error(w, "Go back to the main page", 500)
		return
	}
}
