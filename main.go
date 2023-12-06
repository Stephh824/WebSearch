package main

import (
	"sync"
	"database/sql"
	//"fmt"
	"time"
)

var stopwords map[string]struct{}

type IIndex struct {
	index        map[string]map[string]int // word: url: freq
	dbase     *sql.DB // database
	rules     []RobotsTxt // holds policies from robots.txt
	sitemap string
	mutex     sync.Mutex
}

func indexInIt() IIndex {
	stopwords = stopWordsInit()
	return IIndex{
		index:        make(map[string]map[string]int),
		mutex:        sync.Mutex{}}
}

func main() {
	idx := indexInIt()
	idx.Open() // remvoe after
	dservers(&idx)
	//dcrawl("https://openai.com/robots.txt", &idx)
	//dcrawl("https://openai.com/customer-stories/duolingo", &idx) // remove after
	/*if list, found := idx.imgSearch("duolingo"); found { // remove after
		for _, val := range list {
			fmt.Printf("url: %v, src: %v, altTxt: %v\n", val.url, val.src, val.altTxt)
		}
	} else {
		fmt.Println("Not found")
	}*/

	/*if list, found := dlookup("computer science", &idx, false); found {
		for _, val := range list {
			fmt.Printf("title: %v, tfidf: %v\n", val.title, val.tfidf)
		}
	} else {
		fmt.Println("Not found")
	}*/
	
	for {
		time.Sleep(100 * time.Millisecond)
	}

	/*idx.Open()
	urls, _ := dlookup("rob", &idx, true)
	for _, url := range urls {
		fmt.Printf("{\"%v\", %v},\n", url.url, url.tfidf)
	}*/
}
