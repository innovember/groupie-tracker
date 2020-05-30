package cmd

import (
	internal "../internal"
	"fmt"
	"log"
	"net/http"
	"os"
)

func Execute() {
	http.HandleFunc("/artists", internal.ArtistsPageHandler)
	http.HandleFunc("/artist", internal.ArtistPageHandler)

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
