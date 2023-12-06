package main

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

func (index *IIndex) Open() {
	index.dbase, _ = sql.Open("sqlite3", "example.db")
	_, err := index.dbase.Exec(`
		PRAGMA foreign_keys = ON;
	`)
	if err != nil {
		log.Fatalf("Could not turn foreign keys on: %v\n", err)
	}
}

func (index *IIndex) Close() {
	index.dbase.Close()
}

func (index *IIndex) getSize() (int64, error) {
	info, err := os.Stat("example.db")
	if err != nil {
		log.Fatalf("err getting db size: %v\n", err)
		return 0, err
	}
	return info.Size(), nil
}

func dbInIt(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS terms (
			id integer PRIMARY KEY,
			name string UNIQUE
		);
	`)
	if err != nil {
		log.Fatalf("Could not create TERMS table: %v\n", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
			id integer PRIMARY KEY,
			url string UNIQUE,
			title string,
			altw_count integer,
			w_count integer
		);
	`)
	if err != nil {
		log.Fatalf("Could not create URLS table: %v\n", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS hits (
			id integer PRIMARY KEY,
			term_id integer,
			url_id integer,
			freq integer,
			FOREIGN KEY(term_id) REFERENCES terms(id),
			FOREIGN KEY(url_id) REFERENCES urls(id),
			CONSTRAINT unique_combination UNIQUE (term_id, url_id)
		);
	`)
	if err != nil {
		log.Fatalf("Could not create HITS table: %v\n", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS bigrams (
			id integer PRIMARY KEY,
			term_id1 integer,
			term_id2 integer,
			url_id integer,
			freq integer,
			FOREIGN KEY(term_id1) REFERENCES terms(id),
			FOREIGN KEY(term_id2) REFERENCES terms(id),
			FOREIGN KEY(url_id) REFERENCES urls(id),
			CONSTRAINT unique_combination UNIQUE (term_id1, term_id2, url_id)
		);
	`)
	if err != nil {
		log.Fatalf("Could not create BIGRAMS table: %v\n", err)
	}

	_, err = db.Exec(`
		CREATE INDEX IF NOT EXISTS u_url on urls(url);
	`)
	if err != nil {
		log.Fatalf("Could not create urls(url) INDEX: %v\n", err)
	}

	_, err = db.Exec(`
		CREATE INDEX IF NOT EXISTS n_terms on terms(name);
	`)
	if err != nil {
		log.Fatalf("Could not create n_terms(name) INDEX: %v\n", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS alt_terms (
			id integer PRIMARY KEY,
			name string UNIQUE
		);
	`)
	if err != nil {
		log.Fatalf("Could not create alt_terms table: %v\n", err)
	}
	
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS imgs (
			id integer PRIMARY KEY,
			src string UNIQUE,
			alt_txt string
		);
	`)
	if err != nil {
		log.Fatalf("Could not create imgs table: %v\n", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS alt_hits (
			id integer PRIMARY KEY,
			term_id integer,
			url_id integer,
			img_id integer,
			freq integer,
			FOREIGN KEY(term_id) REFERENCES alt_terms(id),
			FOREIGN KEY(url_id) REFERENCES urls(id),
			FOREIGN KEY(img_id) REFERENCES imgs(id),
			CONSTRAINT unique_combination UNIQUE (term_id, img_id)
		);
	`)
	if err != nil {
		log.Fatalf("Could not create alt_hits table: %v\n", err)
	}
}

func (idx *IIndex) insertUrl(url string) {
	_, err := idx.dbase.Exec(`
		INSERT OR IGNORE INTO urls(url, w_count, altw_count) VALUES(?, 0, 0);`, url)
	if err != nil {
		log.Fatalf("Could not insert URL %v: %v\n", url, err)
	}
}

func (idx *IIndex) addWordCount(url string, wcount int) {
	var id int
	err := idx.dbase.QueryRow("SELECT id from urls WHERE url = ?", url).Scan(&id)
	if err != nil {
		log.Fatalf("Could not select URL %v: %v\n", url, err)
	}
	_, err = idx.dbase.Exec(`
		UPDATE urls SET w_count = ? WHERE id = ?;`, wcount, id)
	if err != nil {
		log.Fatalf("Could not update word count %v: %v\n", url, err)
	}
}

func (idx *IIndex) getWordId(word string) int {
	var term_id int
	err := idx.dbase.QueryRow("SELECT id from terms WHERE name = ?", word).Scan(&term_id)
	if err != nil {
		log.Fatalf("Could not select term_id (getWordId): %v\n", err)
	}
	return term_id
}

func (idx *IIndex) getAltWordId(word string) int {
	var term_id int
	err := idx.dbase.QueryRow("SELECT id from alt_terms WHERE name = ?", word).Scan(&term_id)
	if err != nil {
		log.Fatalf("Could not select term_id (getAltWordId): %v\n", err)
	}
	return term_id
}

func (idx *IIndex) getImgId(src string) int {
	var img_id int
	err := idx.dbase.QueryRow("SELECT id from imgs WHERE src = ?", src).Scan(&img_id)
	if err != nil {
		log.Fatalf("Could not select img_id (getImgId): %v\n", err)
	}
	return img_id
}

func (idx *IIndex) getUrlId(url string) int {
	var url_id int
	err := idx.dbase.QueryRow("SELECT id from urls WHERE url = ?", url).Scan(&url_id)
	if err != nil {
		log.Fatalf("Could not select url_id (getUrlId): %v\n", err)
	}
	return url_id
}
