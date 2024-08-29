package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"groupie-tracker/data"
	"groupie-tracker/funcs"
)

func checkErrors(w http.ResponseWriter, r *http.Request) bool {
	if r.URL.Path != "/" {
		ErrorHandler(w, "Page Not Found", http.StatusNotFound)
		return false
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return false
	}

	return true
}

func parseFiles(w http.ResponseWriter) (*template.Template, bool) {
	tmp, err := template.ParseFiles("./pages/index.html")
	if err != nil {
		ErrorHandler(w, "Internal Server error", http.StatusInternalServerError)
		fmt.Println("When we parse the index.html")
		return nil, false
	}

	tmp, err = tmp.ParseGlob("./templates/*.html")
	if err != nil {
		ErrorHandler(w, "Internal Server error", http.StatusInternalServerError)
		fmt.Println("When we parse all templates")
		return nil, false
	}

	return tmp, true
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if !checkErrors(w, r) {
		return
	}

	tmp, ok := parseFiles(w)

	if !ok {
		return
	}

	var artists []data.ArtistType

	err := funcs.GetAndParse(data.MainData.Artists, &artists)
	if err != nil {
		ErrorHandler(w, "Internal server Error", http.StatusInternalServerError)
		return
	}

	data.InsertArtists(data.SearchTrie, artists)

	filterParams, err := data.GetFilterParams(artists, r.URL.Query())
	if err != nil {
		ErrorHandler(w, "Internal server Error", http.StatusInternalServerError)
		return
	}

	if len(r.URL.Query()) != 0 {
		artists = data.FilterArtists(artists, r.URL.Query())
	}

	homeData := struct {
		Artists []data.ArtistType
		Filter  data.FilterType
	}{
		Artists: artists,
		Filter:  filterParams,
	}

	err = tmp.Execute(w, homeData)
	if err != nil {
		fmt.Println("When we excute the html", err)
		return
	}
}
