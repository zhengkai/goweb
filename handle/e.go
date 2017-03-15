package handle

import (
	"github.com/zhengkai/sigo/handle"
)

type ErrorPage struct {
	Base
}

func (this ErrorPage) New() handle.Handle {
	c := this
	c.Head = new(CommonHead).New()
	return &c
}

func (this *ErrorPage) Parse() {
	this.Error = `test error`
	this.ErrorMsg = `error message`
}

func init() {
	handle.Register(`/e`, &ErrorPage{})
}
