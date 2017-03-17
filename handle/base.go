package handle

import (
	// "fmt"
	// "net/http"

	"bytes"
	"strings"

	"github.com/zhengkai/goweb/module"
	"github.com/zhengkai/sigo/handle"
)

type Base struct {
	handle.BaseHandle
	Session *module.Session
}

func (this *Base) Prepare() bool {

	s := module.SessionParse(this.W, this.R)
	if s != nil {
		this.Session = s
		return true
	}

	if strings.HasPrefix(this.Uri, `/passport`) {
		return true
	}

	this.Redirect(`/passport/login`)
	return false
}

func (this *Base) Output() *bytes.Buffer {
	if this.ContentType == handle.Html && this.Session != nil {
		if this.Data == nil {
			this.Data = make(map[string]interface{})
		}
		this.Data.(map[string]interface{})[`_session`] = this.Session.GetUser()
	}
	return this.BaseHandle.Output()
}
