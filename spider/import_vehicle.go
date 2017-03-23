package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/zhengkai/goweb/typelist"
)

func import_vehicle() {
	var b []byte
	// b = getFile(`vehicles.json`)
	b = getHtml(`https://api.worldoftanks.ru/wot/encyclopedia/vehicles/?application_id=demo&language=zh-cn`)
	fmt.Println(string(b[:100]))

	data, err := simplejson.NewJson(b)
	if err != nil {
		fmt.Println(`JSON decode error: `, err)
		return
	}
	data = data.GetPath(`data`)

	for sid, info := range data.MustMap() {
		row := info.(map[string]interface{})

		// fmt.Println(sid, info)

		id, _ := strconv.Atoi(sid)
		name := row[`name`].(string)
		is_premium := row[`is_premium`].(bool)
		tier, _ := row[`tier`].(json.Number).Int64()
		stype := row[`type`].(string)
		nation := row[`nation`].(string)
		// fmt.Println(id, name, is_premium, tier, stype, nation)

		iType := 0
		if val, ok := typelist.VehicleTypeSearch[stype]; ok {
			iType = val
		}
		iNation := 0

		if val, ok := typelist.NationSearch[nation]; ok {
			iNation = val
		}

		query := `REPLACE INTO vehicle SET ` +
			`id = ?, is_premium = ?, tier = ?, type = ?, nation = ?, ` +
			`name = ?, ts_update = ?`
		db.Exec(
			query,
			id,
			is_premium,
			tier,
			iType,
			iNation,
			name,
			time.Now().Unix())
	}
}
