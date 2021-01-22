package obj

type LeagueData struct {
	LeagueID           int64           `json:"leagueId,string"`        // the ID of the league (see enums/rankingLeague.go)
	GroupID            int64           `json:"groupId,string"`         // ???
	NumUp              int64           `json:"numUp,string"`           // number of people that will be promoted
	NumDown            int64           `json:"numDown,string"`         // number of people that will be relegated
	NumGroupMember     int64           `json:"numGroupMember,string"`  // number of members in group
	NumLeagueMember    int64           `json:"numLeagueMember,string"` // number of members in league
	HighScoreOperator  []OperatorScore `json:"highScoreOpe"`           // high score league rewards
	TotalScoreOperator []OperatorScore `json:"totalScoreOpe"`          // total score league rewards
}

func NewLeagueData(lid, gid, nup, ndown, ngm, nlm int64, hso, tso []OperatorScore) LeagueData {
	return LeagueData{
		lid,
		gid,
		nup,
		ndown,
		ngm,
		nlm,
		hso,
		tso,
	}
}
