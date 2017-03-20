package handle

import (
	"net/http"

	"github.com/zhengkai/goweb/module"
	"github.com/zhengkai/sigo/handle"
)

type PassportLogoutDo struct {
	Base
}

func (this PassportLogoutDo) New() handle.Handle {
	c := this
	c.Head = new(CommonHead).New()
	return &c
}

func (this *PassportLogoutDo) Parse() {

	if !this.CheckPost() {
		return
	}

	this.ContentType = handle.Json
	this.Data = LoginJSON{
		Success: true,
	}
	http.SetCookie(this.W, &module.CookieClean)
}

func init() {
	handle.Register(`/passport/logout.do`, &PassportLogoutDo{})
}
