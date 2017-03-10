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
	return &c
}

func (this *Abc) Get(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	data[`Foo`] = atomic.AddInt64(&connectCount, 1)
	data[`Bar`] = 123

	this.Data = data
}

func init() {
	handle.Register(`/abc`, &Abc{})
}
