package search

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

var matchers = make(map[string]Matcher)

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

func Run(searchTerm string) {
	feeds, err := RetriveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	results := make(chan *Result)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exits := matchers[feed.Type]
		if !exits {
			matcher = matchers["default"]
		}
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
		}(matcher, feed)
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()

	Display(results)

}
