package main

import (
	"net/http"
	"strings"
	"log"
	"strconv"
	"net/url"
	"regexp"
	"io"
	"encoding/xml"
	"io/ioutil"
)

type RobotsTxt struct {
	UserAgent string
	CrawlDelay int
	Allow []string
	Disallow []string
}

type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	Urls []string `xml:"url>loc"`
}

var SiteMap string

func isDisallowed(url string, rules []RobotsTxt) (bool, RobotsTxt) {
	var ourUser RobotsTxt
	for _, rule := range rules {
		if rule.UserAgent == ".*" {
			ourUser = rule
			break
		}
	}

	for _, disa := range ourUser.Disallow {
		// err check
		if matched, _ := regexp.MatchString(disa, url); matched {
			return true, ourUser
		}
	}
	return false, ourUser
}

func robots(Url string) (error, []RobotsTxt) {
	var rules []RobotsTxt
	var currUA string
	var currRules RobotsTxt

	pu, err := url.Parse(Url)
	if err != nil {
		log.Fatal(err)
		return err, rules
	}

	// adding /robots.txt to hostname
	pu.Path = "robots.txt"
	Url = pu.String()

	rsp, err := http.Get(Url)
	if err != nil {
		log.Fatal(err)
		return err, rules
	}
	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Fatalf("Could not get robots body: %v\n", err)
	}

	for _, line := range strings.Split(string(body), "\n") {
		if strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}
		line = strings.ReplaceAll(line, "*", ".*")
		words := strings.SplitN(line, ":", 2)
		if len(words) == 2 {
			field := strings.ToLower(words[0])
			value := strings.TrimSpace(words[1])
			switch field {
				case "user-agent":
					if currUA != "" {
						rules = append(rules, currRules)
						currRules = RobotsTxt{}
					}
					currUA = value
					currRules.UserAgent = value
					currRules.CrawlDelay = 100
				case "crawl-delay":
					currRules.CrawlDelay, _ = strconv.Atoi(value)
				case "disallow":
					currRules.Disallow = append(currRules.Disallow, value)
				case "allow":
					currRules.Allow = append(currRules.Allow, value)
				case "sitemap":
					SiteMap = value
			}
		}
	}

	// to append last policy
	if currUA != "" {
		rules = append(rules, currRules)
	}

	return err, rules
}

func readSitemap() []string {
	rsp, err := http.Get(SiteMap)
	if err != nil {
		log.Fatalf("Could not get sitemap: %v\n", err)
	}
	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Fatalf("Could not read sitemap: %v\n", err)
	}

	var sitemap Sitemap

	err = xml.Unmarshal(data, &sitemap)
	if err != nil {
		log.Fatalf("Could not unmarshal: %v\n", err)
	}

	return sitemap.Urls
}
