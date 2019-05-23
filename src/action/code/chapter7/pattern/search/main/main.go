package main

import (
	"../../search"
	"log"
)

func main() {
	results := search.Submit("golang", search.OnlyFirst, search.Google, search.Bing, search.Yahoo)
	for _, result := range results {
		log.Printf("main: Results: Info %v \n", result)
	}

	results = search.Submit("golang", search.Google, search.Bing, search.Yahoo)
	for _, result := range results {
		log.Printf("main: Results: Info %v \n", result)
	}
}
