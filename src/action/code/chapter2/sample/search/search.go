package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

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

func Register(feetType string, matcher Matcher) {
	if _, exists := matchers[feetType]; exists {
		log.Fatalln(feetType, "Matcher already registered")
	}

	log.Println("Register", feetType, "matcher")
	matchers[feetType] = matcher
}
