package search

import "log"

type Result struct {
	Field   string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// 参数中的results chan <- *Result  表示results 是一个 只写*Result 的channel（不能读数据，即使用 <-results 会报错。），
// 同理，还可以定义一个 只读*Result 的 channel : results <-chan *Result，这个results 是只读的，只能从中读数据，不能写数据
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println("err:", err)
		return
	}
	for _, result := range searchResults {
		results <- result
	}
}

func Display(results chan *Result) {
	for result := range results {
		log.Printf("%s:\n %s\n", result.Field, result.Content)
	}
}
