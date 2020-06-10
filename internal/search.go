package internal

import (
	//"fmt"
	"net/http"
	"strconv"
)

var (
	resultSearchText = map[string]string{
		"album":    "Band released their first album in ",
		"member":   " is member of band: ",
		"creation": "Band is created in ",
		"location": "Band's concert location in ",
		"date":     "Band's concert date will be in ",
	}
	artists   []Artist
	relations Relation
)

func SearchPageHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		SearchPageGetHandler(w, req)
	case "POST":
		SearchResultPageHandler(w, req)
	default:
		http.Error(w, "Only GET/POST method allowed, return to main page", 405)
	}
}

func SearchPageGetHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/search" {
		http.Error(w, "Search page error, go back to the main page", 404)
		return
	}
	artistsInfo, errArtists := GetAllArtists()
	if errArtists != nil {
		http.Error(w, "Go back to the main page", 400)
		return
	}
	relationsInfo, errRelations := GetAllRelations()
	if errRelations != nil {
		http.Error(w, "Go back to the main page", 400)
		return
	}
	searchText := &Search{Artists: artistsInfo, Relations: relationsInfo}
	err := templates.ExecuteTemplate(w, "search.html", searchText)
	if err != nil {
		http.Error(w, "Go back to the main page", 500)
		return
	}
}

func SearchResultPageHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		SearchResultPostPageHandler(w, req)
	default:
		http.Error(w, "Only POST method allowed, return to main page", 405)
	}
}

func SearchResultPostPageHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/search_result" {
		http.Error(w, "SearchResuult page error, go back to the main page", 404)
	}
	if errForm := req.ParseForm(); errForm != nil {
		http.Error(w, "Go back to the main page", 400)
		return
	}
	text := req.Form.Get("searchText")
	choice := req.Form.Get("searchChoice")
	artistsInfo, errArtists := GetAllArtists()
	if errArtists != nil {
		http.Error(w, "Go back to the main page", 400)
		return
	}
	relationsInfo, errRelations := GetAllRelations()
	if errRelations != nil {
		http.Error(w, "Go back to the main page", 400)
		return
	}
	isFound := false
	artistIndexArr := []int{}
	cdate, _ := strconv.Atoi(text)
	creationDate := float32(cdate)
	if choice == "album" || choice == "member" || choice == "band" || choice == "creation" {
		for index, artist := range artistsInfo {
			if choice == "member" && hasWord(artist.Members, text) {
				isFound = true
				artistIndexArr = append(artistIndexArr, index)
			} else if choice == "album" && artist.FirstAlbum == text {
				isFound = true
				artistIndexArr = append(artistIndexArr, index)
			} else if choice == "band" && artist.Name == text {
				isFound = true
				redirectArtistID(w, req, artist.ID)
			} else if choice == "creation" && artist.CreationDate == creationDate {
				isFound = true
				artistIndexArr = append(artistIndexArr, index)
			}
		}
	} else if choice == "location" {
		for index, subrelation := range relationsInfo.Index {
			_, found := subrelation.DatesLocations[text]
			if found {
				isFound = true
				artistIndexArr = append(artistIndexArr, index)
			}
		}
	} else if choice == "date" {
		for index, subrelation := range relationsInfo.Index {
			for _, dates := range subrelation.DatesLocations {
				if hasWord(dates, text) {
					isFound = true
					artistIndexArr = append(artistIndexArr, index)
				}
			}
		}
	} else {
		http.Error(w, "Go back to main page", 400)
		return
	}
	output := ""
	if choice == "member" {
		output = text + resultSearchText[choice]
	} else {
		output = resultSearchText[choice] + text + ": "
	}
	for index, value := range artistIndexArr {
		if index != len(artistIndexArr)-1 {
			output += artistsInfo[value].Name + ";" + "\n"
		} else {
			output += artistsInfo[value].Name
		}
	}
	if !isFound {
		output = "No results for you search"
	}
	err := templates.ExecuteTemplate(w, "result_search.html", struct{ Info string }{output})
	if err != nil {
		http.Error(w, "Go back to the main page", 500)
		return
	}
}

func redirectArtistID(w http.ResponseWriter, req *http.Request, id float32) {
	http.Redirect(w, req, "/artist_info/"+strconv.Itoa(int(id)), 302)
}

func ArtistInfoPageHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		ArtistInfoGetPage(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func ArtistInfoGetPage(w http.ResponseWriter, req *http.Request) {
	if len(req.URL.Path) <= len("/artist_info/") {
		http.Error(w, "Artist ID is missing", 404)
		return
	}

	id, _ := strconv.Atoi(req.URL.Path[len("/artist_info/"):])
	if id < 1 && id > 52 {
		http.Error(w, "There is no such artist, Return to main page", 404)
		return
	}
	artistsInfo, errArtists := GetAllArtists()
	if errArtists != nil {
		http.Error(w, "Go back to the main page", 400)
		return
	}
	relationsInfo, errRelations := GetRelation(strconv.Itoa(id))
	if errRelations != nil {
		http.Error(w, "Go back to the main page", 400)
		return
	}
	index := id - 1
	artist := artistsInfo[index]
	render := &ArtistInfo{ID: float32(id),
		CreationDate:  artist.CreationDate,
		FirstAlbum:    artist.FirstAlbum,
		Image:         artist.Image,
		Members:       artist.Members,
		Name:          artist.Name,
		LocationDates: relationsInfo,
	}
	err := templates.ExecuteTemplate(w, "artist_info.html", render)
	if err != nil {
		http.Error(w, "Go back to the main page", 500)
		return
	}
}
