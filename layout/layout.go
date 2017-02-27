package layout

import (
	"bytes"
	"fmt"
	"html/template"
)

type Layout struct {
	path string
}

func (this *Layout) SetPath(s string) {
	this.path = s
}

func (this *Layout) Parse() []byte {

	fmt.Println(`this SetTpl =`, this.path)

	buf := new(bytes.Buffer)

	tpl, _ := template.ParseFiles(`tpl/head.html`)
	tpl.Execute(buf, nil)

	tpl, _ = template.ParseFiles(`tpl/nav.html`)
	tpl.Execute(buf, nil)

	tpl, _ = template.ParseFiles(`tpl/foot.html`)
	tpl.Execute(buf, nil)

	return buf.Bytes()
}
