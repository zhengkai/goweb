package typelist

import (
	"encoding/json"
	"strconv"
)

type DictRow struct {
	Short string
	Key   string
}

type TypeSearch map[string]int
type TypeList map[int]DictRow

func convertSearch(in TypeList) TypeSearch {
	out := make(TypeSearch)
	for i, v := range in {
		out[v.Key] = i
	}
	return out
}

func ConvertJSON(in TypeList) []byte {
	j := make(map[string]string)
	for i, v := range in {
		j[strconv.Itoa(i)] = v.Short
	}
	b, _ := json.Marshal(j)
	return b
}

func init() {
	VehicleTypeSearch = convertSearch(VehicleType)
	NationSearch = convertSearch(Nation)
}
