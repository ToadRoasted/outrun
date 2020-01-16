package netobj

import "encoding/json"

type PlayerInfo struct {
	Username          string      `json:"username" db:"username"`
	Password          string      `json:"password" db:"password"`
	MigrationPassword string      `json:"migrationPassword" db:"migrate_password"` // used in migration
	UserPassword      string      `json:"userPassword" db:"user_password"`         // used in migration
	Key               string      `json:"key" db:"player_key"`
	LastLogin         int64       `json:"lastLogin" db:"last_login"`
	CharacterState    []Character `json:"characters" db:"characters"`
	ChaoState         []Chao      `json:"chao" db:"chao"`
}

type StoredPlayerInfo struct {
	Username          string `json:"username" db:"username"`
	Password          string `json:"password" db:"password"`
	MigrationPassword string `json:"migrationPassword" db:"migrate_password"` // used in migration
	UserPassword      string `json:"userPassword" db:"user_password"`         // used in migration
	Key               string `json:"key" db:"player_key"`
	LastLogin         int64  `json:"lastLogin" db:"last_login"`
	CharacterState    []byte `json:"characters" db:"characters"`
	ChaoState         []byte `json:"chao" db:"chao"`
}

func PlayerInfoToStoredPlayerInfo(pli PlayerInfo) StoredPlayerInfo {
	characterstate, _ := json.Marshal(pli.CharacterState)
	chaostate, _ := json.Marshal(pli.ChaoState)
	return StoredPlayerInfo{
		pli.Username,
		pli.Password,
		pli.MigrationPassword,
		pli.UserPassword,
		pli.Key,
		pli.LastLogin,
		characterstate,
		chaostate,
	}
}

func StoredPlayerInfoToPlayerInfo(pli StoredPlayerInfo) PlayerInfo {
	var characterstate []Character
	var chaostate []Chao
	json.Unmarshal(pli.CharacterState, characterstate)
	json.Unmarshal(pli.ChaoState, chaostate)
	return PlayerInfo{
		pli.Username,
		pli.Password,
		pli.MigrationPassword,
		pli.UserPassword,
		pli.Key,
		pli.LastLogin,
		characterstate,
		chaostate,
	}
}
