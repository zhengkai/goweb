package main

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/bitly/go-simplejson"
)

func updateInfo() {

	var b []byte
	b = getHtml(`http://rank.kongzhong.com/Data/tanks.js`)
	b = b[13 : len(b)-3]
	b = bytes.Replace(b, []byte(`'`), []byte(`"`), -1)
	fmt.Println(string(b[:]))

	json, err := simplejson.NewJson(b)
	if err != nil {
		fmt.Println(`JSON decode error: `, err)
	}

	for sid, info := range json.MustMap() {

		infoF := info.(map[string]interface{})

		id, _ := strconv.Atoi(sid)
		en_id := infoF[`name`].(string)
		stype := infoF[`entype`].(string)
		name := infoF[`alias`].(string)

		fmt.Println(id, en_id, stype, name)

		query := `UPDATE tank SET tank_id = ?, type = ? WHERE en_name = ?`
		db.Exec(query, id, stype, en_id)
		fmt.Println(id, stype, en_id)
	}
}
