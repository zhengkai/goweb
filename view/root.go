package view

import (
	"bytes"
	"fmt"
	"html/template"
)

func Show() string {

	tpl, err := template.New(`head.html`).ParseFiles(`tpl/head.html`)
	if err != nil {
		fmt.Println(err)
	}

	noItems := struct {
		Title string
	}{
		Title: "My another page",
	}

	buf := new(bytes.Buffer)
	err = tpl.Execute(buf, noItems)
	if err != nil {
		fmt.Println(err)
	}
	return buf.String()
}
