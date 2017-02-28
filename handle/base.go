package handle

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/zhengkai/goweb/layout"
)

type Handle struct {
	uri    string
	layout layout.Layout
	data   interface{}
}

func (this Handle) setUri(s string) {
	this.uri = s
}

func (this Handle) Get(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL.Query())

	this.data = make(map[string]interface{})

	var p layout.Layout
	p.SetPath(`Handle`)
	this.layout = p
}

func (this Handle) Parse() *bytes.Buffer {
	return this.layout.Parse()
}

func (this Handle) Test() string {
	return `Yes RPG`
}

type handle interface {
	setUri(string)
	Get(w http.ResponseWriter, r *http.Request)
	Parse() *bytes.Buffer
}

func register(uri string, data handle) {

	http.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {

		d := data
		d.setUri(uri)

		// request := &data{}
		// fmt.Printf(`%T`, request)

		d.Get(w, r)
		w.Write(d.Parse().Bytes())
	})
}
