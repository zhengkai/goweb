package handle

import (
	"net/http"
)

func init() {
	http.Handle(`/res/`, http.StripPrefix(`/res/`, http.FileServer(http.Dir(`./static/res`))))
}
