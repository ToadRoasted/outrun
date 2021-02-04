package obj

const (
	DefaultMapDistance   = 0
	DefaultNumBossAttack = 0 // Boss HP gets subtracted by this number
	DefaultStageDistance = 0
	DefaultStageMaxScore = 0
)

type MapInfo struct {
	MapDistance   int64 `json:"mapDistance" db:"map_distance"`      // used sparingly in game...?
	NumBossAttack int64 `json:"numBossAttack" db:"num_boss_attack"` // number of hits done on the boss so far?
	StageDistance int64 `json:"stageDistance" db:"stage_distance"`  // TODO: discover use
	StageMaxScore int64 `json:"stageMaxScore" db:"stage_max_score"` // TODO: discover use
}

func DefaultMapInfo() MapInfo {
	return MapInfo{
		DefaultMapDistance,
		DefaultNumBossAttack,
		DefaultStageDistance,
		DefaultStageMaxScore,
	}
}
