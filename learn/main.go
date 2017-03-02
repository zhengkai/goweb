package main

import "fmt"

func main() {

	var a Layout
	a = &BaseLayout{}
	fmt.Println(a)
	a.Set(`foo`)
	fmt.Println(`s =`, a.Get())

	b := BaseLayout{}
	b.Foo(`x`)
	fmt.Println(`x =`, b.Bar())
}
