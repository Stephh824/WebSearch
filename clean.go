package main

import(
	"net/url"
	"log"
	"strings"
)

func clean(host, href string) (string, bool) {
	hrefu, err := url.Parse(href)
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(hrefu.Path, "images") {
		return "", false
	}
	
	hn := hrefu.Hostname()
	if len(hn) == 0 { // no host name
		hostu, err := url.Parse(host)
		if err != nil {
			log.Fatal(err)
		}
		
		if len(hrefu.Path) == 0 && len(hrefu.Fragment) != 0 { // skipping fragments
			hrefu.Path = hostu.Path
			hrefu.Fragment = ""
		}
		hrefu.Scheme = hostu.Scheme // combining valid hrefs
		hrefu.Host = hostu.Host
	} else if hn != host { // if part of different host, dont crawl
		return "", false
	}
	return hrefu.String(), true
}

