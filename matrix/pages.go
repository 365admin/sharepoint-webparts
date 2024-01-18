package matrix

import (
	"encoding/json"
	"io"
	"os"
	"path"

	"github.com/spf13/viper"
)

type Page struct {
	SortOrder  string `json:"SortOrder"`
	ValueChain string `json:"ValueChain"`
	Title      string `json:"Title"`
	FileRef    string `json:"FileRef"`
	ID         int    `json:"ID"`
}

func ReadPages() ([]Page, error) {
	// Read the places from a json file
	filePath := path.Join(viper.GetString("WORKDIR"), "matrix-pages.json")
	jsonFile, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	// defer the closing of our jsonFile so that we can parse it later on

	pages := []Page{}

	json.Unmarshal([]byte(data), &pages)
	defer jsonFile.Close()
	return pages, nil

}

type Row struct {
	SortOrder string `json:"SortOrder"`
	Title     string `json:"Title"`
	FileRef   string `json:"FileRef"`
	ID        int    `json:"ID"`
}
type Column struct {
	SortOrder int    `json:"SortOrder"`
	Title     string `json:"Title"`
	Rows      []Row  `json:"Rows"`
}

func AnalysePages() (map[string][]Row, error) {
	pages, err := ReadPages()
	if err != nil {
		return nil, err
	}

	//columns := []Column{}

	columns := map[string][]Row{}

	for _, page := range pages {
		sortOrder := page.SortOrder
		if err != nil {
			sortOrder = ""
		}
		if page.ValueChain == "" {
			continue
		}

		if (columns[page.ValueChain]) != nil {
			columns[page.ValueChain] = append(columns[page.ValueChain], Row{Title: page.Title, FileRef: page.FileRef, ID: page.ID, SortOrder: sortOrder})
		} else {
			columns[page.ValueChain] = []Row{{Title: page.Title, FileRef: page.FileRef, ID: page.ID, SortOrder: sortOrder}}

		}

	}
	return columns, nil
}
