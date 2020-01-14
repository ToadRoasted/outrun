package consts

const (
	DBFileName = "outrun.db"

	DBBucketSessionIDs = "sessionIDs"
	DBBucketPlayers    = "players"
	DBBucketAnalytics  = "analytics"

	DBSessionExpiryTime = 3600

	// TODO: Add more tables as needed.
	DBMySQLTableAnalytics          = "analytics"
	DBMySQLTableCorePlayerInfo     = "player_info"
	DBMySQLTablePlayerStates       = "player_states"
	DBMySQLTableCharacterStates    = "player_characters"
	DBMySQLTableChaoStates         = "player_chao"
	DBMySQLTableMileageMapStates   = "player_mileage"
	DBMySQLTableOptionUserResults  = "player_user_results"
	DBMySQLTableLastWheelOptions   = "player_roulette_options"
	DBMySQLTableRouletteInfos      = "player_item_roulette_data"
	DBMySQLTableChaoRouletteGroups = "player_chao_roulette_data"
	DBMySQLTableLoginBonusStates   = "player_login_bonus_states"
	DBMySQLTablePersonalEvents     = "player_personal_events"
	DBMySQLTableMessages           = "player_messages"
	DBMySQLTableOperatorMessages   = "player_operator_messages"
	DBMySQLTableSessionIDs         = "session_ids"
)
