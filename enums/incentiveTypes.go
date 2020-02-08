package enums

const (
	IncentiveTypeNone    = iota // DO NOT USE; will crash the game!
	IncentiveTypePoint          // For rewards obtained at each point
	IncentiveTypeChapter        // For rewards obtained at the end of a chapter (NOT for the whole episode)
	IncentiveTypeEpisode        // For rewards obtained at the end of an episode
	IncentiveTypeFriend         // Possibly for rewards for passing a Facebook friend on the Story Mode map? Needs more research.
)
