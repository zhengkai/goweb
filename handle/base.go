package handle

import (
	// "fmt"
	"net/http"

	"github.com/zhengkai/sigo/handle"
)

type Base struct {
	handle.BaseHandle
}

func (this *Base) Prepare(w http.ResponseWriter, r *http.Request) bool {

	// http.Redirect(w, r, `https://soulogic.com`, 302)
	return true
}
