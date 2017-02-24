package handle

import (
	"../layout"
	"fmt"
	"net/http"
	// "github.com/zhengkai/goweb"
)

func Yesrpg() []byte {
	return layout.Parse()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	w.Write(Yesrpg())
}

func init() {
	fmt.Println(`controller init`)
	http.HandleFunc("/", handler)
}
