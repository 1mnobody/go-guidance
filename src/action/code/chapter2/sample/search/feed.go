package search

import (
	"encoding/json"
	"os"
)

// 先使用绝对路径
const dataFile = "E:\\work\\go\\idea\\go-guidance\\src\\action\\code\\chapter2\\sample\\data\\data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetriveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}
