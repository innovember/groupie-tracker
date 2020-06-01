package internal

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

var (
	templates *template.Template
)

func init() {
	templates = template.Must(template.ParseGlob("./templates/*.html"))
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
		http.Error(w, "Go back to the main page", 404)
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

	// Uncomment it when index.html will be ready , to render data to html
	// errTemplate := templates.ExecuteTemplate(w, "index.html", artists)
	// if errTemplate != nil {
	// 	http.Error(w, "Go back to the main page", 500)
	// 	return
	// }
}

func ArtistPageHandler(w http.ResponseWriter, req *http.Request) {
	id, errID := GetQueryID(w, req)
	if errID != nil {
		http.Error(w, errID.Error(), 400)
		return
	}
	if id < 1 || id > 52 {
		http.Error(w, "There is no such group, Return to main page", 404)
		return
	}
	_id := strconv.Itoa(id)
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
