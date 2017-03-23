package handle

import (
	"fmt"
	"sync/atomic"

	"github.com/zhengkai/goweb/module"
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

func (this *Abc) Parse() {
	data := make(map[string]interface{})
	this.Data = data

	sUser := this.R.FormValue(`user`)
	fmt.Println(`user`, sUser)

	sPassword := this.R.FormValue(`password`)
	fmt.Println(`password`, sPassword)

	this.SetTplFunc(`yesrpg`, func(s string) string {
		return `yes!` + s + `!rpg!`
	})

	data[`Foo`] = atomic.AddInt64(&connectCount, 1)
	data[`Bar`] = 123

	uid, err := module.UserCreate(sUser, sPassword)
	if err != nil {
		fmt.Println(`new user error:`, err)
		return
	}
	fmt.Println(`new user`, uid)
}

func init() {
	handle.Register(`/abc`, &Abc{})
}
