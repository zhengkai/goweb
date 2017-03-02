package main

type Layout interface {
	Set(string)
	Get() string
}

type BaseLayout struct {
	s string
	x string
}

func (this *BaseLayout) Set(s string) {
	this.s = s
}

func (this *BaseLayout) Get() string {
	return this.s
}

func (this BaseLayout) Foo(s string) {
	this.x = s
}

func (this BaseLayout) Bar() string {
	return this.x
}
