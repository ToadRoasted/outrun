package constobjs

import (
	"github.com/Mtbcooler/outrun/enums"
	"github.com/Mtbcooler/outrun/obj"
)

// TODO: Remove! No longer used
var DefaultLeagueDataMode0 = obj.NewLeagueData(
	0,
	0,
	40,
	0,
	0,
	0,
	[]obj.OperatorScore{
		obj.NewOperatorScore(2, 50, []obj.Item{}),
		obj.NewOperatorScore(2, 40, []obj.Item{obj.NewItem("910000", 3000)}),
		obj.NewOperatorScore(2, 30, []obj.Item{obj.NewItem("910000", 5000)}),
		obj.NewOperatorScore(2, 20, []obj.Item{obj.NewItem("910000", 7000)}),
		obj.NewOperatorScore(2, 10, []obj.Item{obj.NewItem("910000", 10000)}),
		obj.NewOperatorScore(2, 1, []obj.Item{obj.NewItem("910000", 15000)}),
		obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem("910000", 20000)}),
	},
	[]obj.OperatorScore{
		obj.NewOperatorScore(2, 50, []obj.Item{}),
		obj.NewOperatorScore(2, 40, []obj.Item{obj.NewItem("910000", 3000)}),
		obj.NewOperatorScore(2, 30, []obj.Item{obj.NewItem("910000", 5000)}),
		obj.NewOperatorScore(2, 20, []obj.Item{obj.NewItem("910000", 7000)}),
		obj.NewOperatorScore(2, 10, []obj.Item{obj.NewItem("910000", 10000)}),
		obj.NewOperatorScore(2, 1, []obj.Item{obj.NewItem("910000", 15000)}),
		obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem("910000", 20000)}),
	},
)

// TODO: Remove! No longer used
var DefaultLeagueDataMode1 = obj.NewLeagueData(
	0,
	0,
	40,
	0,
	0,
	0,
	[]obj.OperatorScore{
		obj.NewOperatorScore(2, 50, []obj.Item{}),
		obj.NewOperatorScore(2, 40, []obj.Item{obj.NewItem("900000", 5)}),
		obj.NewOperatorScore(2, 30, []obj.Item{obj.NewItem("900000", 10)}),
		obj.NewOperatorScore(2, 20, []obj.Item{obj.NewItem("900000", 15)}),
		obj.NewOperatorScore(2, 10, []obj.Item{obj.NewItem("900000", 15)}),
		obj.NewOperatorScore(2, 1, []obj.Item{obj.NewItem("900000", 20)}),
		obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem("900000", 20)}),
	},
	[]obj.OperatorScore{
		obj.NewOperatorScore(2, 50, []obj.Item{}),
		obj.NewOperatorScore(2, 40, []obj.Item{obj.NewItem("900000", 3)}),
		obj.NewOperatorScore(2, 30, []obj.Item{obj.NewItem("900000", 5)}),
		obj.NewOperatorScore(2, 20, []obj.Item{obj.NewItem("900000", 5)}),
		obj.NewOperatorScore(2, 10, []obj.Item{obj.NewItem("900000", 10)}),
		obj.NewOperatorScore(2, 1, []obj.Item{obj.NewItem("900000", 10)}),
		obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem("900000", 10)}),
	},
)

var LeagueDataDefinitions = map[int64]obj.LeagueData{
	// F *
	enums.RankingLeagueF_M: obj.NewLeagueData(
		0,
		0,
		40,
		0,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(2, 40, []obj.Item{obj.NewItem("910000", 3000)}),
			obj.NewOperatorScore(2, 30, []obj.Item{obj.NewItem("910000", 5000)}),
			obj.NewOperatorScore(2, 20, []obj.Item{obj.NewItem("910000", 7000)}),
			obj.NewOperatorScore(2, 10, []obj.Item{obj.NewItem("910000", 10000)}),
			obj.NewOperatorScore(2, 1, []obj.Item{obj.NewItem("910000", 15000)}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem("910000", 20000)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(2, 40, []obj.Item{obj.NewItem("900000", 1)}),
			obj.NewOperatorScore(2, 30, []obj.Item{obj.NewItem("900000", 3)}),
			obj.NewOperatorScore(2, 20, []obj.Item{obj.NewItem("900000", 5)}),
			obj.NewOperatorScore(2, 10, []obj.Item{obj.NewItem("900000", 10)}),
			obj.NewOperatorScore(2, 1, []obj.Item{obj.NewItem("900000", 15)}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem("900000", 25)}),
		},
	),
	// F **
	enums.RankingLeagueF: obj.NewLeagueData(
		1,
		0,
		20,
		20,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRing, 54321)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRedRing, 123)}),
		},
	),
	// F ***
	enums.RankingLeagueF_P: obj.NewLeagueData(
		2,
		0,
		20,
		20,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRing, 54321)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRedRing, 123)}),
		},
	),
	// E *
	enums.RankingLeagueE_M: obj.NewLeagueData(
		3,
		0,
		20,
		20,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRing, 54321)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRedRing, 123)}),
		},
	),
	// E **
	enums.RankingLeagueE: obj.NewLeagueData(
		4,
		0,
		20,
		20,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRing, 54321)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRedRing, 123)}),
		},
	),
	// E ***
	enums.RankingLeagueE_P: obj.NewLeagueData(
		5,
		0,
		20,
		20,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRing, 54321)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRedRing, 123)}),
		},
	),
	// D *
	enums.RankingLeagueD_M: obj.NewLeagueData(
		6,
		0,
		20,
		20,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRing, 54321)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRedRing, 123)}),
		},
	),
	// D **
	enums.RankingLeagueD: obj.NewLeagueData(
		7,
		0,
		20,
		20,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRing, 54321)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRedRing, 123)}),
		},
	),
	// D ***
	enums.RankingLeagueD_P: obj.NewLeagueData(
		8,
		0,
		20,
		20,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRing, 54321)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRedRing, 123)}),
		},
	),
	// C *
	enums.RankingLeagueC_M: obj.NewLeagueData(
		9,
		0,
		20,
		20,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRing, 54321)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRedRing, 123)}),
		},
	),
	// C **
	enums.RankingLeagueC: obj.NewLeagueData(
		10,
		0,
		20,
		20,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRing, 54321)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRedRing, 123)}),
		},
	),
	// C ***
	enums.RankingLeagueC_P: obj.NewLeagueData(
		11,
		0,
		20,
		20,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRing, 54321)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRedRing, 123)}),
		},
	),
	// B *
	enums.RankingLeagueB_M: obj.NewLeagueData(
		12,
		0,
		20,
		20,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRing, 54321)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRedRing, 123)}),
		},
	),
	// B **
	enums.RankingLeagueB: obj.NewLeagueData(
		13,
		0,
		20,
		20,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRing, 54321)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRedRing, 123)}),
		},
	),
	// B ***
	enums.RankingLeagueB_P: obj.NewLeagueData(
		14,
		0,
		20,
		20,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRing, 54321)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRedRing, 123)}),
		},
	),
	// A *
	enums.RankingLeagueA_M: obj.NewLeagueData(
		15,
		0,
		20,
		20,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRing, 54321)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRedRing, 123)}),
		},
	),
	// A **
	enums.RankingLeagueA: obj.NewLeagueData(
		16,
		0,
		20,
		20,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRing, 54321)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRedRing, 123)}),
		},
	),
	// A ***
	enums.RankingLeagueA_P: obj.NewLeagueData(
		17,
		0,
		20,
		20,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRing, 54321)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRedRing, 123)}),
		},
	),
	// S *
	enums.RankingLeagueS_M: obj.NewLeagueData(
		18,
		0,
		20,
		20,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRing, 54321)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRedRing, 123)}),
		},
	),
	// S **
	enums.RankingLeagueS: obj.NewLeagueData(
		19,
		0,
		20,
		20,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRing, 54321)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRedRing, 123)}),
		},
	),
	// S ***
	enums.RankingLeagueS_P: obj.NewLeagueData(
		20,
		0,
		20,
		20,
		0,
		0,
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRing, 54321)}),
		},
		[]obj.OperatorScore{
			obj.NewOperatorScore(2, 50, []obj.Item{}),
			obj.NewOperatorScore(0, 1, []obj.Item{obj.NewItem(enums.ItemIDStrRedRing, 123)}),
		},
	),
}
