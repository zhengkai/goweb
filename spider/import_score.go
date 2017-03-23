package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zhengkai/goweb/typelist"
)

var (
	scoreQuery = [...]string{
		`INSERT IGNORE INTO score SET ` +
			`vehicle_id = ?, date = NOW(), ` +
			`g1max = ?, g2max = ?, g3max = ?, ` +
			`g1min = ?, g2min = ?, g3min = ?, ` +
			`m1max = ?, m2max = ?, m3max = ?, m4max = ?, ` +
			`m1min = ?, m2min = ?, m3min = ?, m4min = ?`,
		`REPLACE INTO score_current SET ` +
			`vehicle_id = ?, date = NOW(), ` +
			`g1max = ?, g2max = ?, g3max = ?, ` +
			`g1min = ?, g2min = ?, g3min = ?, ` +
			`m1max = ?, m2max = ?, m3max = ?, m4max = ?, ` +
			`m1min = ?, m2min = ?, m3min = ?, m4min = ?`,
	}
)

func import_score_all() {
	query := `SELECT id FROM vehicle`
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for row.Next() {
		var id int
		row.Scan(&id)
		time.Sleep(3 * time.Second)
		import_score(id)
	}
}

func import_score(tank int) {

	if tank < 1 {
		return
	}

	var b []byte

	url := fmt.Sprintf(
		`http://hd.wot.kongzhong.com/tanksMarkRecordServlet?jsonpcallback=&tkid=%d`,
		tank)
	b = getHtml(url)

	b = bytes.Trim(b, `()[]`)
	if len(b) < 20 {
		fmt.Println(`error tank id`, tank)
		return
	}

	// fmt.Println(string(b[:]))

	v := new(typelist.Score)
	err := json.Unmarshal(b, &v)
	if err != nil {
		fmt.Println("json error:", err)
		return
	}

	for _, query := range scoreQuery {
		db.Exec(
			query, tank,
			v.G1max, v.G2max, v.G3max,
			v.G1min, v.G2min, v.G3min,
			v.M1max, v.M2max, v.M3max, v.M4max,
			v.M1min, v.M2min, v.M3min, v.M4min)
	}
}
