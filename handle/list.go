package handle

import (
	"fmt"

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

	for row.Next() {
		v := new(VehicleRow)
		err = row.Scan(&v.Id, &v.IsPremium, &v.Tier, &v.Type, &v.Nation, &v.Name)
		if err != nil {
			continue
		}
		vehicle[v.Id] = v
	}

	data[`vehicle_type`] = typelist.VehicleType
	data[`nation`] = typelist.Nation
}

func init() {
	handle.Register(`/list`, &List{})
}
