package handle

import (
	"fmt"
	"net/http"

	"github.com/zhengkai/goweb/layout"
)

func handlerABC(w http.ResponseWriter, r *http.Request) (interface{}, layout.Layout) {

	data := make(map[string]interface{})

	fmt.Println(r.URL.Query())

	var p layout.Layout

	p.SetPath(`abc`)

	return data, p
}

func parseABC(w http.ResponseWriter, r *http.Request) {

	_, layout := handlerABC(w, r)
	fmt.Println(layout)
	w.Write(layout.Parse())
}

func register(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, handler)
}

func init() {
	register(`/abc`, parseABC)
}
