package obj

type LeagueData struct {
	LeagueID           int64           `json:"leagueId,string"`
	GroupID            int64           `json:"groupId,string"`
	NumUp              int64           `json:"numUp,string"`
	NumDown            int64           `json:"numDown,string"`
	NumGroupMember     int64           `json:"numGroupMember,string"`
	NumLeagueMember    int64           `json:"numLeagueMember,string"`
	HighScoreOperator  []OperatorScore `json:"highScoreOpe"`
	TotalScoreOperator []OperatorScore `json:"totalScoreOpe"`
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
