package data

import (
	"strconv"
	"strings"

	"groupie-tracker/funcs"
)

func InsertArtists(trie ArtistsTrie, artists []ArtistType) {
	for _, ar := range artists {
		artistName := strings.TrimSpace(strings.ToLower(ar.Name))

		trie.Insert(nil, artistName, ar.Id, "artist", ar.Name)
		trie.Insert(nil, ar.FirstAlbum, ar.Id, "first album", ar.Name)
		trie.Insert(nil, strconv.Itoa(ar.CreationDate), ar.Id, "creation date", ar.Name)
		for _, m := range ar.Members {
			memberName := funcs.RemoveSpace(m)

			trie.Insert(nil, memberName, ar.Id, "member", ar.Name)
		}
	}
}

func InsertLocations(trie ArtistsTrie, locationStruct []LocationsType, artists []string) {
	for _, locations := range locationStruct {
		for _, loc := range locations.Locations {
			trie.Insert(nil, funcs.RemoveSpace(loc), locations.Id, "location", artists[locations.Id-1])
		}
	}
}

func Search(artists []ArtistType, value string) []ArtistType {
	suggetions := SearchTrie.Suggest(nil, funcs.RemoveSpace(value))

	ids := map[int]bool{}
	for _, s := range suggetions {
		for _, artist := range s.Data {
			ids[artist.Id] = true
		}
	}

	res := []ArtistType{}

	for _, artist := range artists {
		if ids[artist.Id] {
			res = append(res, artist)
		}
	}

	return res
}
