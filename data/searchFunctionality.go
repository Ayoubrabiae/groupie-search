package data

import (
	"groupie-tracker/funcs"
	"strconv"
	"strings"
)

func FillSearchStorage() error {
	err := funcs.GetAndParse(MainData.Artists, &SearchStorage.Artists)
	if err != nil {
		return err
	}

	err = funcs.GetAndParse(MainData.Locations, &SearchStorage.Locations)
	if err != nil {
		return err
	}

	return nil
}

func (s *SearchStorageType) Suggest(value string) SearchResult {
	res := SearchResult{}

	for _, item := range s.Artists {
		// Bands
		if strings.Contains(funcs.HandleString(item.Name), funcs.HandleString(value)) {
			res.Result = append(res.Result, SearchItem{Id: item.Id, Name: item.Name, Kind: "band", Value: item.Name})
		}

		// Members
		for _, member := range item.Members {
			if strings.Contains(funcs.HandleString(member), funcs.HandleString(value)) {
				res.Result = append(res.Result, SearchItem{Id: item.Id, Name: item.Name, Kind: "member", Value: member})
			}
		}

		// First Album
		if strings.Contains(funcs.HandleString(item.FirstAlbum), funcs.HandleString(value)) {
			res.Result = append(res.Result, SearchItem{Id: item.Id, Name: item.Name, Kind: "first-album", Value: item.FirstAlbum})
		}

		// Creation Date
		if strings.Contains(funcs.HandleString(strconv.Itoa(item.CreationDate)), funcs.HandleString(value)) {
			res.Result = append(res.Result, SearchItem{Id: item.Id, Name: item.Name, Kind: "creation-date", Value: strconv.Itoa(item.CreationDate)})
		}
	}

	// Locations
	for _, item := range s.Locations.Index {
		for _, location := range item.Locations {
			if strings.Contains(funcs.HandleString(location), funcs.HandleString(value)) {
				res.Result = append(res.Result, SearchItem{Id: item.Id, Name: s.Artists[item.Id-1].Name, Value: location, Kind: "location"})
			}
		}
	}

	return res
}

func (s *SearchStorageType) Search(artists []ArtistType, value string) []ArtistType {
	res := []ArtistType{}
	suggestions := s.Suggest(value)
	searchMap := map[int]bool{}

	for _, sug := range suggestions.Result {
		searchMap[sug.Id] = true
	}

	for _, artist := range artists {
		if searchMap[artist.Id] {
			res = append(res, artist)
		}
	}

	return artists
}
