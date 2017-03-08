package main

import "fmt"

type IF interface {
	Clone() IF
	Run()
}

type Foo struct {
	j int
}

func (this Foo) Clone() IF {
	c := this
	return &c
}

func (this Foo) Run() {
	x := &this
	x.jump(3)
}

func (this *Foo) jump(i int) {
	this.j = i
}

type Bar struct {
	Foo
}

func main() {
	a := Bar{}
	a.Run()
	a.jump(4)
	fmt.Println(a)
}
