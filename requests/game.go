package requests

type QuickPostGameResultsRequest struct {
	Base
	Score                  int64  `json:"score,string"`
	Rings                  int64  `json:"numRings,string"`
	FailureRings           int64  `json:"numFailureRings,string"`
	RedRings               int64  `json:"numRedStarRings,string"`
	Distance               int64  `json:"distance,string"`
	DailyChallengeValue    int64  `json:"dailyChallengeValue,string"`
	DailyChallengeComplete int64  `json:"dailyChallengeComplete"`
	Animals                int64  `json:"numAnimals,string"`
	MaxCombo               int64  `json:"maxCombo,string"`
	Closed                 int64  `json:"closed"`
	CheatResult            string `json:"cheatResult"`
}

type PostGameResultsRequest struct {
	QuickPostGameResultsRequest
	BossDestroyed int64  `json:"bossDestroyed"`
	ChapterClear  int64  `json:"chapterClear"`
	GetChaoEgg    int64  `json:"getChaoEgg"`
	NumBossAttack int64  `json:"numBossAttack,string"`
	ReachPoint    int64  `json:"reachPoint,string"`
	EventID       string `json:"eventId,omitempty"`
	EventValue    int64  `json:"eventValue,string,omitempty"`
}

type QuickActStartRequest struct {
	Base
	Modifier []int64 `json:"modifire"`           // Seems to be list of item IDs.
	Tutorial int64   `json:"tutorial,omitempty"` // will omit the field if not found
}

type ActStartRequest struct {
	QuickActStartRequest
	DistanceFriendList []string `json:"distanceFriendList"` // list of IDs for facebook friends (converted to a DistanceFriendEntry list for response)
	EventID            string   `json:"eventId,omitempty"`
}

type MileageRewardRequest struct {
	Base
	Episode int64 `json:"episode,string"`
	Chapter int64 `json:"chapter,string"`
}
