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

	data[`_head`] = make(map[string]interface{})
	data[`_head`].(map[string]interface{})[`css`] = `abc`

	// fmt.Println(`data`, data)
	// fmt.Println(`data`, data[`Foo`])

	// fmt.Println(`this.Head`, this.Head)
	this.Head.AddJS(`/res/script.js`)
	this.Head.AddJS(`/res/go.js`)
	this.Head.AddCSS(`/res/yes.css`)
	this.Head.AddCSS(`/res/a.less`)
	this.Head.AddCSS(`/res/a.css`)
	this.Head.AddCSS(`/res/b.less`)

	this.Data = data
}

func init() {
	handle.Register(`/abc`, &Abc{})
}
