package search

import (
	"encoding/json"
	"fmt"
	"os"
)

const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeed() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var feeds []*Feed
	errs := json.NewDecoder(file).Decode(&feeds)

	return feeds, errs
}
