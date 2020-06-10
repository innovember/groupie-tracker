package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

func GetAllRelations() (Relation, error) {
	RelationAPI := "https://groupietrackers.herokuapp.com/api/relation"
	output, err := GetJsonFromAPI(RelationAPI)
	var relations Relation
	if err != nil {
		return relations, err
	}
	errJSON := json.Unmarshal(output, &relations)
	if errJSON != nil {
		return relations, errJSON
	}
	return relations, nil
}

func parseRelations(relations Relation) AllRelationsData {
	var allRelations = make([]AllRelations, 333)
	j := 0
	for i := range relations.Index {
		id := relations.Index[i].ID
		valueStr := ""
		item := make(map[string]string)
		for key, value := range relations.Index[i].DatesLocations {
			for i := range value {
				valueStr += formatDate(value[i])
				if len(value) > 1 && i != len(value)-1 {
					valueStr += " | "
				}
			}
			item[key] = valueStr
			allRelations[j].ID = id
			allRelations[j].Location = key
			allRelations[j].Date = item[key]
			j++
			item = map[string]string{}
			valueStr = ""
		}
	}
	var allRelationsData = AllRelationsData{
		Data:            allRelations,
		Draw:            1,
		RecordsTotal:    len(allRelations),
		RecordsFiltered: len(allRelations),
	}
	return allRelationsData
}

func parseMapToStr(datesLocations map[string][]string) map[string]string {
	item := make(map[string]string)
	valueStr := ""
	for key, value := range datesLocations {
		for i := range value {
			valueStr += formatDate(value[i]) + "\n"
		}
		item[formatLocation(key)] = valueStr
	}
	return item
}

func formatLocation(str string) string {
	str = strings.ReplaceAll(str, "-", ", ")
	str = strings.ReplaceAll(str, "_", " ")
	str = strings.Title(str)
	return str
}

func formatDate(str string) string {
	return strings.ReplaceAll(str, "-", ".")
}

func hasWord(arr []string, word string) bool {
	for _, value := range arr {
		if value == word {
			return true
		}
	}
	return false
}
