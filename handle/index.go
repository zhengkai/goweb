package handle

import "github.com/zhengkai/sigo/handle"

type Index struct {
	Base
}

func (this Index) New() handle.Handle {
	c := this
	c.Head = new(CommonHead).New()
	return &c
}

func (this *Index) Parse() {
	uri := this.R.URL.RequestURI()
	if uri != `/` {
		this.Uri = `/error/404`
		e := make(map[string]string)
		this.Data = make(map[string]interface{})
		this.Data.(map[string]interface{})[`_error`] = e
		e[`uri`] = uri
		return
	}

	this.SetUri(`/index`)

	/*
		if this.Error != `` {
			this.Uri = `/error/500`
			e := make(map[string]string)
			this.Data.(map[string]interface{})[`_error`] = e

			e[`title`] = this.Error
			e[`msg`] = this.ErrorMsg
		}
	*/
}

func init() {
	handle.Register(`/`, &Index{})
}
