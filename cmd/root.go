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
	fmt.Println("Server is listening...")

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
