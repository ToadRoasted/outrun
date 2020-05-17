package rpcobj

const (
	StatusOK = iota
	StatusUnknownError
	StatusOtherError
	StatusLeagueNotStarted
	StatusLeagueStillOngoing
	StatusNonexistantPlayer
)
