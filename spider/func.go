package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func getHtml(url string) (b []byte) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, _ = ioutil.ReadAll(resp.Body)
	//s = string(b[:])
	return
}

func getFile(file string) (b []byte) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()
	b, _ = ioutil.ReadAll(f)
	return
}
