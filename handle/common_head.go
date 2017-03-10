package handle

import (
	"github.com/zhengkai/sigo/handle"
)

type CommonHead struct {
	handle.BaseHead
}

func (this CommonHead) New() handle.Head {
	c := this.BaseHead.New()
	c.AddMeta(`<meta name="theme-color" content="#375EAB">`)
	return c
}
