package netobj

type PlayerInfo struct {
	Username          string `json:"username" db:"username"`
	Password          string `json:"password" db:"password"`
	MigrationPassword string `json:"migrationPassword" db:"migrate_password"` // used in migration
	UserPassword      string `json:"userPassword" db:"user_password"`         // used in migration
	Key               string `json:"key" db:"key"`
	LastLogin         int64  `json:"lastLogin" db:"last_login"`
}
