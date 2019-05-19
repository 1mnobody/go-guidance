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
			// Match 会查找数据，将结果数据发送到results中，Display会一直读取results中的数据
			Match(matcher, feed, searchTerm, results)
			// Done，减少waitGroup的计数器，为0时，waitGroup会从Wait()处醒来继续执行
			waitGroup.Done()
		}(matcher, feed)
	}

	go func() {
		// 等待所有的 Matcher 处理完毕
		waitGroup.Wait()
		// 关闭results，结束Display中的 for range 循环（没有close，for range会一直执行）
		//close(results)
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
