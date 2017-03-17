package handle

import (
	"net/http"

	"github.com/zhengkai/goweb/module"
	"github.com/zhengkai/sigo/handle"
)

type PassportRegisterDo struct {
	Base
}

func (this PassportRegisterDo) New() handle.Handle {
	c := this
	c.Head = new(CommonHead).New()
	return &c
}

func (this *PassportRegisterDo) Parse() {
	this.ContentType = handle.Json
	aReturn := LoginJSON{
		Success: true,
	}
	this.Data = &aReturn

	name := this.R.FormValue(`name`)
	password := this.R.FormValue(`password`)

	user, err := module.UserCreate(name, password)
	if err != nil {
		aReturn.Success = false
		aReturn.Error = err.Error()
		return
	}

	s := &module.Session{
		Uid:     user.Id,
		Session: 3,
	}
	http.SetCookie(this.W, s.MakeCookie())

	aReturn.Success = true
	aReturn.Uid = user.Id
}

func init() {
	handle.Register(`/passport/register.do`, &PassportRegisterDo{})
}
