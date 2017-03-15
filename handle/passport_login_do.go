package handle

import (
	"github.com/zhengkai/sigo/handle"
)

type PassportLoginDo struct {
	Base
}

func (this PassportLoginDo) New() handle.Handle {
	c := this
	c.Head = new(CommonHead).New()
	return &c
}

func (this *PassportLoginDo) Parse() {
	this.ContentType = handle.Json

	aReturn := struct {
		Success bool   `json:"success"`
		Uid     int64  `json:"uid,omitempty"`
		Error   string `json:"error,omitempty"`
	}{
		Success: true,
	}

	this.Data = aReturn
}

func init() {
	handle.Register(`/passport/login.do`, &PassportLoginDo{})
}
