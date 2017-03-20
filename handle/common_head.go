package handle

import (
	"github.com/zhengkai/sigo/handle"
)

type CommonHead struct {
	handle.BaseHead
}

func (this CommonHead) New() handle.Head {
	c := this.BaseHead.New()
	c.AddJS(`/res/less.2.7.2.min.js`)
	c.AddJS(`/res/jquery-3.1.1/jquery.min.js`)
	c.AddJS(`/res/tether-1.3.3/js/tether.min.js`)
	c.AddJS(`/res/bootstrap-4.0.0-alpha6/js/bootstrap.js`)
	c.AddJS(`/res/script.js`)
	c.AddCSS(`/res/bootstrap-4.0.0-alpha6/css/bootstrap.min.css`)
	c.AddCSS(`/res/tether-1.3.3/css/tether.min.css`)
	c.AddCSS(`/res/font-awesome-4.7.0/css/font-awesome.min.css`)
	c.AddCSS(`/res/style.less`)
	c.AddCSS(`https://fonts.googleapis.com/css?family=Droid+Sans+Mono|Droid+Sans:400,700|Roboto:400,900`)
	return c
}
