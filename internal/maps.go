package internal

import (
	"net/http"
)

func MapPageHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		MapPostPageHandler(w, req)
	default:
		http.Error(w, "Only POST method allowed, return to main page", 405)
	}
}

func MapPostPageHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/map" {
		http.Error(w, "MapPage page error, go back to the main page", 404)
	}
	if errForm := req.ParseForm(); errForm != nil {
		http.Error(w, "Go back to the main page", 400)
		return
	}
	id := req.Form.Get("map-data")
	relation, errRelation := GetRelation(id)
	if errRelation != nil {
		http.Error(w, "Relation API error, return to main page", 400)
		return
	}
	locations := []string{}
	for key := range relation.DatesLocations {
		locations = append(locations, formatLocation(key))
	}
	locationStruct := &MapData{Count: len(locations), LocationsArr: locations}
	err := templates.ExecuteTemplate(w, "mapresult.html", locationStruct)
	if err != nil {
		http.Error(w, "Go back to the main page", 500)
		return
	}
}
