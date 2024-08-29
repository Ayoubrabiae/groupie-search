package data

import (
	"slices"
)

type ArtistNodeData struct {
	Id   int
	Kind string
	Name string
}

type ArtistNode struct {
	Childs map[byte]*ArtistNode
	Data   []ArtistNodeData
	End    bool
}

type ArtistsTrie struct {
	Root *ArtistNode
}

type ArtistLeaf struct {
	Data  []ArtistNodeData
	Value string
}

func NewTrie() *ArtistsTrie {
	return &ArtistsTrie{
		Root: &ArtistNode{
			Childs: make(map[byte]*ArtistNode),
		},
	}
}

func (t *ArtistsTrie) Insert(node *ArtistNode, word string, id int, kind string, name string) {
	if word == "" {
		node.End = true
		if !(slices.Contains(node.Data, ArtistNodeData{id, kind, name})) {
			node.Data = append(node.Data, ArtistNodeData{id, kind, name})
		}
		return
	}

	if node == nil {
		node = t.Root
	}

	if _, ok := node.Childs[word[0]]; !ok {
		node.Childs[word[0]] = &ArtistNode{Childs: make(map[byte]*ArtistNode)}
	}
	t.Insert(node.Childs[word[0]], word[1:], id, kind, name)
}

func search(node ArtistNode, str string, arr *[]ArtistLeaf) {
	for char := range node.Childs {
		search(*node.Childs[char], str+string(char), arr)
	}
	if node.End {
		*arr = append(*arr, ArtistLeaf{Data: node.Data, Value: str})
	}
}

func (t *ArtistsTrie) Collect(node *ArtistNode) []ArtistLeaf {
	arr := []ArtistLeaf{}

	if node == nil {
		node = t.Root
	}

	search(*node, "", &arr)

	return arr
}

func (t *ArtistsTrie) Suggest(node *ArtistNode, str string) []ArtistLeaf {
	if node == nil {
		node = t.Root
	}

	if str == "" {
		return t.Collect(node)
	} else if node.Childs[str[0]] != nil {
		return t.Suggest(node.Childs[str[0]], str[1:])
	}
	return []ArtistLeaf{}
}

var SearchTrie ArtistsTrie = *NewTrie()
