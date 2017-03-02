package main

import "fmt"

type IF interface {
	Clone() IF
}

type Foo struct {
	i int
}

func (this *Foo) Clone() IF {
	c := *this
	return &c
}

type Bar struct {
	Foo
}

func main() {
	t := &Bar{}
	c := t.Clone()
	fmt.Printf(`%T `, t)
	fmt.Printf(`%T `, c)
}

func (this *Bar) Clone() IF {
	c := *this
	return &c
}
