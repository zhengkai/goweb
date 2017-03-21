package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getHtml(url string) (s string) {

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	s = string(b[:])
	return

}

func main() {

	var s string
	s = getHtml(`http://rank.kongzhong.com/wiki/techtree.html`)

	s = getHtml(`http://rank.kongzhong.com/Data/tankwiki/0.9.15.1/vehicles/japan/list.xml.json`)
	fmt.Println(s)

}

//
