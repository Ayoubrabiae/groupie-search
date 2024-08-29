package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"groupie-tracker/data"
)

func SuggestSearch(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query().Get("q")

	suggestions := data.SearchTrie.Suggest(nil, strings.ToLower(val))

	data, err := json.Marshal(suggestions)
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
