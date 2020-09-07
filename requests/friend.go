package requests

type FacebookIncentiveRequest struct {
	Base
	Type             int64 `json:"type,string"`
	AchievementCount int64 `json:"achievementCount,string"`
}

// Types:
// 0 - Login
// 1 - Review
// 2 - Feed
// 3 - Achievement
// 4 - No-login push notification
// TODO: Research!
