package handle

import (
	"fmt"
	"net/http"
	// "github.com/zhengkai/goweb"
)

var (
	hello = []byte(`hello`)
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	w.Write([]byte(`nothing in /`))
}

func init() {
	http.HandleFunc(`/hello`, echo)
	http.HandleFunc(`/`, root)
}

func echo(w http.ResponseWriter, r *http.Request) {
	w.Write(hello)
}
