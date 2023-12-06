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
	idx.Open() // remove after
	dservers(&idx)
	//dcrawl("https://openai.com/robots.txt", &idx)
	
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
