package data

import (
	"strconv"
	"strings"
)

func InsertArtists(trie ArtistsTrie, artists []ArtistType) {
	for _, ar := range artists {
		trie.Insert(nil, strings.ToLower(ar.Name), ar.Id, "artist", ar.Name)
		trie.Insert(nil, strings.ToLower(ar.FirstAlbum), ar.Id, "first album", ar.Name)
		trie.Insert(nil, strings.ToLower(strconv.Itoa(ar.CreationDate)), ar.Id, "creation date", ar.Name)
		for _, m := range ar.Members {
			trie.Insert(nil, strings.ToLower(m), ar.Id, "member", ar.Name)
		}
	}
}

func InsertLocations(trie ArtistsTrie, locationStruct []LocationsType, artists []string) {
	for _, locations := range locationStruct {
		for _, loc := range locations.Locations {
			trie.Insert(nil, strings.ToLower(loc), locations.Id, "location", artists[locations.Id-1])
		}
	}
}
