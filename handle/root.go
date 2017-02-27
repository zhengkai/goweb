package handle

import (
	"fmt"
	"net/http"
	// "github.com/zhengkai/goweb"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	w.Write([]byte(`nothing in /`))
}

func init() {
	http.HandleFunc(`/`, root)
}
