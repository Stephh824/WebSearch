package main

import (
	"bytes"
	"strings"
	"unicode"
	"log"
	"golang.org/x/net/html"
	"github.com/kljensen/snowball"
)

// to apply same string modification to input and extracted words
func stringMod(sent string) []string {
	isPunc := func (let rune) bool {
		return !unicode.IsLetter(let) && !unicode.IsNumber(let)
	}
	words := strings.FieldsFunc(strings.ToLower(sent), isPunc) // removes punctuation
	for idx, word := range words {
		if stemmed, err := snowball.Stem(word, "english", true); err == nil {
			words[idx] = stemmed
		} else {
			log.Fatalf("Stemmed Failed: %v\n", err)
		}
	}
	return words
}

func extract(dl *DownloadResult, chOut chan ExtractResult) {
	words := []string{}
	hrefs := []string{}
	imgs := make(map[string]string)
	var title string
	
	doc, err := html.Parse(bytes.NewReader(dl.body))
	if err != nil {
		log.Fatalf("Could not parse doc: %v\n", err)
	}	
		var f func(*html.Node)
		f = func(n *html.Node) {
			switch n.Type {
			case html.ElementNode:
				if n.Data == "img" {
					var src string
					for _, attr := range n.Attr {
						if attr.Key == "src" {
							src = attr.Val
						}
						if attr.Key == "alt" {
							imgs[src] = attr.Val
						}
					}
				}
				// extracting hrefs
				for _, attr := range n.Attr {
					if attr.Key == "href" {
						hrefs = append(hrefs, attr.Val)
					}
				}
				// extracting title node
				if n.Data == "title" && n.Parent.Data == "head" {
					title = n.FirstChild.Data
				}
			case html.TextNode:
				p := n.Parent
				if p.Type == html.ElementNode && (p.Data != "style" && p.Data != "script") {
					toks := stringMod(n.Data)
					words = append(words, toks...)
				}
			}
			// go through the child nodes recursively
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
		f(doc)

	urls := []string{}
	for _, href := range hrefs {
		if cleanUrl, ok := clean(dl.url, href); ok {
			urls = append(urls, cleanUrl)
		}
	}

	chOut <- ExtractResult{dl.url, title, words, urls, imgs}
}

