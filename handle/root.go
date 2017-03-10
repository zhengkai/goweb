package handle

import (
	// "fmt"
	"net/http"
	// "github.com/zhengkai/goweb"
)

var (
	hello = []byte(`hello`)
)

func init() {
	http.HandleFunc(`/hello`, echo)

	http.Handle(`/`, http.FileServer(http.Dir(`./static`)))
}

func echo(w http.ResponseWriter, r *http.Request) {
	w.Write(hello)
}
