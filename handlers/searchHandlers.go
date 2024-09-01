package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"groupie-tracker/data"
)

func SuggestSearch(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Sec-Fetch-Mode") != "cors" || r.Header.Get("Sec-Fetch-Site") != "same-origin" {
		ErrorHandler(w, "Access Denied", http.StatusNotFound)
	}
	val := r.URL.Query().Get("q")

	suggestions := data.SearchTrie.Suggest(nil, strings.ToLower(val))

	data, err := json.Marshal(suggestions)
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query().Get("s")
	suggestions := data.SearchTrie.Suggest(nil, strings.ToLower(val))

	fmt.Println(suggestions)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
