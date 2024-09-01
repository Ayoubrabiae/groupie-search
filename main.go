package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/data"
	"groupie-tracker/funcs"
	"groupie-tracker/handlers"
)

func main() {
	port := "8082"

	err := funcs.GetAndParse("https://groupietrackers.herokuapp.com/api", &data.MainData)
	if err != nil {
		fmt.Println(err)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			handlers.ErrorHandler(w, "Internal Server Error!", http.StatusInternalServerError)
		})
		log.Fatal(http.ListenAndServe(":"+port, http.DefaultServeMux))
		return
	}

	var artists []data.ArtistType

	err = funcs.GetAndParse(data.MainData.Artists, &artists)
	if err != nil {
		fmt.Println(err)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			handlers.ErrorHandler(w, "Internal Server Error!", http.StatusInternalServerError)
		})
		log.Fatal(http.ListenAndServe(":"+port, http.DefaultServeMux))
		return
	}

	err = handlers.IntializeTrie(artists)
	if err != nil {
		fmt.Println(err)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			handlers.ErrorHandler(w, "Internal Server Error!", http.StatusInternalServerError)
		})
		log.Fatal(http.ListenAndServe(":"+port, http.DefaultServeMux))
		return
	}

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/static/", http.StripPrefix("/static", handlers.StaticHandler(fileServer)))
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artists/", handlers.ProfileHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/locations/", handlers.LocationsHandler)
	http.HandleFunc("/suggest-search", handlers.SuggestSearch)
	http.HandleFunc("/search", handlers.SuggestSearch)
	fmt.Println("http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, http.DefaultServeMux))
}
