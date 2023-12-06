package main

import (
	"log"
	"net/http"
	"sort"
	"database/sql"
	"text/template"
)

func wordQuery(wcSearch bool) (word_id2, wordIds) 

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
			nurl := TfIdf{Url: url, Title: title, Tfidf: tfidf}
			list = append(list, nurl)
		}
	}
	sort.Sort(list)
	return list, true
}

func (idx *IIndex) imgSearch(word string) (list TfList, found bool) {
	var rows *sql.Rows
	var term_id, url_id, img_id, freq, wordCount, totdocCount, hitdocCount int
	var src, altTxt, url, title string

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
		err = rows.Scan(&url_id, &img_id, &freq)
		err = idx.dbase.QueryRow("SELECT url, title, altw_count from urls WHERE id = ?", url_id).Scan(&url, &title, &wordCount)
		err = idx.dbase.QueryRow("SELECT src, alt_txt from imgs WHERE id = ?", img_id).Scan(&src, &altTxt)
		tf := float64(freq) / float64(wordCount)
		df := float64(hitdocCount) / float64(totdocCount)
		idf := float64(1 / df)
		tfidf := tf * idf
		nurl := TfIdf{url, title, altTxt, src, tfidf}
		list = append(list, nurl)
	}
	sort.Sort(list)
	return list, true
	
}

func renderTemplate(path string, list TfList, w http.ResponseWriter) {
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Fatalf("Could not parse %v template: %v\n", path, err)
	}
	err = t.Execute(w, list)
	if err != nil {
		log.Fatalf("Could not execute list: %v\n", err)
	}
}

func (idx *IIndex) search(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	wcSearch := false
	imgSearch := false
	searchWord := r.Form["term"][0]
	if wc := r.URL.Query().Get("wildcard"); wc != "" {
		wcSearch = true
	}
	if img := r.URL.Query().Get("image"); img != "" {
		imgSearch = true
	}

	switch imgSearch {
		// word searching
		case false:
			list, found := dlookup(searchWord, idx, wcSearch)
			if found {
				renderTemplate("static/urls.html", list, w)
			} else {
				renderTemplate("static/notFound.html", list, w)
			}
		// image searching
		case true:
			list, found := idx.imgSearch(searchWord)
			if found {
				renderTemplate("static/images.html", list, w)
			} else {
				renderTemplate("static/notFound.html", list, w)
			}
	}
}

func dservers(idx *IIndex) {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.Handle("/styles", http.FileServer(http.Dir("./styles")))
	http.HandleFunc("/search", idx.search)

	go http.ListenAndServe(":8080", nil)
}
