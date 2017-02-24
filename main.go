package main

import (
	_ "./handle"
	"./layout"

	"fmt"
	"net/http"
)

func main() {

	layout.Set(`layout`)
	fmt.Println(layout.Get())

	// http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func Test() string {
	return "Test"
}
