package handle

import (
	// "fmt"
	"net/http"
	"sync/atomic"

	"github.com/zhengkai/sigo/handle"
)

var (
	connectCount int64 = 0
)

type Abc struct {
	handle.BaseHandle
}

func (this Abc) New() handle.Handle {
	c := this
	c.Head = new(CommonHead).New()
	// c.Head.AddJS(`/res/script.js`)
	return &c
}

func (this *Abc) Get(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	data[`Foo`] = atomic.AddInt64(&connectCount, 1)
	data[`Bar`] = 123

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
