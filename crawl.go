package main

import (
	"fmt"
	"time"
	"log"
)

type DownloadResult struct {
	url string
	body []byte
	err error
}

type ExtractResult struct {
	url, title string
	words, urls []string
	imgs map[string]string
}

func (idx *IIndex) addImgs(url_id int, ex *ExtractResult) {
	var wcount int
	for img, alt := range ex.imgs {
		//log.Println("image is: " + img + " alt text is: " + alt)
		_, err := idx.dbase.Exec(`INSERT OR IGNORE INTO imgs(src, alt_txt) VALUES(?, ?);`, img, alt)
		if err != nil {
			log.Fatalf("Could not insert %v, %v into imgs table: %v\n", img, alt, err)
		}
		words := stringMod(alt)
		for _, word := range words {
			if !isInStopWords(word, stopwords) {
				wcount++
				_, err := idx.dbase.Exec(`INSERT OR IGNORE INTO alt_terms(name) VALUES(?);`, word)
				if err != nil {
					log.Fatalf("Could not insert %v into alt_terms table: %v\n", word, err)
				}
				img_id := idx.getImgId(img)
				term_id := idx.getAltWordId(word)
				//log.Printf("img_id is: %v term_id is: %v\n", img_id, term_id)
				_, err = idx.dbase.Exec(`INSERT OR IGNORE INTO alt_hits(term_id, url_id, img_id, freq) VALUES (?,?,?,?) ON CONFLICT(term_id, img_id) DO UPDATE SET freq = freq + 1;`,
						int(term_id), url_id, int(img_id), 1)
				if err != nil {
					log.Fatalf("Could not insert alt_hit for %v: %v\n", word, err)
				}
			}
		}
	}
	_, err := idx.dbase.Exec(`UPDATE urls SET altw_count = ? WHERE id = ?`, wcount, url_id)
	if err != nil {
		log.Fatalf("Could not update altw_count: %v\n", err)
	}
}

func (idx *IIndex) dadd(ex *ExtractResult) {
	wcount := 0 //DB
	idx.mutex.Lock()
	defer idx.mutex.Unlock()
	url_id := idx.getUrlId(ex.url)
	go idx.addImgs(url_id, ex)
	_, err := idx.dbase.Exec(`UPDATE urls SET title = ? WHERE id = ?`, ex.title, url_id)
	if err != nil {
		log.Fatalf("Could not insert title to urls table: %v\n", err)
	}
	insertWords := "INSERT OR IGNORE INTO terms(name) VALUES "
	words := []interface{}{}
	insertHits := "INSERT OR IGNORE INTO hits(term_id, url_id, freq) VALUES "
	hits := []interface{}{}
	insertBiGram := "INSERT OR IGNORE INTO bigrams(term_id1, term_id2, url_id, freq) VALUES "
	bigrams := []interface{}{}

	for _, word := range ex.words {
		if !isInStopWords(word, stopwords) { // if word is not in stopwords
			wcount++
			insertWords += "(?),"
			words = append(words, word)
		}
	}
	// for words
	if wcount == 0 {
		return
	}
	// removes last comma
	insertWords = insertWords[0:len(insertWords) - 1]
	stmt, err := idx.dbase.Prepare(insertWords)
	if err != nil {
		log.Fatalf("Could not prepare words stmt: %v\n", err)
	}
	_, err = stmt.Exec(words...)
	if err != nil {
		log.Fatalf("Could not execute words: %v\n", err)
	}
	idx.addWordCount(ex.url, wcount)

	for ix, word := range words {
		// for single-word hit table
		term_id1 := idx.getWordId(word.(string))
		insertHits += "(?,?,?),"
		hits = append(hits, term_id1, url_id, 1)
		// checks if there is a next word for bigrams table
		if ix + 1 < len(words) {
			term_id2 := idx.getWordId(words[ix+1].(string))
			insertBiGram += "(?,?,?,?),"
			bigrams = append(bigrams, term_id1, term_id2, url_id, 1)
		}
	}
	// removes last comma
	insertHits = insertHits[0:len(insertHits) - 1]
	insertHits += "ON CONFLICT(term_id, url_id) DO UPDATE SET freq = freq + 1"
	stmt, err = idx.dbase.Prepare(insertHits)
	if err != nil {
		log.Fatalf("Could not prepare hits stmt: %v\n", err)
	}
	_, err = stmt.Exec(hits...)
	if err != nil {
		log.Fatalf("Could not execute hits: %v\n", err)
	}

	// removes last comma
	insertBiGram = insertBiGram[0:len(insertBiGram) - 1]
	insertBiGram += "ON CONFLICT(term_id1, term_id2, url_id) DO UPDATE SET freq = freq + 1"
	stmt, err = idx.dbase.Prepare(insertBiGram)
	if err != nil {
		log.Fatalf("Could not prepare bigram hits stmt: %v\n", err)
	}
	_, err = stmt.Exec(bigrams...)
	if err != nil {
		log.Fatalf("Could not execute bigram hits: %v\n", err)
	}
	fmt.Println("done")
}

func dcrawl(url string, idx *IIndex) {
	// pass url to crawl robots.txt
	err, agents := robots(url)
	if err == nil {
		idx.rules = agents
	}

	idx.Open()

	// channel initialization
	dlInC := make(chan string, 100000)
	dlOutC := make(chan DownloadResult, 100000)
	exOutC := make(chan ExtractResult, 100000)
	quitC := make(chan bool, 1)

	// to check if crawled
	urls := make(map[string]struct{})
	// initializing tables
	dbInIt(idx.dbase)
	// queue of urls to crawl
	s := []string{}

	// for href crawling
	/*s = append(s, url)
	idx.insertUrl(url)
	urls[url] = struct{}{}
	dlInC <- url*/

	// for sitemap crawling
	sites := readSitemap()
	idx.insertUrl(sites[0])
	dlInC <- sites[0]
	for idx, site := range sites {
		if idx < 5000 {
			s = append(s, site)
		} else {
			break
		}
	}
	s = s[1:]

	go func() {
		for {
			prevSize, _ := idx.getSize()
			time.Sleep(10 * time.Second)
			currSize, _ := idx.getSize()
			if currSize == prevSize && len(dlInC) + len(dlOutC) + len(exOutC) == 0 {
				quitC <- true
				break
			}
		}
	}()

outer:
	for {
		select {
		case url := <-dlInC:
			go download(url, dlOutC)
		case dlRes := <-dlOutC:
			go extract(&dlRes, exOutC)
		case exRes := <-exOutC:
			//log.Println("Adding: " + exRes.url)
			idx.dadd(&exRes)
			// for href crawling to add urls
			// s = append(s, exRes.urls...)
			for i := 0; i < len(s); i++ {
				url := s[i]
				// has not been crawled
				if _, ok := urls[url]; !ok {
					// check if in disallow
					disa, _ := isDisallowed(url, idx.rules)
					// is allowed
					if !disa {
						idx.insertUrl(url)
						urls[url] = struct{}{}
						if idx.rules[0].CrawlDelay == 100 {
							time.Sleep(100 * time.Millisecond)
						} else {
							time.Sleep(time.Duration(idx.rules[0].CrawlDelay) * time.Second)
						}
						dlInC <- url
					}
				}
			}
			case end := <-quitC:
			fmt.Println(end)
			close(dlInC)
			close(dlOutC)
			close(exOutC)
			close(quitC)
			break outer
		}
	}
	fmt.Println("complete")
}
