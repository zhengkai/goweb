package handle

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/zhengkai/goweb/typelist"
	"github.com/zhengkai/sigo/handle"
)

var (
	pGet, _ = regexp.Compile(`^(\d*),(\d*),([\da]*)`)
)

type VehicleRow struct {
	Id        int
	IsPremium int
	Tier      int
	Type      int
	Nation    int
	Name      string
	Show      bool
}

type ListQuery struct {
	Type   []int
	Nation []int
	Tier   []int
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
	c.SetTplFunc(`show_score`, func(i int) string {
		if i <= 0 || i >= 9999 {
			return ``
		}
		s := strconv.Itoa(i)
		if i > 999 {
			s = s[0:1] + `,` + s[1:4]
		}
		return s
	})
	c.SetTplFunc(`in_array`, func(i int, j []int) bool {
		if j == nil {
			return false
		}
		for _, v := range j {
			if i == v {
				return true
			}
		}
		return false
	})
	return &c
}

func (this *List) ParseGet() (q *ListQuery) {

	q = new(ListQuery)

	get := this.R.URL.RawQuery
	if len(get) < 1 {
		return
	}

	r := pGet.FindStringSubmatch(get)

	var conv = func(in string) (q []int) {
		for _, b := range in {
			var n int
			if b == 'a' {
				n = 10
			} else {
				n = int(b - '0')
			}
			q = append(q, n)
		}
		return q
	}

	q.Type = conv(r[1])
	q.Nation = conv(r[2])
	q.Tier = conv(r[3])
	return
}

func (this *List) Parse() {

	data := make(map[string]interface{})
	this.Data = data

	lq := this.ParseGet()
	data[`listQuery`] = lq

	query := `SELECT id, is_premium, tier, type, nation, name FROM vehicle`
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	vehicle := make(map[int]*VehicleRow)
	data[`vehicle`] = vehicle

	vehicle_id := make(map[int]bool)

	inCheck := func(a int, b []int) bool {
		if len(b) < 1 {
			return true
		}
		for _, v := range b {
			if v == a {
				return true
			}
		}
		return false
	}

	checkShow := func(v *VehicleRow) bool {
		if lq == nil {
			return true
		}
		if !inCheck(v.Type, lq.Type) {
			return false
		}
		if !inCheck(v.Nation, lq.Nation) {
			return false
		}
		if !inCheck(v.Tier, lq.Tier) {
			return false
		}
		return true
	}

	for row.Next() {
		v := new(VehicleRow)
		err = row.Scan(&v.Id, &v.IsPremium, &v.Tier, &v.Type, &v.Nation, &v.Name)
		if err != nil {
			continue
		}
		if v.Tier < 7 {
			continue
		}
		v.Show = checkShow(v)
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
