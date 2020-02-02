package dbaccess

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/Mtbcooler/outrun/netobj"

	"github.com/jinzhu/now"

	"github.com/Mtbcooler/outrun/config"
	"github.com/Mtbcooler/outrun/enums"

	"github.com/Mtbcooler/outrun/consts"
	"github.com/Mtbcooler/outrun/obj"
)

// GetHighScores returns the list of high scores, your own entry if applicable, and an error if one is thrown
func GetHighScores(mode, lbtype, offset, limit int64, ownid string, showScoresOfZero bool) ([]obj.LeaderboardEntry, interface{}, error) {
	CheckIfDBSet()
	leagueColumn := "ranking_league"
	if mode == 1 {
		leagueColumn = "quick_ranking_league"
	}
	columnToSortBy := ""
	switch lbtype {
	case 0, 2, 4:
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 1, 3, 5:
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
	additionalQueryOps := ""
	if !showScoresOfZero {
		additionalQueryOps = " WHERE " + columnToSortBy + " > 0"
	}
	rows, err := db.Query("SELECT id, " + columnToSortBy + ", high_score, " + leagueColumn + ", rank, mainchara_id, subchara_id, mainchao_id, subchao_id FROM `" + consts.DBMySQLTablePlayerStates + "`" + additionalQueryOps + " ORDER BY " + columnToSortBy + " DESC LIMIT " + strconv.Itoa(int(limit)) + " OFFSET " + strconv.Itoa(int(offset)))
	if err != nil {
		return []obj.LeaderboardEntry{}, nil, err
	}
	var uid, username, mainchara, subchara, mainchao, subchao, charasjson, chaojson string
	var highscore, maxscore, league, rank, lastlogin, language, maincharalv, subcharalv, mainchaolv, subchaolv int64
	var currentEntry obj.LeaderboardEntry
	var charas []netobj.Character
	var chao []netobj.Chao
	_, resetTime, err := GetStartAndEndTimesForLeague(league, 0)
	if err != nil {
		rows.Close()
		return nil, nil, err
	}
	index := offset
	for rows.Next() {
		err = rows.Scan(&uid, &highscore, &maxscore, &league, &rank, &mainchara, &subchara, &mainchao, &subchao)
		if err != nil {
			rows.Close()
			return []obj.LeaderboardEntry{}, nil, err
		}
		err = db.QueryRow("SELECT username, last_login, language, characters, chao FROM `"+consts.DBMySQLTableCorePlayerInfo+"` WHERE id = ?", uid).Scan(&username, &lastlogin, &language, &charasjson, &chaojson)
		if err != nil {
			rows.Close()
			return []obj.LeaderboardEntry{}, nil, err
		}
		json.Unmarshal([]byte(charasjson), &charas)
		json.Unmarshal([]byte(chaojson), &chao)
		maincharalv = 0
		subcharalv = 0
		for _, char := range charas {
			if char.ID == mainchara {
				maincharalv = char.Level
			}
			if char.ID == subchara {
				subcharalv = char.Level
			}
		}
		mainchaolv = 0
		subchaolv = 0
		for _, c := range chao {
			if c.ID == mainchao {
				mainchaolv = c.Level
			}
			if c.ID == subchao {
				subchaolv = c.Level
			}
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
			maincharalv,
			TryAtoi(subchara),
			subcharalv,
			TryAtoi(mainchao),
			mainchaolv,
			TryAtoi(subchao),
			subchaolv,
			language,
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

func GetOwnLeaderboardEntry(mode, lbtype int64, ownid string, showScoresOfZero bool) (interface{}, error) {
	CheckIfDBSet()
	leagueColumn := "ranking_league"
	if mode == 1 {
		leagueColumn = "quick_ranking_league"
	}
	columnToSortBy := ""
	switch lbtype {
	case 0, 2, 4:
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 1, 3, 5:
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
	additionalQueryOps := ""
	if !showScoresOfZero {
		additionalQueryOps = " WHERE " + columnToSortBy + " > 0"
	}
	var myEntry interface{}
	rows, err := db.Query("SELECT id, " + columnToSortBy + ", high_score, " + leagueColumn + ", rank, mainchara_id, subchara_id, mainchao_id, subchao_id FROM `" + consts.DBMySQLTablePlayerStates + "` ORDER BY " + columnToSortBy + " DESC" + additionalQueryOps)
	if err != nil {
		return nil, err
	}
	var uid, username, mainchara, subchara, mainchao, subchao, charasjson, chaojson string
	var highscore, maxscore, league, rank, lastlogin, language, maincharalv, subcharalv, mainchaolv, subchaolv int64
	var charas []netobj.Character
	var chao []netobj.Chao
	index := 0
	for rows.Next() {
		err = rows.Scan(&uid, &highscore, &maxscore, &league, &rank, &mainchara, &subchara, &mainchao, &subchao)
		if err != nil {
			rows.Close()
			return nil, err
		}
		if uid == ownid {
			err = db.QueryRow("SELECT username, last_login, language, characters, chao FROM `"+consts.DBMySQLTableCorePlayerInfo+"` WHERE id = ?", uid).Scan(&username, &lastlogin, &language, &charasjson, &chaojson)
			if err != nil {
				rows.Close()
				return nil, err
			}
			_, resetTime, err := GetStartAndEndTimesForLeague(league, 0)
			if err != nil {
				rows.Close()
				return nil, err
			}
			json.Unmarshal([]byte(charasjson), &charas)
			json.Unmarshal([]byte(chaojson), &chao)
			maincharalv = 0
			subcharalv = 0
			for _, char := range charas {
				if char.ID == mainchara {
					maincharalv = char.Level
				}
				if char.ID == subchara {
					subcharalv = char.Level
				}
			}
			mainchaolv = 0
			subchaolv = 0
			for _, c := range chao {
				if c.ID == mainchao {
					mainchaolv = c.Level
				}
				if c.ID == subchao {
					subchaolv = c.Level
				}
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
				maincharalv,
				TryAtoi(subchara),
				subcharalv,
				TryAtoi(mainchao),
				mainchaolv,
				TryAtoi(subchao),
				subchaolv,
				enums.LangEnglish,
				league,
				maxscore,
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

func GetNumOfLeaderboardPlayers(mode, lbtype int64) (int64, error) {
	CheckIfDBSet()
	var hsc string
	switch lbtype {
	case 0, 2, 4: // League high score
		if mode == 1 {
			hsc = "quick_league_high_score"
		} else {
			hsc = "league_high_score"
		}
	case 1, 3, 5: // League total score
		if mode == 1 {
			hsc = "quick_league_high_score"
		} else {
			hsc = "league_high_score"
		}
	case 6:
		if mode == 1 {
			hsc = "quick_high_score"
		} else {
			hsc = "high_score"
		}
	case 7:
		if mode == 1 {
			hsc = "quick_high_score"
		} else {
			hsc = "high_score"
		}
	default:
		hsc = "high_score"
	}
	playercount := int64(0)
	err := db.QueryRow("SELECT COUNT(*) FROM `" + consts.DBMySQLTablePlayerStates + "` WHERE " + hsc + " > 0").Scan(&playercount)
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
	result, err := db.Exec("UPDATE `" + consts.DBMySQLTablePlayerStates + "` SET league_high_score = 0, quick_league_high_score = 0")
	if err == nil && config.CFile.DebugPrints {
		rowsAffected, _ = result.RowsAffected()
		log.Printf("[DEBUG] ClearLeagueHighScores operation completed; %v rows affected\n", rowsAffected)
	}
	return err
}
