package dbaccess

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/Mtbcooler/outrun/obj"

	"github.com/Mtbcooler/outrun/consts"

	"github.com/Mtbcooler/outrun/netobj"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/Mtbcooler/outrun/config"
)

var db *sqlx.DB
var DatabaseIsBusy = false

type AnalyticsEntry struct {
	PID   string `db:"pid"`
	Param []byte `db:"param"`
}

func Set(table, column, id string, value interface{}) error {
	CheckIfDBSet()
	result, err := db.Exec("REPLACE INTO `?` (id, ?) VALUES (?, ?)", table, column, id, value)
	if err == nil && config.CFile.DebugPrints {
		rowsAffected, _ := result.RowsAffected()
		log.Printf("[DEBUG] Set operation complete; %v rows affected\n", rowsAffected)
	}
	return err
}

func SetAnalyticsEntry(table, pid string, value []byte) error {
	CheckIfDBSet()
	entry := AnalyticsEntry{
		pid,
		value,
	}
	result, err := db.NamedExec("REPLACE INTO `"+table+"`(pid, param)\n"+
		"VALUES (:pid, :param)",
		entry)
	if err == nil && config.CFile.DebugPrints {
		rowsAffected, _ := result.RowsAffected()
		log.Printf("[DEBUG] SetAnalyticsEntry operation complete; %v rows affected\n", rowsAffected)
	}
	return err
}

func SetPlayerInfo(table, id string, value netobj.PlayerInfo) error {
	CheckIfDBSet()
	convertedValue := netobj.PlayerInfoToStoredPlayerInfo(value)
	result, err := db.NamedExec("REPLACE INTO `"+table+"`(id, username, password, migrate_password, user_password, player_key, last_login, language, characters, chao)\n"+
		"VALUES ("+id+", :username, :password, :migrate_password, :user_password, :player_key, :last_login, :language, :characters, :chao)",
		convertedValue)
	if err == nil && config.CFile.DebugPrints {
		rowsAffected, _ := result.RowsAffected()
		log.Printf("[DEBUG] SetPlayerInfo operation complete; %v rows affected\n", rowsAffected)
	}
	return err
}

func SetPlayerState(table, id string, value netobj.PlayerState) error {
	CheckIfDBSet()
	sqldata := netobj.PlayerStateToSQLCompatiblePlayerState(value)
	result, err := db.NamedExec("REPLACE INTO `"+table+"` "+strings.Replace(consts.SQLPlayerStatesInsertTypeSchema, ":id", id, 1), sqldata)
	if err == nil && config.CFile.DebugPrints {
		rowsAffected, _ := result.RowsAffected()
		log.Printf("[DEBUG] SetPlayerState operation complete; %v rows affected\n", rowsAffected)
	}
	return err
}

func SetMileageMapState(table, id string, value netobj.MileageMapState) error {
	CheckIfDBSet()
	result, err := db.NamedExec("REPLACE INTO `"+table+"`(id, map_distance, num_boss_attack, stage_distance, stage_max_score, episode, chapter, point, stage_total_score, chapter_start_time)\n"+
		"VALUES ("+id+", :map_distance, :num_boss_attack, :stage_distance, :stage_max_score, :episode, :chapter, :point, :stage_total_score, :chapter_start_time)",
		value)
	if err == nil && config.CFile.DebugPrints {
		rowsAffected, _ := result.RowsAffected()
		log.Printf("[DEBUG] SetMileageMapState operation complete; %v rows affected\n", rowsAffected)
	}
	return err
}

func SetOptionUserResult(table, id string, value netobj.OptionUserResult) error {
	CheckIfDBSet()
	result, err := db.NamedExec("REPLACE INTO `"+table+"`(id, high_total_score, high_quick_total_score, total_rings, total_red_rings, chao_roulette_spin_count, roulette_spin_count, num_jackpots, best_jackpot, support)\n"+
		"VALUES ("+id+", :high_total_score, :high_quick_total_score, :total_rings, :total_red_rings, :chao_roulette_spin_count, :roulette_spin_count, :num_jackpots, :best_jackpot, :support)",
		value)
	if err == nil && config.CFile.DebugPrints {
		rowsAffected, _ := result.RowsAffected()
		log.Printf("[DEBUG] SetOptionUserResult operation complete; %v rows affected\n", rowsAffected)
	}
	return err
}

func SetRouletteInfo(table, id string, value netobj.RouletteInfo) error {
	CheckIfDBSet()
	result, err := db.NamedExec("REPLACE INTO `"+table+"`(id, login_roulette_id, roulette_period_end, roulette_count_in_period, got_jackpot_this_period)\n"+
		"VALUES ("+id+", :login_roulette_id, :roulette_period_end, :roulette_count_in_period, :got_jackpot_this_period)",
		value)
	if err == nil && config.CFile.DebugPrints {
		rowsAffected, _ := result.RowsAffected()
		log.Printf("[DEBUG] SetRouletteInfo operation complete; %v rows affected\n", rowsAffected)
	}
	return err
}

func SetLoginBonusState(table, id string, value netobj.LoginBonusState) error {
	CheckIfDBSet()
	result, err := db.NamedExec("REPLACE INTO `"+table+"`(id, current_start_dash_bonus_day, current_login_bonus_day, last_login_bonus_time, next_login_bonus_time, login_bonus_start_time, login_bonus_end_time)\n"+
		"VALUES ("+id+", :current_start_dash_bonus_day, :current_login_bonus_day, :last_login_bonus_time, :next_login_bonus_time, :login_bonus_start_time, :login_bonus_end_time)",
		value)
	if err == nil && config.CFile.DebugPrints {
		rowsAffected, _ := result.RowsAffected()
		log.Printf("[DEBUG] SetLoginBonusState operation complete; %v rows affected\n", rowsAffected)
	}
	return err
}

func SetOperatorInfo(uid, id, param string) error {
	CheckIfDBSet()
	result, err := db.Exec("REPLACE INTO `" + consts.DBMySQLTableOperatorInfos + "`(uid, id, param)\n" +
		"VALUES (" + uid + ", " + id + ", `" + param + "`)")
	if err == nil && config.CFile.DebugPrints {
		rowsAffected, _ := result.RowsAffected()
		log.Printf("[DEBUG] SetOperatorInfo operation complete; %v rows affected\n", rowsAffected)
	}
	return err
}

func Get(table, column, id string) (interface{}, error) {
	CheckIfDBSet()
	var value interface{}
	err := db.QueryRow("SELECT "+column+" FROM `"+table+"` WHERE id = ?", id).Scan(&value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func GetAnalyticsEntry(table, pid string) ([]byte, error) {
	CheckIfDBSet()
	param := []byte{}
	err := db.QueryRow("SELECT `param` FROM `"+table+"` WHERE pid = ?", pid).Scan(&param)
	if err != nil {
		return []byte{}, err
	}
	return param, nil
}

func GetPlayerInfo(table, id string) (netobj.PlayerInfo, error) {
	CheckIfDBSet()
	values := netobj.StoredPlayerInfo{"", "", "", "", "", 0, 0, []byte{}, []byte{}}
	var id2 int64
	err := db.QueryRow("SELECT * FROM `"+table+"` WHERE id = ?", id).Scan(&id2,
		&values.Username,
		&values.Password,
		&values.MigrationPassword,
		&values.UserPassword,
		&values.Key,
		&values.LastLogin,
		&values.Language,
		&values.CharacterState,
		&values.ChaoState,
	)
	if err != nil {
		return netobj.PlayerInfo{"", "", "", "", "", 0, 0, []netobj.Character{}, []netobj.Chao{}}, err
	}
	return netobj.StoredPlayerInfoToPlayerInfo(values), nil
}

func GetPlayerInfoFromMigrationPass(table, pass string) (netobj.PlayerInfo, string, error) {
	CheckIfDBSet()
	values := netobj.StoredPlayerInfo{"", "", "", "", "", 0, 0, []byte{}, []byte{}}
	var pid string
	err := db.QueryRow("SELECT * FROM `"+table+"` WHERE migrate_password = ?", pass).Scan(&pid,
		&values.Username,
		&values.Password,
		&values.MigrationPassword,
		&values.UserPassword,
		&values.Key,
		&values.LastLogin,
		&values.Language,
		&values.CharacterState,
		&values.ChaoState,
	)
	if err != nil {
		return netobj.PlayerInfo{"", "", "", "", "", 0, 0, []netobj.Character{}, []netobj.Chao{}}, "", err
	}
	return netobj.StoredPlayerInfoToPlayerInfo(values), pid, nil
}

func GetPlayerState(table, id string) (netobj.PlayerState, error) {
	CheckIfDBSet()
	values := netobj.PlayerStateToSQLCompatiblePlayerState(netobj.DefaultPlayerState())
	err := db.QueryRow("SELECT * FROM `"+table+"` WHERE id = ?", id).Scan(
		&values.ID,
		&values.Items,
		&values.EquippedItemIDs,
		&values.MainCharaID,
		&values.SubCharaID,
		&values.MainChaoID,
		&values.SubChaoID,
		&values.NumRings,
		&values.NumBuyRings,
		&values.NumRedRings,
		&values.NumBuyRedRings,
		&values.Energy,
		&values.EnergyBuy,
		&values.EnergyRenewsAt,
		&values.MumMessages,
		&values.RankingLeague,
		&values.QuickRankingLeague,
		&values.NumRouletteTicket,
		&values.NumChaoRouletteTicket,
		&values.ChaoEggs,
		&values.HighScore,
		&values.TimedHighScore,
		&values.TotalDistance,
		&values.HighDistance,
		&values.DailyMissionID,
		&values.DailyMissionEndTime,
		&values.DailyChallengeValue,
		&values.DailyChallengeComplete,
		&values.NumDailyChallenge,
		&values.NumPlaying,
		&values.Animals,
		&values.Rank,
		&values.DailyChalCatNum,
		&values.DailyChalSetNum,
		&values.DailyChalPosNum,
		&values.NextNumDailyChallenge,
		&values.LeagueHighScore,
		&values.QuickLeagueHighScore,
		&values.LeagueStartTime,
		&values.LeagueResetTime,
		&values.RankingLeagueGroup,
		&values.QuickRankingLeagueGroup,
		&values.TotalScore,
		&values.TimedTotalScore,
		&values.HighTotalScore,
		&values.TimedHighTotalScore,
	)
	if err != nil {
		return netobj.DefaultPlayerState(), err
	}
	return netobj.SQLCompatiblePlayerStateToPlayerState(values), nil
}

func GetMileageMapState(table, id string) (netobj.MileageMapState, error) {
	CheckIfDBSet()
	values := netobj.DefaultMileageMapState()
	var id2 int64
	err := db.QueryRow("SELECT * FROM `"+table+"` WHERE id = ?", id).Scan(&id2,
		&values.MapDistance,
		&values.NumBossAttack,
		&values.StageDistance,
		&values.StageMaxScore,
		&values.Episode,
		&values.Chapter,
		&values.Point,
		&values.StageTotalScore,
		&values.ChapterStartTime,
	)
	if err != nil {
		return netobj.DefaultMileageMapState(), err
	}
	return values, nil
}

func GetOptionUserResult(table, id string) (netobj.OptionUserResult, error) {
	CheckIfDBSet()
	values := netobj.DefaultOptionUserResult()
	var id2 int64
	err := db.QueryRow("SELECT * FROM `"+table+"` WHERE id = ?", id).Scan(&id2,
		&values.TotalSumHighScore,
		&values.QuickTotalSumHighScore,
		&values.NumTakeAllRings,
		&values.NumTakeAllRedRings,
		&values.NumChaoRoulette,
		&values.NumItemRoulette,
		&values.NumJackpot,
		&values.NumMaximumJackpotRings,
		&values.NumSupport,
	)
	if err != nil {
		return netobj.DefaultOptionUserResult(), err
	}
	return values, nil
}

func GetRouletteInfo(table, id string) (netobj.RouletteInfo, error) {
	CheckIfDBSet()
	values := netobj.DefaultRouletteInfo()
	var id2 int64
	err := db.QueryRow("SELECT * FROM `"+table+"` WHERE id = ?", id).Scan(&id2,
		&values.LoginRouletteID,
		&values.RoulettePeriodEnd,
		&values.RouletteCountInPeriod,
		&values.GotJackpotThisPeriod,
	)
	if err != nil {
		return netobj.DefaultRouletteInfo(), err
	}
	return values, nil
}

func GetLoginBonusState(table, id string) (netobj.LoginBonusState, error) {
	CheckIfDBSet()
	values := netobj.DefaultLoginBonusState(0)
	var id2 int64
	err := db.QueryRow("SELECT * FROM `"+table+"` WHERE id = ?", id).Scan(&id2,
		&values.CurrentFirstLoginBonusDay,
		&values.CurrentLoginBonusDay,
		&values.LastLoginBonusTime,
		&values.NextLoginBonusTime,
		&values.LoginBonusStartTime,
		&values.LoginBonusEndTime,
	)
	if err != nil {
		return netobj.DefaultLoginBonusState(0), err
	}
	return values, nil
}

func GetOperatorInfos(uid string) ([]obj.OperatorInformation, error) {
	CheckIfDBSet()
	values := []obj.OperatorInformation{}
	rows, err := db.Query("SELECT * FROM `"+consts.DBMySQLTableOperatorInfos+"` WHERE uid = ?", uid)
	if err != nil {
		return []obj.OperatorInformation{}, err
	}
	var id int64
	var param string
	for rows.Next() {
		err = rows.Scan(&uid, &id, &param)
		if err != nil {
			rows.Close()
			return []obj.OperatorInformation{}, err
		}
		values = append(values, obj.NewOperatorInformation(id, param))
	}
	rows.Close()
	return values, nil
}

func GetOperatorMessages(uid string) ([]obj.OperatorMessage, error) {
	CheckIfDBSet()
	values := []obj.OperatorMessage{}
	rows, err := db.Query("SELECT * FROM `"+consts.DBMySQLTableOperatorMessages+"` WHERE userid = ?", uid)
	if err != nil {
		return []obj.OperatorMessage{}, err
	}
	var id, userid, contents, itemjson string
	var expiretime int64
	var item obj.MessageItem
	for rows.Next() {
		err = rows.Scan(&id, &userid, &contents, &itemjson, &expiretime)
		if err != nil {
			rows.Close()
			return []obj.OperatorMessage{}, err
		}
		json.Unmarshal([]byte(itemjson), &item)
		values = append(values, obj.OperatorMessage{id, contents, item, expiretime})
	}
	rows.Close()
	return values, nil
}

func GetEventParam(uid string) (int64, error) {
	CheckIfDBSet()
	var param int64
	err := db.QueryRow("SELECT param FROM `"+consts.DBMySQLTableEventStates+"` WHERE uid = ?", uid).Scan(&uid, &param)
	if err != nil {
		return 0, err
	}
	return param, nil
}

func Delete(table, id string) error {
	CheckIfDBSet()
	_, err := db.Exec("DELETE FROM `"+table+"` WHERE id = ?", id)
	return err
}

func PurgeSessionID(sid string) error {
	CheckIfDBSet()
	_, err := db.Exec("DELETE FROM `"+consts.DBMySQLTableSessionIDs+"` WHERE sid = ?", sid)
	return err
}

func PurgeAllExpiredSessionIDs() error {
	CheckIfDBSet()
	rows, err := db.Query("SELECT sid FROM `"+consts.DBMySQLTableSessionIDs+"` WHERE assigned_at_time > ?", time.Now().Unix()-consts.DBSessionExpiryTime)
	if err != nil {
		return err
	}
	sid := ""
	for rows.Next() {
		err = rows.Scan(&sid)
		if err != nil {
			rows.Close()
			return err
		}
		PurgeSessionID(sid)
	}
	err = rows.Err()
	rows.Close()
	return err
}

func CheckIfDBSet() {
	if db == nil {
		log.Println("[INFO] Connecting to MySQL database...")

		sqldb, err := sqlx.Open("mysql", config.CFile.MySQLUsername+":"+config.CFile.MySQLPassword+"@"+config.CFile.MySQLServerAddress+"/"+config.CFile.MySQLDatabaseName)
		if err != nil {
			log.Println("[FATAL] Failed to open a connection! Check your MySQL settings in config.json for any errors.")
			panic(err)
		}
		err = sqldb.Ping()
		if err != nil {
			log.Println("[FATAL] Failed to connect! Please check your MySQL settings in config.json and try again.")
			panic(err)
		}
		db = sqldb
		log.Println("[INFO] Successfully connected to database!")
	}
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return errors.New("cannot close database if it's not set")
}
