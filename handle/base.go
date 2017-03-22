package handle

import (
	// "fmt"
	// "net/http"

	"net/http"
	"strings"

	"github.com/zhengkai/goweb/module"
	"github.com/zhengkai/sigo/handle"
)

var (
	domain        string
	refererPrefix string
)

func SetDomain(s string) {
	domain = s
	refererPrefix = `https://` + s
}

type Base struct {
	handle.BaseHandle
	Session *module.Session
}

func (this *Base) CheckPost() bool {
	if !this.BaseHandle.CheckPost() {
		return false
	}
	if !strings.HasPrefix(this.R.Header.Get(`Referer`), refererPrefix) {
		this.StopByStatus(http.StatusNotAcceptable)
		return false
	}
	return true
}

func (this *Base) Prepare() bool {

	s := module.SessionParse(this.W, this.R)
	if s != nil {
		this.Session = s
		return true
	}

	return true // 暂时允许所有非登录

	if this.Uri == `/` {
		return true
	}
	if strings.HasPrefix(this.Uri, `/passport`) {
		return true
	}

	this.Redirect(`/passport/login`)
	return false
}

func (this *Base) Output() {
	if this.ContentType == handle.Html && this.Session != nil {
		if this.Data == nil {
			this.Data = make(map[string]interface{})
		}
		this.Data.(map[string]interface{})[`_session`] = this.Session.GetUser().Info()
	}
	this.BaseHandle.Output()
}
