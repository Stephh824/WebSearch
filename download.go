package main

import (
	"io"
	"net/http"
	"log"
)

func download(url string, ch chan DownloadResult) {
	rsp, err := http.Get(url)

	if rsp.StatusCode == http.StatusOK && err == nil {
		b, err := io.ReadAll(rsp.Body)
		if err == nil {
			ch <- DownloadResult{url, b, err}
		}
	} else {
		log.Fatalf("err downloading: %v\n", err)
		ch <- DownloadResult{url, []byte{}, err}
	}
}
