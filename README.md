# WebSearch
**Google in Go!**  
A lightweight search engine written in Go that supports web crawling, text extraction, TF-IDF based ranking, bigram search,
wildcard search, stemming, and basic image indexing.

---

### 🚀Features
- Web crawling for downloading HTML page following robots.txt rules
- Content extraction + cleaning
- Tf-IDF scoring for ranking search results
- Snowball stemming to normalize words
- Bigram phrase searching
- Wildcard search
- Image extraction
- Local persistent DB (currently named `example.db`) for storing indexed pages
- Pure Go implementation - no external search frameworks
  
---

### 🧠Concepts Used
This project implements core information-retrieval techniques:  
#### TF-IDF Ranking
Each document is represented by a vector of term weights. Query-document similarity is computed to produce ranked search results.
#### Snowball Stemming
Words are normalized to a common root to improve matching  
Example:  
- Query: *running*
- Document: *ran fast*
After stemming -> both map to **run** and match correctly
#### Bigram Phrase Searching
This means searching for:  
- Machine learning
- Computer vision
- Software development  
returns only documents that contain those two words consecutively, preserving meaning and reducing irrelevant matches.
#### Wildcard Search
Supports prefix wildcard queries:
- computer -> comput*
- emergency -> emerg*
- servive -> serv*
These match any word that starts with the prefix, even after stemming - useful for when the user is unsure of the word form.
#### Image Search
The crawler extracts all image source URLs from each page along with the alternative text attached on each page. Search results
include associated image URLs based off of the alternative text, making this a simple keyword-to-image lookup.

---

### 🛠️Getting Started
#### Install + Clone
```
git clone https://github.com/Stephh824/WebSearch.git
cd WebSearch
```

---

### 💡Running the App
The simplest and cleanest way:  
```go run .```  
Go automatically ignores the test files.  
Then go to `localhost:8080`.  
The current `example.db` includes results from crawling `https://openai.com`. If you wish to change this, in `main.go` simply
uncomment line 28, change to desired site and wait for it to finish crawling. This is concurrent with multiple channels so should
be relatively fast and the site will still be up and running before crawling finishes.

---

### 🔍Searching
Supported query types:  
| Query Type | Example | Behavior |
| ---------- | ------- | -------- |
| Single Word | `emergency` | TF-IDF ranking |
| Two-word bigram | `machine learning` | Exact phrase matching + TF-IDF ranking |
| Wildcard Checkbox | `computer` | Will be processed as `comput*` and show all word forms + TF-IDF ranking |
| Image Checkbox | `screen` | Will search through alternative image texts and display corresponding images + TF-IDF ranking |

---

### 🚧Notes & Limitations
- Wildcard matching currently supports prefix only, not mid-word or suffix wildcarding
- Bigram searching supports *two-word phrases* only
- Image search is text-based not ML-based


