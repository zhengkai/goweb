package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type ProList struct {
	Tank1 []ProListTank `xml:"CZECH>lx>tank"`
	Tank2 []ProListTank `xml:"SWEDEN>lx>tank"`
	Tank3 []ProListTank `xml:"CHINA>lx>tank"`
	Tank4 []ProListTank `xml:"FRANCE>lx>tank"`
	Tank5 []ProListTank `xml:"GERMANY>lx>tank"`
	Tank6 []ProListTank `xml:"USSR>lx>tank"`
	Tank7 []ProListTank `xml:"JAPAN>lx>tank"`
	Tank8 []ProListTank `xml:"USA>lx>tank"`
	Tank9 []ProListTank `xml:"UK>lx>tank"`
	Ver   ProListVer    `xml:"version"`
}

type ProListTank struct {
	Id     int    `xml:"tankId"`
	Name   string `xml:"names"`
	Type   string `xml:"type"`
	IsGold string `xml:"isgold"`
}

type ProListVer struct {
	Date string `xml:"time,attr"`
}

func getList() {
	var b []byte
	// b = getHtml(`http://wot.kongzhong.com/wiki/xml/proList.xml`)

	xmlFile, err := os.Open(`proList.xml`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer xmlFile.Close()
	b, _ = ioutil.ReadAll(xmlFile)

	v := new(ProList)

	err = xml.Unmarshal(b, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	updateList(v.Tank1)
	updateList(v.Tank2)
	updateList(v.Tank3)
	updateList(v.Tank4)
	updateList(v.Tank5)
	updateList(v.Tank6)
	updateList(v.Tank7)
	updateList(v.Tank8)
	updateList(v.Tank9)

	// fmt.Println(string(b[:]))
}

func updateList(list []ProListTank) {
	query := `UPDATE tank SET tank_id = ?, type = ? WHERE name = ?`
	for _, row := range list {
		fmt.Println(row)
		db.Exec(query, row.Id, row.Type, row.IsGold, row.Name)
	}
}
