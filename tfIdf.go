package main

type TfIdf struct {
	url   string
	title string
	altTxt string
	src string
	tfidf float64
}

type TfList []TfIdf

func (list TfList) Len() int {
	return len(list)
}

func (list TfList) Less(i, j int) bool {
	if list[i].tfidf == list[j].tfidf {
		return list[i].url > list[j].url
	}
	return list[i].tfidf > list[j].tfidf
}

func (list TfList) Swap(i, j int) {
	list[i].url, list[i].tfidf, list[j].url, list[j].tfidf = list[j].url, list[j].tfidf, list[i].url, list[i].tfidf
}
