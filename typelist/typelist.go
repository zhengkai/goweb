package typelist

var (
	VehicleTypeSearch = make(TypeSearch)
	NationSearch      = make(TypeSearch)

	VehicleType = TypeList{
		1: DictRow{
			Short: `LT`,
			Key:   `lightTank`,
		},
		2: DictRow{
			Short: `MT`,
			Key:   `mediumTank`,
		},
		3: DictRow{
			Short: `HT`,
			Key:   `heavyTank`,
		},
		4: DictRow{
			Short: `TD`,
			Key:   `AT-SPG`,
		},
		5: DictRow{
			Short: `SPG`,
			Key:   `SPG`,
		},
	}

	Nation = TypeList{
		1: DictRow{
			Short: `德`,
			Key:   `germany`,
		},
		2: DictRow{
			Short: `苏`,
			Key:   `ussr`,
		},
		3: DictRow{
			Short: `美`,
			Key:   `usa`,
		},
		4: DictRow{
			Short: `中`,
			Key:   `china`,
		},
		5: DictRow{
			Short: `英`,
			Key:   `uk`,
		},
		6: DictRow{
			Short: `法`,
			Key:   `france`,
		},
		7: DictRow{
			Short: `捷`,
			Key:   `czech`,
		},
		8: DictRow{
			Short: `瑞`,
			Key:   `sweden`,
		},
		9: DictRow{
			Short: `倭`,
			Key:   `japan`,
		},
	}
)
