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
	
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
