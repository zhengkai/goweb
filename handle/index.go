package handle

import (
	"github.com/zhengkai/sigo/handle"
)

type Index struct {
	Base
}

func (this Index) New() handle.Handle {
	c := this
	c.Head = new(CommonHead).New()
	return &c
}

func (this *Index) Parse() {
	this.SetUri(`/index`)
}

func init() {
	handle.Register(`/`, &Index{})
}
