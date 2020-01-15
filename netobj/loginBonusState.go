package netobj

import (
	"github.com/jinzhu/now"
)

type LoginBonusState struct {
	CurrentFirstLoginBonusDay int64 `json:"currentFirstLoginBonusDay" db:"current_start_dash_bonus_day"` // this doesn't get reset when the login bonus resets
	CurrentLoginBonusDay      int64 `json:"currentLoginBonusDay" db:"current_login_bonus_day"`
	LastLoginBonusTime        int64 `json:"lastLoginBonusTime" db:"last_login_bonus_time"`
	NextLoginBonusTime        int64 `json:"nextLoginBonusTime" db:"next_login_bonus_time"`
	LoginBonusStartTime       int64 `json:"loginBonusStartTime" db:"login_bonus_start_time"`
	LoginBonusEndTime         int64 `json:"loginBonusEndTime" db:"login_bonus_end_time"`
}

func NewLoginBonusState(cflbd, clbd, llbt, nlbt, lbst, lbet int64) LoginBonusState {
	return LoginBonusState{
		cflbd,
		clbd,
		llbt,
		nlbt,
		lbst,
		lbet,
	}
}

func DefaultLoginBonusState(currentFirstLoginBonusDay int64) LoginBonusState {
	currentLoginBonusDay := int64(0)
	lastLoginBonusTime := int64(0)
	nextLoginBonusTime := int64(0)
	loginBonusStartTime := now.BeginningOfWeek().UTC().Unix()
	loginBonusEndTime := now.EndOfWeek().UTC().Unix()
	return NewLoginBonusState(
		currentFirstLoginBonusDay,
		currentLoginBonusDay,
		lastLoginBonusTime,
		nextLoginBonusTime,
		loginBonusStartTime,
		loginBonusEndTime,
	)
}
