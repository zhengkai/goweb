package handle

import (
// "github.com/zhengkai/goweb/layout"
)

type Abc struct {
	Handle
}

func init() {
	register(`/abc`, Abc{})
}
