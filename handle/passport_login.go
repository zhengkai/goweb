package handle

import (
	"net/http"

	"github.com/zhengkai/sigo/handle"
)

type PassportLogin struct {
	Base
}

func (this PassportLogin) New() handle.Handle {
	c := this
	c.Head = new(CommonHead).New()
	return &c
}

func (this *PassportLogin) Get(r *http.Request) {
	data := make(map[string]interface{})
	this.Data = data
}

func init() {
	handle.Register(`/passport/login`, &PassportLogin{})
}
