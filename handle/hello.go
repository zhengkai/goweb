package handle

import (
	"net/http"
)

var (
	hello = []byte(`hello`)
)

func init() {
	http.HandleFunc(`/hello`, func(w http.ResponseWriter, r *http.Request) {
		w.Write(hello)
	})
}
