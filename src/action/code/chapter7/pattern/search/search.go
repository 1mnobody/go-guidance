package search

import "log"

type Result struct {
	Engine      string
	Title       string
	Description string
	Link        string
}

type Searcher interface {
	Search(searchTerm string, searchResult chan<- []Result)
}

type searchSession struct {
	searchers  map[string]Searcher
	first      bool
	resultChan chan []Result
}

func Google(s *searchSession) {
	log.Println("search: Submit: Info: Adding Google")
	s.searchers["google"] = google{}
}

func Bing(s *searchSession) {
	log.Println("search: Submit: Info: Adding Bing")
	s.searchers["bing"] = bing{}
}

func Yahoo(s *searchSession) {
	log.Println("search: Submit: Info: Adding Yahoo")
	s.searchers["yahoo"] = yahoo{}
}

func OnlyFirst(s *searchSession) {
	s.first = true
}

// 接收一个函数列表，这些函数都是对 session进行初始化的
func Submit(query string, options ...func(session *searchSession)) []Result {
	var session = new(searchSession)
	session.searchers = make(map[string]Searcher)
	session.resultChan = make(chan []Result)
	for _, opt := range options {
		opt(session)
	}

	for _, s := range session.searchers {
		go s.Search(query, session.resultChan)
	}

	var results []Result
	for search := 0; search < len(session.searchers); search++ {
		// first 为true时，只关注第一个返回的信息
		if session.first && search > 0 {
			go func() {
				r := <-session.resultChan
				log.Printf("search: Submit: Info: Result Discarded: Results %d\n", len(r))
			}()
			continue
		}

		log.Println("search: Submit: Info: Waiting For Results...")
		res := <-session.resultChan
		log.Printf("search : Submit : Info : Results Used : Results[%d]\n", len(res))
		results = append(results, res...)
	}

	log.Printf("search : Submit : Completed : Found [%d] Results\n", len(results))
	return results
}
