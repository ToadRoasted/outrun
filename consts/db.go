package consts

const (
	DBFileName = "outrun.db"

	DBBucketSessionIDs = "sessionIDs"
	DBBucketPlayers    = "players"
	DBBucketAnalytics  = "analytics"

	DBSessionExpiryTime = 3600

	// TODO: Add more tables as needed.
	DBMySQLTableAnalytics         = "analytics"
	DBMySQLTableCorePlayerInfo    = "player_info"
	DBMySQLTableEventStates       = "player_event_states"
	DBMySQLTablePlayerStates      = "player_states"
	DBMySQLTableMileageMapStates  = "player_mileage"
	DBMySQLTableOptionUserResults = "player_user_results"
	DBMySQLTableLastWheelOptions  = "player_roulette_options"
	DBMySQLTableRouletteInfos     = "player_item_roulette_data"
	DBMySQLTableLoginBonusStates  = "player_login_bonus_states"
	DBMySQLTablePersonalEvents    = "player_personal_events"
	DBMySQLTableMessages          = "player_messages"
	DBMySQLTableOperatorMessages  = "player_operator_messages"
	DBMySQLTableOperatorInfos     = "player_operator_infos"
	DBMySQLTableRankingLeagueData = "ranking_league_data"
	DBMySQLTableSessionIDs        = "session_ids"

	SQLAnalyticsSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableAnalytics + ` (
		pid VARCHAR(20) NOT NULL,
		param JSON,
		PRIMARY KEY (pid)
	) ENGINE = InnoDB;`
	SQLCorePlayerInfoSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableCorePlayerInfo + ` (
		id BIGINT UNSIGNED NOT NULL,
		username VARCHAR(12) NOT NULL,
		password VARCHAR(10) NOT NULL,
		migrate_password VARCHAR(10) NOT NULL,
		user_password TEXT NOT NULL,
		player_key VARCHAR(10) NOT NULL,
		last_login BIGINT UNSIGNED NOT NULL,
		language INTEGER NOT NULL,
		characters JSON,
		chao JSON,
		PRIMARY KEY (id)
	) ENGINE = InnoDB;`
	SQLPlayerStatesSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTablePlayerStates + ` (
		id BIGINT UNSIGNED NOT NULL,
		items JSON NOT NULL,
		equipped_items JSON NOT NULL,
		mainchara_id TEXT NOT NULL,
		subchara_id TEXT NOT NULL,
		mainchao_id TEXT NOT NULL,
		subchao_id TEXT NOT NULL,
		num_rings INTEGER NOT NULL,
		num_buy_rings INTEGER NOT NULL,
		num_red_rings INTEGER NOT NULL,
		num_buy_red_rings INTEGER NOT NULL,
		energy INTEGER NOT NULL,
		energy_buy INTEGER NOT NULL,
		energy_renews_at BIGINT UNSIGNED NOT NULL,
		num_messages INTEGER NOT NULL,
		ranking_league INTEGER NOT NULL,
		quick_ranking_league INTEGER NOT NULL,
		num_roulette_ticket INTEGER NOT NULL,
		num_chao_roulette_ticket INTEGER NOT NULL,
		chao_eggs INTEGER NOT NULL,
		high_score BIGINT NOT NULL,
		quick_high_score BIGINT NOT NULL,
		total_distance BIGINT NOT NULL,
		best_distance BIGINT NOT NULL,
		daily_mission_id INTEGER UNSIGNED NOT NULL,
		daily_mission_end_time BIGINT UNSIGNED NOT NULL,
		daily_challenge_value INTEGER,
		daily_challenge_complete TINYINT UNSIGNED NOT NULL,
		num_daily_chal_cont INTEGER NOT NULL,
		num_plays INTEGER NOT NULL,
		num_animals INTEGER NOT NULL,
		rank INTEGER UNSIGNED NOT NULL,
		dm_cat INTEGER NOT NULL,
		dm_set INTEGER NOT NULL,
		dm_pos INTEGER NOT NULL,
		dm_nextcont INTEGER NOT NULL,
		league_high_score BIGINT NOT NULL,
		quick_league_high_score BIGINT NOT NULL,
		league_start_time BIGINT UNSIGNED NOT NULL,
		league_reset_time BIGINT UNSIGNED NOT NULL,
		ranking_league_group INTEGER NOT NULL,
		quick_ranking_league_group INTEGER NOT NULL,
		total_score BIGINT NOT NULL,
		quick_total_score BIGINT NOT NULL,
		best_total_score BIGINT NOT NULL,
		best_quick_total_score BIGINT NOT NULL,
		PRIMARY KEY (id)
	) ENGINE = InnoDB;`
	SQLMileageMapStatesSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableMileageMapStates + ` (
		id BIGINT UNSIGNED NOT NULL,
		map_distance INTEGER,
		num_boss_attack INTEGER,
		stage_distance INTEGER,
		stage_max_score INTEGER,
		episode INTEGER,
		chapter INTEGER,
		point INTEGER,
		stage_total_score BIGINT,
		chapter_start_time BIGINT UNSIGNED NOT NULL,
		PRIMARY KEY (id)
	) ENGINE = InnoDB;`
	SQLOptionUserResultsSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableOptionUserResults + ` (
		id BIGINT UNSIGNED NOT NULL,
		high_total_score INTEGER,
		high_quick_total_score INTEGER,
		total_rings INTEGER,
		total_red_rings INTEGER,
		chao_roulette_spin_count INTEGER,
		roulette_spin_count INTEGER,
		num_jackpots INTEGER,
		best_jackpot INTEGER,
		support INTEGER,
		PRIMARY KEY (id)
	) ENGINE = InnoDB;`
	SQLRouletteInfosSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableRouletteInfos + ` (
		id BIGINT UNSIGNED NOT NULL,
		login_roulette_id INTEGER,
		roulette_period_end BIGINT UNSIGNED NOT NULL,
		roulette_count_in_period INTEGER,
		got_jackpot_this_period INTEGER,
		PRIMARY KEY (id)
	) ENGINE = InnoDB;`
	SQLLoginBonusStatesSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableLoginBonusStates + ` (
		id BIGINT UNSIGNED NOT NULL,
		current_start_dash_bonus_day INTEGER,
		current_login_bonus_day INTEGER,
		last_login_bonus_time BIGINT UNSIGNED NOT NULL,
		next_login_bonus_time BIGINT UNSIGNED NOT NULL,
		login_bonus_start_time BIGINT UNSIGNED NOT NULL,
		login_bonus_end_time BIGINT UNSIGNED NOT NULL,
		PRIMARY KEY (id)
	) ENGINE = InnoDB;`
	SQLOperatorMessagesSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableOperatorMessages + ` (
		id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
		userid BIGINT UNSIGNED NOT NULL,
		contents TEXT,
		item JSON,
		expire_time BIGINT UNSIGNED NOT NULL,
		PRIMARY KEY (id)
	) ENGINE = InnoDB;`
	SQLRankingLeagueDataSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableRankingLeagueData + ` (
		league_id INT UNSIGNED NOT NULL,
		group_id INT UNSIGNED NOT NULL,
		start_time BIGINT UNSIGNED NOT NULL,
		reset_time BIGINT UNSIGNED NOT NULL,
		league_player_count INTEGER,
		group_player_count INTEGER,
		PRIMARY KEY (league_id, group_id)
	) ENGINE = InnoDB;`
	SQLSessionIDsSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableSessionIDs + ` (
		sid VARCHAR(48) NOT NULL,
		uid BIGINT UNSIGNED NOT NULL,
		assigned_at_time BIGINT UNSIGNED NOT NULL,
		PRIMARY KEY (sid)
	) ENGINE = InnoDB;`
	SQLOperatorInfosSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableOperatorInfos + ` (
		uid BIGINT UNSIGNED NOT NULL,
		id INTEGER NOT NULL,
		param TEXT,
		PRIMARY KEY (uid, id)
	) ENGINE = InnoDB;`
	SQLEventStatesSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableEventStates + ` (
		uid BIGINT UNSIGNED NOT NULL,
		param INTEGER NOT NULL,
		PRIMARY KEY (uid)
	) ENGINE = InnoDB;`
	SQLPlayerStatesInsertTypeSchema = `(
		id,
		items,
		equipped_items,
		mainchara_id,
		subchara_id,
		mainchao_id,
		subchao_id,
		num_rings,
		num_buy_rings,
		num_red_rings,
		num_buy_red_rings,
		energy,
		energy_buy,
		energy_renews_at,
		num_messages,
		ranking_league,
		quick_ranking_league,
		num_roulette_ticket,
		num_chao_roulette_ticket,
		chao_eggs,
		high_score,
		quick_high_score,
		total_distance,
		best_distance,
		daily_mission_id,
		daily_mission_end_time,
		daily_challenge_value,
		daily_challenge_complete,
		num_daily_chal_cont,
		num_plays,
		num_animals,
		rank,
		dm_cat,
		dm_set,
		dm_pos,
		dm_nextcont,
		league_high_score,
		quick_league_high_score,
		league_start_time,
		league_reset_time,
		ranking_league_group,
		quick_ranking_league_group,
		total_score,
		quick_total_score,
		best_total_score,
		best_quick_total_score
	)
	VALUES (
		:id,
		:items,
		:equipped_items,
		:mainchara_id,
		:subchara_id,
		:mainchao_id,
		:subchao_id,
		:num_rings,
		:num_buy_rings,
		:num_red_rings,
		:num_buy_red_rings,
		:energy,
		:energy_buy,
		:energy_renews_at,
		:num_messages,
		:ranking_league,
		:quick_ranking_league,
		:num_roulette_ticket,
		:num_chao_roulette_ticket,
		:chao_eggs,
		:high_score,
		:quick_high_score,
		:total_distance,
		:best_distance,
		:daily_mission_id,
		:daily_mission_end_time,
		:daily_challenge_value,
		:daily_challenge_complete,
		:num_daily_chal_cont,
		:num_plays,
		:num_animals,
		:rank,
		:dm_cat,
		:dm_set,
		:dm_pos,
		:dm_nextcont,
		:league_high_score,
		:quick_league_high_score,
		:league_start_time,
		:league_reset_time,
		:ranking_league_group,
		:quick_ranking_league_group,
		:total_score,
		:quick_total_score,
		:best_total_score,
		:best_quick_total_score
	)`
)
