package handle

import (
	"fmt"
	"strconv"

	"github.com/zhengkai/goweb/typelist"
	"github.com/zhengkai/sigo/handle"
)

type VehicleRow struct {
	Id        int
	IsPremium int
	Tier      int
	Type      int
	Nation    int
	Name      string
}

type List struct {
	Base
}

func (this List) New() handle.Handle {
	c := this
	c.Head = new(CommonHead).New()
	c.Head.AddJS(`/res/tablesorter-2.28.5/js/jquery.tablesorter.min.js`)
	c.Head.AddJS(`/res/list.js`)
	c.Head.AddCSS(`/res/tablesorter-2.28.5/css/theme.materialize.min.css`)
	return &c
}

func (this *List) Parse() {

	this.SetTplFunc(`show_score`, func(i int) string {
		if i <= 0 || i >= 9999 {
			return ``
		}
		return strconv.Itoa(i)
	})

	data := make(map[string]interface{})
	this.Data = data

	query := `SELECT id, is_premium, tier, type, nation, name FROM vehicle`
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	vehicle := make(map[int]*VehicleRow)
	data[`vehicle`] = vehicle

	vehicle_id := make(map[int]bool)

	for row.Next() {
		v := new(VehicleRow)
		err = row.Scan(&v.Id, &v.IsPremium, &v.Tier, &v.Type, &v.Nation, &v.Name)
		if err != nil {
			continue
		}
		if v.Tier < 7 {
			continue
		}
		vehicle[v.Id] = v
		vehicle_id[v.Id] = true
	}

	score := make(map[int]*typelist.Score)
	data[`score`] = score

	query = `SELECT vehicle_id, ` +
		`g1max, g2max, g3max, ` +
		`g1min, g2min, g3min, ` +
		`m1max, m2max, m3max, m4max, ` +
		`m1min, m2min, m3min, m4min ` +
		`FROM score_current`
	row, err = db.Query(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	for row.Next() {
		v := new(typelist.Score)
		id := 0
		err = row.Scan(
			&id,
			&v.G1max, &v.G2max, &v.G3max,
			&v.G1min, &v.G2min, &v.G3min,
			&v.M1max, &v.M2max, &v.M3max, &v.M4max,
			&v.M1min, &v.M2min, &v.M3min, &v.M4min)

		if err != nil {
			continue
		}
		score[id] = v
		delete(vehicle_id, id)
	}

	for id, _ := range vehicle_id {
		delete(vehicle, id)
	}

	data[`vehicle_type`] = typelist.VehicleType
	data[`nation`] = typelist.Nation

	data[`tier`] = [...]int{7, 8, 9, 10}
}

func init() {
	handle.Register(`/list`, &List{})
}
