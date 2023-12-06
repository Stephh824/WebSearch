package main

type TfIdf struct {
	Url   string
	Title string
	AltTxt string
	Src string
	Tfidf float64
}

type TfList []TfIdf

func (list TfList) Len() int {
	return len(list)
}

func (list TfList) Less(i, j int) bool {
	if list[i].Tfidf == list[j].Tfidf {
		return list[i].Url > list[j].Url
	}
	return list[i].Tfidf > list[j].Tfidf
}

func (list TfList) Swap(i, j int) {
	list[i].Url, list[i].Tfidf, list[j].Url, list[j].Tfidf = list[j].Url, list[j].Tfidf, list[i].Url, list[i].Tfidf
}
