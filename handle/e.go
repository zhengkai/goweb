package handle

import (
	"net/http"

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

func (this *ErrorPage) Get(r *http.Request) {
	this.Error = `test error`
	this.ErrorMsg = `error message`
}

func init() {
	handle.Register(`/e`, &ErrorPage{})
}
