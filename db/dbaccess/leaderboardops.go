package dbaccess

import (
	"log"
	"strconv"
	"time"

	"github.com/jinzhu/now"

	"github.com/Mtbcooler/outrun/config"
	"github.com/Mtbcooler/outrun/enums"

	"github.com/Mtbcooler/outrun/consts"
	"github.com/Mtbcooler/outrun/obj"
)

// GetHighScores returns the list of high scores, your own entry if applicable, and an error if one is thrown
func GetHighScores(mode, lbtype, offset, limit int64, ownid string) ([]obj.LeaderboardEntry, interface{}, error) {
	CheckIfDBSet()
	leagueColumn := "ranking_league"
	if mode == 1 {
		leagueColumn = "quick_ranking_league"
	}
	columnToSortBy := ""
	switch lbtype {
	case 0: // Friends high score?
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 1: // Friends total score?
		// TODO: Define total_score column
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 2: // World high score
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 3: // World total score
		// TODO: Define total_score column
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 4: // Runners' League high score
		log.Println("[WARN] Please use GetLeagueHighScores() for getting Runners' League high scores!")
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 5: // Runners' League total score
		log.Println("[WARN] Please use GetLeagueHighScores() for getting Runners' League high scores!")
		// TODO: Define total_score column
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 6: // Historical high score
		if mode == 1 {
			columnToSortBy = "quick_high_score"
		} else {
			columnToSortBy = "high_score"
		}
	case 7: // Historical total score
		// TODO: Define high_total_score column
		if mode == 1 {
			columnToSortBy = "quick_high_score"
		} else {
			columnToSortBy = "high_score"
		}
	default:
		log.Printf("[WARN] Unknown leaderboard type %v", lbtype)
		columnToSortBy = "high_score"
	}
	var myEntry interface{}
	leaderboardEntries := []obj.LeaderboardEntry{}
	rows, err := db.Query("SELECT id, " + columnToSortBy + ", " + leagueColumn + ", rank, mainchara_id, subchara_id, mainchao_id, subchao_id FROM `" + consts.DBMySQLTablePlayerStates + "` ORDER BY " + columnToSortBy + " DESC LIMIT " + strconv.Itoa(int(limit)) + " OFFSET " + strconv.Itoa(int(offset)))
	if err != nil {
		return []obj.LeaderboardEntry{}, nil, err
	}
	var uid, username, mainchara, subchara, mainchao, subchao, charasjson, chaojson string
	var highscore, league, rank, lastlogin int64
	var currentEntry obj.LeaderboardEntry
	_, resetTime, err := GetStartAndEndTimesForLeague(league, 0)
	if err != nil {
		rows.Close()
		return nil, nil, err
	}
	index := offset
	for rows.Next() {
		err = rows.Scan(&uid, &highscore, &league, &rank, &mainchara, &subchara, &mainchao, &subchao)
		if err != nil {
			rows.Close()
			return []obj.LeaderboardEntry{}, nil, err
		}
		err = db.QueryRow("SELECT username, last_login, characters, chao FROM `"+consts.DBMySQLTableCorePlayerInfo+"` WHERE id = ?", uid).Scan(&username, &lastlogin, &charasjson, &chaojson)
		if err != nil {
			rows.Close()
			return []obj.LeaderboardEntry{}, nil, err
		}
		currentEntry = obj.NewLeaderboardEntry(
			uid,
			username,
			"",
			int64(index+1),
			0,
			highscore,
			0,
			0,
			resetTime,
			rank,
			lastlogin,
			TryAtoi(mainchara),
			0,
			TryAtoi(subchara),
			0,
			TryAtoi(mainchao),
			0,
			TryAtoi(subchao),
			0,
			enums.LangEnglish,
			league,
			highscore,
			0,
		)
		if uid == ownid {
			myEntry = currentEntry
		}
		leaderboardEntries = append(leaderboardEntries, currentEntry)
		index++
	}
	err = rows.Err()
	rows.Close()
	return leaderboardEntries, myEntry, err
}

func GetOwnLeaderboardEntry(mode, lbtype int64, ownid string) (interface{}, error) {
	CheckIfDBSet()
	leagueColumn := "ranking_league"
	if mode == 1 {
		leagueColumn = "quick_ranking_league"
	}
	columnToSortBy := ""
	switch lbtype {
	case 0: // Friends high score?
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 1: // Friends total score?
		// TODO: Define total_score column
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 2: // World high score
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 3: // World total score
		// TODO: Define total_score column
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 4: // Runners' League high score
		log.Println("[WARN] Please use GetLeagueHighScores() for getting Runners' League high scores!")
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 5: // Runners' League total score
		log.Println("[WARN] Please use GetLeagueHighScores() for getting Runners' League high scores!")
		// TODO: Define total_score column
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 6: // Historical high score
		if mode == 1 {
			columnToSortBy = "quick_high_score"
		} else {
			columnToSortBy = "high_score"
		}
	case 7: // Historical total score
		// TODO: Define high_total_score column
		if mode == 1 {
			columnToSortBy = "quick_high_score"
		} else {
			columnToSortBy = "high_score"
		}
	default:
		log.Printf("[WARN] Unknown leaderboard type %v", lbtype)
		columnToSortBy = "high_score"
	}
	var myEntry interface{}
	rows, err := db.Query("SELECT id, " + columnToSortBy + ", " + leagueColumn + ", rank, mainchara_id, subchara_id, mainchao_id, subchao_id FROM `" + consts.DBMySQLTablePlayerStates + "` ORDER BY " + columnToSortBy + " DESC")
	if err != nil {
		return nil, err
	}
	var uid, username, mainchara, subchara, mainchao, subchao, charasjson, chaojson string
	var highscore, league, rank, lastlogin int64
	index := 0
	for rows.Next() {
		err = rows.Scan(&uid, &highscore, &league, &rank, &mainchara, &subchara, &mainchao, &subchao)
		if err != nil {
			rows.Close()
			return nil, err
		}
		if uid == ownid {
			err = db.QueryRow("SELECT username, last_login, characters, chao FROM `"+consts.DBMySQLTableCorePlayerInfo+"` WHERE id = ?", uid).Scan(&username, &lastlogin, &charasjson, &chaojson)
			if err != nil {
				rows.Close()
				return nil, err
			}
			_, resetTime, err := GetStartAndEndTimesForLeague(league, 0)
			if err != nil {
				rows.Close()
				return nil, err
			}
			myEntry = obj.NewLeaderboardEntry(
				uid,
				username,
				"",
				int64(index+1),
				0,
				highscore,
				0,
				0,
				resetTime,
				rank,
				lastlogin,
				TryAtoi(mainchara),
				0,
				TryAtoi(subchara),
				0,
				TryAtoi(mainchao),
				0,
				TryAtoi(subchao),
				0,
				enums.LangEnglish,
				league,
				highscore,
				0,
			)
		}
	}
	err = rows.Err()
	rows.Close()
	return myEntry, err
}

func GetNumOfPlayers() (int64, error) {
	CheckIfDBSet()
	playercount := int64(0)
	err := db.QueryRow("SELECT COUNT(*) FROM `" + consts.DBMySQLTablePlayerStates + "`").Scan(&playercount)
	if err != nil {
		return -1, err
	}
	return playercount, nil
}

func TryAtoi(toconvert string) int64 {
	if toconvert == "empty" {
		return -1
	}
	result, _ := strconv.Atoi(toconvert)
	return int64(result)
}

func GetStartAndEndTimesForLeague(leagueid, groupid int64) (int64, int64, error) {
	CheckIfDBSet()
	var starttime, endtime int64
	err := db.QueryRow("SELECT start_time, reset_time FROM `"+consts.DBMySQLTableRankingLeagueData+"` WHERE league_id = ? AND group_id = ?", leagueid, groupid).Scan(&starttime, &endtime)
	if err != nil {
		return 0, 0, err
	}
	return starttime, endtime, nil
}

func SetRankingLeagueData(leagueid, groupid, starttime, endtime, leagueplayercount, groupplayercount int64) error {
	CheckIfDBSet()
	result, err := db.Exec("REPLACE INTO `"+consts.DBMySQLTableRankingLeagueData+"`(league_id, group_id, start_time, reset_time, league_player_count, group_player_count)\n"+
		"VALUES (?,?,?,?,?,?)",
		leagueid, groupid, starttime, endtime, leagueplayercount, groupplayercount)
	if err == nil && config.CFile.DebugPrints {
		rowsAffected, _ := result.RowsAffected()
		log.Printf("[DEBUG] SetRankingLeagueData operation complete; %v rows affected\n", rowsAffected)
	}
	return err
}

func ResetAllRankingLeagueData() error {
	CheckIfDBSet()
	rowsAffected := int64(0)
	result, err := db.Exec("DROP TABLE `" + consts.DBMySQLTableRankingLeagueData + "`")
	if err == nil && config.CFile.DebugPrints {
		rowsAffected, _ = result.RowsAffected()
		log.Printf("[DEBUG] Ranking League data wiped; %v rows affected\n", rowsAffected)
	}
	if err != nil {
		return err
	}
	result, err = db.Exec(consts.SQLRankingLeagueDataSchema)
	if err == nil && config.CFile.DebugPrints {
		rowsAffected, _ = result.RowsAffected()
		log.Printf("[DEBUG] Ranking League data table created; %v rows affected\n", rowsAffected)
	}
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueF_M, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueF, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueF_P, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueE_M, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueE, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueE_P, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueD_M, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueD, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueD_P, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueC_M, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueC, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueC_P, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueB_M, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueB, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueB_P, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueA_M, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueA, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueA_P, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueS_M, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueS, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	if err != nil {
		return err
	}
	err = SetRankingLeagueData(enums.RankingLeagueS_P, 0, time.Now().UTC().Unix(), now.EndOfWeek().UTC().Unix(), 50, 50)
	return err
}

func ClearLeagueHighScores() error {
	CheckIfDBSet()
	rowsAffected := int64(0)
	result, err := db.Exec("UPDATE `" + consts.DBMySQLTablePlayerStates + "` league_high_scores = 0, quick_league_high_scores = 0")
	if err == nil && config.CFile.DebugPrints {
		rowsAffected, _ = result.RowsAffected()
		log.Printf("[DEBUG] ClearLeagueHighScores operation completed; %v rows affected\n", rowsAffected)
	}
	return err
}
