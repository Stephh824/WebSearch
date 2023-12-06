package main

import (
	"sync"
	"database/sql"
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
	dservers(&idx)
	//dcrawl("https://openai.com", &idx)
}
