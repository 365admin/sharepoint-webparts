package matrix

import (
	"encoding/json"
	"io"
	"os"
	"path"

	"github.com/365admin/sharepoint-webparts/.app/sharedcommands"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Page struct {
	SortOrder  any    `json:"SortOrder"`
	ValueChain any    `json:"ValueChain"`
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
	SortOrder int    `json:"SortOrder"`
	Title     string `json:"Title"`
	FileRef   string `json:"FileRef"`
	ID        int    `json:"ID"`
}
type Column struct {
	SortOrder int    `json:"SortOrder"`
	Title     string `json:"Title"`
	Rows      []Row  `json:"Rows"`
}

func AnalysePages() ([]Column, error) {
	pages, err := ReadPages()
	if err != nil {
		return nil, err
	}

	columns := []Column{}

	for _, page := range pages {
		sortOrder := page.SortOrder.(int)
		foundColumn := false
		for _, column := range columns {
			if column.Title == page.Title {
				foundColumn = true

				column.Rows = append(column.Rows, Row{Title: page.Title, FileRef: page.FileRef, ID: page.ID, SortOrder: sortOrder})
			}

		}
		if !foundColumn {
			columns = append(columns, Column{Title: page.Title, Rows: []Row{{Title: page.Title, FileRef: page.FileRef, ID: page.ID, SortOrder: sortOrder}}})
		}
	}
	return columns, nil
}

func init() {

	sharedcommands.JobsCmd.AddCommand(&cobra.Command{
		Use:   "matrix",
		Short: "Generate matrix",
		//Args:  cobra.MinimumNArgs(1),
		Long: ``,

		Run: func(cmd *cobra.Command, args []string) {
			AnalysePages()
			// argument := args[0]
			// switch argument {
			// case "places":

			// default:
			// 	log.Fatalf("Unknown argument %s", argument)
			// }
		},
	})
}
