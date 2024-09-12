package data

type SearchStorageType struct {
	Artists   []ArtistType
	Locations struct {
		Index []LocationsType `json:"index"`
	}
}

type SearchItem struct {
	Id    int
	Name  string
	Value string
	Kind  string
}

type SearchResult struct {
	Result []SearchItem
}

var SearchStorage SearchStorageType
