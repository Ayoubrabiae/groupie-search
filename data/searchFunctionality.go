package data

import (
	"strconv"
	"strings"
)

func InsertArtists(trie ArtistsTrie, artists []ArtistType) {
	for _, ar := range artists {
		artistName := strings.TrimSpace(strings.ToLower(ar.Name))

		trie.Insert(nil, artistName, ar.Id, "artist", ar.Name)
		trie.Insert(nil, ar.FirstAlbum, ar.Id, "first album", ar.Name)
		trie.Insert(nil, strconv.Itoa(ar.CreationDate), ar.Id, "creation date", ar.Name)
		for _, m := range ar.Members {
			memberName := strings.TrimSpace(strings.ToLower(m))

			trie.Insert(nil, memberName, ar.Id, "member", ar.Name)
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
