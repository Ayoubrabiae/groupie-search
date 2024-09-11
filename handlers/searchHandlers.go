package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"groupie-tracker/data"
	"groupie-tracker/funcs"
)

func SuggestSearch(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Sec-Fetch-Mode") != "cors" || r.Header.Get("Sec-Fetch-Site") != "same-origin" {
		ErrorHandler(w, "Access Denied", http.StatusNotFound)
	}
	val := r.URL.Query().Get("q")

	if funcs.IsSpace(val) {
		return
	}

	suggestions := data.SearchTrie.Suggest(nil, strings.ToLower(val))

	data, err := json.Marshal(suggestions)
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
