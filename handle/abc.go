package handle

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/zhengkai/goweb/model"
	"github.com/zhengkai/sigo/handle"
)

var (
	connectCount int64 = 0
)

type Abc struct {
	Base
}

func (this Abc) New() handle.Handle {
	c := this
	c.Head = new(CommonHead).New()
	return &c
}

func (this *Abc) Get(r *http.Request) {
	data := make(map[string]interface{})
	this.Data = data

	sUser := r.FormValue(`user`)
	fmt.Println(`user`, sUser)

	sPassword := r.FormValue(`password`)
	fmt.Println(`password`, sPassword)

	data[`Foo`] = atomic.AddInt64(&connectCount, 1)
	data[`Bar`] = 123

	uid, err := model.UserCreate(sUser, sPassword)
	if err != nil {
		fmt.Println(`new user error:`, err)
		return
	}
	fmt.Println(`new user`, uid)
}

func init() {
	handle.Register(`/abc`, &Abc{})
}
