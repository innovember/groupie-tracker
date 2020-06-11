package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	internal "../internal"
)

func Execute() {

	// API handlers
	http.HandleFunc("/artists", internal.ArtistsPageHandler)
	http.HandleFunc("/artist/", internal.ArtistPageHandler)
	http.HandleFunc("/relations", internal.RelationPageHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	// User pages handlers
	http.HandleFunc("/", internal.ShowArtistsHandler)
	http.HandleFunc("/relation", internal.ShowRelationHandler)

	// Search page handlers
	http.HandleFunc("/search", internal.SearchPageHandler)
	http.HandleFunc("/search_result", internal.SearchResultPageHandler)
	http.HandleFunc("/artist_info/", internal.ArtistInfoPageHandler)

	// Map page handler
	http.HandleFunc("/map", internal.MapPageHandler)
	fmt.Println("Server is listening port 8181...")

	err := http.ListenAndServe(getPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8181"
	}
	return ":" + port
}
