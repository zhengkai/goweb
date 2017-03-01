package handle

import (
	// "fmt"
	"net/http"

	"github.com/zhengkai/sigo/handle"
)

type Abc struct {
	handle.BaseHandle
}

func (this *Abc) Get(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data[`Foo`] = 123
	data[`Bar`] = 123

	// fmt.Println(`data`, data)

	this.Data = data
}

func init() {
	handle.Register(`/abc`, &Abc{})
}
