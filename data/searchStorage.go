package data

type SearchStorageType struct {
	Artists   []ArtistType
	Locations LocationsType
	Relations RelationsType
	Dates     DatesType
}

type SearchItem struct {
	Id   int
	Name string
	Kind string
}

type SearchResult struct {
	Result []SearchItem
}

var SearchStorage SearchStorageType
