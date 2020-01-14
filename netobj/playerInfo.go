package netobj

type PlayerInfo struct {
	ID                string `json:"userID"`
	Username          string `json:"username"`
	Password          string `json:"password"`
	MigrationPassword string `json:"migrationPassword"` // used in migration
	UserPassword      string `json:"userPassword"`      // used in migration
	Key               string `json:"key"`
	LastLogin         int64  `json:"lastLogin"`
}
