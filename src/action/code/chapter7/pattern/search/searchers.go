package search

import (
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type google struct{}

type bing struct{}

type yahoo struct{}

// 注意，这里使用的是值作为接收者，会新建一个type
func (g google) Search(searchTerm string, searchResult chan<- []Result) {
	log.Printf("Google: Search: Started: Search Term: %s \n", searchTerm)

	var results []Result
	time.Sleep(time.Microsecond * time.Duration(rand.Int63n(900)))
	results = append(results, Result{
		Engine:      "Google",
		Title:       "The Go Programming Language",
		Description: "The Go Programming Language",
		Link:        "www.google.com",
	})

	log.Printf("Google: Search: Completed: Found %d \n", len(results))
	searchResult <- results
}

func (b bing) Search(searchTerm string, searchResult chan<- []Result) {
	log.Printf("Bing: Search: Started: Search Term: %s \n", searchTerm)

	var results []Result
	time.Sleep(time.Microsecond * time.Duration(rand.Int63n(900)))
	results = append(results, Result{
		Engine:      "Bing",
		Title:       "A Tour of Go",
		Description: "Welcome to a tour of the Go programming language",
		Link:        "www.bing.com",
	})

	log.Printf("Bing: Search: Completed: Found %d \n", len(results))
	searchResult <- results
}

func (y yahoo) Search(searchTerm string, searchResult chan<- []Result) {
	log.Printf("Yahoo: Search: Started: Search Term: %s \n", searchTerm)

	var results []Result
	time.Sleep(time.Microsecond * time.Duration(rand.Int63n(900)))
	results = append(results, Result{
		Engine:      "Yahoo",
		Title:       "Go Playground",
		Description: "The Go Playground is a web service that runs on golang.org's servers",
		Link:        "www.yahoo.com",
	})

	log.Printf("Yahoo: Search: Completed: Found %d \n", len(results))
	searchResult <- results
}
