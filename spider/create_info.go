package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/bitly/go-simplejson"
)

// 暂时不用

func createInfo() {
	var b []byte

	// 获得国家列表

	b = getHtml(`http://rank.kongzhong.com/wiki/techtree.html`)

	pCountry, _ := regexp.Compile(`<div class="country" country="([a-z]+)">`)
	match := pCountry.FindAllSubmatch(b, -1)

	var lCountry = make([]string, len(match))

	for i, row := range match {
		lCountry[i] = string(row[1][:])
	}
	//fmt.Println(lCountry)

	// 获得每个国家的车辆列表

	for _, country := range lCountry {

		url := `http://rank.kongzhong.com/Data/tankwiki/` + gameVer + `/vehicles/` + country + `/list.xml.json`
		fmt.Println(url)
		b = getHtml(url)

		// fmt.Println(string(b[:]))

		json, err := simplejson.NewJson(b)
		if err != nil {
			fmt.Println(country, `JSON decode error: `, err)
			return
		}
		json = json.GetPath(`list.xml`)

		for sid, info := range json.MustMap() {

			infoF := info.(map[string]interface{})

			// fmt.Println(infoF)

			// id, _ := strconv.Atoi(infoF[`id`].(string))
			level, _ := strconv.Atoi(infoF[`level`].(string))

			short_name := ``
			if infoF[`shortUserString`] != nil {
				short_name = infoF[`shortUserString`].(string)
			}

			price := 0
			is_gold := 0

			priceF := infoF[`price`]

			switch priceF.(type) {
			case string:
				price, _ = strconv.Atoi(priceF.(string))
			case map[string]interface{}:
				price, _ = strconv.Atoi(priceF.(map[string]interface{})[`#text`].(string))
				is_gold = 1
			}

			query := `REPLACE INTO tank ` +
				`SET country = ?, name = ?, short_name = ?, ` +
				`price = ?, description = ?, tag = ?, level = ?, ` +
				`en_name = ?, is_gold = ?, ts_update = ?`
			_, err := db.Exec(
				query,
				country,
				infoF[`userString`].(string),
				short_name,
				price,
				infoF[`description`].(string),
				infoF[`tags`].(string),
				level,
				sid,
				is_gold,
				time.Now().Unix())
			if err != nil {
				fmt.Println(`db error`, err)
				return
			}

			fmt.Println(level, infoF[`userString`].(string))
		}
	}
}
