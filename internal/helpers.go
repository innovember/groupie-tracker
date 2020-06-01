package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//API urls
//ArtistsAPI := "https://groupietrackers.herokuapp.com/api/artists"
// LocationsAPI := "https://groupietrackers.herokuapp.com/api/locations"
// DatesAPI := "https://groupietrackers.herokuapp.com/api/dates"
// RelationAPI := "https://groupietrackers.herokuapp.com/api/relation"

func GetAllArtists() ([]Artist, error) {
	ArtistsAPI := "https://groupietrackers.herokuapp.com/api/artists"
	// Get Artist structure from API
	artists, errArtists := GetArtists(ArtistsAPI)
	if errArtists != nil {
		fmt.Println(errArtists.Error())
		return nil, errArtists
	}
	// // Get Locations structure from API
	// locations, errLocations := GetLocations(LocationsAPI)
	// if errLocations != nil {
	// 	fmt.Println(errArtists.Error())
	// 	return nil, errLocations
	// }
	// // Get Dates structure from API
	// dates, errDates := GetDates(DatesAPI)
	// if errDates != nil {
	// 	fmt.Println(errArtists.Error())
	// 	return nil, errDates
	// }
	// // Get Relation structure from API
	// relation, errRelation := GetRelation(RelationAPI)
	// if errRelation != nil {
	// 	fmt.Println(errArtists.Error())
	// 	return nil, errRelation
	// }
	//
	return artists, nil
}

func GetArtists(url string) ([]Artist, error) {
	output, err := GetJsonFromAPI(url)
	if err != nil {
		return nil, err
	}
	artists := make([]Artist, 0)
	errJSON := json.Unmarshal(output, &artists)
	if errJSON != nil {
		return nil, errJSON
	}
	return artists, nil
}

// func GetLocations(url string) ([]Artist, error) {
// 	// return locations, nil
// 	return nil, nil
// }

// func GetDates(url string) ([]Artist, error) {
// 	// return dates, nil
// 	return nil, nil
// }

// func GetRelation(url string) ([]Artist, error) {
// 	// return relation, nil
// 	return nil, nil
// }

func GetJsonFromAPI(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	}
	defer response.Body.Close()
	output, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		fmt.Printf("%s", err)
		return nil, err
	}
	return output, nil
}

func GetArtist(id string) (Artist, error) {
	var artist Artist
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%s", id)
	body, err := GetJsonFromAPI(url)
	if err != nil {
		return artist, err
	}
	err2 := json.Unmarshal(body, &artist)
	if err2 != nil {
		return artist, err2
	}
	return artist, nil
}

func GetRelation(id string) (SubRelation, error) {
	var subRelation SubRelation
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%s", id)

	body, err := GetJsonFromAPI(url)
	if err != nil {
		return subRelation, err
	}
	err2 := json.Unmarshal(body, &subRelation)
	if err2 != nil {
		return subRelation, err2
	}
	return subRelation, nil
}

func GetQueryID(w http.ResponseWriter, req *http.Request) (int, error) {
	keys, exist := req.URL.Query()["id"]
	if !exist || len(keys) != 1 {
		return 0, errors.New("URL Param 'id' is missing")
	}
	key := keys[0]
	id, err := strconv.Atoi(key)
	if err != nil {
		return 0, err
	}
	return id, nil
}
