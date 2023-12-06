package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"database/sql"
	"text/template"
)

type Book struct {
	Title, Author string
}

func dlookup(word string, idx *IIndex, wcSearch bool) (list TfList, found bool) {
	var word_id int
	var word_id2 int
	var totdocCount int
	var hitdocCount int
	wcWords := []int{}

	words := stringMod(word)
	word = words[0]
	
	err := idx.dbase.QueryRow("SELECT count() FROM urls").Scan(&totdocCount)
	if err != nil {
		return list, false
	}

	// if is not a wildcard search, select only terms
	if !wcSearch {
		err = idx.dbase.QueryRow("SELECT id from terms WHERE name = ?", word).Scan(&word_id)
		if err != nil {
			return list, false
		}
		wcWords = append(wcWords, word_id)

		// selecting second word if its 2 words
		if len(words) > 1 {
			err = idx.dbase.QueryRow("SELECT id from terms WHERE name = ?", words[1]).Scan(&word_id2)
			if err != nil {
				return list, false
			}
		}

	// if is a wildcard search, select all like terms
	} else {
		word = word + "%"
		stmt, err := idx.dbase.Prepare("SELECT id from terms WHERE name LIKE ?")
		if err != nil {
			log.Fatalf("Could not prep wildcard stmt: %v\n", err)
		}
		rows, err := stmt.Query(word)
		if err != nil {
			return list, false
		}
		for rows.Next() {
			err = rows.Scan(&word_id)
			if err != nil {
				log.Fatalf("Could not scan into term_id: %v\n", err)
			}
			wcWords = append(wcWords, word_id)
		}
	}
	for _, term_id := range wcWords {
		var rows *sql.Rows
		if len(words) == 1 { // single-word hit
			err = idx.dbase.QueryRow("SELECT count() FROM hits WHERE term_id = ?", term_id).Scan(&hitdocCount)
			if err != nil {
				return list, false
			}
			rows, err = idx.dbase.Query("SELECT url_id, freq from hits WHERE term_id = ?", term_id)
			if err != nil {
				return list, false
			}
		} else { // bi-gram hit
			err = idx.dbase.QueryRow("SELECT count() from bigrams WHERE (term_id1, term_id2) = (?, ?)", term_id, word_id2).Scan(&hitdocCount)
			if err != nil {
				return list, false
			}
			rows, err = idx.dbase.Query("SELECT url_id, freq from bigrams WHERE (term_id1, term_id2) = (?, ?)", term_id, word_id2)
			if err != nil {
				return list, false
			}
		}
		for rows.Next() {
			var url_id int
			var freq int
			var title string
			var wordCount int
			var url string
			err = rows.Scan(&url_id, &freq)
			if err != nil {
				log.Fatalf("Could not scan freq row: %v\n", err)
			}
			err := idx.dbase.QueryRow("SELECT url, title, w_count from urls WHERE id = ?", url_id).Scan(&url, &title, &wordCount)
			if err != nil {
				log.Fatalf("Word lookup: failed to get url info %v: %v\n", url_id, err)
			}
			/*fmt.Printf("word: %v, url: %v, hitdocCount: %v, totdocCount: %v, freq: %v, wordCount: %v\n",
			word, url, hitdocCount, totdocCount, freq, wordCount)*/
			tf := float64(freq) / float64(wordCount)
			df := float64(hitdocCount) / float64(totdocCount)
			idf := float64(1 / df)
			tfidf := tf * idf
			nurl := TfIdf{url: url, title: title, tfidf: tfidf}
			list = append(list, nurl)
		}
	}
	sort.Sort(list)
	return list, true
}

func (idx *IIndex) imgSearch(word string) (list TfList, found bool) {
	var rows *sql.Rows
	var term_id int // yes
	var url_id int // yes
	var img_id int // yes
	var src string // yes, not used, returning
	var altTxt string // yes, not used, returning
	var url string // yes, not used, returning
	var freq int // yes, not used
	var wordCount int // yes, not used
	var title string // yes, not used, returning
	var totdocCount int // yes, not used
	var hitdocCount int // yes, not used

	words := stringMod(word)
	word = words[0]

	err := idx.dbase.QueryRow("SELECT count() FROM urls").Scan(&totdocCount)
	if err != nil {
		return list, false
	}

	err = idx.dbase.QueryRow("SELECT id from alt_terms WHERE name = ?", word).Scan(&term_id)
	if err != nil {
		return list, false
	}

	err = idx.dbase.QueryRow("SELECT count() FROM alt_hits WHERE term_id = ?", term_id).Scan(&hitdocCount)
	if err != nil {
		return list, false
	}
	
	rows, err = idx.dbase.Query("SELECT url_id, img_id, freq from alt_hits WHERE term_id = ?", term_id)

	for rows.Next() {
		// with url_id, need wordCount, title
		err = rows.Scan(&url_id, &img_id, &freq)
		// log.Printf("url_id: %v, img_id: %v, freq: %v\n", url_id, img_id, freq)
		err = idx.dbase.QueryRow("SELECT url, title, altw_count from urls WHERE id = ?", url_id).Scan(&url, &title, &wordCount)
		//log.Printf("title: %v, wordCount: %v\n", title, wordCount)
		err = idx.dbase.QueryRow("SELECT src, alt_txt from imgs WHERE id = ?", img_id).Scan(&src, &altTxt)
		//log.Printf("src: %v, altTxt: %v\n", src, altTxt)
		tf := float64(freq) / float64(wordCount)
		df := float64(hitdocCount) / float64(totdocCount)
		idf := float64(1 / df)
		tfidf := tf * idf
		//log.Println(tfidf)
		nurl := TfIdf{url, title, altTxt, src, tfidf}
		list = append(list, nurl)
	}
	sort.Sort(list)
	return list, true
	
}

func dsearch(idx *IIndex) func(w http.ResponseWriter, r *http.Request) {
	searchserv := func(w http.ResponseWriter, r *http.Request) {
		wcSearch := false
		if wc := r.URL.Query().Get("wildcard"); wc != "" {
			wcSearch = true
		}
		if img := r.URL.Query().Get("image"); img != "" {
			log.Println("image search")
		}
		r.ParseForm()
		searchw := r.Form["term"][0]                   // retrieve search term
		if img := r.URL.Query().Get("image"); img != "" {
			idx.imgSearch(searchw)
		} else {
			if keys, found := dlookup(searchw, idx, wcSearch); found { // lookup term; was found keys = map[string]int
				for _, value := range keys { // range through word's url:freq map
					stri := fmt.Sprintf("%s: %g\n", value.url, value.tfidf) // string format
					io.WriteString(w, stri)                                 // write to browser
				}
			} else { // was not found
				io.WriteString(w, "Not found")
			}
		}
	}
	return searchserv
}

func (idx *IIndex) search(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	wcSearch := false
	searchWord := r.Form["term"][0]
	if wc := r.URL.Query().Get("wildcard"); wc != "" {
		wcSearch = true
	}
	log.Printf("wildcard: %v word: %v\n", wcSearch, searchWord)

	

	t, err := template.ParseFiles("static/search.html")
	if err != nil {
		log.Fatalf("FarseFiles: ", err)
	}

	books := []Book {
		{"The Iliad", "Homer"},
		{"Dracula", "Bram Stoker"},
	}

	err = t.Execute(w, books)
	if err != nil {
		log.Fatalf("Temp Execute: ", err)
	}
}

func dservers(idx *IIndex) {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/search", idx.search)
	//http.HandleFunc("/search", dsearch(idx))

	go http.ListenAndServe(":8080", nil)
}
