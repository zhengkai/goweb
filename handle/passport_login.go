package handle

import (
	"github.com/zhengkai/sigo/handle"
)

type PassportLogin struct {
	Base
}

func (this PassportLogin) New() handle.Handle {
	c := this
	c.Head = new(CommonHead).New()
	c.Head.AddJS(`/res/login.js`)
	return &c
}

func init() {
	handle.Register(`/passport/login`, &PassportLogin{})
}
