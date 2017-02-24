package layout

import (
	"bytes"
	"html/template"
)

var (
	layout = `default`
)

func Set(s string) {
	layout = s
}

func Get() string {
	return layout
}

func Parse() []byte {

	buf := new(bytes.Buffer)

	tpl, _ := template.ParseFiles(`tpl/head.html`)
	tpl.Execute(buf, nil)

	tpl, _ = template.ParseFiles(`tpl/nav.html`)
	tpl.Execute(buf, nil)

	tpl, _ = template.ParseFiles(`tpl/foot.html`)
	tpl.Execute(buf, nil)

	return buf.Bytes()
}
