package responses

import (
	"github.com/Mtbcooler/outrun/enums"
	"github.com/Mtbcooler/outrun/logic/conversion"
	"github.com/Mtbcooler/outrun/netobj"
	"github.com/Mtbcooler/outrun/obj"
	"github.com/Mtbcooler/outrun/obj/constobjs"
	"github.com/Mtbcooler/outrun/responses/responseobjs"
	"github.com/jinzhu/now"
)

type WeeklyLeaderboardOptionsResponse struct {
	BaseResponse
	Mode      int64 `json:"mode"`      // 0 == ENDLESS, 1 == QUICK
	Type      int64 `json:"type"`      // 0 == RankingScoreType.HIGH_SCORE, else == RankingScoreType.TOTAL_SCORE
	Param     int64 `json:"param"`     // seemingly unused
	StartTime int64 `json:"startTime"` // both times are also seemingly unused...
	ResetTime int64 `json:"resetTime"`
}

func WeeklyLeaderboardOptions(base responseobjs.BaseInfo, mode, ltype, param, startTime, resetTime int64) WeeklyLeaderboardOptionsResponse {
	baseResponse := NewBaseResponse(base)
	return WeeklyLeaderboardOptionsResponse{
		baseResponse,
		mode,
		ltype,
		param,
		startTime,
		resetTime,
	}
}

func DefaultWeeklyLeaderboardOptions(base responseobjs.BaseInfo, mode int64) WeeklyLeaderboardOptionsResponse {
	startTime := now.BeginningOfWeek().UTC().Unix()
	resetTime := now.EndOfWeek().UTC().Unix()
	//ltype := int64(1)
	ltype := int64(0)
	//param := int64(0)
	param := int64(5)
	return WeeklyLeaderboardOptions(base, mode, ltype, param, startTime, resetTime)
}

type WeeklyLeaderboardEntriesResponse struct {
	BaseResponse
	PlayerEntry  obj.LeaderboardEntry   `json:"playerEntry"`
	LastOffset   int64                  `json:"lastOffset"`
	StartTime    int64                  `json:"startTime"`
	ResetTime    int64                  `json:"resetTime"`
	StartIndex   int64                  `json:"startIndex"`
	Mode         int64                  `json:"mode"`
	TotalEntries int64                  `json:"totalEntries"`
	EntriesList  []obj.LeaderboardEntry `json:"entriesList"`
}

func WeeklyLeaderboardEntries(base responseobjs.BaseInfo, pe obj.LeaderboardEntry, lo, st, rt, si, m, te int64, el []obj.LeaderboardEntry) WeeklyLeaderboardEntriesResponse {
	baseResponse := NewBaseResponse(base)
	out := WeeklyLeaderboardEntriesResponse{
		baseResponse,
		pe,
		lo,
		st,
		rt,
		si,
		m,
		te,
		el,
	}
	return out
}

func DefaultWeeklyLeaderboardEntries(base responseobjs.BaseInfo, player netobj.Player, mode, ltype int64) WeeklyLeaderboardEntriesResponse {
	startTime := now.BeginningOfWeek().UTC().Unix()
	resetTime := now.EndOfWeek().UTC().Unix()
	myEntry := conversion.PlayerToLeaderboardEntry(player, mode)
	return WeeklyLeaderboardEntries(
		base,
		//obj.DefaultLeaderboardEntry(uid),
		myEntry,
		5,
		startTime,
		resetTime,
		1,
		mode,
		1,
		[]obj.LeaderboardEntry{
			myEntry,
		},
	)
}

type LeagueDataResponse struct {
	BaseResponse
	LeagueData obj.LeagueData `json:"leagueData"`
	Mode       int64          `json:"mode"`
}

func LeagueData(base responseobjs.BaseInfo, leagueData obj.LeagueData, mode int64) LeagueDataResponse {
	baseResponse := NewBaseResponse(base)
	out := LeagueDataResponse{
		baseResponse,
		leagueData,
		mode,
	}
	return out
}

func DefaultLeagueData(base responseobjs.BaseInfo, mode int64) LeagueDataResponse {
	var leagueData obj.LeagueData
	if mode == 0 {
		leagueData = constobjs.DefaultLeagueDataMode0
	} else if mode == 1 {
		leagueData = constobjs.DefaultLeagueDataMode1
	}
	return LeagueData(base, leagueData, mode)
}

type LeagueOperatorDataResponse struct {
	BaseResponse
	LeagueList []obj.LeagueData `json:"leagueOperatorList"`
	LeagueID   int64            `json:"leagueId"`
}

func LeagueOperatorData(base responseobjs.BaseInfo, leagueList []obj.LeagueData, leagueId int64) LeagueOperatorDataResponse {
	baseResponse := NewBaseResponse(base)
	out := LeagueOperatorDataResponse{
		baseResponse,
		leagueList,
		leagueId,
	}
	return out
}

func DefaultLeagueOperatorData(base responseobjs.BaseInfo, leagueId int64) LeagueOperatorDataResponse {
	leagueList := []obj.LeagueData{
		constobjs.LeagueDataDefinitions[enums.RankingLeagueF_M],
		constobjs.LeagueDataDefinitions[enums.RankingLeagueF],
		constobjs.LeagueDataDefinitions[enums.RankingLeagueF_P],
		constobjs.LeagueDataDefinitions[enums.RankingLeagueE_M],
		constobjs.LeagueDataDefinitions[enums.RankingLeagueE],
		constobjs.LeagueDataDefinitions[enums.RankingLeagueE_P],
		constobjs.LeagueDataDefinitions[enums.RankingLeagueD_M],
		constobjs.LeagueDataDefinitions[enums.RankingLeagueD],
		constobjs.LeagueDataDefinitions[enums.RankingLeagueD_P],
		constobjs.LeagueDataDefinitions[enums.RankingLeagueC_M],
		constobjs.LeagueDataDefinitions[enums.RankingLeagueC],
		constobjs.LeagueDataDefinitions[enums.RankingLeagueC_P],
		constobjs.LeagueDataDefinitions[enums.RankingLeagueB_M],
		constobjs.LeagueDataDefinitions[enums.RankingLeagueB],
		constobjs.LeagueDataDefinitions[enums.RankingLeagueB_P],
		constobjs.LeagueDataDefinitions[enums.RankingLeagueA_M],
		constobjs.LeagueDataDefinitions[enums.RankingLeagueA],
		constobjs.LeagueDataDefinitions[enums.RankingLeagueA_P],
		constobjs.LeagueDataDefinitions[enums.RankingLeagueS_M],
		constobjs.LeagueDataDefinitions[enums.RankingLeagueS],
		constobjs.LeagueDataDefinitions[enums.RankingLeagueS_P],
	}
	return LeagueOperatorData(base, leagueList, leagueId)
}
